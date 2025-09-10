package cmd

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"sync"

	"github.com/areias03/metagen/api/db"
	"github.com/spf13/cobra"
)

var wg sync.WaitGroup

func defineQuery(term string, url string) string {
	var query string = strings.ReplaceAll(url, "item", term)
	return query
}

func processQuery(query string) {
	resp, err := http.Get(query)
	if err != nil {
		log.Fatal(err)
	}
	if resp.StatusCode != http.StatusOK || resp.ContentLength != -1 {
		fmt.Println(query, "\t", "Not found!")
	} else {
		fmt.Println(query, "\t", "Found!")
	}
	wg.Done()
}

func searchDBs(item string, dbs db.Databases) {
	for _, v := range dbs.Databases {
		var query string = defineQuery(item, v.Url)
		wg.Add(1)
		go processQuery(query)
	}
	wg.Wait()
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
		jsonFile, err := os.Open("config/databases.json")
		if err != nil {
			log.Fatal(err)
		}
		defer jsonFile.Close()

		byteValue, err := io.ReadAll(jsonFile)
		if err != nil {
			log.Fatal(err)
		}

		var dbs db.Databases
		// we unmarshal our byteArray which contains our
		// jsonFile's content into 'users' which we defined above
		json.Unmarshal(byteValue, &dbs)
		searchDBs("SAMN07510030", dbs)
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
