package igsnGraph

import (
	"fmt"
	"log"
	"os"
	"reflect"
	"strings"

	"github.com/knakk/rdf"
	"opencoredata.org/ocdGarden/CSDCO/GraphBuilder/utils"
)

func BuildGraph() {
	log.Println("In CSV proc")

	csvdata := utils.ReadIGSNData()
	tr := []rdf.Triple{}

	for _, item := range csvdata {
		// test:  reflection approach to remove line count sadness..

		e := reflect.ValueOf(&item).Elem()

		for i := 0; i < e.NumField(); i++ {
			varName := e.Type().Field(i).Name
			varType := e.Type().Field(i).Type
			varValue := e.Field(i).Interface()
			fmt.Printf("%v %v %v\n", varName, varType, varValue)
		}

		// end test

		// TODO   location and archive are missing...

		newsub, _ := rdf.NewIRI(fmt.Sprintf("http://opencoredata.org/id/resource/csdco/feature/%s", strings.ToLower(item.Holeid))) // Sprintf a correct URI here

		newpred1, _ := rdf.NewIRI("http://opencoredata.org/id/voc/csdco/v1/ngdc_ship")
		newpred2, _ := rdf.NewIRI("http://opencoredata.org/id/voc/csdco/v1/ngdc_cruise")
		newpred3, _ := rdf.NewIRI("http://opencoredata.org/id/voc/csdco/v1/ngdc_sample")
		newpred4, _ := rdf.NewIRI("http://opencoredata.org/id/voc/csdco/v1/ngdc_serial")
		newpred5, _ := rdf.NewIRI("http://opencoredata.org/id/voc/csdco/v1/ngdc_comment")
		newpred6, _ := rdf.NewIRI("http://opencoredata.org/id/voc/csdco/v1/location")
		newpred7, _ := rdf.NewIRI("http://opencoredata.org/id/voc/csdco/v1/locationtype")
		newpred8, _ := rdf.NewIRI("http://opencoredata.org/id/voc/csdco/v1/expedition")
		newpred9, _ := rdf.NewIRI("http://opencoredata.org/id/voc/csdco/v1/locationid")
		newpred10, _ := rdf.NewIRI("http://opencoredata.org/id/voc/csdco/v1/site")
		newpred11, _ := rdf.NewIRI("http://opencoredata.org/id/voc/csdco/v1/hole")
		newpred12, _ := rdf.NewIRI("http://opencoredata.org/id/voc/csdco/v1/sitehole")
		newpred13, _ := rdf.NewIRI("http://opencoredata.org/id/voc/csdco/v1/originalid")
		newpred14, _ := rdf.NewIRI("http://opencoredata.org/id/voc/csdco/v1/holeid")
		newpred15, _ := rdf.NewIRI("http://opencoredata.org/id/voc/csdco/v1/platform")
		newpred16, _ := rdf.NewIRI("http://opencoredata.org/id/voc/csdco/v1/date")
		newpred17, _ := rdf.NewIRI("http://opencoredata.org/id/voc/csdco/v1/waterdepthm")
		newpred18, _ := rdf.NewIRI("http://opencoredata.org/id/voc/csdco/v1/country")
		newpred19, _ := rdf.NewIRI("http://opencoredata.org/id/voc/csdco/v1/state_province")
		newpred20, _ := rdf.NewIRI("http://opencoredata.org/id/voc/csdco/v1/county_region")
		newpred21, _ := rdf.NewIRI("http://opencoredata.org/id/voc/csdco/v1/pi")
		newpred22, _ := rdf.NewIRI("http://www.w3.org/2003/01/geo/wgs84_pos#lat")
		newpred23, _ := rdf.NewIRI("http://www.w3.org/2003/01/geo/wgs84_pos#long")
		newpred24, _ := rdf.NewIRI("http://opencoredata.org/id/voc/csdco/v1/elevation")
		newpred25, _ := rdf.NewIRI("http://opencoredata.org/id/voc/csdco/v1/position")
		newpred26, _ := rdf.NewIRI("http://opencoredata.org/id/voc/csdco/v1/sampletype")
		newpred27, _ := rdf.NewIRI("http://opencoredata.org/id/voc/csdco/v1/comment")
		newpred28, _ := rdf.NewIRI("http://opencoredata.org/id/voc/csdco/v1/mblft")
		newpred29, _ := rdf.NewIRI("http://opencoredata.org/id/voc/csdco/v1/mblfb")
		newpred30, _ := rdf.NewIRI("http://opencoredata.org/id/voc/csdco/v1/metadatasource")
		newpred31, _ := rdf.NewIRI("http://opencoredata.org/id/voc/csdco/v1/googleearthcommentfield")
		newpred32, _ := rdf.NewIRI("http://opencoredata.org/id/voc/csdco/v1/platformname")
		newpred33, _ := rdf.NewIRI("http://opencoredata.org/id/voc/csdco/v1/platformtype")
		newpred34, _ := rdf.NewIRI("http://opencoredata.org/id/voc/csdco/v1/azimuth")
		newpred35, _ := rdf.NewIRI("http://opencoredata.org/id/voc/csdco/v1/dip")
		newpred36, _ := rdf.NewIRI("http://pid.geoscience.gov.au/def/ont/ga/igsn#Sample")
		newpred37, _ := rdf.NewIRI("http://www.w3.org/2000/01/rdf-schema#label")
		newpred38, _ := rdf.NewIRI("http://www.w3.org/2000/01/rdf-schema#comment")

		// For sample class, we need a resolvable URI then..
		// Ref http://ldweb.ga.gov.au/def/ont/ga/igsn/igsn.html#classes for cardinality issues
		// Is there a classic resolve like dx?
		// Use a schema description model like in DataCite K4 to set the triples..

		newobj1, _ := rdf.NewLiteral(item.Ngdcship)
		newobj2, _ := rdf.NewLiteral(item.Ngdccruise)
		newobj3, _ := rdf.NewLiteral(item.Ngdcsample)
		newobj4, _ := rdf.NewLiteral(item.Ngdcserial)
		newobj5, _ := rdf.NewLiteral(item.Ngdccomment)
		newobj6, _ := rdf.NewLiteral(utils.StripCtlAndExtFromUnicode(item.Location))
		newobj7, _ := rdf.NewLiteral(item.Locationtype)
		newobj8, _ := rdf.NewLiteral(item.Expedition)
		newobj9, _ := rdf.NewLiteral(item.Locationid)
		newobj10, _ := rdf.NewLiteral(item.Site)
		newobj11, _ := rdf.NewLiteral(item.Hole)
		newobj12, _ := rdf.NewLiteral(item.Sitehole)
		newobj13, _ := rdf.NewLiteral(item.Originalid)
		newobj14, _ := rdf.NewLiteral(item.Holeid)
		newobj15, _ := rdf.NewLiteral(item.Platform)
		newobj16, _ := rdf.NewLiteral(item.Date)
		newobj17, _ := rdf.NewLiteral(item.Waterdepthm)
		newobj18, _ := rdf.NewLiteral(item.Country)
		newobj19, _ := rdf.NewLiteral(utils.StripCtlAndExtFromUnicode(item.State_province))
		newobj20, _ := rdf.NewLiteral(utils.StripCtlAndExtFromUnicode(item.County_region))
		newobj21, _ := rdf.NewLiteral(utils.StripCtlAndExtFromUnicode(item.Pi))
		newobj22, _ := rdf.NewLiteral(item.Lat)
		newobj23, _ := rdf.NewLiteral(item.Long)
		newobj24, _ := rdf.NewLiteral(item.Elevation)
		newobj25, _ := rdf.NewLiteral(item.Position)
		newobj26, _ := rdf.NewLiteral(item.Sampletype)
		newobj27, _ := rdf.NewLiteral(item.Comment)
		newobj28, _ := rdf.NewLiteral(item.Mblft)
		newobj29, _ := rdf.NewLiteral(item.Mblfb)
		newobj30, _ := rdf.NewLiteral(item.Metadatasource)
		newobj31, _ := rdf.NewLiteral(item.Googleearthcommentfield)
		newobj32, _ := rdf.NewLiteral(item.Platformname)
		newobj33, _ := rdf.NewLiteral(item.Platformtype)
		newobj34, _ := rdf.NewLiteral(item.Azimuth)
		newobj35, _ := rdf.NewLiteral(item.Dip)
		newobj36, _ := rdf.NewLiteral(item.Igsn)
		newobj37, _ := rdf.NewLiteral(item.Holeid)
		newobj38, _ := rdf.NewLiteral(fmt.Sprintf("Project %s and Hole %s in country %s, state %s, county %s,  location %s ( %s ), with PIs %s",
			item.Expedition, item.Holeid, item.Country,
			utils.StripCtlAndExtFromUnicode(item.State_province), utils.StripCtlAndExtFromUnicode(item.County_region), item.Locationid,
			utils.StripCtlAndExtFromUnicode(item.Location), utils.StripCtlAndExtFromUnicode(item.Pi)))

		// Blazegraph specific typed literal for spatial search...  (sadly a vender BIF)
		nt39, _ := rdf.NewIRI("http://www.bigdata.com/rdf/geospatial/literals/v1#geoliteral:lat-lon")
		newobj39 := rdf.NewTypedLiteral(fmt.Sprintf("%s#%s", item.Lat, item.Long), nt39)

		newtriple1 := rdf.Triple{Subj: newsub, Pred: newpred1, Obj: newobj1}
		newtriple2 := rdf.Triple{Subj: newsub, Pred: newpred2, Obj: newobj2}
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
		newtriple29 := rdf.Triple{Subj: newsub, Pred: newpred29, Obj: newobj29}
		newtriple30 := rdf.Triple{Subj: newsub, Pred: newpred30, Obj: newobj30}
		newtriple31 := rdf.Triple{Subj: newsub, Pred: newpred31, Obj: newobj31}
		newtriple32 := rdf.Triple{Subj: newsub, Pred: newpred32, Obj: newobj32}
		newtriple33 := rdf.Triple{Subj: newsub, Pred: newpred33, Obj: newobj33}
		newtriple34 := rdf.Triple{Subj: newsub, Pred: newpred34, Obj: newobj34}
		newtriple35 := rdf.Triple{Subj: newsub, Pred: newpred35, Obj: newobj35}
		newtriple36 := rdf.Triple{Subj: newsub, Pred: newpred36, Obj: newobj36}
		newtriple37 := rdf.Triple{Subj: newsub, Pred: newpred37, Obj: newobj37}
		newtriple38 := rdf.Triple{Subj: newsub, Pred: newpred38, Obj: newobj38}
		newtriple39 := rdf.Triple{Subj: newsub, Pred: nt39, Obj: newobj39}

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
		if newtriple29.Obj.String() != "" {
			tr = append(tr, newtriple29)
		}
		if newtriple30.Obj.String() != "" {
			tr = append(tr, newtriple30)
		}
		if newtriple31.Obj.String() != "" {
			tr = append(tr, newtriple31)
		}
		if newtriple32.Obj.String() != "" {
			tr = append(tr, newtriple32)
		}
		if newtriple33.Obj.String() != "" {
			tr = append(tr, newtriple33)
		}
		if newtriple34.Obj.String() != "" {
			tr = append(tr, newtriple34)
		}
		if newtriple35.Obj.String() != "" {
			tr = append(tr, newtriple35)
		}
		if newtriple36.Obj.String() != "" {
			tr = append(tr, newtriple36)
		}
		if newtriple37.Obj.String() != "" {
			tr = append(tr, newtriple37)
		}
		if newtriple38.Obj.String() != "" {
			tr = append(tr, newtriple38)
		}
		if newtriple39.Obj.String() != "" {
			tr = append(tr, newtriple39)
		}
	}

	// Create the output file
	outFile, err := os.Create("csdcoIGSN.nt")
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
