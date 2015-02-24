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

/* TODO:
create DAL
*/

func getJob(w http.ResponseWriter,
	r *http.Request) string {

	log.Println(r.URL.Query())

	params := r.URL.Query()
	customer := params.Get("customer")
	from := params.Get("from")
	to := params.Get("to")

	client, err := elastic.NewClient(
		http.DefaultClient,
		esurl)

	if err != nil {

		http.Error(w,
			err.Error(),
			http.StatusInternalServerError)

		return ""
	}

	if len(customer) == 0 {

		http.Error(w,
			"Could you please specify the customer",
			http.StatusBadRequest)
		return ""

	}

	customerQuery := elastic.NewTermQuery("customer", customer)

	postDateFilter := elastic.NewRangeFilter("@timestamp").
		From(from).
		To(to)

	nested := elastic.NewFilteredQuery(customerQuery)
	nested = nested.Filter(postDateFilter)

	sizeSumAggr := elastic.NewSumAggregation().Field("size")

	searchResult, err := client.Search().
		Index(index).
		Query(&nested).
		Aggregation("sum", sizeSumAggr).
		Debug(true).
		Pretty(true).
		Do()

	if err != nil {

		http.Error(w,
			err.Error(),
			http.StatusInternalServerError)

		return ""
	}

	if searchResult.Hits != nil {

		var sumResult map[string]interface{}

		err = json.Unmarshal(
			searchResult.Aggregations["sum"],
			&sumResult)

		if err != nil {

			http.Error(w,
				err.Error(),
				http.StatusInternalServerError)

			return ""
		}

		size := sumResult["value"]

		return fmt.Sprintf(`{"count":%d, "size": %f}`,
			searchResult.Hits.TotalHits, size)

	}

	return `{"count": 0, "size": 0}`

}

func main() {

	m := martini.Classic()
	m.Use(martini.Logger())

	m.Get("/job", getJob)

	m.Run()
}
