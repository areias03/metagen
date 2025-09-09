/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

type Databases struct {
	databases []Database `json:"databases"`
}

type Database struct {
	name string `json:"name"`
	url  string `json:"url"`
}

func loadDatabase(config string) Databases {
	byteValue, err := os.ReadFile(config)
	if err != nil {
		log.Fatal(err)
	}
	var dbs Databases
	json.Unmarshal(byteValue, &dbs)
	return dbs

}

func defineQuery(term string, url string) string {
	var query string = strings.ReplaceAll(url, "item", term)
	return query
}

func searchDBs(item string, dbs Databases) {
	for i := 0; i < len(dbs.databases); i++ {
		var query string = defineQuery(item, dbs.databases[i].url)
		go http.Get(query)
	}
}

// searchCmd represents the search command
var searchCmd = &cobra.Command{
	Use:   "search",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("search called")
	},
}

func init() {
	rootCmd.AddCommand(searchCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// searchCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// searchCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
