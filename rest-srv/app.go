package main

import (
	"log"
	"net/http"
	"fmt"

	"github.com/go-martini/martini"
	"github.com/olivere/elastic"
)

var (
	esurl string = "http://localhost:8080" 
	index string = "s3data"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func getJob(r *http.Request) string {
	log.Println(r.URL.Query())
	params := r.URL.Query()
	customer := params.Get("customer")
	client, err := elastic.NewClient(http.DefaultClient, esurl)
	check(err)

	//todo:find a frameworkk that supports declarative fields description
	if(len(customer) < 1){
		panic("customer is required")
	}

	termQuery := elastic.NewTermQuery("customer", customer)

	searchResult, err := client.Search().
		Index(index).
		Query(&termQuery).
		From(0).Size(10). 
		Debug(true). 
		Pretty(true). 
		Do() 
	check(err)

	if searchResult.Hits != nil {
		//todo:use json.marshal
		log.Printf("Found a total of %d tweets\n", searchResult.Hits.TotalHits)
		return fmt.Sprintf(`{"count":%d, "size": 220}`,  searchResult.Hits.TotalHits)
	}

	panic("todo:return 500")
}

func main() {

	m := martini.Classic()
	m.Use(martini.Logger())

	m.Get("/job", getJob)

	m.Run()

}
