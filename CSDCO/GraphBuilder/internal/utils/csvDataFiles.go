package utils

import (
	"encoding/csv"
	"os"
	"strings"
	"unicode/utf8"
	"log"

	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
)

type CSDCOv2 struct {
	LocationName   string
	LocationType   string
	Project        string
	LocationID     string
	Site           string
	Hole           string
	SiteHole       string
	OriginalID     string
	HoleID         string
	Platform       string
	Date           string
	WaterDepthM    string
	Country        string
	State_Province string
	County_Region  string
	PI             string
	Lat            string
	Long           string
	Elevation      string
	Position       string
	SampleType     string
	MblfT          string
	MblfB          string
}

// the above is missing  StorageLocationWorking StorageLocationArchive Comment MetadataSource

type CSDCO struct {
	LocationName           string
	LocationType           string
	Project                string
	LocationID             string
	Site                   string
	Hole                   string
	SiteHole               string
	OriginalID             string
	HoleID                 string
	Platform               string
	Date                   string
	WaterDepthM            string
	Country                string
	State_Province         string
	County_Region          string
	PI                     string
	Lat                    string
	Long                   string
	Elevation              string
	Position               string
	StorageLocationWorking string
	StorageLocationArchive string
	SampleType             string
	Comment                string
	MblfT                  string
	MblfB                  string
	MetadataSource         string
}

type CSDCOIGSN struct {
	Ngdcship                string
	Ngdccruise              string
	Ngdcsample              string
	Ngdcserial              string
	Ngdccomment             string
	Location                string
	Locationtype            string
	Expedition              string
	Locationid              string
	Site                    string
	Hole                    string
	Sitehole                string
	Originalid              string
	Holeid                  string
	Platform                string
	Date                    string
	Waterdepthm             string
	Country                 string
	State_province          string
	County_region           string
	Pi                      string
	Lat                     string
	Long                    string
	Elevation               string
	Position                string
	StorageLocationWorking  string
	StorageLocationArchive  string
	Sampletype              string
	Comment                 string
	Mblft                   string
	Mblfb                   string
	Metadatasource          string
	Googleearthcommentfield string
	Platformname            string
	Platformtype            string
	Azimuth                 string
	Dip                     string
	Igsn                    string
}

// ReadIGSNData function pulls in CSV file and generates resources from it
func ReadIGSNData() []CSDCOIGSN {
	// /Users/dfils/src/go/src/opencoredata.org/ocdBulk/CSDCOProjects/LacCore_Holes_20161010_trim_corrected.xlsx
	csvFile, err := os.Open("./data/input/LacCore_Holes_20180910.csv") // note file name.. LacCore_Holes.csv    LacCore_Holes_20160915_trim.csv
	defer csvFile.Close()
	if err != nil {
		panic(err)
	}
	r := csv.NewReader(csvFile)
	//r.FieldsPerRecord = -1 //  don't require given field count since some may not have hole (fix with sparql query update too?)

	lines, err := r.ReadAll()
	if err != nil {
		log.Fatalf("Error reading all lines: %v", err)
	}

	commentLines := 1
	observations := make([]CSDCOIGSN, len(lines)-commentLines) //  why did I do this and not a slice?  then just append?

	for i, line := range lines {
		if i < commentLines {
			continue // skip header lines
		}

		//line0, err := strconv.ParseFloat(line[0], 64)
		if err != nil {
			panic(err)
		}

		latval := StripCtlAndExtFromUnicode(line[21])
		longval := StripCtlAndExtFromUnicode(line[22])

		ob := CSDCOIGSN{Ngdcship: line[0], Ngdccruise: line[1], Ngdcsample: line[2], Ngdcserial: line[3],
			Ngdccomment: line[4], Location: line[5], Locationtype: line[6], Expedition: line[7],
			Locationid: line[8], Site: line[9], Hole: line[10], Sitehole: line[11],
			Originalid: line[12], Holeid: line[13], Platform: line[14], Date: line[15],
			Waterdepthm: line[16], Country: line[17], State_province: line[18],
			County_region: line[19], Pi: line[20],
			Lat: strings.TrimSpace(latval), Long: strings.TrimSpace(longval),
			Elevation: line[23], Position: line[24],
			StorageLocationWorking: line[25], StorageLocationArchive: line[26], Sampletype: line[27],
			Comment: line[28], Mblft: line[29], Mblfb: line[30],
			Metadatasource: line[31], Googleearthcommentfield: line[32],
			Platformname: line[33], Platformtype: line[34], Azimuth: line[35],
			Dip: line[36], Igsn: line[37]}

		observations[i-commentLines] = ob
	}

	return observations
}

