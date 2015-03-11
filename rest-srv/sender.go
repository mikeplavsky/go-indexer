package main

import (
	"encoding/json"
	"log"
	"strings"

	"go-indexer/go-send/sender"

	"net/http"

	"gopkg.in/olivere/elastic.v1"
)

func sendJob(j job) {

	log.Println("Sending", j)

	q, err := sender.GetQueue()

	if err != nil {
		log.Println(err)
	}

	client, _ := elastic.NewClient(http.DefaultClient,
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
			Debug(debug).
			Pretty(debug).
			Do()

		if err != nil {
			log.Println(err)
		}

		total = out.Hits.TotalHits
		skip += take

		log.Println(total, skip)

		i := 0

		for _, hit := range out.Hits.Hits {

			go func(h *elastic.SearchHit) {

				item := make(map[string]interface{})

				json.Unmarshal(*h.Source,
					&item)

				uri := strings.TrimPrefix(item["uri"].(string),
					"https://s3.amazonaws.com/")

				sender.Send(uri, q)

			}(hit)

			i = (i + 1) % runtime.NumCPU()

		}
	}

	log.Println(j, "Done")

}
