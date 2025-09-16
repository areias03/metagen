package db

import (
	"log"
	"net/http"
	"sync"
)

var wg sync.WaitGroup

func processQuery(db Database, query string) {
	resp, err := http.Get(query)
	if err != nil {
		log.Fatal(err)
	}
	if resp.StatusCode != http.StatusOK || resp.ContentLength != -1 {
		db.Match = "Not found!"
	} else {
		db.Match = "Found!"
	}
	wg.Done()
}

func SearchDBs(item string, dbs Databases) {
	for _, v := range dbs.Databases {
		var query string = v.defineQuery(item)
		wg.Add(1)
		go processQuery(v, query)
	}
	wg.Wait()
}
