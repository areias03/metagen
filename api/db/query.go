package db

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

var wg sync.WaitGroup

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

func SearchDBs(item string, dbs Databases) {
	for _, v := range dbs.Databases {
		var query string = v.defineQuery(item)
		wg.Add(1)
		go processQuery(query)
	}
	wg.Wait()
}
