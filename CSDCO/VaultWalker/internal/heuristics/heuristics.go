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
		HTest{DirPattern: "/",
			FilePattern: []string{".-metadata"},
			FileExts:    []string{},
			BlackList:   []string{},
			Comment:     "Project Metadata file",
			URI:         "http://opencoredata.org/voc/csdco/v1/Metadata"},
		HTest{DirPattern: "/",
			FilePattern: []string{"metadata format Dtube Label_"},
			FileExts:    []string{},
			BlackList:   []string{},
			Comment:     "Dtube metadata",
			URI:         "http://opencoredata.org/voc/csdco/v1/DtubeMetadata"},
		HTest{DirPattern: "/",
			FilePattern: []string{"SRF"},
			FileExts:    []string{""},
			BlackList:   []string{},
			Comment:     "SRF",
			URI:         "http://opencoredata.org/voc/csdco/v1/SRF"},
		HTest{DirPattern: "/",
			FilePattern: []string{},
			FileExts:    []string{".cml"},
			BlackList:   []string{},
			Comment:     "Corewall session files",
			URI:         "http://opencoredata.org/voc/csdco/v1/CML"},
		HTest{DirPattern: "/",
			FilePattern: []string{},
			FileExts:    []string{".car"},
			BlackList:   []string{},
			Comment:     "Corewall package files",
			URI:         "http://opencoredata.org/voc/csdco/v1/Car"},
		HTest{DirPattern: "/Images",
			FilePattern: []string{},
			FileExts:    []string{".bmp", ".jpeg", ".jpg", "tif", "tiff"},
			BlackList:   []string{},
			Comment:     "Images",
			URI:         "http://opencoredata.org/voc/csdco/v1/Image"},
		HTest{DirPattern: "/Images/rgb",
			FilePattern: []string{},
			FileExts:    []string{".csv"},
			BlackList:   []string{},
			Comment:     "RGB Image Data",
			URI:         "http://opencoredata.org/voc/csdco/v1/RGBData"},
		HTest{DirPattern: "Geotek Data/whole-core data",
			FilePattern:   []string{"_MSCL"},
			IgnorePattern: []string{"other data"},
			FileExts:      []string{".xls", ".xlsx", ".csv"}, // what is the point of a black list?  I only validate on FileExts found???
			BlackList:     []string{".raw", ".dat", ".out", ".cal"},
			Comment:       "GEOTEK WhCr",
			URI:           "http://opencoredata.org/voc/csdco/v1/WholeCoreData"},
		HTest{DirPattern: "Geotek Data/high-resolution MS data",
			FilePattern: []string{"_HRMS", "_XYZ"},
			FileExts:    []string{".xls", ".xlsx", ".csv"},
			BlackList:   []string{},
			Comment:     "GEOTEK HiRez",
			URI:         "http://opencoredata.org/voc/csdco/v1/SplitCoreData"},
		HTest{DirPattern: "ICD/",
			FilePattern: []string{"ICD sheet.pdf"},
			FileExts:    []string{".pdf"},
			BlackList:   []string{},
			Comment:     "ICD",
			URI:         "http://opencoredata.org/voc/csdco/v1/ICDFiles"}}

	return ht
}
