package report

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"html/template"

	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/boltdb/bolt"
)

// Values hold the values of our report
type Values struct {
	Message      string
	Projname     string
	Filename     string
	FacilityType string
}

const rt string = `
{{define "T"}}
<div class="{{.Message}}">
<div class="pn"> {{.Projname}} </div>
<div class="fn"> {{.Filename}} </div>
<div class="ft"> {{.FacilityType}} </div>
</div>
{{end}}

`

func (v *Values) save(db *bolt.DB) error {
	// Store the user model in the user bucket using the username as the key.
	err := db.Update(func(tx *bolt.Tx) error {
		b, err := tx.CreateBucketIfNotExists([]byte("filedata"))
		if err != nil {
			return err
		}

		encoded, err := json.Marshal(v)
		if err != nil {
			return err
		}
		return b.Put([]byte(v.Filename), encoded) // TODO..  can't use filename as key..  NOT unique
	})
	return err
}

// GenReport will generate the report type we wish
func GenReport(message, projname, filename, facilityType string) {
	v := Values{Message: message, Projname: projname, Filename: filename, FacilityType: facilityType}
	// fmt.Printf("%s:%s type:%s status:%s \n", projname, filename, facilityType, message)

	t, err := template.New("t").Parse(rt)
	if err != nil {
		log.Println("Error in template parse")
	}
	err = t.ExecuteTemplate(os.Stdout, "T", v)
}

func InitNotebook() *excelize.File {
	xlsx := excelize.NewFile()
	//index := xlsx.NewSheet("Sheet1")

	return xlsx
}

func SaveNotebook(file *excelize.File) error {
	err := file.SaveAs("./Book1.xlsx")
	if err != nil {
		fmt.Println(err)
	}

	return err
}

// WriteNotebookRow
// TODO..   encode into here...  what I need to build a package with...
// make a packageID that notes all the files in a given package  based on the projname likely
func WriteNotebookRow(row int, xlsx *excelize.File, message, projname, filename, facilityType string) (int, error) {
	v := Values{Message: message, Projname: projname, Filename: filename, FacilityType: facilityType}
	fmt.Println(v)

	//nr := len(xlsx.GetRows("Sheet1")) + 1
	sc := fmt.Sprintf("A%d", row)
	fmt.Println(sc)

	xlsx.SetSheetRow("Sheet1", sc, &[]interface{}{v.Filename, v.FacilityType, v.Message, v.Projname})

	//	return nil
	return row + 1, nil
}
