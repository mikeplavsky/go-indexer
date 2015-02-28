package main

import (
	"encoding/json"
	"log"
	"strings"

	"go-indexer/go-send/sender"

	"github.com/olivere/elastic"
)

func sendJob(j job) {

	log.Println("Sending", j)

	q, err := sender.GetQueue()

	if err != nil {
		log.Println(err)
	}

	client, _ := elastic.NewClient(
		esurl)

	filteredQuery := getFilteredQuery(j)

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
			Pretty(debug).
			Do()

		if err != nil {
			log.Println(err)
		}

		total = out.Hits.TotalHits
		skip += take

		log.Println(total, skip)

		for _, hit := range out.Hits.Hits {

			go func(*elastic.SearchHit) {

				item := make(map[string]interface{})

				json.Unmarshal(*hit.Source,
					&item)

				uri := strings.TrimPrefix(item["uri"].(string),
					"https://s3.amazonaws.com/")

				sender.Send(uri, q)

			}(hit)

		}
	}

	log.Println(j, "Done")

}
