package main

import (
	"github.com/frictionlessdata/tableschema-go/csv"
	"github.com/frictionlessdata/tableschema-go/schema"
)

type user struct {
	ID   int
	Age  int
	Name string
}

func main() {
	tab, err := csv.NewTable(csv.FromFile("data_infer_utf8.csv"), csv.SetHeaders("id", "age", "name"))
	if err != nil {
		panic(err)
	}
	sch, err := schema.Infer(tab) // infer the table schema
	if err != nil {
		panic(err)
	}
	sch.SaveToFile("schema.json") // save inferred schema to file
	var users []user
	sch.DecodeTable(tab, &users) // unmarshals the table data into the slice.
}
