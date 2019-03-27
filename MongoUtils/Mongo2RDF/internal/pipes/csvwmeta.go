package pipes

import (
	"fmt"

	"github.com/knakk/rdf"
	mgo "gopkg.in/mgo.v2"

	"opencoredata.org/ocdGarden/MongoUtils/Mongo2RDF/internal/common"
	ocdstructs "opencoredata.org/ocdGarden/MongoUtils/Mongo2RDF/internal/structs"
)

/*
{
	"_id": "562d259814df3782006b2c37",
	"context": "context string here",
	"dc_license": {
		"id": "http://opendefinition.org/licenses/cc-by/"
	},
	"dc_modified": {
		"type": "xsd:date",
		"value": "Sunday, 25-Oct-15 14:55:20 EDT"
	},
	"dc_publisher": {
		"schema_name": "Open Core Data",
		"schema_url": {
			"id": "http://opencoredata.org"
		}
	},
	"dc_title": "101_628A_JanusAgeDatapoint_CGcYexOw.csv",
	"dcat_keyword": ["DSDP", "ODP", "IODP", "JanusAgeDatapoint"],
	"tableschema": {
		"abouturl": "",
		"columns": [{
			"datatype": "int64",
			"dc_description": "",
			"name": "Leg",
			"required": false,
			"titles": []
		}, {
			"datatype": "int64",
			"dc_description": "",
			"name": "Site",
			"required": false,
			"titles": []
		}, {
			"datatype": "string",
			"dc_description": "",
			"name": "Hole",
			"required": false,
			"titles": []
		}, {
			"datatype": "string",
			"dc_description": "",
			"name": "Age_model_type",
			"required": false,
			"titles": []
		}, {
			"datatype": "float64",
			"dc_description": "",
			"name": "Depth_mbsf",
			"required": false,
			"titles": []
		}, {
			"datatype": "sql.NullFloat64",
			"dc_description": "",
			"name": "Age_ma",
			"required": false,
			"titles": []
		}, {
			"datatype": "sql.NullString",
			"dc_description": "",
			"name": "Control_point_comment",
			"required": false,
			"titles": []
		}],
		"primarykey": ""
	},
	"url": "http://opencoredata.org/id/dataset/da39147d-deda-44ac-879d-684491a110fe"
}
*/

// Csvwmeta update with material from jsongold
func Csvwmeta(session *mgo.Session) {

	// connect and get documents
	csvw := session.DB("test").C("csvwmeta")
	var csvwDocs []ocdstructs.CSVWMeta
	err := csvw.Find(nil).All(&csvwDocs)
	if err != nil {
		fmt.Printf("this is error %v \n", err)
	}

	// RDF item
	tr := []rdf.Triple{}

	// Loop on documents
	for _, item := range csvwDocs {
		newsub, _ := rdf.NewIRI(fmt.Sprintf("http://opencoredata/id/resource/janus/query/%s", item.URL)) // Sprintf a correct URI here

		// title
		newpred0, _ := rdf.NewIRI("http://www.w3.org/1999/02/22-rdf-syntax-ns#title")
		newobj0, _ := rdf.NewIRI("http://opencoredata.org/id/voc/janus/v1/JanusQuery")
		newtriple0 := rdf.Triple{Subj: newsub, Pred: newpred0, Obj: newobj0}
		tr = append(tr, newtriple0)
	}

	common.WriteFile("./output/csvw.nt", tr)

}
