package utils

import "fmt"

// FileTypes is a file type struct (duh)
type FileTypes struct {
	Name         string
	Regexpattern string
	URI          string
	Description  string
}

// TODO:  resolve these items into the struct
// _HRMS or _XYZ"
// _HRMS or _XYZ
// ICD/
// ICD sheet.pdf

// GetFileTypes is a set of metadata for CSDCO file types
// Will be made into a simple vocabulary
func GetFileTypes() {
	fmt.Println("vim-go")

	fta := []FileTypes{}

	fta = append(fta, FileTypes{Name: " ", Regexpattern: "-metadata", URI: "http://opencoredata.org/id/voc/csdco/v1/metadata", Description: "Metadata for boreholes (tab 1) and core sections (tab 2) in each project."})
	fta = append(fta, FileTypes{Name: " ", Regexpattern: "metadata format Dtube Label_", URI: "http://opencoredata.org/id/voc/csdco/v1/dtubeMetadata", Description: "Metadata for core sections in each project. A superset of the base metadata file."})
	fta = append(fta, FileTypes{Name: " ", Regexpattern: "SRF", URI: "http://opencoredata.org/id/voc/csdco/v1/srf", Description: "Spreadsheet with metadata for subsamples extracted from core section halves."})
	fta = append(fta, FileTypes{Name: " ", Regexpattern: "", URI: "http://opencoredata.org/id/voc/csdco/v1/cml", Description: "Corelyzer session file: XML-type text that describes the workspace for a Corelyzer session."})
	fta = append(fta, FileTypes{Name: " ", Regexpattern: "", URI: "http://opencoredata.org/id/voc/csdco/v1/car", Description: "Corelyzer archive file: large packaged file with all images, data, and other files needed for complete Corelyzer session portability."})
	fta = append(fta, FileTypes{Name: " ", Regexpattern: "Images", URI: "http://opencoredata.org/id/voc/csdco/v1/image", Description: "mages of core section halves or core sections."})
	fta = append(fta, FileTypes{Name: " ", Regexpattern: "Images/rgb", URI: "http://opencoredata.org/id/voc/csdco/v1/rgbData", Description: "RGB color data extracted from the core image, in a 5mm wide strip down the middle of the image."})
	fta = append(fta, FileTypes{Name: " ", Regexpattern: "Geotek Data/whole-core data", URI: "http://opencoredata.org/id/voc/csdco/v1/wholeCoreData", Description: "Multisensor core logger data, determined by a Geotek MSCL-S."})
	fta = append(fta, FileTypes{Name: " ", Regexpattern: "Geotek Data/high-resolution MS data", URI: "http://opencoredata.org/id/voc/csdco/v1/geotekHighResMSdata", Description: "Multisensor core logger data, determined by a Geotek MSCL-XYZ."})
	fta = append(fta, FileTypes{Name: " ", Regexpattern: "ICD/ ", URI: "http://opencoredata.org/id/voc/csdco/v1/icdFiles", Description: "Core lithologic descriptions."})
}
