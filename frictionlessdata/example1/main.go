package main

import "github.com/frictionlessdata/tableschema-go/csv"

// struct representing each row of the table.
type person struct {
	Name string
	Age  uint16
}

func main() {
	t, _ := csv.New(csv.FromFile("data.csv"), csv.LoadHeaders()) // load table
	t.Infer()                                                    // infer the table schema
	t.Schema.SaveToFile("schema.json")                           // save inferred schema to file
	data := []person{}
	t.CastAll(&data) // casts the table data into the data slice.
}
