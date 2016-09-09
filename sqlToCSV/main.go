package main

// Simple call to Oracle in a self contained packge
// be sure to do:
// export JANUS_USERNAME=X
// export JANUS_PASSWORD=X
// export JANUS_HOST=X
// export JANUS_SERVICENAME=X
//
// also needs Oracle client SDK
//
// ref:  https://github.com/StabbyCutyou/sqltocsv

import (
	"encoding/csv"
	"fmt"
	// "github.com/kisielk/sqlstruct"
	"log"
	"os"
	"strconv"
	"time"

	"gopkg.in/rana/ora.v3"
	// "text/template"
	// "strings"
	"github.com/jmoiron/sqlx"
)

// init function to register oracle driver for SQL driver
func init() {
	ora.Register(nil)
}

func check(e error) {
	if e != nil {
		log.Print(e)
	}
}

// main simple Oracle call test
func main() {
	err := JanusPaleoSampleFunc()
	if err != nil {
		log.Printf(`Error: "%s"`, err)
	}
}

// GetJanusCon returns a Oracle DB connection
func GetJanusCon() (*sqlx.DB, error) {
	username := os.Getenv("JANUS_USERNAME")
	password := os.Getenv("JANUS_PASSWORD")
	host := os.Getenv("JANUS_HOST")
	servicename := os.Getenv("JANUS_SERVICENAME")
	connectionString := fmt.Sprintf("%s/%s@(DESCRIPTION=(ADDRESS_LIST=(ADDRESS=(PROTOCOL=TCP)(HOST=%s)(PORT=1521)))(CONNECT_DATA=(SERVICE_NAME=%s)))", username, password, host, servicename)
	return sqlx.Open("ora", connectionString)
}

// JanusPaleoSampleFunc calls to database and issues the select from ocd_paleo_sample call
func JanusPaleoSampleFunc() error {

	// the query
	qry := "SELECT * FROM ocd_paleo_sample WHERE rownum < 101"
	//qry := "SELECT * FROM ocd_paleo_sample"

	// get the Oracle connection
	conn, err := GetJanusCon()
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	// Make the calls the get the rows
	results, err := conn.Queryx(qry)
	if err != nil {
		log.Printf(`Error with "%s": %s`, qry, err)
	}

	// struct to hold all the results
	// allResults := []JanusPaleoSamplejanusRow{}
	i := 1 // rows is not a enumeration..  sigh..  have to do it myself..

	// go through the rows
	output, err := os.Create("test.csv")
	if err != nil {
		log.Fatal(err)
	}

	csvWriter := csv.NewWriter(output)
	csvWriter.Comma = 0x0009
	firstLine := true
	for results.Next() {
		rowscan, err := results.SliceScan()
		if err != nil {
			log.Fatal(err)
		}
		if firstLine {
			firstLine = false
			cols, err := results.Columns()
			if err != nil {
				log.Fatal(err)
			}
			csvWriter.Write(cols)
		}

		rowStrings := make([]string, len(rowscan))
		// It seems for mysql, the case is always []byte of a string?
		for i, col := range rowscan {
			//log.Print(reflect.TypeOf(col))
			switch col.(type) {
			case float64:
				rowStrings[i] = strconv.FormatFloat(col.(float64), 'f', 6, 64)
			case int64:
				rowStrings[i] = strconv.FormatInt(col.(int64), 10)
			case bool:
				rowStrings[i] = strconv.FormatBool(col.(bool))
			case []byte:
				rowStrings[i] = string(col.([]byte))
			case string:
				rowStrings[i] = col.(string)
			case time.Time:
				rowStrings[i] = col.(time.Time).String()
			case nil:
				rowStrings[i] = "NULL"
			default:
				log.Print(col)
			}
		}
		csvWriter.Write(rowStrings)
	}

	csvWriter.Flush()
	output.Close()

	return nil
}
