package main

import (
	"log"
	"net/http"
	"fmt"

	"github.com/go-martini/martini"
	"github.com/olivere/elastic"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func getJob(r *http.Request) string {
	#todo:dehartdcode
	url := "http://localhost:8080"

	log.Println(r.URL.Query())

	client, err := elastic.NewClient(http.DefaultClient, url)
	check(err)

	searchResult, err := client.Search().
		Index("s3data").  // search in index "twitter"
		From(0).Size(10). // take documents 0-9
		Debug(true).      // print request and response to stdout
		Pretty(true).     // pretty print request and response JSON
		Do()              // execute
	check(err)

	if searchResult.Hits != nil {
		#todo:use json.marchal
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
