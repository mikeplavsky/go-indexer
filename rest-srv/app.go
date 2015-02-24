package main

import (
	"fmt"
	"log"
	"net/http"

	"encoding/json"
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


/* TODO:
create DAL
*/

func getJob(r *http.Request) string {
	log.Println(r.URL.Query())
	params := r.URL.Query()
	customer := params.Get("customer")
	client, err := elastic.NewClient(http.DefaultClient, esurl)
	check(err)

	//todo: find a frameworkk that supports declarative fields description
//	if len(customer) < 1 {
//		panic("customer is required")
//	}

	customerQuery := elastic.NewTermQuery("customer", customer)
	sizeSumAggr := elastic.NewSumAggregation().Field("size")
	searchResult, err := client.Search().
		Index(index).
		Query(&customerQuery).
		Aggregation("sum", sizeSumAggr).
		Debug(true).
		Pretty(true).
		Do()
	check(err)

	if searchResult.Hits != nil {
		var sumResult map[string]interface{}
		err = json.Unmarshal(searchResult.Aggregations["sum"], &sumResult) 
		check(err)
		size := sumResult["value"]
		return fmt.Sprintf(`{"count":%d, "size": %f}`, searchResult.Hits.TotalHits, size)
	}

	panic("todo:return 500")
}

func main() {

	m := martini.Classic()
	m.Use(martini.Logger())

	m.Get("/job", getJob)

	m.Run()
}
