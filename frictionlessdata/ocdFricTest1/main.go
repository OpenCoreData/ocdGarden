package main

import (
	"fmt"
	"log"
	"os"

	"github.com/frictionlessdata/tableschema-go/csv"
	"github.com/frictionlessdata/tableschema-go/schema"
)

type chemicalCarbon struct {
	Leg        string
	Site       string
	Hole       string
	Depth_mbsf int
}

func main() {
	// inferSchema1()
	validateSchema1()
}

func inferSchema1() {

	fmt.Println("A simple test of the frictionless package")

	// tab, err := csv.NewTable(csv.FromFile("../testdata/208_1262C_JanusGraSection_tALVlUMg.csv"), csv.SetHeaders("Leg", "Site", "Hole", "Core", "Core_type", "Section_number", "Section_type", "Top_cm", "Bot_cm", "Depth_mbsf", "Inor_c_wt_pct", "Caco3_wt_pct", "Tot_c_wt_pct", "Org_c_wt_pct", "Nit_wt_pct", "Sul_wt_pct", "H_wt_pct"))
	tab, err := csv.NewTable(csv.FromFile("../testdata/data1.csv"), csv.LoadHeaders())
	if err != nil {
		log.Fatal(err)
	}
	sch, err := schema.Infer(tab) // infer the table schema
	if err != nil {
		log.Fatal(err)
	}
	sch.SaveToFile("schema.json") // save inferred schema to file
	var cc []chemicalCarbon
	sch.DecodeTable(tab, &cc) // unmarshals the table data into the slice.

}

func validateSchema1() {
	// Reading schema.
	chemCarbSchema, err := schema.LoadFromFile("schema.json")
	if err != nil {
		log.Fatal(err)
	}
	// Validate schema.  What does this really do?
	if err := chemCarbSchema.Validate(); err != nil {
		log.Fatal(err)
	}

	// Printing schema fields names.
	log.Println("Fields:")
	for i, f := range chemCarbSchema.Fields {
		log.Printf("%d - %s\n", i, f.Name)
	}

	// Working with schema fields.
	if chemCarbSchema.HasField("Section_number") {
		log.Println("Field Section_number exists in schema")
	} else {
		log.Fatalf("Schema must have the field Section_number")
	}

	// Get a schema field and test a value against it
	field, _ := chemCarbSchema.GetField("Leg")
	if field.TestString("123") {
		value, err := field.Decode("123")
		log.Printf("Unmarshal to value: %v\n", value)
		if err != nil {
			log.Fatalf("Error casting value: %q", err)
		}
	} else {
		log.Fatalf("Value 123 must fit in field Leg.")
	}

	// Dealing with tabular data associated with the schema.
	table, err := csv.NewTable(csv.FromFile("../testdata/data1.csv"), csv.LoadHeaders())

	var cc []chemicalCarbon

	// iter, _ := table.Iter()
	// for iter.Next() {
	// 	if err := chemCarbSchema.Decode(iter.Row(), &chemicalCarbon); err != nil {
	// 		log.Fatalf("Couldn't unmarshal row:%v err:%q", iter.Row(), err)
	// 	}
	// 	log.Printf("Unmarshal Row: %+v\n", chemicalCarbon)
	// }

	// Don't iterate..  data is small enough..
	chemCarbSchema.DecodeTable(table, &cc)

	// play with output
	f, _ := os.Open("./testout.csv")
	defer f.Close()

	w := csv.NewWriter(f)
	defer w.Flush()

	w.Write([]string{"Leg", "Site", "Hole", "depth_mbsf"})
	err = w.Error()
	if err != nil {
		log.Fatal(err)
	}
	for _, row := range cc {
		fmt.Println(row)
		row, _ := chemCarbSchema.Encode(row)
		w.Write(row)
		err = w.Error()
		if err != nil {
			log.Fatal(err)
		}
	}

}
