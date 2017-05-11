package janus

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
		top_cm NUMBER(6,3) NULL,
		bot_cm NUMBER(6,3) NULL,
		depth_mbsf NUMBER NULL,
		inor_c_wt_pct NUMBER NULL,
		caco3_wt_pct NUMBER NULL,
		tot_c_wt_pct NUMBER NULL,
		org_c_wt_pct NUMBER NULL,
		nit_wt_pct NUMBER NULL,
		sul_wt_pct NUMBER NULL,
		h_wt_pct NUMBER(6,3) NULL
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

	testDb.Exec(`DROP TABLE ocd_hole_test`)
	testDb.Exec(`DROP TABLE ocd_section_test`)
	testDb.Exec(`DROP TABLE ocd_sample_test`)
	testDb.Exec(`DROP TABLE ocd_chem_carb_sample_test`)
	testDb.Exec(`DROP TABLE ocd_chem_carb_analysis_test`)

	if _, err := testDb.Exec(`CREATE TABLE ocd_hole_test (
 LEG    NUMBER(5) NOT NULL,
 SITE   NUMBER(6) NOT NULL,
 HOLE   VARCHAR2(1) NOT NULL
)`); err != nil {
		t.Fatal(err)
	}

	if _, err := testDb.Exec(`CREATE TABLE ocd_section_test (
 SECTION_ID       NUMBER(7) NOT NULL,
 SECTION_NUMBER   NUMBER(2) NOT NULL,
 SECTION_TYPE     VARCHAR2(2),
 LEG              NUMBER(5) NOT NULL,
 SITE             NUMBER(6) NOT NULL,
 HOLE             VARCHAR2(1) NOT NULL,
 CORE             NUMBER(5) NOT NULL,
 CORE_TYPE        VARCHAR2(1) NOT NULL
)`); err != nil {
		t.Fatal(err)
	}

	if _, err := testDb.Exec(`CREATE TABLE ocd_sample_test (
 SAMPLE_ID               NUMBER(9) NOT NULL,
 LOCATION                VARCHAR2(3) NOT NULL,
 SAM_SECTION_ID          NUMBER(7),
 TOP_INTERVAL            NUMBER(6,3),
 BOTTOM_INTERVAL         NUMBER(6,3)
)`); err != nil {
		t.Fatal(err)
	}

	if _, err := testDb.Exec(`CREATE TABLE ocd_chem_carb_sample_test (
 RUN_ID                  NUMBER(9) NOT NULL,
 SAMPLE_ID               NUMBER(9) NOT NULL,
 LOCATION                VARCHAR2(3) NOT NULL
)`); err != nil {
		t.Fatal(err)
	}

	if _, err := testDb.Exec(`CREATE TABLE ocd_chem_carb_analysis_test (
 RUN_ID                  NUMBER(9) NOT NULL,
 ANALYSIS_CODE           VARCHAR2(15) NOT NULL,
 METHOD_CODE             VARCHAR2(10) NOT NULL,
 ANALYSIS_RESULT         NUMBER(15,5)
)`); err != nil {
		t.Fatal(err)
	}

	// create the views

	testDb.Exec(`DROP PUBLIC SYNONYM ocd_chem_carb_test`)
	testDb.Exec(`DROP VIEW ocd_chem_carb_test_v`)

	if _, err := testDb.Exec(`Insert into OCD_CHEM_CARB_ANALYSIS_TEST (RUN_ID,ANALYSIS_CODE,METHOD_CODE,ANALYSIS_RESULT) values (42285,'CaCO3','C',1)`); err != nil {
		t.Fatal(err)
	}
	if _, err := testDb.Exec(`Insert into OCD_CHEM_CARB_ANALYSIS_TEST (RUN_ID,ANALYSIS_CODE,METHOD_CODE,ANALYSIS_RESULT) values (42285,'INOR_C','C',0.12)`); err != nil {
		t.Fatal(err)
	}
	if _, err := testDb.Exec(`Insert into OCD_CHEM_CARB_ANALYSIS_TEST (RUN_ID,ANALYSIS_CODE,METHOD_CODE,ANALYSIS_RESULT) values (42290,'CaCO3','C',0.45)`); err != nil {
		t.Fatal(err)
	}
	if _, err := testDb.Exec(`Insert into OCD_CHEM_CARB_ANALYSIS_TEST (RUN_ID,ANALYSIS_CODE,METHOD_CODE,ANALYSIS_RESULT) values (42290,'H','CNS',0.62)`); err != nil {
		t.Fatal(err)
	}
	if _, err := testDb.Exec(`Insert into OCD_CHEM_CARB_ANALYSIS_TEST (RUN_ID,ANALYSIS_CODE,METHOD_CODE,ANALYSIS_RESULT) values (42290,'HI','RE',187)`); err != nil {
		t.Fatal(err)
	}
	if _, err := testDb.Exec(`Insert into OCD_CHEM_CARB_ANALYSIS_TEST (RUN_ID,ANALYSIS_CODE,METHOD_CODE,ANALYSIS_RESULT) values (42290,'INOR_C','C',0.054)`); err != nil {
		t.Fatal(err)
	}
	if _, err := testDb.Exec(`Insert into OCD_CHEM_CARB_ANALYSIS_TEST (RUN_ID,ANALYSIS_CODE,METHOD_CODE,ANALYSIS_RESULT) values (42290,'NIT','CNS',0)`); err != nil {
		t.Fatal(err)
	}
	if _, err := testDb.Exec(`Insert into OCD_CHEM_CARB_ANALYSIS_TEST (RUN_ID,ANALYSIS_CODE,METHOD_CODE,ANALYSIS_RESULT) values (42290,'OI','RE',56)`); err != nil {
		t.Fatal(err)
	}
	if _, err := testDb.Exec(`Insert into OCD_CHEM_CARB_ANALYSIS_TEST (RUN_ID,ANALYSIS_CODE,METHOD_CODE,ANALYSIS_RESULT) values (42290,'ORG_C','CNS',0.6)`); err != nil {
		t.Fatal(err)
	}
	if _, err := testDb.Exec(`Insert into OCD_CHEM_CARB_ANALYSIS_TEST (RUN_ID,ANALYSIS_CODE,METHOD_CODE,ANALYSIS_RESULT) values (42290,'PC','RE',0.07)`); err != nil {
		t.Fatal(err)
	}
	if _, err := testDb.Exec(`Insert into OCD_CHEM_CARB_ANALYSIS_TEST (RUN_ID,ANALYSIS_CODE,METHOD_CODE,ANALYSIS_RESULT) values (42290,'PI','RE',0.1)`); err != nil {
		t.Fatal(err)
	}
	if _, err := testDb.Exec(`Insert into OCD_CHEM_CARB_ANALYSIS_TEST (RUN_ID,ANALYSIS_CODE,METHOD_CODE,ANALYSIS_RESULT) values (42290,'S1','RE',0.09)`); err != nil {
		t.Fatal(err)
	}
	if _, err := testDb.Exec(`Insert into OCD_CHEM_CARB_ANALYSIS_TEST (RUN_ID,ANALYSIS_CODE,METHOD_CODE,ANALYSIS_RESULT) values (42290,'S2','RE',0.77)`); err != nil {
		t.Fatal(err)
	}
	if _, err := testDb.Exec(`Insert into OCD_CHEM_CARB_ANALYSIS_TEST (RUN_ID,ANALYSIS_CODE,METHOD_CODE,ANALYSIS_RESULT) values (42290,'S3','RE',0.23)`); err != nil {
		t.Fatal(err)
	}
	if _, err := testDb.Exec(`Insert into OCD_CHEM_CARB_ANALYSIS_TEST (RUN_ID,ANALYSIS_CODE,METHOD_CODE,ANALYSIS_RESULT) values (42290,'SUL','CNS',0.06)`); err != nil {
		t.Fatal(err)
	}
	if _, err := testDb.Exec(`Insert into OCD_CHEM_CARB_ANALYSIS_TEST (RUN_ID,ANALYSIS_CODE,METHOD_CODE,ANALYSIS_RESULT) values (42290,'TMX','RE',460)`); err != nil {
		t.Fatal(err)
	}
	if _, err := testDb.Exec(`Insert into OCD_CHEM_CARB_ANALYSIS_TEST (RUN_ID,ANALYSIS_CODE,METHOD_CODE,ANALYSIS_RESULT) values (42290,'TOC','RE',0.41)`); err != nil {
		t.Fatal(err)
	}
	if _, err := testDb.Exec(`Insert into OCD_CHEM_CARB_ANALYSIS_TEST (RUN_ID,ANALYSIS_CODE,METHOD_CODE,ANALYSIS_RESULT) values (42290,'TOT_C','CNS',0.65)`); err != nil {
		t.Fatal(err)
	}
	if _, err := testDb.Exec(`Insert into OCD_CHEM_CARB_ANALYSIS_TEST (RUN_ID,ANALYSIS_CODE,METHOD_CODE,ANALYSIS_RESULT) values (42295,'CaCO3','C',0.45)`); err != nil {
		t.Fatal(err)
	}
	if _, err := testDb.Exec(`Insert into OCD_CHEM_CARB_ANALYSIS_TEST (RUN_ID,ANALYSIS_CODE,METHOD_CODE,ANALYSIS_RESULT) values (42295,'INOR_C','C',0.054)`); err != nil {
		t.Fatal(err)
	}
	if _, err := testDb.Exec(`Insert into OCD_CHEM_CARB_ANALYSIS_TEST (RUN_ID,ANALYSIS_CODE,METHOD_CODE,ANALYSIS_RESULT) values (3295,'CaCO3','C',73.91)`); err != nil {
		t.Fatal(err)
	}
	if _, err := testDb.Exec(`Insert into OCD_CHEM_CARB_ANALYSIS_TEST (RUN_ID,ANALYSIS_CODE,METHOD_CODE,ANALYSIS_RESULT) values (3295,'H','CNS',0.27)`); err != nil {
		t.Fatal(err)
	}
	if _, err := testDb.Exec(`Insert into OCD_CHEM_CARB_ANALYSIS_TEST (RUN_ID,ANALYSIS_CODE,METHOD_CODE,ANALYSIS_RESULT) values (3295,'INOR_C','C',8.87)`); err != nil {
		t.Fatal(err)
	}
	if _, err := testDb.Exec(`Insert into OCD_CHEM_CARB_ANALYSIS_TEST (RUN_ID,ANALYSIS_CODE,METHOD_CODE,ANALYSIS_RESULT) values (3295,'ORG_C','CNS',0.11)`); err != nil {
		t.Fatal(err)
	}
	if _, err := testDb.Exec(`Insert into OCD_CHEM_CARB_ANALYSIS_TEST (RUN_ID,ANALYSIS_CODE,METHOD_CODE,ANALYSIS_RESULT) values (3295,'TOT_C','CNS',8.98)`); err != nil {
		t.Fatal(err)
	}
	if _, err := testDb.Exec(`Insert into OCD_CHEM_CARB_ANALYSIS_TEST (RUN_ID,ANALYSIS_CODE,METHOD_CODE,ANALYSIS_RESULT) values (3300,'CaCO3','C',76.07)`); err != nil {
		t.Fatal(err)
	}
	if _, err := testDb.Exec(`Insert into OCD_CHEM_CARB_ANALYSIS_TEST (RUN_ID,ANALYSIS_CODE,METHOD_CODE,ANALYSIS_RESULT) values (3300,'H','CNS',0.24)`); err != nil {
		t.Fatal(err)
	}
	if _, err := testDb.Exec(`Insert into OCD_CHEM_CARB_ANALYSIS_TEST (RUN_ID,ANALYSIS_CODE,METHOD_CODE,ANALYSIS_RESULT) values (3300,'INOR_C','C',9.13)`); err != nil {
		t.Fatal(err)
	}
	if _, err := testDb.Exec(`Insert into OCD_CHEM_CARB_ANALYSIS_TEST (RUN_ID,ANALYSIS_CODE,METHOD_CODE,ANALYSIS_RESULT) values (3300,'ORG_C','CNS',0)`); err != nil {
		t.Fatal(err)
	}
	if _, err := testDb.Exec(`Insert into OCD_CHEM_CARB_ANALYSIS_TEST (RUN_ID,ANALYSIS_CODE,METHOD_CODE,ANALYSIS_RESULT) values (3300,'TOT_C','CNS',9.11)`); err != nil {
		t.Fatal(err)
	}
	if _, err := testDb.Exec(`Insert into OCD_CHEM_CARB_ANALYSIS_TEST (RUN_ID,ANALYSIS_CODE,METHOD_CODE,ANALYSIS_RESULT) values (3240,'CaCO3','C',70.97)`); err != nil {
		t.Fatal(err)
	}
	if _, err := testDb.Exec(`Insert into OCD_CHEM_CARB_ANALYSIS_TEST (RUN_ID,ANALYSIS_CODE,METHOD_CODE,ANALYSIS_RESULT) values (3240,'H','CNS',0.15)`); err != nil {
		t.Fatal(err)
	}
	if _, err := testDb.Exec(`Insert into OCD_CHEM_CARB_ANALYSIS_TEST (RUN_ID,ANALYSIS_CODE,METHOD_CODE,ANALYSIS_RESULT) values (3240,'INOR_C','C',8.52)`); err != nil {
		t.Fatal(err)
	}
	if _, err := testDb.Exec(`Insert into OCD_CHEM_CARB_ANALYSIS_TEST (RUN_ID,ANALYSIS_CODE,METHOD_CODE,ANALYSIS_RESULT) values (3240,'ORG_C','CNS',0)`); err != nil {
		t.Fatal(err)
	}
	if _, err := testDb.Exec(`Insert into OCD_CHEM_CARB_ANALYSIS_TEST (RUN_ID,ANALYSIS_CODE,METHOD_CODE,ANALYSIS_RESULT) values (3240,'TOT_C','CNS',8.24)`); err != nil {
		t.Fatal(err)
	}
	if _, err := testDb.Exec(`Insert into OCD_CHEM_CARB_ANALYSIS_TEST (RUN_ID,ANALYSIS_CODE,METHOD_CODE,ANALYSIS_RESULT) values (3245,'CaCO3','C',82.33)`); err != nil {
		t.Fatal(err)
	}
	if _, err := testDb.Exec(`Insert into OCD_CHEM_CARB_ANALYSIS_TEST (RUN_ID,ANALYSIS_CODE,METHOD_CODE,ANALYSIS_RESULT) values (3245,'H','CNS',0.09)`); err != nil {
		t.Fatal(err)
	}
	if _, err := testDb.Exec(`Insert into OCD_CHEM_CARB_ANALYSIS_TEST (RUN_ID,ANALYSIS_CODE,METHOD_CODE,ANALYSIS_RESULT) values (3245,'INOR_C','C',9.88)`); err != nil {
		t.Fatal(err)
	}
	if _, err := testDb.Exec(`Insert into OCD_CHEM_CARB_ANALYSIS_TEST (RUN_ID,ANALYSIS_CODE,METHOD_CODE,ANALYSIS_RESULT) values (3245,'ORG_C','CNS',0)`); err != nil {
		t.Fatal(err)
	}
	if _, err := testDb.Exec(`Insert into OCD_CHEM_CARB_ANALYSIS_TEST (RUN_ID,ANALYSIS_CODE,METHOD_CODE,ANALYSIS_RESULT) values (3245,'TOT_C','CNS',9.6)`); err != nil {
		t.Fatal(err)
	}
	if _, err := testDb.Exec(`Insert into OCD_CHEM_CARB_ANALYSIS_TEST (RUN_ID,ANALYSIS_CODE,METHOD_CODE,ANALYSIS_RESULT) values (3250,'CaCO3','C',29.66)`); err != nil {
		t.Fatal(err)
	}
	if _, err := testDb.Exec(`Insert into OCD_CHEM_CARB_ANALYSIS_TEST (RUN_ID,ANALYSIS_CODE,METHOD_CODE,ANALYSIS_RESULT) values (3250,'H','CNS',0.5)`); err != nil {
		t.Fatal(err)
	}
	if _, err := testDb.Exec(`Insert into OCD_CHEM_CARB_ANALYSIS_TEST (RUN_ID,ANALYSIS_CODE,METHOD_CODE,ANALYSIS_RESULT) values (3250,'INOR_C','C',3.56)`); err != nil {
		t.Fatal(err)
	}
	if _, err := testDb.Exec(`Insert into OCD_CHEM_CARB_ANALYSIS_TEST (RUN_ID,ANALYSIS_CODE,METHOD_CODE,ANALYSIS_RESULT) values (3250,'ORG_C','CNS',0.08)`); err != nil {
		t.Fatal(err)
	}
	if _, err := testDb.Exec(`Insert into OCD_CHEM_CARB_ANALYSIS_TEST (RUN_ID,ANALYSIS_CODE,METHOD_CODE,ANALYSIS_RESULT) values (3250,'PI','RE',0.25)`); err != nil {
		t.Fatal(err)
	}
	if _, err := testDb.Exec(`Insert into OCD_CHEM_CARB_ANALYSIS_TEST (RUN_ID,ANALYSIS_CODE,METHOD_CODE,ANALYSIS_RESULT) values (3250,'S1','RE',0.02)`); err != nil {
		t.Fatal(err)
	}
	if _, err := testDb.Exec(`Insert into OCD_CHEM_CARB_ANALYSIS_TEST (RUN_ID,ANALYSIS_CODE,METHOD_CODE,ANALYSIS_RESULT) values (3250,'S2','RE',0.1)`); err != nil {
		t.Fatal(err)
	}
	if _, err := testDb.Exec(`Insert into OCD_CHEM_CARB_ANALYSIS_TEST (RUN_ID,ANALYSIS_CODE,METHOD_CODE,ANALYSIS_RESULT) values (3250,'S3','RE',1.92)`); err != nil {
		t.Fatal(err)
	}
	if _, err := testDb.Exec(`Insert into OCD_CHEM_CARB_ANALYSIS_TEST (RUN_ID,ANALYSIS_CODE,METHOD_CODE,ANALYSIS_RESULT) values (3250,'TMX','RE',413)`); err != nil {
		t.Fatal(err)
	}
	if _, err := testDb.Exec(`Insert into OCD_CHEM_CARB_ANALYSIS_TEST (RUN_ID,ANALYSIS_CODE,METHOD_CODE,ANALYSIS_RESULT) values (3250,'TOC','RE',0)`); err != nil {
		t.Fatal(err)
	}
	if _, err := testDb.Exec(`Insert into OCD_CHEM_CARB_ANALYSIS_TEST (RUN_ID,ANALYSIS_CODE,METHOD_CODE,ANALYSIS_RESULT) values (3250,'TOT_C','CNS',3.64)`); err != nil {
		t.Fatal(err)
	}
	if _, err := testDb.Exec(`Insert into OCD_CHEM_CARB_ANALYSIS_TEST (RUN_ID,ANALYSIS_CODE,METHOD_CODE,ANALYSIS_RESULT) values (3250,'CaCO3','C',29.66)`); err != nil {
		t.Fatal(err)
	}
	if _, err := testDb.Exec(`Insert into OCD_CHEM_CARB_ANALYSIS_TEST (RUN_ID,ANALYSIS_CODE,METHOD_CODE,ANALYSIS_RESULT) values (3250,'H','CNS',0.5)`); err != nil {
		t.Fatal(err)
	}
	if _, err := testDb.Exec(`Insert into OCD_CHEM_CARB_ANALYSIS_TEST (RUN_ID,ANALYSIS_CODE,METHOD_CODE,ANALYSIS_RESULT) values (3250,'INOR_C','C',3.56)`); err != nil {
		t.Fatal(err)
	}
	if _, err := testDb.Exec(`Insert into OCD_CHEM_CARB_ANALYSIS_TEST (RUN_ID,ANALYSIS_CODE,METHOD_CODE,ANALYSIS_RESULT) values (3250,'ORG_C','CNS',0.08)`); err != nil {
		t.Fatal(err)
	}
	if _, err := testDb.Exec(`Insert into OCD_CHEM_CARB_ANALYSIS_TEST (RUN_ID,ANALYSIS_CODE,METHOD_CODE,ANALYSIS_RESULT) values (3250,'PI','RE',0.25)`); err != nil {
		t.Fatal(err)
	}
	if _, err := testDb.Exec(`Insert into OCD_CHEM_CARB_ANALYSIS_TEST (RUN_ID,ANALYSIS_CODE,METHOD_CODE,ANALYSIS_RESULT) values (3250,'S1','RE',0.02)`); err != nil {
		t.Fatal(err)
	}
	if _, err := testDb.Exec(`Insert into OCD_CHEM_CARB_ANALYSIS_TEST (RUN_ID,ANALYSIS_CODE,METHOD_CODE,ANALYSIS_RESULT) values (3250,'S2','RE',0.1)`); err != nil {
		t.Fatal(err)
	}
	if _, err := testDb.Exec(`Insert into OCD_CHEM_CARB_ANALYSIS_TEST (RUN_ID,ANALYSIS_CODE,METHOD_CODE,ANALYSIS_RESULT) values (3250,'S3','RE',1.92)`); err != nil {
		t.Fatal(err)
	}
	if _, err := testDb.Exec(`Insert into OCD_CHEM_CARB_ANALYSIS_TEST (RUN_ID,ANALYSIS_CODE,METHOD_CODE,ANALYSIS_RESULT) values (3250,'TMX','RE',413)`); err != nil {
		t.Fatal(err)
	}
	if _, err := testDb.Exec(`Insert into OCD_CHEM_CARB_ANALYSIS_TEST (RUN_ID,ANALYSIS_CODE,METHOD_CODE,ANALYSIS_RESULT) values (3250,'TOC','RE',0)`); err != nil {
		t.Fatal(err)
	}
	if _, err := testDb.Exec(`Insert into OCD_CHEM_CARB_ANALYSIS_TEST (RUN_ID,ANALYSIS_CODE,METHOD_CODE,ANALYSIS_RESULT) values (3250,'TOT_C','CNS',3.64)`); err != nil {
		t.Fatal(err)
	}
	if _, err := testDb.Exec(`Insert into OCD_CHEM_CARB_ANALYSIS_TEST (RUN_ID,ANALYSIS_CODE,METHOD_CODE,ANALYSIS_RESULT) values (3255,'CaCO3','C',63.12)`); err != nil {
		t.Fatal(err)
	}
	if _, err := testDb.Exec(`Insert into OCD_CHEM_CARB_ANALYSIS_TEST (RUN_ID,ANALYSIS_CODE,METHOD_CODE,ANALYSIS_RESULT) values (3255,'H','CNS',0.24)`); err != nil {
		t.Fatal(err)
	}
	if _, err := testDb.Exec(`Insert into OCD_CHEM_CARB_ANALYSIS_TEST (RUN_ID,ANALYSIS_CODE,METHOD_CODE,ANALYSIS_RESULT) values (3255,'HI','RE',50)`); err != nil {
		t.Fatal(err)
	}
	if _, err := testDb.Exec(`Insert into OCD_CHEM_CARB_ANALYSIS_TEST (RUN_ID,ANALYSIS_CODE,METHOD_CODE,ANALYSIS_RESULT) values (3255,'INOR_C','C',7.58)`); err != nil {
		t.Fatal(err)
	}
	if _, err := testDb.Exec(`Insert into OCD_CHEM_CARB_ANALYSIS_TEST (RUN_ID,ANALYSIS_CODE,METHOD_CODE,ANALYSIS_RESULT) values (3255,'OI','RE',1675)`); err != nil {
		t.Fatal(err)
	}
	if _, err := testDb.Exec(`Insert into OCD_CHEM_CARB_ANALYSIS_TEST (RUN_ID,ANALYSIS_CODE,METHOD_CODE,ANALYSIS_RESULT) values (3255,'ORG_C','CNS',0.16)`); err != nil {
		t.Fatal(err)
	}
	if _, err := testDb.Exec(`Insert into OCD_CHEM_CARB_ANALYSIS_TEST (RUN_ID,ANALYSIS_CODE,METHOD_CODE,ANALYSIS_RESULT) values (3255,'PI','RE',0)`); err != nil {
		t.Fatal(err)
	}
	if _, err := testDb.Exec(`Insert into OCD_CHEM_CARB_ANALYSIS_TEST (RUN_ID,ANALYSIS_CODE,METHOD_CODE,ANALYSIS_RESULT) values (3255,'S1','RE',0)`); err != nil {
		t.Fatal(err)
	}
	if _, err := testDb.Exec(`Insert into OCD_CHEM_CARB_ANALYSIS_TEST (RUN_ID,ANALYSIS_CODE,METHOD_CODE,ANALYSIS_RESULT) values (3255,'S2','RE',0.04)`); err != nil {
		t.Fatal(err)
	}
	if _, err := testDb.Exec(`Insert into OCD_CHEM_CARB_ANALYSIS_TEST (RUN_ID,ANALYSIS_CODE,METHOD_CODE,ANALYSIS_RESULT) values (3255,'S3','RE',1.34)`); err != nil {
		t.Fatal(err)
	}
	if _, err := testDb.Exec(`Insert into OCD_CHEM_CARB_ANALYSIS_TEST (RUN_ID,ANALYSIS_CODE,METHOD_CODE,ANALYSIS_RESULT) values (3255,'TMX','RE',410)`); err != nil {
		t.Fatal(err)
	}
	if _, err := testDb.Exec(`Insert into OCD_CHEM_CARB_ANALYSIS_TEST (RUN_ID,ANALYSIS_CODE,METHOD_CODE,ANALYSIS_RESULT) values (3255,'TOC','RE',0.08)`); err != nil {
		t.Fatal(err)
	}
	if _, err := testDb.Exec(`Insert into OCD_CHEM_CARB_ANALYSIS_TEST (RUN_ID,ANALYSIS_CODE,METHOD_CODE,ANALYSIS_RESULT) values (3255,'TOT_C','CNS',7.74)`); err != nil {
		t.Fatal(err)
	}
	if _, err := testDb.Exec(`Insert into OCD_CHEM_CARB_ANALYSIS_TEST (RUN_ID,ANALYSIS_CODE,METHOD_CODE,ANALYSIS_RESULT) values (3255,'CaCO3','C',63.12)`); err != nil {
		t.Fatal(err)
	}
	if _, err := testDb.Exec(`Insert into OCD_CHEM_CARB_ANALYSIS_TEST (RUN_ID,ANALYSIS_CODE,METHOD_CODE,ANALYSIS_RESULT) values (3255,'H','CNS',0.24)`); err != nil {
		t.Fatal(err)
	}
	if _, err := testDb.Exec(`Insert into OCD_CHEM_CARB_ANALYSIS_TEST (RUN_ID,ANALYSIS_CODE,METHOD_CODE,ANALYSIS_RESULT) values (3255,'HI','RE',50)`); err != nil {
		t.Fatal(err)
	}
	if _, err := testDb.Exec(`Insert into OCD_CHEM_CARB_ANALYSIS_TEST (RUN_ID,ANALYSIS_CODE,METHOD_CODE,ANALYSIS_RESULT) values (3255,'INOR_C','C',7.58)`); err != nil {
		t.Fatal(err)
	}
	if _, err := testDb.Exec(`Insert into OCD_CHEM_CARB_ANALYSIS_TEST (RUN_ID,ANALYSIS_CODE,METHOD_CODE,ANALYSIS_RESULT) values (3255,'OI','RE',1675)`); err != nil {
		t.Fatal(err)
	}
	if _, err := testDb.Exec(`Insert into OCD_CHEM_CARB_ANALYSIS_TEST (RUN_ID,ANALYSIS_CODE,METHOD_CODE,ANALYSIS_RESULT) values (3255,'ORG_C','CNS',0.16)`); err != nil {
		t.Fatal(err)
	}
	if _, err := testDb.Exec(`Insert into OCD_CHEM_CARB_ANALYSIS_TEST (RUN_ID,ANALYSIS_CODE,METHOD_CODE,ANALYSIS_RESULT) values (3255,'PI','RE',0)`); err != nil {
		t.Fatal(err)
	}
	if _, err := testDb.Exec(`Insert into OCD_CHEM_CARB_ANALYSIS_TEST (RUN_ID,ANALYSIS_CODE,METHOD_CODE,ANALYSIS_RESULT) values (3255,'S1','RE',0)`); err != nil {
		t.Fatal(err)
	}
	if _, err := testDb.Exec(`Insert into OCD_CHEM_CARB_ANALYSIS_TEST (RUN_ID,ANALYSIS_CODE,METHOD_CODE,ANALYSIS_RESULT) values (3255,'S2','RE',0.04)`); err != nil {
		t.Fatal(err)
	}
	if _, err := testDb.Exec(`Insert into OCD_CHEM_CARB_ANALYSIS_TEST (RUN_ID,ANALYSIS_CODE,METHOD_CODE,ANALYSIS_RESULT) values (3255,'S3','RE',1.34)`); err != nil {
		t.Fatal(err)
	}
	if _, err := testDb.Exec(`Insert into OCD_CHEM_CARB_ANALYSIS_TEST (RUN_ID,ANALYSIS_CODE,METHOD_CODE,ANALYSIS_RESULT) values (3255,'TMX','RE',410)`); err != nil {
		t.Fatal(err)
	}
	if _, err := testDb.Exec(`Insert into OCD_CHEM_CARB_ANALYSIS_TEST (RUN_ID,ANALYSIS_CODE,METHOD_CODE,ANALYSIS_RESULT) values (3255,'TOC','RE',0.08)`); err != nil {
		t.Fatal(err)
	}
	if _, err := testDb.Exec(`Insert into OCD_CHEM_CARB_ANALYSIS_TEST (RUN_ID,ANALYSIS_CODE,METHOD_CODE,ANALYSIS_RESULT) values (3255,'TOT_C','CNS',7.74)`); err != nil {
		t.Fatal(err)
	}
	if _, err := testDb.Exec(`Insert into OCD_CHEM_CARB_ANALYSIS_TEST (RUN_ID,ANALYSIS_CODE,METHOD_CODE,ANALYSIS_RESULT) values (3260,'CaCO3','C',55.44)`); err != nil {
		t.Fatal(err)
	}
	if _, err := testDb.Exec(`Insert into OCD_CHEM_CARB_ANALYSIS_TEST (RUN_ID,ANALYSIS_CODE,METHOD_CODE,ANALYSIS_RESULT) values (3260,'H','CNS',0.46)`); err != nil {
		t.Fatal(err)
	}
	if _, err := testDb.Exec(`Insert into OCD_CHEM_CARB_ANALYSIS_TEST (RUN_ID,ANALYSIS_CODE,METHOD_CODE,ANALYSIS_RESULT) values (3260,'HI','RE',543)`); err != nil {
		t.Fatal(err)
	}
	if _, err := testDb.Exec(`Insert into OCD_CHEM_CARB_ANALYSIS_TEST (RUN_ID,ANALYSIS_CODE,METHOD_CODE,ANALYSIS_RESULT) values (3260,'INOR_C','C',6.66)`); err != nil {
		t.Fatal(err)
	}
	if _, err := testDb.Exec(`Insert into OCD_CHEM_CARB_ANALYSIS_TEST (RUN_ID,ANALYSIS_CODE,METHOD_CODE,ANALYSIS_RESULT) values (3260,'NIT','CNS',0.014)`); err != nil {
		t.Fatal(err)
	}
	if _, err := testDb.Exec(`Insert into OCD_CHEM_CARB_ANALYSIS_TEST (RUN_ID,ANALYSIS_CODE,METHOD_CODE,ANALYSIS_RESULT) values (3260,'OI','RE',152)`); err != nil {
		t.Fatal(err)
	}
	if _, err := testDb.Exec(`Insert into OCD_CHEM_CARB_ANALYSIS_TEST (RUN_ID,ANALYSIS_CODE,METHOD_CODE,ANALYSIS_RESULT) values (3260,'ORG_C','CNS',1.68)`); err != nil {
		t.Fatal(err)
	}
	if _, err := testDb.Exec(`Insert into OCD_CHEM_CARB_ANALYSIS_TEST (RUN_ID,ANALYSIS_CODE,METHOD_CODE,ANALYSIS_RESULT) values (3260,'PI','RE',0.01)`); err != nil {
		t.Fatal(err)
	}
	if _, err := testDb.Exec(`Insert into OCD_CHEM_CARB_ANALYSIS_TEST (RUN_ID,ANALYSIS_CODE,METHOD_CODE,ANALYSIS_RESULT) values (3260,'S1','RE',0.08)`); err != nil {
		t.Fatal(err)
	}
	if _, err := testDb.Exec(`Insert into OCD_CHEM_CARB_ANALYSIS_TEST (RUN_ID,ANALYSIS_CODE,METHOD_CODE,ANALYSIS_RESULT) values (3260,'S2','RE',6.6)`); err != nil {
		t.Fatal(err)
	}
	if _, err := testDb.Exec(`Insert into OCD_CHEM_CARB_ANALYSIS_TEST (RUN_ID,ANALYSIS_CODE,METHOD_CODE,ANALYSIS_RESULT) values (3260,'S3','RE',1.86)`); err != nil {
		t.Fatal(err)
	}
	if _, err := testDb.Exec(`Insert into OCD_CHEM_CARB_ANALYSIS_TEST (RUN_ID,ANALYSIS_CODE,METHOD_CODE,ANALYSIS_RESULT) values (3260,'TMX','RE',403)`); err != nil {
		t.Fatal(err)
	}
	if _, err := testDb.Exec(`Insert into OCD_CHEM_CARB_ANALYSIS_TEST (RUN_ID,ANALYSIS_CODE,METHOD_CODE,ANALYSIS_RESULT) values (3260,'TOC','RE',1.22)`); err != nil {
		t.Fatal(err)
	}
	if _, err := testDb.Exec(`Insert into OCD_CHEM_CARB_ANALYSIS_TEST (RUN_ID,ANALYSIS_CODE,METHOD_CODE,ANALYSIS_RESULT) values (3260,'TOT_C','CNS',8.34)`); err != nil {
		t.Fatal(err)
	}
	if _, err := testDb.Exec(`Insert into OCD_CHEM_CARB_ANALYSIS_TEST (RUN_ID,ANALYSIS_CODE,METHOD_CODE,ANALYSIS_RESULT) values (3260,'CaCO3','C',55.44)`); err != nil {
		t.Fatal(err)
	}
	if _, err := testDb.Exec(`Insert into OCD_CHEM_CARB_ANALYSIS_TEST (RUN_ID,ANALYSIS_CODE,METHOD_CODE,ANALYSIS_RESULT) values (3260,'H','CNS',0.46)`); err != nil {
		t.Fatal(err)
	}
	if _, err := testDb.Exec(`Insert into OCD_CHEM_CARB_ANALYSIS_TEST (RUN_ID,ANALYSIS_CODE,METHOD_CODE,ANALYSIS_RESULT) values (3260,'HI','RE',543)`); err != nil {
		t.Fatal(err)
	}
	if _, err := testDb.Exec(`Insert into OCD_CHEM_CARB_ANALYSIS_TEST (RUN_ID,ANALYSIS_CODE,METHOD_CODE,ANALYSIS_RESULT) values (3260,'INOR_C','C',6.66)`); err != nil {
		t.Fatal(err)
	}
	if _, err := testDb.Exec(`Insert into OCD_CHEM_CARB_ANALYSIS_TEST (RUN_ID,ANALYSIS_CODE,METHOD_CODE,ANALYSIS_RESULT) values (3260,'NIT','CNS',0.014)`); err != nil {
		t.Fatal(err)
	}
	if _, err := testDb.Exec(`Insert into OCD_CHEM_CARB_ANALYSIS_TEST (RUN_ID,ANALYSIS_CODE,METHOD_CODE,ANALYSIS_RESULT) values (3260,'OI','RE',152)`); err != nil {
		t.Fatal(err)
	}
	if _, err := testDb.Exec(`Insert into OCD_CHEM_CARB_ANALYSIS_TEST (RUN_ID,ANALYSIS_CODE,METHOD_CODE,ANALYSIS_RESULT) values (3260,'ORG_C','CNS',1.68)`); err != nil {
		t.Fatal(err)
	}
	if _, err := testDb.Exec(`Insert into OCD_CHEM_CARB_ANALYSIS_TEST (RUN_ID,ANALYSIS_CODE,METHOD_CODE,ANALYSIS_RESULT) values (3260,'PI','RE',0.01)`); err != nil {
		t.Fatal(err)
	}
	if _, err := testDb.Exec(`Insert into OCD_CHEM_CARB_ANALYSIS_TEST (RUN_ID,ANALYSIS_CODE,METHOD_CODE,ANALYSIS_RESULT) values (3260,'S1','RE',0.08)`); err != nil {
		t.Fatal(err)
	}
	if _, err := testDb.Exec(`Insert into OCD_CHEM_CARB_ANALYSIS_TEST (RUN_ID,ANALYSIS_CODE,METHOD_CODE,ANALYSIS_RESULT) values (3260,'S2','RE',6.6)`); err != nil {
		t.Fatal(err)
	}
	if _, err := testDb.Exec(`Insert into OCD_CHEM_CARB_ANALYSIS_TEST (RUN_ID,ANALYSIS_CODE,METHOD_CODE,ANALYSIS_RESULT) values (3260,'S3','RE',1.86)`); err != nil {
		t.Fatal(err)
	}
	if _, err := testDb.Exec(`Insert into OCD_CHEM_CARB_ANALYSIS_TEST (RUN_ID,ANALYSIS_CODE,METHOD_CODE,ANALYSIS_RESULT) values (3260,'TMX','RE',403)`); err != nil {
		t.Fatal(err)
	}
	if _, err := testDb.Exec(`Insert into OCD_CHEM_CARB_ANALYSIS_TEST (RUN_ID,ANALYSIS_CODE,METHOD_CODE,ANALYSIS_RESULT) values (3260,'TOC','RE',1.22)`); err != nil {
		t.Fatal(err)
	}
	if _, err := testDb.Exec(`Insert into OCD_CHEM_CARB_ANALYSIS_TEST (RUN_ID,ANALYSIS_CODE,METHOD_CODE,ANALYSIS_RESULT) values (3260,'TOT_C','CNS',8.34)`); err != nil {
		t.Fatal(err)
	}
	if _, err := testDb.Exec(`Insert into OCD_CHEM_CARB_ANALYSIS_TEST (RUN_ID,ANALYSIS_CODE,METHOD_CODE,ANALYSIS_RESULT) values (3265,'CaCO3','C',51.53)`); err != nil {
		t.Fatal(err)
	}
	if _, err := testDb.Exec(`Insert into OCD_CHEM_CARB_ANALYSIS_TEST (RUN_ID,ANALYSIS_CODE,METHOD_CODE,ANALYSIS_RESULT) values (3265,'H','CNS',0.64)`); err != nil {
		t.Fatal(err)
	}
	if _, err := testDb.Exec(`Insert into OCD_CHEM_CARB_ANALYSIS_TEST (RUN_ID,ANALYSIS_CODE,METHOD_CODE,ANALYSIS_RESULT) values (3265,'INOR_C','C',6.19)`); err != nil {
		t.Fatal(err)
	}
	if _, err := testDb.Exec(`Insert into OCD_CHEM_CARB_ANALYSIS_TEST (RUN_ID,ANALYSIS_CODE,METHOD_CODE,ANALYSIS_RESULT) values (3265,'NIT','CNS',0.04)`); err != nil {
		t.Fatal(err)
	}
	if _, err := testDb.Exec(`Insert into OCD_CHEM_CARB_ANALYSIS_TEST (RUN_ID,ANALYSIS_CODE,METHOD_CODE,ANALYSIS_RESULT) values (3265,'ORG_C','CNS',2.97)`); err != nil {
		t.Fatal(err)
	}
	if _, err := testDb.Exec(`Insert into OCD_CHEM_CARB_ANALYSIS_TEST (RUN_ID,ANALYSIS_CODE,METHOD_CODE,ANALYSIS_RESULT) values (3265,'TOT_C','CNS',9.16)`); err != nil {
		t.Fatal(err)
	}
	if _, err := testDb.Exec(`Insert into OCD_CHEM_CARB_ANALYSIS_TEST (RUN_ID,ANALYSIS_CODE,METHOD_CODE,ANALYSIS_RESULT) values (3265,'CaCO3','C',51.53)`); err != nil {
		t.Fatal(err)
	}
	if _, err := testDb.Exec(`Insert into OCD_CHEM_CARB_ANALYSIS_TEST (RUN_ID,ANALYSIS_CODE,METHOD_CODE,ANALYSIS_RESULT) values (3265,'H','CNS',0.64)`); err != nil {
		t.Fatal(err)
	}
	if _, err := testDb.Exec(`Insert into OCD_CHEM_CARB_ANALYSIS_TEST (RUN_ID,ANALYSIS_CODE,METHOD_CODE,ANALYSIS_RESULT) values (3265,'INOR_C','C',6.19)`); err != nil {
		t.Fatal(err)
	}
	if _, err := testDb.Exec(`Insert into OCD_CHEM_CARB_ANALYSIS_TEST (RUN_ID,ANALYSIS_CODE,METHOD_CODE,ANALYSIS_RESULT) values (3265,'NIT','CNS',0.04)`); err != nil {
		t.Fatal(err)
	}
	if _, err := testDb.Exec(`Insert into OCD_CHEM_CARB_ANALYSIS_TEST (RUN_ID,ANALYSIS_CODE,METHOD_CODE,ANALYSIS_RESULT) values (3265,'ORG_C','CNS',2.97)`); err != nil {
		t.Fatal(err)
	}
	if _, err := testDb.Exec(`Insert into OCD_CHEM_CARB_ANALYSIS_TEST (RUN_ID,ANALYSIS_CODE,METHOD_CODE,ANALYSIS_RESULT) values (3265,'TOT_C','CNS',9.16)`); err != nil {
		t.Fatal(err)
	}
	if _, err := testDb.Exec(`Insert into OCD_CHEM_CARB_ANALYSIS_TEST (RUN_ID,ANALYSIS_CODE,METHOD_CODE,ANALYSIS_RESULT) values (3270,'CaCO3','C',74.48)`); err != nil {
		t.Fatal(err)
	}
	if _, err := testDb.Exec(`Insert into OCD_CHEM_CARB_ANALYSIS_TEST (RUN_ID,ANALYSIS_CODE,METHOD_CODE,ANALYSIS_RESULT) values (3270,'H','CNS',0.34)`); err != nil {
		t.Fatal(err)
	}
	if _, err := testDb.Exec(`Insert into OCD_CHEM_CARB_ANALYSIS_TEST (RUN_ID,ANALYSIS_CODE,METHOD_CODE,ANALYSIS_RESULT) values (3270,'HI','RE',605)`); err != nil {
		t.Fatal(err)
	}
	if _, err := testDb.Exec(`Insert into OCD_CHEM_CARB_ANALYSIS_TEST (RUN_ID,ANALYSIS_CODE,METHOD_CODE,ANALYSIS_RESULT) values (3270,'INOR_C','C',8.94)`); err != nil {
		t.Fatal(err)
	}
	if _, err := testDb.Exec(`Insert into OCD_CHEM_CARB_ANALYSIS_TEST (RUN_ID,ANALYSIS_CODE,METHOD_CODE,ANALYSIS_RESULT) values (3270,'NIT','CNS',0.012)`); err != nil {
		t.Fatal(err)
	}
	if _, err := testDb.Exec(`Insert into OCD_CHEM_CARB_ANALYSIS_TEST (RUN_ID,ANALYSIS_CODE,METHOD_CODE,ANALYSIS_RESULT) values (3270,'OI','RE',85)`); err != nil {
		t.Fatal(err)
	}
	if _, err := testDb.Exec(`Insert into OCD_CHEM_CARB_ANALYSIS_TEST (RUN_ID,ANALYSIS_CODE,METHOD_CODE,ANALYSIS_RESULT) values (3270,'ORG_C','CNS',1.69)`); err != nil {
		t.Fatal(err)
	}
	if _, err := testDb.Exec(`Insert into OCD_CHEM_CARB_ANALYSIS_TEST (RUN_ID,ANALYSIS_CODE,METHOD_CODE,ANALYSIS_RESULT) values (3270,'PI','RE',0.01)`); err != nil {
		t.Fatal(err)
	}
	if _, err := testDb.Exec(`Insert into OCD_CHEM_CARB_ANALYSIS_TEST (RUN_ID,ANALYSIS_CODE,METHOD_CODE,ANALYSIS_RESULT) values (3270,'S1','RE',0.1)`); err != nil {
		t.Fatal(err)
	}
	if _, err := testDb.Exec(`Insert into OCD_CHEM_CARB_ANALYSIS_TEST (RUN_ID,ANALYSIS_CODE,METHOD_CODE,ANALYSIS_RESULT) values (3270,'S2','RE',8.4)`); err != nil {
		t.Fatal(err)
	}
	if _, err := testDb.Exec(`Insert into OCD_CHEM_CARB_ANALYSIS_TEST (RUN_ID,ANALYSIS_CODE,METHOD_CODE,ANALYSIS_RESULT) values (3270,'S3','RE',1.19)`); err != nil {
		t.Fatal(err)
	}
	if _, err := testDb.Exec(`Insert into OCD_CHEM_CARB_ANALYSIS_TEST (RUN_ID,ANALYSIS_CODE,METHOD_CODE,ANALYSIS_RESULT) values (3270,'SUL','CNS',0.1)`); err != nil {
		t.Fatal(err)
	}
	if _, err := testDb.Exec(`Insert into OCD_CHEM_CARB_ANALYSIS_TEST (RUN_ID,ANALYSIS_CODE,METHOD_CODE,ANALYSIS_RESULT) values (3270,'TMX','RE',395)`); err != nil {
		t.Fatal(err)
	}
	if _, err := testDb.Exec(`Insert into OCD_CHEM_CARB_ANALYSIS_TEST (RUN_ID,ANALYSIS_CODE,METHOD_CODE,ANALYSIS_RESULT) values (3270,'TOC','RE',1.39)`); err != nil {
		t.Fatal(err)
	}
	if _, err := testDb.Exec(`Insert into OCD_CHEM_CARB_ANALYSIS_TEST (RUN_ID,ANALYSIS_CODE,METHOD_CODE,ANALYSIS_RESULT) values (3270,'TOT_C','CNS',10.63)`); err != nil {
		t.Fatal(err)
	}
	if _, err := testDb.Exec(`Insert into OCD_CHEM_CARB_ANALYSIS_TEST (RUN_ID,ANALYSIS_CODE,METHOD_CODE,ANALYSIS_RESULT) values (3270,'CaCO3','C',74.48)`); err != nil {
		t.Fatal(err)
	}
	if _, err := testDb.Exec(`Insert into OCD_CHEM_CARB_ANALYSIS_TEST (RUN_ID,ANALYSIS_CODE,METHOD_CODE,ANALYSIS_RESULT) values (3270,'H','CNS',0.34)`); err != nil {
		t.Fatal(err)
	}
	if _, err := testDb.Exec(`Insert into OCD_CHEM_CARB_ANALYSIS_TEST (RUN_ID,ANALYSIS_CODE,METHOD_CODE,ANALYSIS_RESULT) values (3270,'HI','RE',605)`); err != nil {
		t.Fatal(err)
	}
	if _, err := testDb.Exec(`Insert into OCD_CHEM_CARB_ANALYSIS_TEST (RUN_ID,ANALYSIS_CODE,METHOD_CODE,ANALYSIS_RESULT) values (3270,'INOR_C','C',8.94)`); err != nil {
		t.Fatal(err)
	}
	if _, err := testDb.Exec(`Insert into OCD_CHEM_CARB_ANALYSIS_TEST (RUN_ID,ANALYSIS_CODE,METHOD_CODE,ANALYSIS_RESULT) values (3270,'NIT','CNS',0.012)`); err != nil {
		t.Fatal(err)
	}
	if _, err := testDb.Exec(`Insert into OCD_CHEM_CARB_ANALYSIS_TEST (RUN_ID,ANALYSIS_CODE,METHOD_CODE,ANALYSIS_RESULT) values (3270,'OI','RE',85)`); err != nil {
		t.Fatal(err)
	}
	if _, err := testDb.Exec(`Insert into OCD_CHEM_CARB_ANALYSIS_TEST (RUN_ID,ANALYSIS_CODE,METHOD_CODE,ANALYSIS_RESULT) values (3270,'ORG_C','CNS',1.69)`); err != nil {
		t.Fatal(err)
	}
	if _, err := testDb.Exec(`Insert into OCD_CHEM_CARB_ANALYSIS_TEST (RUN_ID,ANALYSIS_CODE,METHOD_CODE,ANALYSIS_RESULT) values (3270,'PI','RE',0.01)`); err != nil {
		t.Fatal(err)
	}
	if _, err := testDb.Exec(`Insert into OCD_CHEM_CARB_ANALYSIS_TEST (RUN_ID,ANALYSIS_CODE,METHOD_CODE,ANALYSIS_RESULT) values (3270,'S1','RE',0.1)`); err != nil {
		t.Fatal(err)
	}
	if _, err := testDb.Exec(`Insert into OCD_CHEM_CARB_ANALYSIS_TEST (RUN_ID,ANALYSIS_CODE,METHOD_CODE,ANALYSIS_RESULT) values (3270,'S2','RE',8.4)`); err != nil {
		t.Fatal(err)
	}
	if _, err := testDb.Exec(`Insert into OCD_CHEM_CARB_ANALYSIS_TEST (RUN_ID,ANALYSIS_CODE,METHOD_CODE,ANALYSIS_RESULT) values (3270,'S3','RE',1.19)`); err != nil {
		t.Fatal(err)
	}
	if _, err := testDb.Exec(`Insert into OCD_CHEM_CARB_ANALYSIS_TEST (RUN_ID,ANALYSIS_CODE,METHOD_CODE,ANALYSIS_RESULT) values (3270,'SUL','CNS',0.1)`); err != nil {
		t.Fatal(err)
	}
	if _, err := testDb.Exec(`Insert into OCD_CHEM_CARB_ANALYSIS_TEST (RUN_ID,ANALYSIS_CODE,METHOD_CODE,ANALYSIS_RESULT) values (3270,'TMX','RE',395)`); err != nil {
		t.Fatal(err)
	}
	if _, err := testDb.Exec(`Insert into OCD_CHEM_CARB_ANALYSIS_TEST (RUN_ID,ANALYSIS_CODE,METHOD_CODE,ANALYSIS_RESULT) values (3270,'TOC','RE',1.39)`); err != nil {
		t.Fatal(err)
	}
	if _, err := testDb.Exec(`Insert into OCD_CHEM_CARB_ANALYSIS_TEST (RUN_ID,ANALYSIS_CODE,METHOD_CODE,ANALYSIS_RESULT) values (3270,'TOT_C','CNS',10.63)`); err != nil {
		t.Fatal(err)
	}
	if _, err := testDb.Exec(`Insert into OCD_CHEM_CARB_ANALYSIS_TEST (RUN_ID,ANALYSIS_CODE,METHOD_CODE,ANALYSIS_RESULT) values (3275,'CaCO3','C',41.65)`); err != nil {
		t.Fatal(err)
	}
	if _, err := testDb.Exec(`Insert into OCD_CHEM_CARB_ANALYSIS_TEST (RUN_ID,ANALYSIS_CODE,METHOD_CODE,ANALYSIS_RESULT) values (3275,'H','CNS',0.73)`); err != nil {
		t.Fatal(err)
	}
	if _, err := testDb.Exec(`Insert into OCD_CHEM_CARB_ANALYSIS_TEST (RUN_ID,ANALYSIS_CODE,METHOD_CODE,ANALYSIS_RESULT) values (3275,'HI','RE',485)`); err != nil {
		t.Fatal(err)
	}
	if _, err := testDb.Exec(`Insert into OCD_CHEM_CARB_ANALYSIS_TEST (RUN_ID,ANALYSIS_CODE,METHOD_CODE,ANALYSIS_RESULT) values (3275,'INOR_C','C',5)`); err != nil {
		t.Fatal(err)
	}
	if _, err := testDb.Exec(`Insert into OCD_CHEM_CARB_ANALYSIS_TEST (RUN_ID,ANALYSIS_CODE,METHOD_CODE,ANALYSIS_RESULT) values (3275,'NIT','CNS',0.038)`); err != nil {
		t.Fatal(err)
	}
	if _, err := testDb.Exec(`Insert into OCD_CHEM_CARB_ANALYSIS_TEST (RUN_ID,ANALYSIS_CODE,METHOD_CODE,ANALYSIS_RESULT) values (3275,'OI','RE',79)`); err != nil {
		t.Fatal(err)
	}
	if _, err := testDb.Exec(`Insert into OCD_CHEM_CARB_ANALYSIS_TEST (RUN_ID,ANALYSIS_CODE,METHOD_CODE,ANALYSIS_RESULT) values (3275,'ORG_C','CNS',3.99)`); err != nil {
		t.Fatal(err)
	}
	if _, err := testDb.Exec(`Insert into OCD_CHEM_CARB_ANALYSIS_TEST (RUN_ID,ANALYSIS_CODE,METHOD_CODE,ANALYSIS_RESULT) values (3275,'PI','RE',0.01)`); err != nil {
		t.Fatal(err)
	}
	if _, err := testDb.Exec(`Insert into OCD_CHEM_CARB_ANALYSIS_TEST (RUN_ID,ANALYSIS_CODE,METHOD_CODE,ANALYSIS_RESULT) values (3275,'S1','RE',0.18)`); err != nil {
		t.Fatal(err)
	}
	if _, err := testDb.Exec(`Insert into OCD_CHEM_CARB_ANALYSIS_TEST (RUN_ID,ANALYSIS_CODE,METHOD_CODE,ANALYSIS_RESULT) values (3275,'S2','RE',15.6)`); err != nil {
		t.Fatal(err)
	}
	if _, err := testDb.Exec(`Insert into OCD_CHEM_CARB_ANALYSIS_TEST (RUN_ID,ANALYSIS_CODE,METHOD_CODE,ANALYSIS_RESULT) values (3275,'S3','RE',2.55)`); err != nil {
		t.Fatal(err)
	}
	if _, err := testDb.Exec(`Insert into OCD_CHEM_CARB_ANALYSIS_TEST (RUN_ID,ANALYSIS_CODE,METHOD_CODE,ANALYSIS_RESULT) values (3275,'TMX','RE',406)`); err != nil {
		t.Fatal(err)
	}
	if _, err := testDb.Exec(`Insert into OCD_CHEM_CARB_ANALYSIS_TEST (RUN_ID,ANALYSIS_CODE,METHOD_CODE,ANALYSIS_RESULT) values (3275,'TOC','RE',3.2)`); err != nil {
		t.Fatal(err)
	}
	if _, err := testDb.Exec(`Insert into OCD_CHEM_CARB_ANALYSIS_TEST (RUN_ID,ANALYSIS_CODE,METHOD_CODE,ANALYSIS_RESULT) values (3275,'TOT_C','CNS',8.99)`); err != nil {
		t.Fatal(err)
	}
	if _, err := testDb.Exec(`Insert into OCD_CHEM_CARB_ANALYSIS_TEST (RUN_ID,ANALYSIS_CODE,METHOD_CODE,ANALYSIS_RESULT) values (3275,'CaCO3','C',41.65)`); err != nil {
		t.Fatal(err)
	}
	if _, err := testDb.Exec(`Insert into OCD_CHEM_CARB_ANALYSIS_TEST (RUN_ID,ANALYSIS_CODE,METHOD_CODE,ANALYSIS_RESULT) values (3275,'H','CNS',0.73)`); err != nil {
		t.Fatal(err)
	}
	if _, err := testDb.Exec(`Insert into OCD_CHEM_CARB_ANALYSIS_TEST (RUN_ID,ANALYSIS_CODE,METHOD_CODE,ANALYSIS_RESULT) values (3275,'HI','RE',485)`); err != nil {
		t.Fatal(err)
	}
	if _, err := testDb.Exec(`Insert into OCD_CHEM_CARB_ANALYSIS_TEST (RUN_ID,ANALYSIS_CODE,METHOD_CODE,ANALYSIS_RESULT) values (3275,'INOR_C','C',5)`); err != nil {
		t.Fatal(err)
	}
	if _, err := testDb.Exec(`Insert into OCD_CHEM_CARB_ANALYSIS_TEST (RUN_ID,ANALYSIS_CODE,METHOD_CODE,ANALYSIS_RESULT) values (3275,'NIT','CNS',0.038)`); err != nil {
		t.Fatal(err)
	}
	if _, err := testDb.Exec(`Insert into OCD_CHEM_CARB_ANALYSIS_TEST (RUN_ID,ANALYSIS_CODE,METHOD_CODE,ANALYSIS_RESULT) values (3275,'OI','RE',79)`); err != nil {
		t.Fatal(err)
	}
	if _, err := testDb.Exec(`Insert into OCD_CHEM_CARB_ANALYSIS_TEST (RUN_ID,ANALYSIS_CODE,METHOD_CODE,ANALYSIS_RESULT) values (3275,'ORG_C','CNS',3.99)`); err != nil {
		t.Fatal(err)
	}
	if _, err := testDb.Exec(`Insert into OCD_CHEM_CARB_ANALYSIS_TEST (RUN_ID,ANALYSIS_CODE,METHOD_CODE,ANALYSIS_RESULT) values (3275,'PI','RE',0.01)`); err != nil {
		t.Fatal(err)
	}
	if _, err := testDb.Exec(`Insert into OCD_CHEM_CARB_ANALYSIS_TEST (RUN_ID,ANALYSIS_CODE,METHOD_CODE,ANALYSIS_RESULT) values (3275,'S1','RE',0.18)`); err != nil {
		t.Fatal(err)
	}
	if _, err := testDb.Exec(`Insert into OCD_CHEM_CARB_ANALYSIS_TEST (RUN_ID,ANALYSIS_CODE,METHOD_CODE,ANALYSIS_RESULT) values (3275,'S2','RE',15.6)`); err != nil {
		t.Fatal(err)
	}
	if _, err := testDb.Exec(`Insert into OCD_CHEM_CARB_ANALYSIS_TEST (RUN_ID,ANALYSIS_CODE,METHOD_CODE,ANALYSIS_RESULT) values (3275,'S3','RE',2.55)`); err != nil {
		t.Fatal(err)
	}
	if _, err := testDb.Exec(`Insert into OCD_CHEM_CARB_ANALYSIS_TEST (RUN_ID,ANALYSIS_CODE,METHOD_CODE,ANALYSIS_RESULT) values (3275,'TMX','RE',406)`); err != nil {
		t.Fatal(err)
	}
	if _, err := testDb.Exec(`Insert into OCD_CHEM_CARB_ANALYSIS_TEST (RUN_ID,ANALYSIS_CODE,METHOD_CODE,ANALYSIS_RESULT) values (3275,'TOC','RE',3.2)`); err != nil {
		t.Fatal(err)
	}
	if _, err := testDb.Exec(`Insert into OCD_CHEM_CARB_ANALYSIS_TEST (RUN_ID,ANALYSIS_CODE,METHOD_CODE,ANALYSIS_RESULT) values (3275,'TOT_C','CNS',8.99)`); err != nil {
		t.Fatal(err)
	}
	if _, err := testDb.Exec(`Insert into OCD_CHEM_CARB_ANALYSIS_TEST (RUN_ID,ANALYSIS_CODE,METHOD_CODE,ANALYSIS_RESULT) values (3280,'CaCO3','C',49.76)`); err != nil {
		t.Fatal(err)
	}
	if _, err := testDb.Exec(`Insert into OCD_CHEM_CARB_ANALYSIS_TEST (RUN_ID,ANALYSIS_CODE,METHOD_CODE,ANALYSIS_RESULT) values (3280,'H','CNS',1.54)`); err != nil {
		t.Fatal(err)
	}
	if _, err := testDb.Exec(`Insert into OCD_CHEM_CARB_ANALYSIS_TEST (RUN_ID,ANALYSIS_CODE,METHOD_CODE,ANALYSIS_RESULT) values (3280,'HI','RE',699)`); err != nil {
		t.Fatal(err)
	}
	if _, err := testDb.Exec(`Insert into OCD_CHEM_CARB_ANALYSIS_TEST (RUN_ID,ANALYSIS_CODE,METHOD_CODE,ANALYSIS_RESULT) values (3280,'INOR_C','C',5.97)`); err != nil {
		t.Fatal(err)
	}
	if _, err := testDb.Exec(`Insert into OCD_CHEM_CARB_ANALYSIS_TEST (RUN_ID,ANALYSIS_CODE,METHOD_CODE,ANALYSIS_RESULT) values (3280,'NIT','CNS',0.16)`); err != nil {
		t.Fatal(err)
	}
	if _, err := testDb.Exec(`Insert into OCD_CHEM_CARB_ANALYSIS_TEST (RUN_ID,ANALYSIS_CODE,METHOD_CODE,ANALYSIS_RESULT) values (3280,'OI','RE',45)`); err != nil {
		t.Fatal(err)
	}
	if _, err := testDb.Exec(`Insert into OCD_CHEM_CARB_ANALYSIS_TEST (RUN_ID,ANALYSIS_CODE,METHOD_CODE,ANALYSIS_RESULT) values (3280,'ORG_C','CNS',11.45)`); err != nil {
		t.Fatal(err)
	}
	if _, err := testDb.Exec(`Insert into OCD_CHEM_CARB_ANALYSIS_TEST (RUN_ID,ANALYSIS_CODE,METHOD_CODE,ANALYSIS_RESULT) values (3280,'PI','RE',0.02)`); err != nil {
		t.Fatal(err)
	}
	if _, err := testDb.Exec(`Insert into OCD_CHEM_CARB_ANALYSIS_TEST (RUN_ID,ANALYSIS_CODE,METHOD_CODE,ANALYSIS_RESULT) values (3280,'S1','RE',1.35)`); err != nil {
		t.Fatal(err)
	}
	if _, err := testDb.Exec(`Insert into OCD_CHEM_CARB_ANALYSIS_TEST (RUN_ID,ANALYSIS_CODE,METHOD_CODE,ANALYSIS_RESULT) values (3280,'S2','RE',70.9)`); err != nil {
		t.Fatal(err)
	}
	if _, err := testDb.Exec(`Insert into OCD_CHEM_CARB_ANALYSIS_TEST (RUN_ID,ANALYSIS_CODE,METHOD_CODE,ANALYSIS_RESULT) values (3280,'S3','RE',4.63)`); err != nil {
		t.Fatal(err)
	}
	if _, err := testDb.Exec(`Insert into OCD_CHEM_CARB_ANALYSIS_TEST (RUN_ID,ANALYSIS_CODE,METHOD_CODE,ANALYSIS_RESULT) values (3280,'SUL','CNS',0.62)`); err != nil {
		t.Fatal(err)
	}
	if _, err := testDb.Exec(`Insert into OCD_CHEM_CARB_ANALYSIS_TEST (RUN_ID,ANALYSIS_CODE,METHOD_CODE,ANALYSIS_RESULT) values (3280,'TMX','RE',393)`); err != nil {
		t.Fatal(err)
	}
	if _, err := testDb.Exec(`Insert into OCD_CHEM_CARB_ANALYSIS_TEST (RUN_ID,ANALYSIS_CODE,METHOD_CODE,ANALYSIS_RESULT) values (3280,'TOC','RE',10.14)`); err != nil {
		t.Fatal(err)
	}
	if _, err := testDb.Exec(`Insert into OCD_CHEM_CARB_ANALYSIS_TEST (RUN_ID,ANALYSIS_CODE,METHOD_CODE,ANALYSIS_RESULT) values (3280,'TOT_C','CNS',17.42)`); err != nil {
		t.Fatal(err)
	}
	if _, err := testDb.Exec(`Insert into OCD_CHEM_CARB_ANALYSIS_TEST (RUN_ID,ANALYSIS_CODE,METHOD_CODE,ANALYSIS_RESULT) values (3280,'CaCO3','C',49.76)`); err != nil {
		t.Fatal(err)
	}
	if _, err := testDb.Exec(`Insert into OCD_CHEM_CARB_ANALYSIS_TEST (RUN_ID,ANALYSIS_CODE,METHOD_CODE,ANALYSIS_RESULT) values (3280,'H','CNS',1.54)`); err != nil {
		t.Fatal(err)
	}
	if _, err := testDb.Exec(`Insert into OCD_CHEM_CARB_ANALYSIS_TEST (RUN_ID,ANALYSIS_CODE,METHOD_CODE,ANALYSIS_RESULT) values (3280,'HI','RE',699)`); err != nil {
		t.Fatal(err)
	}
	if _, err := testDb.Exec(`Insert into OCD_CHEM_CARB_ANALYSIS_TEST (RUN_ID,ANALYSIS_CODE,METHOD_CODE,ANALYSIS_RESULT) values (3280,'INOR_C','C',5.97)`); err != nil {
		t.Fatal(err)
	}
	if _, err := testDb.Exec(`Insert into OCD_CHEM_CARB_ANALYSIS_TEST (RUN_ID,ANALYSIS_CODE,METHOD_CODE,ANALYSIS_RESULT) values (3280,'NIT','CNS',0.16)`); err != nil {
		t.Fatal(err)
	}
	if _, err := testDb.Exec(`Insert into OCD_CHEM_CARB_ANALYSIS_TEST (RUN_ID,ANALYSIS_CODE,METHOD_CODE,ANALYSIS_RESULT) values (3280,'OI','RE',45)`); err != nil {
		t.Fatal(err)
	}
	if _, err := testDb.Exec(`Insert into OCD_CHEM_CARB_ANALYSIS_TEST (RUN_ID,ANALYSIS_CODE,METHOD_CODE,ANALYSIS_RESULT) values (3280,'ORG_C','CNS',11.45)`); err != nil {
		t.Fatal(err)
	}
	if _, err := testDb.Exec(`Insert into OCD_CHEM_CARB_ANALYSIS_TEST (RUN_ID,ANALYSIS_CODE,METHOD_CODE,ANALYSIS_RESULT) values (3280,'PI','RE',0.02)`); err != nil {
		t.Fatal(err)
	}
	if _, err := testDb.Exec(`Insert into OCD_CHEM_CARB_ANALYSIS_TEST (RUN_ID,ANALYSIS_CODE,METHOD_CODE,ANALYSIS_RESULT) values (3280,'S1','RE',1.35)`); err != nil {
		t.Fatal(err)
	}
	if _, err := testDb.Exec(`Insert into OCD_CHEM_CARB_ANALYSIS_TEST (RUN_ID,ANALYSIS_CODE,METHOD_CODE,ANALYSIS_RESULT) values (3280,'S2','RE',70.9)`); err != nil {
		t.Fatal(err)
	}
	if _, err := testDb.Exec(`Insert into OCD_CHEM_CARB_ANALYSIS_TEST (RUN_ID,ANALYSIS_CODE,METHOD_CODE,ANALYSIS_RESULT) values (3280,'S3','RE',4.63)`); err != nil {
		t.Fatal(err)
	}
	if _, err := testDb.Exec(`Insert into OCD_CHEM_CARB_ANALYSIS_TEST (RUN_ID,ANALYSIS_CODE,METHOD_CODE,ANALYSIS_RESULT) values (3280,'SUL','CNS',0.62)`); err != nil {
		t.Fatal(err)
	}
	if _, err := testDb.Exec(`Insert into OCD_CHEM_CARB_ANALYSIS_TEST (RUN_ID,ANALYSIS_CODE,METHOD_CODE,ANALYSIS_RESULT) values (3280,'TMX','RE',393)`); err != nil {
		t.Fatal(err)
	}
	if _, err := testDb.Exec(`Insert into OCD_CHEM_CARB_ANALYSIS_TEST (RUN_ID,ANALYSIS_CODE,METHOD_CODE,ANALYSIS_RESULT) values (3280,'TOC','RE',10.14)`); err != nil {
		t.Fatal(err)
	}
	if _, err := testDb.Exec(`Insert into OCD_CHEM_CARB_ANALYSIS_TEST (RUN_ID,ANALYSIS_CODE,METHOD_CODE,ANALYSIS_RESULT) values (3280,'TOT_C','CNS',17.42)`); err != nil {
		t.Fatal(err)
	}
	if _, err := testDb.Exec(`Insert into OCD_CHEM_CARB_ANALYSIS_TEST (RUN_ID,ANALYSIS_CODE,METHOD_CODE,ANALYSIS_RESULT) values (3285,'CaCO3','C',51.2)`); err != nil {
		t.Fatal(err)
	}
	if _, err := testDb.Exec(`Insert into OCD_CHEM_CARB_ANALYSIS_TEST (RUN_ID,ANALYSIS_CODE,METHOD_CODE,ANALYSIS_RESULT) values (3285,'H','CNS',0.85)`); err != nil {
		t.Fatal(err)
	}
	if _, err := testDb.Exec(`Insert into OCD_CHEM_CARB_ANALYSIS_TEST (RUN_ID,ANALYSIS_CODE,METHOD_CODE,ANALYSIS_RESULT) values (3285,'INOR_C','C',6.15)`); err != nil {
		t.Fatal(err)
	}
	if _, err := testDb.Exec(`Insert into OCD_CHEM_CARB_ANALYSIS_TEST (RUN_ID,ANALYSIS_CODE,METHOD_CODE,ANALYSIS_RESULT) values (3285,'NIT','CNS',0.085)`); err != nil {
		t.Fatal(err)
	}
	if _, err := testDb.Exec(`Insert into OCD_CHEM_CARB_ANALYSIS_TEST (RUN_ID,ANALYSIS_CODE,METHOD_CODE,ANALYSIS_RESULT) values (3285,'ORG_C','CNS',5.39)`); err != nil {
		t.Fatal(err)
	}
	if _, err := testDb.Exec(`Insert into OCD_CHEM_CARB_ANALYSIS_TEST (RUN_ID,ANALYSIS_CODE,METHOD_CODE,ANALYSIS_RESULT) values (3285,'TOT_C','CNS',11.54)`); err != nil {
		t.Fatal(err)
	}
	if _, err := testDb.Exec(`Insert into OCD_CHEM_CARB_ANALYSIS_TEST (RUN_ID,ANALYSIS_CODE,METHOD_CODE,ANALYSIS_RESULT) values (3285,'CaCO3','C',51.2)`); err != nil {
		t.Fatal(err)
	}
	if _, err := testDb.Exec(`Insert into OCD_CHEM_CARB_ANALYSIS_TEST (RUN_ID,ANALYSIS_CODE,METHOD_CODE,ANALYSIS_RESULT) values (3285,'H','CNS',0.85)`); err != nil {
		t.Fatal(err)
	}
	if _, err := testDb.Exec(`Insert into OCD_CHEM_CARB_ANALYSIS_TEST (RUN_ID,ANALYSIS_CODE,METHOD_CODE,ANALYSIS_RESULT) values (3285,'INOR_C','C',6.15)`); err != nil {
		t.Fatal(err)
	}
	if _, err := testDb.Exec(`Insert into OCD_CHEM_CARB_ANALYSIS_TEST (RUN_ID,ANALYSIS_CODE,METHOD_CODE,ANALYSIS_RESULT) values (3285,'NIT','CNS',0.085)`); err != nil {
		t.Fatal(err)
	}
	if _, err := testDb.Exec(`Insert into OCD_CHEM_CARB_ANALYSIS_TEST (RUN_ID,ANALYSIS_CODE,METHOD_CODE,ANALYSIS_RESULT) values (3285,'ORG_C','CNS',5.39)`); err != nil {
		t.Fatal(err)
	}
	if _, err := testDb.Exec(`Insert into OCD_CHEM_CARB_ANALYSIS_TEST (RUN_ID,ANALYSIS_CODE,METHOD_CODE,ANALYSIS_RESULT) values (3285,'TOT_C','CNS',11.54)`); err != nil {
		t.Fatal(err)
	}
	if _, err := testDb.Exec(`Insert into OCD_CHEM_CARB_ANALYSIS_TEST (RUN_ID,ANALYSIS_CODE,METHOD_CODE,ANALYSIS_RESULT) values (3290,'CaCO3','C',88.42)`); err != nil {
		t.Fatal(err)
	}
	if _, err := testDb.Exec(`Insert into OCD_CHEM_CARB_ANALYSIS_TEST (RUN_ID,ANALYSIS_CODE,METHOD_CODE,ANALYSIS_RESULT) values (3290,'H','CNS',0.007)`); err != nil {
		t.Fatal(err)
	}
	if _, err := testDb.Exec(`Insert into OCD_CHEM_CARB_ANALYSIS_TEST (RUN_ID,ANALYSIS_CODE,METHOD_CODE,ANALYSIS_RESULT) values (3290,'INOR_C','C',10.61)`); err != nil {
		t.Fatal(err)
	}
	if _, err := testDb.Exec(`Insert into OCD_CHEM_CARB_ANALYSIS_TEST (RUN_ID,ANALYSIS_CODE,METHOD_CODE,ANALYSIS_RESULT) values (3290,'ORG_C','CNS',0)`); err != nil {
		t.Fatal(err)
	}
	if _, err := testDb.Exec(`Insert into OCD_CHEM_CARB_ANALYSIS_TEST (RUN_ID,ANALYSIS_CODE,METHOD_CODE,ANALYSIS_RESULT) values (3290,'TOT_C','CNS',10.43)`); err != nil {
		t.Fatal(err)
	}
	if _, err := testDb.Exec(`Insert into OCD_CHEM_CARB_ANALYSIS_TEST (RUN_ID,ANALYSIS_CODE,METHOD_CODE,ANALYSIS_RESULT) values (3290,'CaCO3','C',88.42)`); err != nil {
		t.Fatal(err)
	}
	if _, err := testDb.Exec(`Insert into OCD_CHEM_CARB_ANALYSIS_TEST (RUN_ID,ANALYSIS_CODE,METHOD_CODE,ANALYSIS_RESULT) values (3290,'H','CNS',0.007)`); err != nil {
		t.Fatal(err)
	}
	if _, err := testDb.Exec(`Insert into OCD_CHEM_CARB_ANALYSIS_TEST (RUN_ID,ANALYSIS_CODE,METHOD_CODE,ANALYSIS_RESULT) values (3290,'INOR_C','C',10.61)`); err != nil {
		t.Fatal(err)
	}
	if _, err := testDb.Exec(`Insert into OCD_CHEM_CARB_ANALYSIS_TEST (RUN_ID,ANALYSIS_CODE,METHOD_CODE,ANALYSIS_RESULT) values (3290,'ORG_C','CNS',0)`); err != nil {
		t.Fatal(err)
	}
	if _, err := testDb.Exec(`Insert into OCD_CHEM_CARB_ANALYSIS_TEST (RUN_ID,ANALYSIS_CODE,METHOD_CODE,ANALYSIS_RESULT) values (3290,'TOT_C','CNS',10.43)`); err != nil {
		t.Fatal(err)
	}
	if _, err := testDb.Exec(`Insert into OCD_CHEM_CARB_ANALYSIS_TEST (RUN_ID,ANALYSIS_CODE,METHOD_CODE,ANALYSIS_RESULT) values (8171,'HI','RE',451)`); err != nil {
		t.Fatal(err)
	}
	if _, err := testDb.Exec(`Insert into OCD_CHEM_CARB_ANALYSIS_TEST (RUN_ID,ANALYSIS_CODE,METHOD_CODE,ANALYSIS_RESULT) values (8171,'OI','RE',70)`); err != nil {
		t.Fatal(err)
	}
	if _, err := testDb.Exec(`Insert into OCD_CHEM_CARB_ANALYSIS_TEST (RUN_ID,ANALYSIS_CODE,METHOD_CODE,ANALYSIS_RESULT) values (8171,'PI','RE',0.01)`); err != nil {
		t.Fatal(err)
	}
	if _, err := testDb.Exec(`Insert into OCD_CHEM_CARB_ANALYSIS_TEST (RUN_ID,ANALYSIS_CODE,METHOD_CODE,ANALYSIS_RESULT) values (8171,'S1','RE',0.26)`); err != nil {
		t.Fatal(err)
	}
	if _, err := testDb.Exec(`Insert into OCD_CHEM_CARB_ANALYSIS_TEST (RUN_ID,ANALYSIS_CODE,METHOD_CODE,ANALYSIS_RESULT) values (8171,'S2','RE',20.6)`); err != nil {
		t.Fatal(err)
	}
	if _, err := testDb.Exec(`Insert into OCD_CHEM_CARB_ANALYSIS_TEST (RUN_ID,ANALYSIS_CODE,METHOD_CODE,ANALYSIS_RESULT) values (8171,'S3','RE',3.24)`); err != nil {
		t.Fatal(err)
	}
	if _, err := testDb.Exec(`Insert into OCD_CHEM_CARB_ANALYSIS_TEST (RUN_ID,ANALYSIS_CODE,METHOD_CODE,ANALYSIS_RESULT) values (8171,'TMX','RE',407)`); err != nil {
		t.Fatal(err)
	}
	if _, err := testDb.Exec(`Insert into OCD_CHEM_CARB_ANALYSIS_TEST (RUN_ID,ANALYSIS_CODE,METHOD_CODE,ANALYSIS_RESULT) values (8171,'TOC','RE',4.57)`); err != nil {
		t.Fatal(err)
	}
	if _, err := testDb.Exec(`Insert into OCD_CHEM_CARB_ANALYSIS_TEST (RUN_ID,ANALYSIS_CODE,METHOD_CODE,ANALYSIS_RESULT) values (8176,'OI','RE',1016)`); err != nil {
		t.Fatal(err)
	}
	if _, err := testDb.Exec(`Insert into OCD_CHEM_CARB_ANALYSIS_TEST (RUN_ID,ANALYSIS_CODE,METHOD_CODE,ANALYSIS_RESULT) values (8176,'PI','RE',0.04)`); err != nil {
		t.Fatal(err)
	}
	if _, err := testDb.Exec(`Insert into OCD_CHEM_CARB_ANALYSIS_TEST (RUN_ID,ANALYSIS_CODE,METHOD_CODE,ANALYSIS_RESULT) values (8176,'S1','RE',0.03)`); err != nil {
		t.Fatal(err)
	}
	if _, err := testDb.Exec(`Insert into OCD_CHEM_CARB_ANALYSIS_TEST (RUN_ID,ANALYSIS_CODE,METHOD_CODE,ANALYSIS_RESULT) values (8176,'S2','RE',0.8)`); err != nil {
		t.Fatal(err)
	}
	if _, err := testDb.Exec(`Insert into OCD_CHEM_CARB_ANALYSIS_TEST (RUN_ID,ANALYSIS_CODE,METHOD_CODE,ANALYSIS_RESULT) values (8176,'S3','RE',0.61)`); err != nil {
		t.Fatal(err)
	}
	if _, err := testDb.Exec(`Insert into OCD_CHEM_CARB_ANALYSIS_TEST (RUN_ID,ANALYSIS_CODE,METHOD_CODE,ANALYSIS_RESULT) values (8176,'TMX','RE',445)`); err != nil {
		t.Fatal(err)
	}
	if _, err := testDb.Exec(`Insert into OCD_CHEM_CARB_ANALYSIS_TEST (RUN_ID,ANALYSIS_CODE,METHOD_CODE,ANALYSIS_RESULT) values (8176,'TOC','RE',0.06)`); err != nil {
		t.Fatal(err)
	}
	if _, err := testDb.Exec(`Insert into OCD_CHEM_CARB_SAMPLE_TEST (RUN_ID,SAMPLE_ID,LOCATION) values (42285,114942,'SHI')`); err != nil {
		t.Fatal(err)
	}
	if _, err := testDb.Exec(`Insert into OCD_CHEM_CARB_SAMPLE_TEST (RUN_ID,SAMPLE_ID,LOCATION) values (42290,114943,'SHI')`); err != nil {
		t.Fatal(err)
	}
	if _, err := testDb.Exec(`Insert into OCD_CHEM_CARB_SAMPLE_TEST (RUN_ID,SAMPLE_ID,LOCATION) values (42295,114944,'SHI')`); err != nil {
		t.Fatal(err)
	}
	if _, err := testDb.Exec(`Insert into OCD_CHEM_CARB_SAMPLE_TEST (RUN_ID,SAMPLE_ID,LOCATION) values (3295,25277,'SHI')`); err != nil {
		t.Fatal(err)
	}
	if _, err := testDb.Exec(`Insert into OCD_CHEM_CARB_SAMPLE_TEST (RUN_ID,SAMPLE_ID,LOCATION) values (3300,25263,'SHI')`); err != nil {
		t.Fatal(err)
	}
	if _, err := testDb.Exec(`Insert into OCD_CHEM_CARB_SAMPLE_TEST (RUN_ID,SAMPLE_ID,LOCATION) values (3240,25061,'SHI')`); err != nil {
		t.Fatal(err)
	}
	if _, err := testDb.Exec(`Insert into OCD_CHEM_CARB_SAMPLE_TEST (RUN_ID,SAMPLE_ID,LOCATION) values (3245,25063,'SHI')`); err != nil {
		t.Fatal(err)
	}
	if _, err := testDb.Exec(`Insert into OCD_CHEM_CARB_SAMPLE_TEST (RUN_ID,SAMPLE_ID,LOCATION) values (3250,25107,'SHI')`); err != nil {
		t.Fatal(err)
	}
	if _, err := testDb.Exec(`Insert into OCD_CHEM_CARB_SAMPLE_TEST (RUN_ID,SAMPLE_ID,LOCATION) values (3250,25107,'SHI')`); err != nil {
		t.Fatal(err)
	}
	if _, err := testDb.Exec(`Insert into OCD_CHEM_CARB_SAMPLE_TEST (RUN_ID,SAMPLE_ID,LOCATION) values (3255,25106,'SHI')`); err != nil {
		t.Fatal(err)
	}
	if _, err := testDb.Exec(`Insert into OCD_CHEM_CARB_SAMPLE_TEST (RUN_ID,SAMPLE_ID,LOCATION) values (3255,25106,'SHI')`); err != nil {
		t.Fatal(err)
	}
	if _, err := testDb.Exec(`Insert into OCD_CHEM_CARB_SAMPLE_TEST (RUN_ID,SAMPLE_ID,LOCATION) values (3260,25105,'SHI')`); err != nil {
		t.Fatal(err)
	}
	if _, err := testDb.Exec(`Insert into OCD_CHEM_CARB_SAMPLE_TEST (RUN_ID,SAMPLE_ID,LOCATION) values (3260,25105,'SHI')`); err != nil {
		t.Fatal(err)
	}
	if _, err := testDb.Exec(`Insert into OCD_CHEM_CARB_SAMPLE_TEST (RUN_ID,SAMPLE_ID,LOCATION) values (3265,25102,'SHI')`); err != nil {
		t.Fatal(err)
	}
	if _, err := testDb.Exec(`Insert into OCD_CHEM_CARB_SAMPLE_TEST (RUN_ID,SAMPLE_ID,LOCATION) values (3265,25102,'SHI')`); err != nil {
		t.Fatal(err)
	}
	if _, err := testDb.Exec(`Insert into OCD_CHEM_CARB_SAMPLE_TEST (RUN_ID,SAMPLE_ID,LOCATION) values (3270,25104,'SHI')`); err != nil {
		t.Fatal(err)
	}
	if _, err := testDb.Exec(`Insert into OCD_CHEM_CARB_SAMPLE_TEST (RUN_ID,SAMPLE_ID,LOCATION) values (3270,25104,'SHI')`); err != nil {
		t.Fatal(err)
	}
	if _, err := testDb.Exec(`Insert into OCD_CHEM_CARB_SAMPLE_TEST (RUN_ID,SAMPLE_ID,LOCATION) values (3275,25103,'SHI')`); err != nil {
		t.Fatal(err)
	}
	if _, err := testDb.Exec(`Insert into OCD_CHEM_CARB_SAMPLE_TEST (RUN_ID,SAMPLE_ID,LOCATION) values (3275,25103,'SHI')`); err != nil {
		t.Fatal(err)
	}
	if _, err := testDb.Exec(`Insert into OCD_CHEM_CARB_SAMPLE_TEST (RUN_ID,SAMPLE_ID,LOCATION) values (3280,25101,'SHI')`); err != nil {
		t.Fatal(err)
	}
	if _, err := testDb.Exec(`Insert into OCD_CHEM_CARB_SAMPLE_TEST (RUN_ID,SAMPLE_ID,LOCATION) values (3280,25101,'SHI')`); err != nil {
		t.Fatal(err)
	}
	if _, err := testDb.Exec(`Insert into OCD_CHEM_CARB_SAMPLE_TEST (RUN_ID,SAMPLE_ID,LOCATION) values (3285,25100,'SHI')`); err != nil {
		t.Fatal(err)
	}
	if _, err := testDb.Exec(`Insert into OCD_CHEM_CARB_SAMPLE_TEST (RUN_ID,SAMPLE_ID,LOCATION) values (3285,25100,'SHI')`); err != nil {
		t.Fatal(err)
	}
	if _, err := testDb.Exec(`Insert into OCD_CHEM_CARB_SAMPLE_TEST (RUN_ID,SAMPLE_ID,LOCATION) values (3290,25099,'SHI')`); err != nil {
		t.Fatal(err)
	}
	if _, err := testDb.Exec(`Insert into OCD_CHEM_CARB_SAMPLE_TEST (RUN_ID,SAMPLE_ID,LOCATION) values (3290,25099,'SHI')`); err != nil {
		t.Fatal(err)
	}
	if _, err := testDb.Exec(`Insert into OCD_CHEM_CARB_SAMPLE_TEST (RUN_ID,SAMPLE_ID,LOCATION) values (8171,227155,'SHI')`); err != nil {
		t.Fatal(err)
	}
	if _, err := testDb.Exec(`Insert into OCD_CHEM_CARB_SAMPLE_TEST (RUN_ID,SAMPLE_ID,LOCATION) values (8176,227154,'SHI')`); err != nil {
		t.Fatal(err)
	}
	if _, err := testDb.Exec(`Insert into OCD_HOLE_TEST (LEG,SITE,HOLE) values (171,1049,'B')`); err != nil {
		t.Fatal(err)
	}
	if _, err := testDb.Exec(`Insert into OCD_SAMPLE_TEST (SAMPLE_ID,LOCATION,SAM_SECTION_ID,TOP_INTERVAL,BOTTOM_INTERVAL) values (25099,'SHI',42830,0.21,0.22)`); err != nil {
		t.Fatal(err)
	}
	if _, err := testDb.Exec(`Insert into OCD_SAMPLE_TEST (SAMPLE_ID,LOCATION,SAM_SECTION_ID,TOP_INTERVAL,BOTTOM_INTERVAL) values (25100,'SHI',42830,0.19,0.21)`); err != nil {
		t.Fatal(err)
	}
	if _, err := testDb.Exec(`Insert into OCD_SAMPLE_TEST (SAMPLE_ID,LOCATION,SAM_SECTION_ID,TOP_INTERVAL,BOTTOM_INTERVAL) values (25101,'SHI',42830,0.175,0.19)`); err != nil {
		t.Fatal(err)
	}
	if _, err := testDb.Exec(`Insert into OCD_SAMPLE_TEST (SAMPLE_ID,LOCATION,SAM_SECTION_ID,TOP_INTERVAL,BOTTOM_INTERVAL) values (25102,'SHI',42830,0.125,0.22)`); err != nil {
		t.Fatal(err)
	}
	if _, err := testDb.Exec(`Insert into OCD_SAMPLE_TEST (SAMPLE_ID,LOCATION,SAM_SECTION_ID,TOP_INTERVAL,BOTTOM_INTERVAL) values (25103,'SHI',42830,0.16,0.175)`); err != nil {
		t.Fatal(err)
	}
	if _, err := testDb.Exec(`Insert into OCD_SAMPLE_TEST (SAMPLE_ID,LOCATION,SAM_SECTION_ID,TOP_INTERVAL,BOTTOM_INTERVAL) values (25104,'SHI',42830,0.135,0.16)`); err != nil {
		t.Fatal(err)
	}
	if _, err := testDb.Exec(`Insert into OCD_SAMPLE_TEST (SAMPLE_ID,LOCATION,SAM_SECTION_ID,TOP_INTERVAL,BOTTOM_INTERVAL) values (25105,'SHI',42830,0.125,0.135)`); err != nil {
		t.Fatal(err)
	}
	if _, err := testDb.Exec(`Insert into OCD_SAMPLE_TEST (SAMPLE_ID,LOCATION,SAM_SECTION_ID,TOP_INTERVAL,BOTTOM_INTERVAL) values (25106,'SHI',42830,0.085,0.125)`); err != nil {
		t.Fatal(err)
	}
	if _, err := testDb.Exec(`Insert into OCD_SAMPLE_TEST (SAMPLE_ID,LOCATION,SAM_SECTION_ID,TOP_INTERVAL,BOTTOM_INTERVAL) values (25107,'SHI',42830,0.07,0.085)`); err != nil {
		t.Fatal(err)
	}
	if _, err := testDb.Exec(`Insert into OCD_SAMPLE_TEST (SAMPLE_ID,LOCATION,SAM_SECTION_ID,TOP_INTERVAL,BOTTOM_INTERVAL) values (227154,'SHI',42830,0.205,0.22)`); err != nil {
		t.Fatal(err)
	}
	if _, err := testDb.Exec(`Insert into OCD_SAMPLE_TEST (SAMPLE_ID,LOCATION,SAM_SECTION_ID,TOP_INTERVAL,BOTTOM_INTERVAL) values (227155,'SHI',42830,0.19,0.205)`); err != nil {
		t.Fatal(err)
	}
	if _, err := testDb.Exec(`Insert into OCD_SECTION_TEST (SECTION_ID,SECTION_NUMBER,SECTION_TYPE,LEG,SITE,HOLE,CORE,CORE_TYPE) values (42730,1,'S',171,1049,'B',8,'H')`); err != nil {
		t.Fatal(err)
	}
	if _, err := testDb.Exec(`Insert into OCD_SECTION_TEST (SECTION_ID,SECTION_NUMBER,SECTION_TYPE,LEG,SITE,HOLE,CORE,CORE_TYPE) values (42730,1,'S',171,1049,'B',8,'H')`); err != nil {
		t.Fatal(err)
	}
	if _, err := testDb.Exec(`Insert into OCD_SECTION_TEST (SECTION_ID,SECTION_NUMBER,SECTION_TYPE,LEG,SITE,HOLE,CORE,CORE_TYPE) values (42830,3,'C',171,1049,'B',11,'X')`); err != nil {
		t.Fatal(err)
	}
	if _, err := testDb.Exec(`Insert into OCD_SECTION_TEST (SECTION_ID,SECTION_NUMBER,SECTION_TYPE,LEG,SITE,HOLE,CORE,CORE_TYPE) values (42740,3,'S',171,1049,'B',8,'H')`); err != nil {
		t.Fatal(err)
	}
	if _, err := testDb.Exec(`Insert into OCD_SECTION_TEST (SECTION_ID,SECTION_NUMBER,SECTION_TYPE,LEG,SITE,HOLE,CORE,CORE_TYPE) values (42830,3,'C',171,1049,'B',11,'X')`); err != nil {
		t.Fatal(err)
	}
	if _, err := testDb.Exec(`Insert into OCD_SECTION_TEST (SECTION_ID,SECTION_NUMBER,SECTION_TYPE,LEG,SITE,HOLE,CORE,CORE_TYPE) values (42740,3,'S',171,1049,'B',8,'H')`); err != nil {
		t.Fatal(err)
	}
	if _, err := testDb.Exec(`Insert into OCD_SECTION_TEST (SECTION_ID,SECTION_NUMBER,SECTION_TYPE,LEG,SITE,HOLE,CORE,CORE_TYPE) values (42830,3,'C',171,1049,'B',11,'X')`); err != nil {
		t.Fatal(err)
	}
	if _, err := testDb.Exec(`Insert into OCD_SECTION_TEST (SECTION_ID,SECTION_NUMBER,SECTION_TYPE,LEG,SITE,HOLE,CORE,CORE_TYPE) values (42740,3,'S',171,1049,'B',8,'H')`); err != nil {
		t.Fatal(err)
	}
	if _, err := testDb.Exec(`Insert into OCD_SECTION_TEST (SECTION_ID,SECTION_NUMBER,SECTION_TYPE,LEG,SITE,HOLE,CORE,CORE_TYPE) values (42830,3,'C',171,1049,'B',11,'X')`); err != nil {
		t.Fatal(err)
	}
	if _, err := testDb.Exec(`Insert into OCD_SECTION_TEST (SECTION_ID,SECTION_NUMBER,SECTION_TYPE,LEG,SITE,HOLE,CORE,CORE_TYPE) values (42745,4,'S',171,1049,'B',8,'H')`); err != nil {
		t.Fatal(err)
	}
	if _, err := testDb.Exec(`Insert into OCD_SECTION_TEST (SECTION_ID,SECTION_NUMBER,SECTION_TYPE,LEG,SITE,HOLE,CORE,CORE_TYPE) values (42830,3,'C',171,1049,'B',11,'X')`); err != nil {
		t.Fatal(err)
	}
	if _, err := testDb.Exec(`Insert into OCD_SECTION_TEST (SECTION_ID,SECTION_NUMBER,SECTION_TYPE,LEG,SITE,HOLE,CORE,CORE_TYPE) values (42745,4,'S',171,1049,'B',8,'H')`); err != nil {
		t.Fatal(err)
	}
	if _, err := testDb.Exec(`Insert into OCD_SECTION_TEST (SECTION_ID,SECTION_NUMBER,SECTION_TYPE,LEG,SITE,HOLE,CORE,CORE_TYPE) values (42830,3,'C',171,1049,'B',11,'X')`); err != nil {
		t.Fatal(err)
	}
	if _, err := testDb.Exec(`Insert into OCD_SECTION_TEST (SECTION_ID,SECTION_NUMBER,SECTION_TYPE,LEG,SITE,HOLE,CORE,CORE_TYPE) values (42745,4,'S',171,1049,'B',8,'H')`); err != nil {
		t.Fatal(err)
	}
	if _, err := testDb.Exec(`Insert into OCD_SECTION_TEST (SECTION_ID,SECTION_NUMBER,SECTION_TYPE,LEG,SITE,HOLE,CORE,CORE_TYPE) values (42830,3,'C',171,1049,'B',11,'X')`); err != nil {
		t.Fatal(err)
	}
	if _, err := testDb.Exec(`Insert into OCD_SECTION_TEST (SECTION_ID,SECTION_NUMBER,SECTION_TYPE,LEG,SITE,HOLE,CORE,CORE_TYPE) values (42745,4,'S',171,1049,'B',8,'H')`); err != nil {
		t.Fatal(err)
	}
	if _, err := testDb.Exec(`Insert into OCD_SECTION_TEST (SECTION_ID,SECTION_NUMBER,SECTION_TYPE,LEG,SITE,HOLE,CORE,CORE_TYPE) values (42830,3,'C',171,1049,'B',11,'X')`); err != nil {
		t.Fatal(err)
	}
	if _, err := testDb.Exec(`Insert into OCD_SECTION_TEST (SECTION_ID,SECTION_NUMBER,SECTION_TYPE,LEG,SITE,HOLE,CORE,CORE_TYPE) values (42745,4,'S',171,1049,'B',8,'H')`); err != nil {
		t.Fatal(err)
	}
	if _, err := testDb.Exec(`Insert into OCD_SECTION_TEST (SECTION_ID,SECTION_NUMBER,SECTION_TYPE,LEG,SITE,HOLE,CORE,CORE_TYPE) values (42830,3,'C',171,1049,'B',11,'X')`); err != nil {
		t.Fatal(err)
	}
	if _, err := testDb.Exec(`Insert into OCD_SECTION_TEST (SECTION_ID,SECTION_NUMBER,SECTION_TYPE,LEG,SITE,HOLE,CORE,CORE_TYPE) values (42745,4,'S',171,1049,'B',8,'H')`); err != nil {
		t.Fatal(err)
	}
	if _, err := testDb.Exec(`Insert into OCD_SECTION_TEST (SECTION_ID,SECTION_NUMBER,SECTION_TYPE,LEG,SITE,HOLE,CORE,CORE_TYPE) values (42735,2,'S',171,1049,'B',8,'H')`); err != nil {
		t.Fatal(err)
	}
	if _, err := testDb.Exec(`Insert into OCD_SECTION_TEST (SECTION_ID,SECTION_NUMBER,SECTION_TYPE,LEG,SITE,HOLE,CORE,CORE_TYPE) values (42735,2,'S',171,1049,'B',8,'H')`); err != nil {
		t.Fatal(err)
	}
	if _, err := testDb.Exec(`Insert into OCD_SECTION_TEST (SECTION_ID,SECTION_NUMBER,SECTION_TYPE,LEG,SITE,HOLE,CORE,CORE_TYPE) values (42820,1,'S',171,1049,'B',11,'X')`); err != nil {
		t.Fatal(err)
	}
	if _, err := testDb.Exec(`Insert into OCD_SECTION_TEST (SECTION_ID,SECTION_NUMBER,SECTION_TYPE,LEG,SITE,HOLE,CORE,CORE_TYPE) values (42820,1,'S',171,1049,'B',11,'X')`); err != nil {
		t.Fatal(err)
	}
	if _, err := testDb.Exec(`Insert into OCD_SECTION_TEST (SECTION_ID,SECTION_NUMBER,SECTION_TYPE,LEG,SITE,HOLE,CORE,CORE_TYPE) values (42820,1,'S',171,1049,'B',11,'X')`); err != nil {
		t.Fatal(err)
	}
	if _, err := testDb.Exec(`Insert into OCD_SECTION_TEST (SECTION_ID,SECTION_NUMBER,SECTION_TYPE,LEG,SITE,HOLE,CORE,CORE_TYPE) values (42830,3,'C',171,1049,'B',11,'X')`); err != nil {
		t.Fatal(err)
	}
	if _, err := testDb.Exec(`Insert into OCD_SECTION_TEST (SECTION_ID,SECTION_NUMBER,SECTION_TYPE,LEG,SITE,HOLE,CORE,CORE_TYPE) values (42830,3,'C',171,1049,'B',11,'X')`); err != nil {
		t.Fatal(err)
	}

	if _, err := testDb.Exec(`CREATE VIEW ocd_chem_carb_test_v AS
SELECT
    x.leg, x.site, x.hole
  , x.core, x.core_type
  , x.section_number, x.section_type
  , s.top_interval*100.0 top_cm
  , s.bottom_interval*100.0 bot_cm
  , AVG(DECODE(cca.analysis_code,'INOR_C',cca.analysis_result)) INOR_C_wt_pct
  , AVG(DECODE(cca.analysis_code,'CaCO3', cca.analysis_result)) CaCO3_wt_pct
  , AVG(DECODE(cca.analysis_code,'TOT_C', cca.analysis_result)) TOT_C_wt_pct
  , AVG(DECODE(cca.analysis_code,'ORG_C', cca.analysis_result)) ORG_C_wt_pct
  , AVG(DECODE(cca.analysis_code,'NIT',   cca.analysis_result)) NIT_wt_pct
  , AVG(DECODE(cca.analysis_code,'SUL',   cca.analysis_result)) SUL_wt_pct
  , AVG(DECODE(cca.analysis_code,'H',     cca.analysis_result)) H_wt_pct
FROM
    ocd_hole_test h, ocd_section_test x, ocd_sample_test s
  , ocd_chem_carb_sample_test ccs, ocd_chem_carb_analysis_test cca
WHERE
        h.leg = x.leg
    AND h.site = x.site
    AND h.hole = x.hole
    AND x.section_id = s.sam_section_id
    AND s.sample_id = ccs.sample_id
    AND s.location = ccs.location
    AND ccs.run_id = cca.run_id
GROUP BY x.leg, x.site, x.hole, x.core, x.core_type, x.section_number, x.section_type, s.top_interval, s.bottom_interval
ORDER BY x.leg, x.site, x.hole, x.core, x.core_type, x.section_number, s.top_interval
`); err != nil {
		t.Fatal(err)
	}

	//testDb.Exec(`DROP PUBLIC SYNONYM fils_ocd_chem_carb_test`)
	testDb.Exec(`DROP PUBLIC SYNONYM ocd_chem_carb_test`)

	if _, err := testDb.Exec(`GRANT SELECT ON ocd_chem_carb_test_v TO PUBLIC`); err != nil {
		t.Fatal(err)
	}
	// if _, err := testDb.Exec(`CREATE PUBLIC SYNONYM ocd_chem_carb_test FOR fils.ocd_chem_carb_test_v`); err != nil {
	// 	t.Fatal(err)
	// }
	if _, err := testDb.Exec(`CREATE PUBLIC SYNONYM ocd_chem_carb_test FOR fils.ocd_chem_carb_test_v`); err != nil {
		t.Fatal(err)
	}

	if _, err := testDb.Exec(`DROP TABLE ocd_chem_carb_test_table`); err != nil {
		t.Log(err)
	}

	if _, err := testDb.Exec(`CREATE TABLE ocd_chem_carb_test_table AS SELECT * FROM fils.ocd_chem_carb_test_v`); err != nil {
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
	   FROM test_janus_v
	WHERE
	       leg = 171
	    AND site = 1049
	   AND hole = 'B'
	ORDER BY leg, site, hole, core, section_number, top_cm
	`

	qry3 := `SELECT
            x.leg, x.site, x.hole
          , x.core, x.core_type
          , x.section_number, x.section_type
          , s.top_interval*100.0 top_cm
          , s.bottom_interval*100.0 bot_cm
          , AVG(DECODE(cca.analysis_code,'INOR_C',cca.analysis_result)) INOR_C_wt_pct
          , AVG(DECODE(cca.analysis_code,'CaCO3', cca.analysis_result)) CaCO3_wt_pct
          , AVG(DECODE(cca.analysis_code,'TOT_C', cca.analysis_result)) TOT_C_wt_pct
          , AVG(DECODE(cca.analysis_code,'ORG_C', cca.analysis_result)) ORG_C_wt_pct
          , AVG(DECODE(cca.analysis_code,'NIT',   cca.analysis_result)) NIT_wt_pct
          , AVG(DECODE(cca.analysis_code,'SUL',   cca.analysis_result)) SUL_wt_pct
          , AVG(DECODE(cca.analysis_code,'H',     cca.analysis_result)) H_wt_pct
        FROM
            ocd_hole_test h, ocd_section_test x, ocd_sample_test s
          , ocd_chem_carb_sample_test ccs, ocd_chem_carb_analysis_test cca
        WHERE
                h.leg = x.leg
            AND h.site = x.site
            AND h.hole = x.hole
            AND x.section_id = s.sam_section_id
            AND s.sample_id = ccs.sample_id
            AND s.location = ccs.location
            AND ccs.run_id = cca.run_id
            AND x.leg = 171
            AND x.site = 1049
            AND x.hole = upper('B')
        GROUP BY x.leg, x.site, x.hole, x.core, x.core_type, x.section_number, x.section_type, s.top_interval, s.bottom_interval
        ORDER BY x.leg, x.site, x.hole, x.core, x.core_type, x.section_number, s.top_interval
`

	qry4 := `SELECT
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
	   FROM  ocd_chem_carb_test
	WHERE
	       leg = 171
	    AND site = 1049
	   AND hole = 'B'
	ORDER BY leg, site, hole, core, section_number, top_cm
	`

	qry5 := `SELECT
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
	   FROM  ocd_chem_carb_test_v
	WHERE
	       leg = 171
	    AND site = 1049
	   AND hole = 'B'
	ORDER BY leg, site, hole, core, section_number, top_cm
	`

	qry6 := `SELECT
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
	   FROM  ocd_chem_carb_test_table
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

	log.Printf("Run query 1\n")

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

	log.Printf("Run query 2\n")

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

		log.Printf("Results: %v %v %v %v %v %v %v %v %v %v %v %v %v %v %v %v", Leg, Site, Hole, Core, Core_type, Section_number, Section_type, Top_cm, Bot_cm, Inor_c_wt_pct, Caco3_wt_pct, Tot_c_wt_pct, Org_c_wt_pct, Nit_wt_pct, Sul_wt_pct, H_wt_pct)

	}
	if err := rows2.Err(); err != nil {
		t.Error(err)
	}

	log.Printf("Run query 3\n")

	rows3, err := testDb.Query(qry3)
	if err != nil {
		t.Errorf(`Error with "%s": %s`, qry3, err)
		return
	}
	defer rows3.Close()

	iii := 0
	for rows3.Next() {
		iii++
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

		if err := rows3.Scan(&Leg, &Site, &Hole, &Core, &Core_type, &Section_number, &Section_type, &Top_cm, &Bot_cm, &Inor_c_wt_pct, &Caco3_wt_pct, &Tot_c_wt_pct, &Org_c_wt_pct, &Nit_wt_pct, &Sul_wt_pct, &H_wt_pct); err != nil {
			t.Fatalf("scan %d. record: %v", iii, err)
		}

		log.Printf("Results: %v %v %v %v %v %v %v %v %v %v %v %v %v %v %v %v", Leg, Site, Hole, Core, Core_type, Section_number, Section_type, Top_cm, Bot_cm, Inor_c_wt_pct, Caco3_wt_pct, Tot_c_wt_pct, Org_c_wt_pct, Nit_wt_pct, Sul_wt_pct, H_wt_pct)

	}
	if err := rows3.Err(); err != nil {
		t.Error(err)
	}

	log.Printf("Run query 4\n")

	rows4, err := testDb.Query(qry4)
	if err != nil {
		t.Errorf(`Error with "%s": %s`, qry4, err)
		return
	}
	defer rows4.Close()

	iiii := 0
	for rows4.Next() {
		iiii++
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

		if err := rows4.Scan(&Leg, &Site, &Hole, &Core, &Core_type, &Section_number, &Section_type, &Top_cm, &Bot_cm, &Inor_c_wt_pct, &Caco3_wt_pct, &Tot_c_wt_pct, &Org_c_wt_pct, &Nit_wt_pct, &Sul_wt_pct, &H_wt_pct); err != nil {
			t.Fatalf("scan %d. record: %v", iiii, err)
		}

		log.Printf("Results: %v %v %v %v %v %v %v %v %v %v %v %v %v %v %v %v", Leg, Site, Hole, Core, Core_type, Section_number, Section_type, Top_cm, Bot_cm, Inor_c_wt_pct, Caco3_wt_pct, Tot_c_wt_pct, Org_c_wt_pct, Nit_wt_pct, Sul_wt_pct, H_wt_pct)

	}
	if err := rows4.Err(); err != nil {
		t.Error(err)
	}

	log.Printf("Run query 5\n")

	rows5, err := testDb.Query(qry5)
	if err != nil {
		t.Errorf(`Error with "%s": %s`, qry5, err)
		return
	}
	defer rows5.Close()

	v := 0
	for rows5.Next() {
		v++
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

		if err := rows5.Scan(&Leg, &Site, &Hole, &Core, &Core_type, &Section_number, &Section_type, &Top_cm, &Bot_cm, &Inor_c_wt_pct, &Caco3_wt_pct, &Tot_c_wt_pct, &Org_c_wt_pct, &Nit_wt_pct, &Sul_wt_pct, &H_wt_pct); err != nil {
			t.Fatalf("scan %d. record: %v", v, err)
		}

		log.Printf("Results: %v %v %v %v %v %v %v %v %v %v %v %v %v %v %v %v", Leg, Site, Hole, Core, Core_type, Section_number, Section_type, Top_cm, Bot_cm, Inor_c_wt_pct, Caco3_wt_pct, Tot_c_wt_pct, Org_c_wt_pct, Nit_wt_pct, Sul_wt_pct, H_wt_pct)

	}
	if err := rows5.Err(); err != nil {
		t.Error(err)
	}

	log.Printf("Run query 6\n")

	rows6, err := testDb.Query(qry6)
	if err != nil {
		t.Errorf(`Error with "%s": %s`, qry6, err)
		return
	}
	defer rows6.Close()

	vi := 0
	for rows6.Next() {
		vi++
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

		if err := rows6.Scan(&Leg, &Site, &Hole, &Core, &Core_type, &Section_number, &Section_type, &Top_cm, &Bot_cm, &Inor_c_wt_pct, &Caco3_wt_pct, &Tot_c_wt_pct, &Org_c_wt_pct, &Nit_wt_pct, &Sul_wt_pct, &H_wt_pct); err != nil {
			t.Fatalf("scan %d. record: %v", v, err)
		}

		log.Printf("Results: %v %v %v %v %v %v %v %v %v %v %v %v %v %v %v %v", Leg, Site, Hole, Core, Core_type, Section_number, Section_type, Top_cm, Bot_cm, Inor_c_wt_pct, Caco3_wt_pct, Tot_c_wt_pct, Org_c_wt_pct, Nit_wt_pct, Sul_wt_pct, H_wt_pct)

	}
	if err := rows6.Err(); err != nil {
		t.Error(err)
	}

}
