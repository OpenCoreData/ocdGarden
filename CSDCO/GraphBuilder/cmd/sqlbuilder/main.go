package main

import (
	"database/sql"
	"fmt"
	"log"

	"opencoredata.org/ocdGarden/CSDCO/GraphBuilder/internal/query"
)

func main() {
	fmt.Println("SQLite based CSDCO Digital Objectv builder")

	db, err := sql.Open("sqlite3", "./data/input/CSDCO.sqlite3")
	if err != nil {
		log.Panic(err)
	}

	//query.Boreholes(db)
	query.Projects(db)
}
