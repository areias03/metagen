package db

import (
	"fmt"
	"github.com/spf13/viper"
	"strings"
)

type Database struct {
	Name  string `json:"name"`
	Url   string `json:"url"`
	Match string
}

type Databases struct {
	Databases []Database `json:"databases"`
}

func (db Database) defineQuery(term string) string {
	var query string = strings.ReplaceAll(db.Url, "item", term)
	return query
}

// Create a new config instance.
var (
	conf *Databases
)

// Read the config file from the current directory and marshal
// into the conf config struct.
func getConf() *Databases {
	viper.AddConfigPath("config")
	viper.SetConfigName("databases")
	err := viper.ReadInConfig()

	if err != nil {
		fmt.Printf("%v", err)
	}

	conf := &Databases{}
	err = viper.Unmarshal(conf)
	if err != nil {
		fmt.Printf("unable to decode into config struct, %v", err)
	}

	return conf
}

// Initialization routine.
func init() {
	// Retrieve config options.
	conf = getConf()
}

var Conf = &conf
