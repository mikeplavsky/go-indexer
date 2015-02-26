package main

import (
	"go-indexer/go-send/sender"
	"log"
	"net/http"
)

func getEta() (int, string) {

	q, err := sender.GetQueue()

	if err != nil {

		return http.StatusInternalServerError,
			err.Error()

	}

	res, err := q.GetQueueAttributes("ApproximateNumberOfMessages")

	if err != nil {

		return http.StatusInternalServerError,
			err.Error()

	}

	log.Println(res)

	return http.StatusOK, `{"files": 12, "time" : "2h"}`

}
