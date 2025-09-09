package db

type Databases struct {
	Databases []Database `json:"databases"`
}

type Database struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}
