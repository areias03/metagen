package db

import (
	"strings"
)

type Database struct {
	Name  string `json:"name"`
	Url   string `json:"url"`
	Match int
}

type Databases struct {
	Databases []Database `json:"databases"`
}

func (db Database) defineQuery(term string) string {
	var query string = strings.ReplaceAll(db.Url, "item", term)
	return query
}

// Create a new config instance.
var DatabaseConfig Databases = Databases{[]Database{
	{
		Name: "SPIRE Study",
		Url:  "https://spire.embl.de/api/study/item",
	},
	{
		Name: "SPIRE Sample",
		Url:  "https://spire.embl.de/api/sample/item",
	},
	{
		Name: "ENA",
		Url:  "https://www.ebi.ac.uk/ena/browser/api/xml/item",
	},
	{
		Name: "SRA Biosample",
		Url:  "https://www.ncbi.nlm.nih.gov/biosample/item",
	},
	{
		Name: "SRA Bioproject",
		Url:  "https://www.ncbi.nlm.nih.gov/bioproject/item",
	}}}
