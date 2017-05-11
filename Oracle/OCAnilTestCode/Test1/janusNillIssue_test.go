package main

import (
	"database/sql"
	orahlp "github.com/tgulacsi/go/orahlp"
	"gopkg.in/rana/ora.v3"
	"log"
	"testing"
)

func init() {
	ora.Register(nil)
}

func TestFils(t *testing.T) {

	log.Println("Ready to connect")

	testDb, err := sql.Open("ora", "X/X@(DESCRIPTION=(ADDRESS_LIST=(ADDRESS=(PROTOCOL=TCP)(HOST=X)(PORT=1521)))(CONNECT_DATA=(SERVICE_NAME=X)))")
	if err != nil {
		log.Printf(`Error with	: %s`, err)
		return
	}

	defer testDb.Close()
	log.Println("Connected")

	testDb.Exec(`DROP TABLE test_janus`)
	testDb.Exec(`DROP VIEW test_janus_v`)
	if _, err := testDb.Exec(`CREATE TABLE test_janus (
		leg NUMBER(5),
		site NUMBER(6),
		hole VARCHAR2(1),
		core NUMBER(5),
		core_type VARCHAR2(1),
		section_number NUMBER(2),
		section_Type VARCHAR2(2) NULL,
		top_cm NUMBER NULL,
		bot_cm NUMBER NULL,
		depth_mbsf NUMBER NULL,
		inor_c_wt_pct NUMBER NULL,
		caco3_wt_pct NUMBER NULL,
		tot_c_wt_pct NUMBER NULL,
		org_c_wt_pct NUMBER NULL,
		nit_wt_pct NUMBER NULL,
		sul_wt_pct NUMBER NULL,
		h_wt_pct NUMBER NULL
	)`); err != nil {
		t.Fatal(err)
	}

	if _, err := testDb.Exec(`INSERT INTO test_janus (
		leg, site, hole, core, core_type, section_number,
		section_type, top_cm, bot_cm, depth_mbsf,
		inor_c_wt_pct, caco3_wt_pct, tot_c_wt_pct,
		org_c_wt_pct, nit_wt_pct, sul_wt_pct, h_wt_pct)
	VALUES (207, 1259, 'C', 3, 'B', 4, '@', 5.2, NULL, 7.6, 8., 9., 10., 11., NULL , 13., 14.)`,
	); err != nil {
		t.Fatal(err)
	}

	if _, err := testDb.Exec(`INSERT INTO test_janus (
		leg, site, hole, core, core_type, section_number,
		section_type, top_cm, bot_cm, depth_mbsf,
		inor_c_wt_pct, caco3_wt_pct, tot_c_wt_pct,
		org_c_wt_pct, nit_wt_pct, sul_wt_pct, h_wt_pct)
	VALUES (171, 1049, 'B', 3, 'B', 4.2, '@', NULL, 6.12, 7.12, 8, 9.99, NULL, 11., NULL , 0.8, 0.42)`,
	); err != nil {
		t.Fatal(err)
	}

	if _, err := testDb.Exec(`CREATE VIEW test_janus_v AS SELECT * FROM test_janus`); err != nil {
		t.Fatal(err)
	}

	qry := `SELECT
	   leg, site, hole, core, core_type
	 , section_number, section_type
	 , top_cm, bot_cm
	  , depth_mbsf
	 , inor_c_wt_pct
	 , caco3_wt_pct
	 , tot_c_wt_pct
	 , org_c_wt_pct
	 , nit_wt_pct
	 , sul_wt_pct
	 , h_wt_pct
	FROM
	   test_janus
	WHERE
	       leg = 171
	    AND site = 1049
	   AND hole = 'B'
	ORDER BY leg, site, hole, core, section_number, top_cm
`

	qry2 := `SELECT
	   leg, site, hole, core, core_type
	 , section_number, section_type
	 , top_cm, bot_cm
	 , inor_c_wt_pct
	 , caco3_wt_pct
	 , tot_c_wt_pct
	 , org_c_wt_pct
	 , nit_wt_pct
	 , sul_wt_pct
	 , h_wt_pct
	   FROM  ocd_chem_carb
	WHERE
	       leg = 171
	    AND site = 1049
	   AND hole = 'B'
	ORDER BY leg, site, hole, core, section_number, top_cm
	`

	// qry2 := `SELECT
	//     x.leg, x.site, x.hole
	//   , x.core, x.core_type
	//   , x.section_number, x.section_type
	//   , s.top_interval*100.0 top_cm
	//   , s.bottom_interval*100.0 bot_cm
	//   , AVG(DECODE(cca.analysis_code,'INOR_C',cca.analysis_result)) INOR_C_wt_pct
	//   , AVG(DECODE(cca.analysis_code,'CaCO3', cca.analysis_result)) CaCO3_wt_pct
	//   , AVG(DECODE(cca.analysis_code,'TOT_C', cca.analysis_result)) TOT_C_wt_pct
	//   , AVG(DECODE(cca.analysis_code,'ORG_C', cca.analysis_result)) ORG_C_wt_pct
	//   , AVG(DECODE(cca.analysis_code,'NIT',   cca.analysis_result)) NIT_wt_pct
	//   , AVG(DECODE(cca.analysis_code,'SUL',   cca.analysis_result)) SUL_wt_pct
	//   , AVG(DECODE(cca.analysis_code,'H',     cca.analysis_result)) H_wt_pct
	// FROM
	//     hole h, section x, sample s
	//   , chem_carb_sample ccs, chem_carb_analysis cca
	// WHERE
	//         h.leg = x.leg
	//     AND h.site = x.site
	//     AND h.hole = x.hole
	//     AND x.section_id = s.sam_section_id
	//     AND s.sample_id = ccs.sample_id
	//     AND s.location = ccs.location
	//     AND ccs.run_id = cca.run_id
	//     AND h.leg = 171
	//     AND h.site = 1049
	//     AND h.hole = 'B'
	// GROUP BY x.leg, x.site, x.hole, x.core, x.core_type, x.section_number, x.section_type, s.top_interval, s.bottom_interval
	// ORDER BY x.leg, x.site, x.hole, x.core, x.core_type, x.section_number, s.top_interval
	// `

	desc, err := orahlp.DescribeQuery(testDb, qry2)
	if err != nil {
		t.Errorf(`Error with : %s`, err)
		//return nil, errgo.Newf("error getting description for %q: %s", qry, err)
	}
	log.Printf("desc: %#v", desc)

	rows, err := testDb.Query(qry)
	if err != nil {
		t.Errorf(`Error with "%s": %s`, qry, err)
		return
	}
	defer rows.Close()

	i := 0
	for rows.Next() {
		i++
		var (
			Leg            int
			Site           int
			Hole           string
			Core           int
			Core_type      string
			Section_number int
			Section_type   string
			Top_cm         sql.NullFloat64
			Bot_cm         sql.NullFloat64
			Depth_mbsf     sql.NullFloat64
			Inor_c_wt_pct  sql.NullFloat64
			Caco3_wt_pct   sql.NullFloat64
			Tot_c_wt_pct   sql.NullFloat64
			Org_c_wt_pct   sql.NullFloat64
			Nit_wt_pct     sql.NullFloat64
			Sul_wt_pct     sql.NullFloat64
			H_wt_pct       sql.NullFloat64
		)

		if err := rows.Scan(&Leg, &Site, &Hole, &Core, &Core_type, &Section_number, &Section_type, &Top_cm, &Bot_cm, &Depth_mbsf, &Inor_c_wt_pct, &Caco3_wt_pct, &Tot_c_wt_pct, &Org_c_wt_pct, &Nit_wt_pct, &Sul_wt_pct, &H_wt_pct); err != nil {
			t.Fatalf("scan %d. record: %v", i, err)
		}

		log.Printf("Results: %v %v %v %v %v %v %v %v %v %v %v %v %v %v %v %v %v", Leg, Site, Hole, Core, Core_type, Section_number, Section_type, Top_cm, Bot_cm, Depth_mbsf, Inor_c_wt_pct, Caco3_wt_pct, Tot_c_wt_pct, Org_c_wt_pct, Nit_wt_pct, Sul_wt_pct, H_wt_pct)

	}
	if err := rows.Err(); err != nil {
		t.Error(err)
	}

	rows2, err := testDb.Query(qry2)
	if err != nil {
		t.Errorf(`Error with "%s": %s`, qry2, err)
		return
	}
	defer rows2.Close()

	ii := 0
	for rows2.Next() {
		ii++
		var (
			Leg            int
			Site           int
			Hole           string
			Core           int
			Core_type      string
			Section_number int
			Section_type   string
			Top_cm         sql.NullFloat64
			Bot_cm         sql.NullFloat64
			Inor_c_wt_pct  sql.NullFloat64
			Caco3_wt_pct   sql.NullFloat64
			Tot_c_wt_pct   sql.NullFloat64
			Org_c_wt_pct   sql.NullFloat64
			Nit_wt_pct     sql.NullFloat64
			Sul_wt_pct     sql.NullFloat64
			H_wt_pct       sql.NullFloat64
		)

		if err := rows2.Scan(&Leg, &Site, &Hole, &Core, &Core_type, &Section_number, &Section_type, &Top_cm, &Bot_cm, &Inor_c_wt_pct, &Caco3_wt_pct, &Tot_c_wt_pct, &Org_c_wt_pct, &Nit_wt_pct, &Sul_wt_pct, &H_wt_pct); err != nil {
			t.Fatalf("scan %d. record: %v", ii, err)
		}

		log.Printf("Results: %v %v %v %v %v %v %v %v %v %v %v %v %v %v %v %v %v", Leg, Site, Hole, Core, Core_type, Section_number, Section_type, Top_cm, Bot_cm, Inor_c_wt_pct, Caco3_wt_pct, Tot_c_wt_pct, Org_c_wt_pct, Nit_wt_pct, Sul_wt_pct, H_wt_pct)

	}
	if err := rows2.Err(); err != nil {
		t.Error(err)
	}

}
