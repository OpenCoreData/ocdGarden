package heuristics

type HTest struct {
	DirPattern    string
	FilePattern   []string
	IgnorePattern []string
	FileExts      []string
	BlackList     []string
	Comment       string
	URI           string
}

// CSDCOHTs a set of tests to do on directory and file path/extensions.
func CSDCOHTs() []HTest {
	ht := []HTest{
		// HTest{DirPattern: "/Data/Sampling",
		// 	FilePattern: []string{"SRF"},
		// 	FileExts:    []string{""},
		// 	BlackList:   []string{},
		// 	Comment:     "SRF",
		// 	URI:         "http://opencoredata.org/voc/csdco/v1/SRF"},
		HTest{DirPattern: "Data/Corelyzer/",
			FilePattern: []string{},
			FileExts:    []string{".cml", ".xml"},
			BlackList:   []string{},
			Comment:     "Corelyzer files",
			URI:         "http://opencoredata.org/voc/csdco/v1/CML"},
		HTest{DirPattern: "Data/Corelyzer/",
			FilePattern: []string{},
			FileExts:    []string{".car"},
			BlackList:   []string{},
			Comment:     "Corelyzer archive files",
			URI:         "http://opencoredata.org/voc/csdco/v1/Car"},
		HTest{DirPattern: "Images/",
			FilePattern:   []string{},
			FileExts:      []string{".bmp", ".jpeg", ".jpg", "tif", "tiff"},
			IgnorePattern: []string{"tiff"},
			BlackList:     []string{},
			Comment:       "Images",
			URI:           "http://opencoredata.org/voc/csdco/v1/Image"},
		HTest{DirPattern: "Images/rgb/",
			FilePattern: []string{},
			FileExts:    []string{".csv"},
			BlackList:   []string{},
			Comment:     "RGB Image Data",
			URI:         "http://opencoredata.org/voc/csdco/v1/RGBData"},
		HTest{DirPattern: "MSCL/MSCL-S/",
			FilePattern:   []string{"_MSCL-S"},
			IgnorePattern: []string{"Other data"},
			FileExts:      []string{".xls", ".xlsx", ".csv"}, // what is the point of a black list?  I only validate on FileExts found???
			BlackList:     []string{".raw", ".dat", ".out", ".cal"},
			Comment:       "Geotek MSCL-S",
			URI:           "http://opencoredata.org/voc/csdco/v1/WholeCoreData"},
		HTest{DirPattern: "MSCL/MSCL-S_split/",
			FilePattern:   []string{"_MSCL-S_split"},
			IgnorePattern: []string{"Other data"},
			FileExts:      []string{".xls", ".xlsx", ".csv"}, // what is the point of a black list?  I only validate on FileExts found???
			BlackList:     []string{".raw", ".dat", ".out", ".cal"},
			Comment:       "Geotek MSCL-S Split-core",
			URI:           "http://opencoredata.org/voc/csdco/v1/WholeCoreData"},
		HTest{DirPattern: "MSCL/MSCL-XYZ/",
			FilePattern:   []string{"_MSCL-XYZ"},
			IgnorePattern: []string{"Other data"},
			FileExts:      []string{".xls", ".xlsx", ".csv"}, // what is the point of a black list?  I only validate on FileExts found???
			BlackList:     []string{".raw", ".dat", ".out", ".cal"},
			Comment:       "Geotek MSCL-XYZ",
			URI:           "http://opencoredata.org/voc/csdco/v1/SplitCoreData"},
		HTest{DirPattern: "ICD/",
			FilePattern: []string{},
			FileExts:    []string{".pdf"},
			BlackList:   []string{},
			Comment:     "ICD",
			URI:         "http://opencoredata.org/voc/csdco/v1/ICDFiles"},
		HTest{DirPattern: "ICD/",
			FilePattern: []string{"ICD_tabular"},
			FileExts:    []string{".xls", ".xlsx", ".csv"},
			BlackList:   []string{},
			Comment:     "ICD",
			URI:         "http://opencoredata.org/voc/csdco/v1/ICDFiles"}}

	return ht
}
