package csvproc

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/knakk/rdf"
	"opencoredata.org/ocdGarden/CSDCO/GraphBuilder/utils"
)

func BuildGraph() {
	log.Println("In CSV proc")

	csvdata := utils.ReadMetaData()
	tr := []rdf.Triple{}

	for _, item := range csvdata {
		newsub, _ := rdf.NewIRI(fmt.Sprintf("http://opencoredata.org/id/resource/csdco/feature/%s", strings.ToLower(item.HoleID))) // Sprintf a correct URI here

		newpred1, _ := rdf.NewIRI("http://opencoredata.org/id/voc/csdco/v1/locationname")
		newpred2, _ := rdf.NewIRI("http://opencoredata.org/id/voc/csdco/v1/locationtype")
		newpred3, _ := rdf.NewIRI("http://opencoredata.org/id/voc/csdco/v1/project")
		newpred4, _ := rdf.NewIRI("http://opencoredata.org/id/voc/csdco/v1/locationid")
		newpred5, _ := rdf.NewIRI("http://opencoredata.org/id/voc/csdco/v1/site")
		newpred6, _ := rdf.NewIRI("http://opencoredata.org/id/voc/csdco/v1/hole")
		newpred7, _ := rdf.NewIRI("http://opencoredata.org/id/voc/csdco/v1/sitehole")
		newpred8, _ := rdf.NewIRI("http://opencoredata.org/id/voc/csdco/v1/originalid")
		newpred9, _ := rdf.NewIRI("http://opencoredata.org/id/voc/csdco/v1/holeid")
		newpred10, _ := rdf.NewIRI("http://opencoredata.org/id/voc/csdco/v1/platform")
		newpred11, _ := rdf.NewIRI("http://opencoredata.org/id/voc/csdco/v1/date") // DCAT:issued
		newpred12, _ := rdf.NewIRI("http://opencoredata.org/id/voc/csdco/v1/waterdepthm")
		newpred13, _ := rdf.NewIRI("http://opencoredata.org/id/voc/csdco/v1/country")
		newpred14, _ := rdf.NewIRI("http://opencoredata.org/id/voc/csdco/v1/state_province")
		newpred15, _ := rdf.NewIRI("http://opencoredata.org/id/voc/csdco/v1/county_region")
		newpred16, _ := rdf.NewIRI("http://opencoredata.org/id/voc/csdco/v1/pi")
		newpred17, _ := rdf.NewIRI("http://www.w3.org/2003/01/geo/wgs84_pos#lat")
		newpred18, _ := rdf.NewIRI("http://www.w3.org/2003/01/geo/wgs84_pos#long")
		newpred19, _ := rdf.NewIRI("http://opencoredata.org/id/voc/csdco/v1/elevation")
		newpred20, _ := rdf.NewIRI("http://opencoredata.org/id/voc/csdco/v1/position")
		newpred21, _ := rdf.NewIRI("http://opencoredata.org/id/voc/csdco/v1/storagelocation")
		newpred22, _ := rdf.NewIRI("http://opencoredata.org/id/voc/csdco/v1/storagelocation")
		newpred23, _ := rdf.NewIRI("http://opencoredata.org/id/voc/csdco/v1/sampletype")
		newpred24, _ := rdf.NewIRI("http://opencoredata.org/id/voc/csdco/v1/comment")
		newpred25, _ := rdf.NewIRI("http://opencoredata.org/id/voc/csdco/v1/mblft")
		newpred26, _ := rdf.NewIRI("http://opencoredata.org/id/voc/csdco/v1/mblfb")
		newpred27, _ := rdf.NewIRI("http://opencoredata.org/id/voc/csdco/v1/metadatasource")
		newpred28, _ := rdf.NewIRI("http://opencoredata.org/id/voc/csdco/v1/latlong")

		newobj1, _ := rdf.NewLiteral(utils.StripCtlAndExtFromUnicode(item.LocationName))
		newobj2, _ := rdf.NewLiteral(item.LocationType)
		newobj3, _ := rdf.NewLiteral(item.Project)
		newobj4, _ := rdf.NewLiteral(item.LocationID)
		newobj5, _ := rdf.NewLiteral(item.Site)
		newobj6, _ := rdf.NewLiteral(item.Hole)
		newobj7, _ := rdf.NewLiteral(item.SiteHole)
		newobj8, _ := rdf.NewLiteral(item.OriginalID)
		newobj9, _ := rdf.NewLiteral(item.HoleID)
		newobj10, _ := rdf.NewLiteral(item.Platform)
		newobj11, _ := rdf.NewLiteral(item.Date)
		newobj12, _ := rdf.NewLiteral(item.WaterDepthM)
		newobj13, _ := rdf.NewLiteral(item.Country)
		newobj14, _ := rdf.NewLiteral(utils.StripCtlAndExtFromUnicode(item.State_Province))
		newobj15, _ := rdf.NewLiteral(utils.StripCtlAndExtFromUnicode(item.County_Region))
		newobj16, _ := rdf.NewLiteral(utils.StripCtlAndExtFromUnicode(item.PI))
		newobj17, _ := rdf.NewLiteral(item.Lat)
		newobj18, _ := rdf.NewLiteral(item.Long)
		newobj19, _ := rdf.NewLiteral(item.Elevation)
		newobj20, _ := rdf.NewLiteral(item.Position)
		newobj21, _ := rdf.NewLiteral(item.StorageLocationWorking)
		newobj22, _ := rdf.NewLiteral(item.StorageLocationArchive)
		newobj23, _ := rdf.NewLiteral(item.SampleType)
		newobj24, _ := rdf.NewLiteral(item.Comment)
		newobj25, _ := rdf.NewLiteral(item.MblfT)
		newobj26, _ := rdf.NewLiteral(item.MblfB)
		newobj27, _ := rdf.NewLiteral(item.MetadataSource)

		// Blazegraph specific typed literal for spatial search...  (sadly a vender BIF)
		nt28, _ := rdf.NewIRI("http://www.bigdata.com/rdf/geospatial/literals/v1#geoliteral:lat-lon")
		newobj28 := rdf.NewTypedLiteral(fmt.Sprintf("%s#%s", item.Lat, item.Long), nt28)

		newtriple2 := rdf.Triple{Subj: newsub, Pred: newpred1, Obj: newobj1}
		newtriple1 := rdf.Triple{Subj: newsub, Pred: newpred2, Obj: newobj2}
		newtriple3 := rdf.Triple{Subj: newsub, Pred: newpred3, Obj: newobj3}
		newtriple4 := rdf.Triple{Subj: newsub, Pred: newpred4, Obj: newobj4}
		newtriple5 := rdf.Triple{Subj: newsub, Pred: newpred5, Obj: newobj5}
		newtriple6 := rdf.Triple{Subj: newsub, Pred: newpred6, Obj: newobj6}
		newtriple7 := rdf.Triple{Subj: newsub, Pred: newpred7, Obj: newobj7}
		newtriple8 := rdf.Triple{Subj: newsub, Pred: newpred8, Obj: newobj8}
		newtriple9 := rdf.Triple{Subj: newsub, Pred: newpred9, Obj: newobj9}
		newtriple10 := rdf.Triple{Subj: newsub, Pred: newpred10, Obj: newobj10}
		newtriple11 := rdf.Triple{Subj: newsub, Pred: newpred11, Obj: newobj11}
		newtriple12 := rdf.Triple{Subj: newsub, Pred: newpred12, Obj: newobj12}
		newtriple13 := rdf.Triple{Subj: newsub, Pred: newpred13, Obj: newobj13}
		newtriple14 := rdf.Triple{Subj: newsub, Pred: newpred14, Obj: newobj14}
		newtriple15 := rdf.Triple{Subj: newsub, Pred: newpred15, Obj: newobj15}
		newtriple16 := rdf.Triple{Subj: newsub, Pred: newpred16, Obj: newobj16}
		newtriple17 := rdf.Triple{Subj: newsub, Pred: newpred17, Obj: newobj17}
		newtriple18 := rdf.Triple{Subj: newsub, Pred: newpred18, Obj: newobj18}
		newtriple19 := rdf.Triple{Subj: newsub, Pred: newpred19, Obj: newobj19}
		newtriple20 := rdf.Triple{Subj: newsub, Pred: newpred20, Obj: newobj20}
		newtriple21 := rdf.Triple{Subj: newsub, Pred: newpred21, Obj: newobj21}
		newtriple22 := rdf.Triple{Subj: newsub, Pred: newpred22, Obj: newobj22}
		newtriple23 := rdf.Triple{Subj: newsub, Pred: newpred23, Obj: newobj23}
		newtriple24 := rdf.Triple{Subj: newsub, Pred: newpred24, Obj: newobj24}
		newtriple25 := rdf.Triple{Subj: newsub, Pred: newpred25, Obj: newobj25}
		newtriple26 := rdf.Triple{Subj: newsub, Pred: newpred26, Obj: newobj26}
		newtriple27 := rdf.Triple{Subj: newsub, Pred: newpred27, Obj: newobj27}
		newtriple28 := rdf.Triple{Subj: newsub, Pred: newpred28, Obj: newobj28}

		newpred0, _ := rdf.NewIRI("http://www.w3.org/1999/02/22-rdf-syntax-ns#type")
		newobj0, _ := rdf.NewIRI("http://opencoredata.org/id/voc/csdco/v1/CSDCOProject")
		newtriple0 := rdf.Triple{Subj: newsub, Pred: newpred0, Obj: newobj0}

		tr = append(tr, newtriple0)

		if newtriple1.Obj.String() != "" {
			tr = append(tr, newtriple1)
		}
		if newtriple2.Obj.String() != "" {
			tr = append(tr, newtriple2)
		}
		if newtriple3.Obj.String() != "" {
			tr = append(tr, newtriple3)
		}
		if newtriple4.Obj.String() != "" {
			tr = append(tr, newtriple4)
		}
		if newtriple5.Obj.String() != "" {
			tr = append(tr, newtriple5)
		}
		if newtriple6.Obj.String() != "" {
			tr = append(tr, newtriple6)
		}
		if newtriple7.Obj.String() != "" {
			tr = append(tr, newtriple7)
		}
		if newtriple8.Obj.String() != "" {
			tr = append(tr, newtriple8)
		}
		if newtriple9.Obj.String() != "" {
			tr = append(tr, newtriple9)
		}
		if newtriple10.Obj.String() != "" {
			tr = append(tr, newtriple10)
		}
		if newtriple11.Obj.String() != "" {
			tr = append(tr, newtriple11)
		}
		if newtriple12.Obj.String() != "" {
			tr = append(tr, newtriple12)
		}
		if newtriple13.Obj.String() != "" {
			tr = append(tr, newtriple13)
		}
		if newtriple14.Obj.String() != "" {
			tr = append(tr, newtriple14)
		}
		if newtriple15.Obj.String() != "" {
			tr = append(tr, newtriple15)
		}
		if newtriple16.Obj.String() != "" {
			tr = append(tr, newtriple16)
		}
		if newtriple17.Obj.String() != "" {
			tr = append(tr, newtriple17)
		}
		if newtriple18.Obj.String() != "" {
			tr = append(tr, newtriple18)
		}
		if newtriple19.Obj.String() != "" {
			tr = append(tr, newtriple19)
		}
		if newtriple20.Obj.String() != "" {
			tr = append(tr, newtriple20)
		}
		if newtriple21.Obj.String() != "" {
			tr = append(tr, newtriple21)
		}
		if newtriple22.Obj.String() != "" {
			tr = append(tr, newtriple22)
		}
		if newtriple23.Obj.String() != "" {
			tr = append(tr, newtriple23)
		}
		if newtriple24.Obj.String() != "" {
			tr = append(tr, newtriple24)
		}
		if newtriple25.Obj.String() != "" {
			tr = append(tr, newtriple25)
		}
		if newtriple26.Obj.String() != "" {
			tr = append(tr, newtriple26)
		}
		if newtriple27.Obj.String() != "" {
			tr = append(tr, newtriple27)
		}
		if newtriple28.Obj.String() != "" {
			tr = append(tr, newtriple28)
		}
	}

	// Create the output file
	outFile, err := os.Create("csdcoProjects.nt")
	if err != nil {
		panic(err)
	}
	defer outFile.Close()

	// Write triples to a file
	var inoutFormat rdf.Format
	inoutFormat = rdf.NTriples // Turtle NQuads

	enc := rdf.NewTripleEncoder(outFile, inoutFormat)
	err = enc.EncodeAll(tr)
	// err = enc.Encode(newtriple)
	enc.Close()
	if err != nil {
		log.Fatal(err)
	}

}
