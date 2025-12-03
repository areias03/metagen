package db

import (
	"fmt"
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
	body, err := io.ReadAll(resp.Request.Response.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(body))

	if resp.StatusCode != http.StatusOK || resp.ContentLength != -1 {
		results <- queryResult{ID: db.Name, Output: 0}
	} else {
		results <- queryResult{ID: db.Name, Struct: db.parseResponse(db.Name, body), Output: 1}
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
		ResultMap[res.ID] = MapVal{res.Output, res.Struct}
	}
}
