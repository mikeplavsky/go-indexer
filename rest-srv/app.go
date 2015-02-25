package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
	"os/exec"

	"github.com/dustin/go-humanize"
	"github.com/go-martini/martini"
	"github.com/olivere/elastic"
)

var (
	esurl string = "http://localhost:8080"
	index string = "s3data"
	debug        = true
)

type Job struct {
	customer, from, to string
}

func parseParams(r *http.Request) (job Job, e error) {
	log.Println(r.URL.Query())
	params := r.URL.Query()
	customer := params.Get("customer")
	from := params.Get("from")
	to := params.Get("to")
	job = Job{customer: customer, from: from, to: to}
	e = nil

	if len(customer) == 0 || len(from) == 0 || len(to) == 0 {
		e = fmt.Errorf("customer, from, to fields are required")
	}
	return
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

//todo:move to DAL
func getFilteredQuery(job Job) elastic.FilteredQuery {
	customerQuery := elastic.NewTermQuery("customer", job.customer)
	filteredQuery := elastic.NewFilteredQuery(customerQuery)
	dateFilter := elastic.NewRangeFilter("@timestamp").From(job.from).To(job.to)
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

		return fmt.Sprintf(
			`{"count":"%v", "size": "%v"}`,
			humanize.Comma(out.Hits.TotalHits),
			humanize.Bytes(uint64(size.(float64))))

	}

	return `{"count": 0, "size": 0}`
}

func startJob(w http.ResponseWriter,
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

	skip := 0
	take := 1000
	var total int64
	total = int64(take)

	for int64(skip) < total {
		out, err := client.Search().
			Index(index).
			From(skip).
			Size(take).
			Query(&filteredQuery).
			Debug(debug).
			Pretty(debug).
			Do()
		total = out.Hits.TotalHits
		skip += take
		if err != nil {
			return showError(w, err)
		}

		cmd := exec.Command("go-send")
		cmdin, err := cmd.StdinPipe()

		if err != nil {
			return showError(w, err)
		}

		for _, hit := range out.Hits.Hits {
			item := make(map[string]interface{})
			json.Unmarshal(*hit.Source, &item)

			uri := strings.TrimLeft(item["uri"].(string), "https://3.amazonaws.com/")
			io.WriteString(cmdin, uri+"\n")
		}

		cmdin.Close()
		cmdout, err := cmd.Output()

		if err != nil {
			return showError(w, err)
		}

		log.Println("go-s3 out: " + string(cmdout))
	}

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

	m := martini.Classic()
	m.Use(martini.Logger())

	m.Get("/job", getJob)
	m.Get("/customers", listCustomers)
	m.Post("/job", startJob)
	//todo:remove this as I understand how to enable post in CUI
	m.Get("/job/create", startJob)
	m.Run()
}
