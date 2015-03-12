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

	doneCh := make(chan bool)
	totalReceived := false
	for int64(skip) < total {

		out, err := getFiles(j, skip, take)

		if err != nil {
			log.Println(err)
		}

		total = out.TotalHits
		skip += take

		if !totalReceived {
			doneCh = make(chan bool, total)
			totalReceived = true
		}

		log.Println(total, skip)
		for _, hit := range out.Hits {
			func(*elastic.SearchHit) {
				item := make(map[string]interface{})

				json.Unmarshal(*hit.Source,
					&item)

				uri := strings.TrimPrefix(item["uri"].(string),
					"https://s3.amazonaws.com/")
				sender.Send(uri, q)
				doneCh <- true
			}(hit)
		}
	}
	for i := 0; i < int(total); i++ {
		<-doneCh
	}
	log.Println(j, "Done")
}
