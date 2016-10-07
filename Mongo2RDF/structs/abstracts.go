package structs

import "time"

type Mdocs struct {
	Title   string `json:"title"`
	Type    string `json:"type"`
	Authors []struct {
		First_Name string `json:"first_name"`
		Last_Name  string `json:"last_name"`
	} `json:"authors"`
	Year        int    `json:"year"`
	Source      string `json:"source"`
	Identifiers struct {
		Doi  string `json:"doi"`
		Issn string `json:"issn"`
	} `json:"identifiers"`
	ID            string    `json:"id"`
	Created       time.Time `json:"created"`
	Profile_ID    string    `json:"profile_id"`
	Group_ID      string    `json:"group_id"`
	Last_Modified time.Time `json:"last_modified"`
	Abstract      string    `json:"abstract"`
}
