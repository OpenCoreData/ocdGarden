package structs

// schema.org Dataset metadata structs
type SchemaOrgMetadata struct {
	Context             []interface{} `json:"@context"`
	Type                string        `json:"@type"`
	Author              Author        `json:"author"`
	Description         string        `json:"description"`
	Distribution        Distribution  `json:"distribution"`
	GlviewDataset       string        `json:"glview:dataset"`
	GlviewKeywords      string        `json:"glview:keywords"`
	OpenCoreLeg         string        `json:"opencore:leg"`
	OpenCoreSite        string        `json:"opencore:site"`
	OpenCoreHole        string        `json:"opencore:hole"`
	OpenCoreMeasurement string        `json:"opencore:measurement"`
	Keywords            string        `json:"keywords"`
	Name                string        `json:"name"`
	Spatial             Spatial       `json:"spatial"`
	URL                 string        `json:"url"`
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
