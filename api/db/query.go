package db

import (
	"io"
	"log"
	"net/http"
	"sync"
)

var (
	wg        sync.WaitGroup
	ResultMap = make(map[string]MapVal)
	results   = make(chan queryResult)
)

type MapVal struct {
	Status int
	Stru   any
}

type queryResult struct {
	ID     string
	Struct any
	Output int
}

func processQuery(db *Database, query string, results chan<- queryResult, wg *sync.WaitGroup) {
	defer wg.Done()
	resp, err := http.Get(query)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	if resp.StatusCode != http.StatusOK || resp.ContentLength != -1 {
		results <- queryResult{ID: db.Name, Output: 0}
	} else {
		responseStruct := parseResponse(db.Name, body)
		results <- queryResult{ID: db.Name, Struct: responseStruct, Output: 1}
	}
}

func SearchDBs(item string, dbs *Databases) {
	for _, v := range dbs.Databases {
		query := v.defineQuery(item)
		wg.Add(1)
		go processQuery(&v, query, results, &wg)
	}
	go func() {
		wg.Wait()
		close(results)
	}()

	for res := range results {
		ResultMap[res.ID] = MapVal{res.Output, res.Struct}
	}
}
