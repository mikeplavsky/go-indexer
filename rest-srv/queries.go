package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"gopkg.in/olivere/elastic.v1"
)

var (
	esurl = "http://localhost:8080"
	index = "s3data"
	debug = false
)

const CUSTOMER_LIMIT = 1000

type job struct {
	customer, from, to string
}

func newConnection() (*elastic.Client, error) {
	return elastic.NewClient(http.DefaultClient, esurl)
}

var getCustomers = func () ([]string, error) {
	conn, err := newConnection()

	customerTermsAggr := elastic.NewTermsAggregation().Field("customer").Size(CUSTOMER_LIMIT)

	out, err := conn.Search().
		Index(index).
		Aggregation("cust_unique", customerTermsAggr).
		Debug(debug).
		Pretty(debug).
		Do()

	if err != nil {
		return nil, err
	}

	if out.Hits != nil {
		var aggrResult map[string]interface{}

		err = json.Unmarshal(
			out.Aggregations["cust_unique"],
			&aggrResult)

		if err != nil {
			return nil, err
		}

		buckets := aggrResult["buckets"].([]interface{})

		ret := make([]string, len(buckets))
		for i, bucket := range buckets {
			item := bucket.(map[string]interface{})
			ret[i] = item["key"].(string)
		}

		return ret, nil
	}
	//todo:check on empty db
	return nil, fmt.Errorf("no hits")
}

func getFilteredQuery(j job) elastic.FilteredQuery {

	customerQuery := elastic.NewTermQuery("customer", j.customer)
	filteredQuery := elastic.NewFilteredQuery(customerQuery)

	dateFilter := elastic.NewRangeFilter("@timestamp").
		From(j.from).
		To(j.to)

	filteredQuery = filteredQuery.Filter(dateFilter)

	return filteredQuery
}

func getJobStats(j job) (map[string]uint64, error) {
	conn, err := newConnection()

	if err != nil {
		return nil, err
	}

	filteredQuery := getFilteredQuery(j)

	sizeSumAggr := elastic.NewSumAggregation().Field("size")

	out, err := conn.Search().
		Index(index).
		Query(&filteredQuery).
		Aggregation("sum", sizeSumAggr).
		Debug(debug).
		Pretty(debug).
		Do()

	if err != nil {
		return nil, err
	}

	if out.Hits != nil {

		var aggrResult map[string]interface{}

		err = json.Unmarshal(
			out.Aggregations["sum"],
			&aggrResult)

		if err != nil {
			return nil, err
		}

		return map[string]uint64{
			"count": uint64(out.Hits.TotalHits),
			"size":  uint64(aggrResult["value"].(float64)),
		}, nil

	}

	return nil, fmt.Errorf("no hits")
}
