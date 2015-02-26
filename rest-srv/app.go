package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"runtime"

	"github.com/dustin/go-humanize"
	"github.com/go-martini/martini"
	"github.com/olivere/elastic"
)

var (
	esurl = "http://localhost:8080"
	index = "s3data"
	debug        = false
)

type job struct {
	customer, from, to string
}

func parseParams(r *http.Request) (job, error) {

	log.Println(r.URL.Query())

	params := r.URL.Query()

	customer := params.Get("customer")
	from := params.Get("from")
	to := params.Get("to")

	if len(customer) == 0 || len(from) == 0 || len(to) == 0 {
		return job{},
			fmt.Errorf("customer, from, to fields are required")
	}

	return job{customer: customer, from: from, to: to}, nil

}

//todo: create DAL

func listCustomers(w http.ResponseWriter,
	r *http.Request) string {
	client, err := elastic.NewClient(
		http.DefaultClient,
		esurl)

	if err != nil {
		http.Error(w,
			err.Error(),
			http.StatusBadRequest)
	}

	customerTermsAggr := elastic.NewTermsAggregation().Field("customer")

	out, err := client.Search().
		Index(index).
		Aggregation("cust_unique", customerTermsAggr).
		Debug(debug).
		Pretty(debug).
		Do()

	if err != nil {
		showError(w, err)
	}
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

//todo:move to DAL
func getFilteredQuery(j job) elastic.FilteredQuery {

	customerQuery := elastic.NewTermQuery("customer", j.customer)
	filteredQuery := elastic.NewFilteredQuery(customerQuery)

	dateFilter := elastic.NewRangeFilter("@timestamp").
		From(j.from).
		To(j.to)

	filteredQuery = filteredQuery.Filter(dateFilter)

	return filteredQuery

}

func getJob(w http.ResponseWriter,
	r *http.Request) string {

	job, err := parseParams(r)

	if err != nil {
		http.Error(w,
			err.Error(),
			http.StatusBadRequest)
		return ""
	}

	client, err := elastic.NewClient(
		http.DefaultClient,
		esurl)

	if err != nil {
		return showError(w, err)
	}

	filteredQuery := getFilteredQuery(job)

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

		res["eta"] = calcEta(float64(out.Hits.TotalHits))

		data, _ := json.Marshal(res)
		return string(data)
	}

	return `{"count": 0, "size": 0}`
}

func startJob(w http.ResponseWriter,
	r *http.Request) string {

	j, err := parseParams(r)

	if err != nil {
		http.Error(w,
			err.Error(),
			http.StatusBadRequest)
		return ""
	}

	go sendJob(j)

	return "started"
}

//todo: show stacktrace error in debug localhost, show empty 500 in production
func showError(w http.ResponseWriter, err error) string {
	http.Error(w,
		err.Error(),
		http.StatusInternalServerError)

	return ""
}

func main() {

	num := runtime.NumCPU()
	runtime.GOMAXPROCS(num)

	m := martini.Classic()
	m.Use(martini.Logger())

	m.Get("/job", getJob)
	m.Get("/eta", getEta)
	m.Get("/customers", listCustomers)
	m.Post("/job", startJob)
	//todo:remove this as I understand how to enable post in CUI
	m.Get("/job/create", startJob)
	m.Run()
}
