package main

import (
	"encoding/json"
	"go-indexer/go-send/sender"
	"log"
	"strings"
	"sync"

	"gopkg.in/olivere/elastic.v1"
)

var (
	// PageSize is max hits size to go-sender per goroutine
	PageSize = 1000
)

type queue interface {
	send(int, string)
	qNum() int
}

type q struct{}

func (q) send(qn int, uri string) {

	q, err := sender.GetQueue(qn)

	if err != nil {
		log.Println(err)
		return
	}

	sender.Send(uri, q)
}

func (q) qNum() int {
	return sender.NQueues
}

func sendJob(job job) {

	log.Println("Sending", job)
	sendJobImpl(job, q{})
}

func sendJobImpl(job job, q queue) error {

	skip := 0
	take := PageSize
	total := int64(take)

	var w sync.WaitGroup

	for int64(skip) < total {

		out, err := getFiles(job, skip, take)

		if err != nil {
			return err
		}

		total = out.TotalHits
		skip += take

		log.Println(total, skip)

		i := 0

		for _, hit := range out.Hits {

			w.Add(1)

			go func(qn int, h *elastic.SearchHit) {

				item := make(map[string]interface{})

				json.Unmarshal(*h.Source,
					&item)

				uri := strings.TrimPrefix(item["uri"].(string),
					"https://s3.amazonaws.com/")

				q.send(qn, uri)

				w.Done()

			}(i, hit)

			i = (i + 1) % q.qNum()

		}
	}

	w.Wait()
	return nil

}
