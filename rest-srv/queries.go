package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"gopkg.in/olivere/elastic.v1"
)

func newConnection() (*elastic.Client, error) {
	return elastic.NewClient(http.DefaultClient, esurl)
}

func getCustomers() ([]string, error) {
	conn, err := newConnection()
	customerTermsAggr := elastic.NewTermsAggregation().Field("customer")

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
