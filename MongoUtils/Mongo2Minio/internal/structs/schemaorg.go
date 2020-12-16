package structs

// schema.org Dataset metadata structs
type SchemaOrgMetadata struct {
	Context             map[string]interface{} `json:"@context"`
	Type                string                 `json:"@type"`
	Author              Author                 `json:"author"`
	Description         string                 `json:"description"`
	Distribution        Distribution           `json:"distribution"`
	GlviewDataset       string                 `json:"glview:dataset"`
	GlviewKeywords      string                 `json:"glview:keywords"`
	OpenCoreLeg         string                 `json:"opencore:leg"`
	OpenCoreSite        string                 `json:"opencore:site"`
	OpenCoreHole        string                 `json:"opencore:hole"`
	OpenCoreMeasurement string                 `json:"opencore:measurement"`
	Keywords            string                 `json:"keywords"`
	Name                string                 `json:"name"`
	Spatial             Spatial                `json:"spatial"`
	URL                 string                 `json:"url"`
}

type Author struct {
	Type        string `json:"@type"`
	Description string `json:"description"`
	Name        string `json:"name"`
	URL         string `json:"url"`
}

type Distribution struct {
	Type           string `json:"@type"`
	ContentURL     string `json:"contentUrl"`
	DatePublished  string `json:"datePublished"`
	EncodingFormat string `json:"encodingFormat"`
	InLanguage     string `json:"inLanguage"`
}

type Spatial struct {
	Type string `json:"@type"`
	Geo  Geo    `json:"geo"`
}

type Geo struct {
	Type      string `json:"@type"`
	Latitude  string `json:"latitude"`
	Longitude string `json:"longitude"`
}

type VariableMeasured struct {
	Name        string
	UnitText    string
	Description string
	URL         string
}

// VoidDataset is a struct to hold items from a VOiD file that
// describe a dataset  https://developers.google.com/search/docs/data-types/datasets
type VoidDataset struct {
	ID                 string
	URL                string // type URL: Location of a page describing the dataset.
	Description        string // A short summary describing a dataset.
	Keywords           string // Keywords summarizing the dataset.
	Name               string // A descriptive name of a dataset (e.g., “Snow depth in Northern Hemisphere”)
	ContentURL         string
	AccrualPeriodicity string
	Issued             string
	License            string
	Publisher          string // Person, Org The name of the dataset creator (person or organization).
	Title              string
	DataDump           string
	Source             string
	LandingPage        string
	DownloadURL        string
	MediaType          string
	SameAs             string             // type URL: Other URLs that can be used to access the dataset page.
	Version            string             // The version number for this dataset.
	VariableMeasured   []VariableMeasured // What does the dataset measure? (e.g., temperature, pressure)
	PublisherDesc      string
	PublisherName      string
	PublisherURL       string
	Latitude           string
	Longitude          string
	OpenCoreLeg        string
	OpenCoreSite       string
	OpenCoreHole       string
}
