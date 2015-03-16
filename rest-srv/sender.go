package main

import (
	"encoding/json"
	"go-indexer/go-send/sender"
	"gopkg.in/olivere/elastic.v1"
	"log"
	"strings"
)

var (
	PageSize = 1000
)

func sendJob(job job) {
	log.Println("Sending", job)

	skip := 0
	take := PageSize
	var total int64
	total = int64(take)

	for int64(skip) < total {

		out, err := getFiles(job, skip, take)

		if err != nil {
			log.Println(err)
		}

		total = out.TotalHits
		skip += take

		log.Println(total, skip)

		i := 0

		for _, hit := range out.Hits {

			go func(qn int, h *elastic.SearchHit) {

				item := make(map[string]interface{})

				json.Unmarshal(*h.Source,
					&item)

				uri := strings.TrimPrefix(item["uri"].(string),
					"https://s3.amazonaws.com/")

				q, err := sender.GetQueue(qn)

				if err != nil {
					log.Println(err)
					return
				}

				sender.Send(uri, q)

			}(i, hit)

			i = (i + 1) % sender.NQueues

		}
	}

	log.Println(job, "Done")
}
