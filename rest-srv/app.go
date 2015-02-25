package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"runtime"
	"time"

	"github.com/dustin/go-humanize"
	"github.com/go-martini/martini"
	"github.com/olivere/elastic"
)

var (
	esurl string = "http://localhost:8080"
	index string = "s3data"
	debug        = true
)

//todo: create DAL

func listCustomers(w http.ResponseWriter,
	r *http.Request) string {
	client, err := elastic.NewClient(
		http.DefaultClient,
		esurl)

	if err != nil {
		return showError(w, err)
	}

	customerTermsAggr := elastic.NewTermsAggregation().Field("customer")

	out, _ := client.Search().
		Index(index).
		Aggregation("cust_unique", customerTermsAggr).
		Debug(debug).
		Pretty(debug).
		Do()

	if out.Hits != nil {
		var aggrResult map[string]interface{}

		err = json.Unmarshal(
			out.Aggregations["cust_unique"],
			&aggrResult)

		if err != nil {
			showError(w, err)
		}

		buckets := aggrResult["buckets"].([]interface{})

		ret := make([]string, len(buckets))
		for i, bucket := range buckets {
			item := bucket.(map[string]interface{})
			ret[i] = item["key"].(string)
		}

		JSON, _ := json.Marshal(map[string]interface{}{
			"result": ret,
		})

		return string(JSON)

	}
	return ""
}

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
		return showError(w, err)
	}

	if len(customer) == 0 {

		http.Error(w,
			"Could you please specify the customer",
			http.StatusBadRequest)
		return ""

	}

	customerQuery := elastic.NewTermQuery("customer", customer)

	dateFilter := elastic.NewRangeFilter("@timestamp").
		From(from).
		To(to)

	filteredQuery := elastic.NewFilteredQuery(customerQuery)
	filteredQuery = filteredQuery.Filter(dateFilter)

	sizeSumAggr := elastic.NewSumAggregation().Field("size")

	out, err := client.Search().
		Index(index).
		Query(&filteredQuery).
		Aggregation("sum", sizeSumAggr).
		Debug(debug).
		Pretty(debug).
		Do()

	if err != nil {
		return showError(w, err)
	}

	if out.Hits != nil {

		var aggrResult map[string]interface{}

		err = json.Unmarshal(
			out.Aggregations["sum"],
			&aggrResult)

		if err != nil {
			showError(w, err)
		}

		size := aggrResult["value"]

		res := map[string]interface{}{}

		res["count"] = humanize.Comma(out.Hits.TotalHits)
		res["size"] = humanize.Bytes(uint64(size.(float64)))

		cpu := runtime.NumCPU()

		fPerSec := 2.5 / 70.0 * float64(cpu)
		secs := float64(out.Hits.TotalHits) / fPerSec

		eta := time.Second * time.Duration(secs)

		res["eta"] = fmt.Sprintf("%v", eta)

		data, _ := json.Marshal(res)
		return string(data)
	}

	return `{"count": 0, "size": 0}`

}

//todo: show stacktrace error in debug localhost, show empty 500 in production
func showError(w http.ResponseWriter, err error) string {
	http.Error(w,
		err.Error(),
		http.StatusInternalServerError)

	return ""
}

func main() {

	m := martini.Classic()
	m.Use(martini.Logger())

	m.Get("/job", getJob)
	m.Get("/customers", listCustomers)

	m.Run()
}
