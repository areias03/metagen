package db

import (
	"log"
	"net/http"
	"sync"
)

var (
	wg        sync.WaitGroup
	ResultMap = make(map[string]int)
	results   = make(chan queryResult)
)

type queryResult struct {
	ID     string
	Output int
}

func processQuery(db *Database, query string, results chan<- queryResult, wg *sync.WaitGroup) {
	defer wg.Done()
	resp, err := http.Get(query)
	if err != nil {
		log.Fatal(err)
	}
	if resp.StatusCode != http.StatusOK || resp.ContentLength != -1 {
		results <- queryResult{ID: db.Name, Output: 0}
	} else {
		results <- queryResult{ID: db.Name, Output: 1}
	}
}

func SearchDBs(item string, dbs *Databases) {
	for _, v := range dbs.Databases {
		var query string = v.defineQuery(item)
		wg.Add(1)
		go processQuery(&v, query, results, &wg)
	}
	go func() {
		wg.Wait()
		close(results)
	}()

	for res := range results {
		ResultMap[res.ID] = res.Output
	}
}