func ReadMetaData() []CSDCO {
	// /Users/dfils/src/go/src/opencoredata.org/ocdBulk/CSDCOProjects/LacCore_Holes_20161010_trim_corrected.xlsx
	csvFile, err := os.Open("./data/input/LacCore_Holes_20161010_trim_corrected.csv") // note file name.. LacCore_Holes.csv    LacCore_Holes_20160915_trim.csv
	defer csvFile.Close()
	if err != nil {
		panic(err)
	}
	r := csv.NewReader(csvFile)
	//r.FieldsPerRecord = -1 //  don't require given field count since some may not have hole (fix with sparql query update too?)

	lines, err := r.ReadAll()
	if err != nil {
		//log.Fatalf("Error reading all lines: %v", err)
	}

	commentLines := 1
	observations := make([]CSDCO, len(lines)-commentLines) //  why did I do this and not a slice?  then just append?

	for i, line := range lines {
		if i < commentLines {
			continue // skip header lines
		}

		//line0, err := strconv.ParseFloat(line[0], 64)
		if err != nil {
			panic(err)
		}

		// line[29]   url with UUDI struct to it..
		//  http://opencoredata.org/collection/d53453fc-1431-4f09-a9dd-4502ed922be8
		//
		// ob := CSDCO{LocationName: line[0], LocationType: line[1], Project: line[2], LocationID: line[3], Site: line[4], Hole: line[5],
		// 	SiteHole: line[6], OriginalID: line[7], HoleID: line[8], Platform: line[9], Date: line[10], WaterDepthM: line[11],
		// 	Country: line[12], State_Province: line[13], County_Region: line[14], PI: line[15], Lat: line[16], Long: line[17],
		// 	Elevation: line[18], Position: line[19], StorageLocationWorking: line[20], StorageLocationArchive: line[21],
		// 	SampleType: line[22], Comment: line[23], MblfT: line[24], MblfB: line[25], MetadataSource: line[26]}

		latval := StripCtlAndExtFromUnicode(line[16])
		longval := StripCtlAndExtFromUnicode(line[17])

		ob := CSDCO{LocationName: line[0], LocationType: line[1], Project: line[2], LocationID: line[3], Site: line[4], Hole: line[5],
			SiteHole: line[6], OriginalID: line[7], HoleID: line[8], Platform: line[9], Date: line[10], WaterDepthM: line[11],
			Country: line[12], State_Province: line[13], County_Region: line[14], PI: line[15],
			Lat: strings.TrimSpace(latval), Long: strings.TrimSpace(longval),
			Elevation: line[18], Position: line[19], SampleType: line[20], MblfT: line[21], MblfB: line[22]}

		observations[i-commentLines] = ob
	}

	return observations
}

// ref:  https://rosettacode.org/wiki/Strip_control_codes_and_extended_characters_from_a_string#Go
func StripCtlAndExtFromUnicode(str string) string {
	isOk := func(r rune) bool {
		return r < 32 || r >= 127
	}
	// The isOk filter is such that there is no need to chain to norm.NFC
	t := transform.Chain(norm.NFKD, transform.RemoveFunc(isOk))
	// This Transformer could also trivially be applied as an io.Reader
	// or io.Writer filter to automatically do such filtering when reading
	// or writing data anywhere.
	str, _, _ = transform.String(t, str)
	return str
}

// get the better version from the gardem in the mongo2rdf
func DONOTUSEutfremove(s string) string {
	//if !utf8.ValidString(s) {
	v := make([]rune, 0, len(s))
	for i, r := range s {
		if r == utf8.RuneError {
			_, size := utf8.DecodeRuneInString(s[i:])
			if size == 1 {
				continue
			}
		}
		v = append(v, r)
	}
	s = string(v)
	//}
	return s
}
