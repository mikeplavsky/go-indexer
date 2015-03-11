package main

import (
	"encoding/json"
	"log"
	"strings"

	"go-indexer/go-send/sender"

	"gopkg.in/olivere/elastic.v1"
)

func sendJob(j job) {
	log.Println("Sending", j)

	q, err := sender.GetQueue()

	if err != nil {
		log.Println(err)
	}

	skip := 0
	take := 1000
	var total int64
	total = int64(take)

	for int64(skip) < total {

		out, err := getFiles(j, skip, take)

		if err != nil {
			log.Println(err)
		}

		total = out.TotalHits
		skip += take

		log.Println(total, skip)

		for _, hit := range out.Hits {

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
