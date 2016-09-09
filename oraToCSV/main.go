/*
   Package main in csvdump represents a cursor->csv dumper

   Copyright 2013 Tamás Gulácsi

   Licensed under the Apache License, Version 2.0 (the "License");
   you may not use this file except in compliance with the License.
   You may obtain a copy of the License at

     http://www.apache.org/licenses/LICENSE-2.0

   Unless required by applicable law or agreed to in writing, software
   distributed under the License is distributed on an "AS IS" BASIS,
   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
   See the License for the specific language governing permissions and
   limitations under the License.
*/
package main

import (
	"bufio"
	"database/sql"
	"fmt"
	"io"
	"log"
	"os"
	// "strings"
	"time"

	"gopkg.in/rana/ora.v3"
	"github.com/tgulacsi/go/orahlp"

	"gopkg.in/errgo.v1"
)

const SQL_RockEval = `select x.leg, x.site, x.hole, x.core, x.core_type,
  x.section_number, x.section_type,
  s.top_interval*100.0, s.bottom_interval*100.0,
  get_depths(x.section_id,'STD',s.top_interval,0,0) mbsf,
  null,
  avg(decode(cca_1.analysis_code,'ORG_C', cca_1.analysis_result)),
  avg(decode(cca_1.analysis_code,'TOC', cca_1.analysis_result)),
  avg(decode(cca_2.analysis_code,'S1', cca_2.analysis_result)),
  avg(decode(cca_2.analysis_code,'S2', cca_2.analysis_result)),
  avg(decode(cca_2.analysis_code,'S3', cca_2.analysis_result)),
  avg(decode(cca_2.analysis_code,'TMX', cca_2.analysis_result)),
  avg(decode(cca_2.analysis_code,'PI', cca_2.analysis_result)),
  avg(decode(cca_2.analysis_code,'PC', cca_2.analysis_result)),
  avg(decode(cca_2.analysis_code,'OI', cca_2.analysis_result)),
  avg(decode(cca_2.analysis_code,'HI', cca_2.analysis_result))
from hole h, section x, sample s,
  chem_carb_sample ccs, chem_carb_analysis cca_1, chem_carb_analysis cca_2
where h.leg = x.leg and h.site = x.site and h.hole = x.hole and
       x.section_id = s.sam_section_id and
       s.sample_id = ccs.sample_id and
       s.location = ccs.location and
       ccs.run_id = cca_1.run_id and
       cca_1.run_id = cca_2.run_id(+)and
       cca_2.method_code = 'RE'
  and x.leg = 178
 group by x.leg, x.site, x.hole, x.core, x.core_type, x.section_type, x.section_number, s.top_interval, s.bottom_interval, x.section_id
order by x.leg, x.site, x.hole, x.core,
  x.core_type, x.section_number, s.top_interval`

// func getQuery(table, where string, columns []string) string {
// 	if strings.HasPrefix(table, "SELECT ") {
// 		return table
// 	}
// 	cols := "*"
// 	if len(columns) > 0 {
// 		cols = strings.Join(columns, ", ")
// 	}
// 	if where == "" {
// 		return "SELECT " + cols + " FROM " + table
// 	}
// 	return "SELECT " + cols + " FROM " + table + " WHERE " + where
// }


// init function to register oracle driver for SQL driver
func init() {
	ora.Register(nil)
}

func GetJanusCon() (*sql.DB, error) {

	username := os.Getenv("JANUS_USERNAME")
	password := os.Getenv("JANUS_PASSWORD")
	host := os.Getenv("JANUS_HOST")
	servicename := os.Getenv("JANUS_SERVICENAME")

	connectionString := fmt.Sprintf("%s/%s@(DESCRIPTION=(ADDRESS_LIST=(ADDRESS=(PROTOCOL=TCP)(HOST=%s)(PORT=1521)))(CONNECT_DATA=(SERVICE_NAME=%s)))", username, password, host, servicename)
	// return sql.Open("goracle", connectionString)
	return sql.Open("ora", connectionString)
}

func dump(w io.Writer, qry string) error {
	// db, err := connect.GetConnection("")
    
    db, err := GetJanusCon()
	if err != nil {
		return errgo.Notef(err, "connect to database")
	}
	defer db.Close()
	columns, err := GetColumns(db, qry)
	if err != nil {
		return errgo.Notef(err, "get column converters", err)
	}
	log.Printf("columns: %#v", columns)
	row := make([]interface{}, len(columns))
	rowP := make([]interface{}, len(row))
	for i := range row {
		rowP[i] = &row[i]
	}

	rows, err := db.Query(qry)
	if err != nil {
		return errgo.Newf("error executing %q: %s", qry, err)
	}
	defer rows.Close()

	bw := bufio.NewWriterSize(w, 65536)
	defer bw.Flush()
	for i, col := range columns {
		if i > 0 {
			bw.Write([]byte{';'})
		}
		bw.Write([]byte{'"'})
		bw.WriteString(col.Name)
		bw.Write([]byte{'"'})
	}
	bw.Write([]byte{'\n'})
	n := 0
	for rows.Next() {
		if err = rows.Scan(rowP...); err != nil {
			return errgo.Notef(err, "scan %d. row", n+1)
		}
		for i, data := range row {
			if i > 0 {
				bw.Write([]byte{';'})
			}
			if data == nil {
				continue
			}
			bw.WriteString(columns[i].String(data))
		}
		bw.Write([]byte{'\n'})
		n++
	}
	log.Printf("written %d rows.", n)
	err = rows.Err()
	if err != nil {
		return err
	}
	return nil
}

type ColConverter func(interface{}) string

type Column struct {
	Name   string
	String ColConverter
}

func GetColumns(db *sql.DB, qry string) (cols []Column, err error) {
	desc, err := orahlp.DescribeQuery(db, qry)   // not small issue of removing SQL from this function name
	if err != nil {
		return nil, errgo.Newf("error getting description for %q: %s", qry, err)
	}
	log.Printf("desc: %#v", desc)
	var ok bool
	cols = make([]Column, len(desc))
	for i, col := range desc {
		cols[i].Name = col.Name
		if cols[i].String, ok = converters[col.Type]; !ok {
			cols[i].String = defaultConverter
			log.Printf("no converter for type %d (column name: %s)", col.Type, col.Name)
		}
	}
	return cols, nil
}

func defaultConverter(data interface{}) string { return fmt.Sprintf("%v", data) }

var converters = map[int]ColConverter{
	1: func(data interface{}) string { //VARCHAR2
		return fmt.Sprintf("%q", data.(string))
	},
	6: func(data interface{}) string { //NUMBER
		return fmt.Sprintf("%v", data)
	},
	96: func(data interface{}) string { //CHAR
		return fmt.Sprintf("%q", data.(string))
	},
	156: func(data interface{}) string { //DATE
		return `"` + data.(time.Time).Format(time.RFC3339) + `"`
	},
}

func main() {
	
	qry := SQL_RockEval
	if err := dump(os.Stdout, qry); err != nil {
		log.Printf("error dumping: %s", err)
		os.Exit(1)
	}
	os.Exit(0)
}