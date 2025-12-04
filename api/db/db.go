package db

import (
	"encoding/json"
	"log"
	"strings"
)

type Database struct {
	Name           string
	Url            string
	ResponseStruct any
}

type Databases struct {
	Databases []Database
}

func (db *Database) defineQuery(term string) string {
	query := strings.ReplaceAll(db.Url, "item", term)
	return query
}

func parseResponse(name string, body []byte) any {
	switch name {
	case "SPIRE Study":
		var resp SpireStudy
		if err := json.Unmarshal(body, &resp.Samples); err != nil {
			log.Fatal(err)
		}
		return resp
	case "SPIRE Sample":
		var resp SpireSampleGenomes
		if err := json.Unmarshal(body, &resp.Mags); err != nil {
			log.Fatal(err)
		}
		return resp
	}
	return nil
}

var DatabaseConfig Databases = Databases{[]Database{
	{
		Name: "SPIRE Study",
		Url:  "https://spire.embl.de/spire/api/study/item",
	},
	{
		Name: "SPIRE Sample",
		Url:  "https://spire.embl.de/spire/api/sample/item",
	},
}}

// {
// 	Name: "ENA",
// 	Url:  "https://www.ebi.ac.uk/ena/browser/api/xml/item",
// },
// {
// 	Name: "SRA Biosample",
// 	Url:  "https://www.ncbi.nlm.nih.gov/biosample/item",
// },
// {
// 	Name: "SRA Bioproject",
// 	Url:  "https://www.ncbi.nlm.nih.gov/bioproject/item",
// }}}
