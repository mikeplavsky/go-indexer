package main

import (
	"encoding/json"
	"go-indexer/go-sync/sqs"
	"log"
)

type S3 struct {
	Bucket map[string]interface{} `json:"bucket"`
	Object map[string]interface{} `json:"object"`
}

type Record struct {
	S3 S3 `json:"s3"`
}

type Message struct {
	Records []Record
}

type Event struct {
	Message string
}

func main() {

	sqs := sqs.Sqs{}
	res, err := sqs.GetMessage()

	if err != nil {
		log.Fatalln(err)
	}

	raw := res.Body

	var msg Event
	err = json.Unmarshal([]byte(raw), &msg)

	if err != nil {
		log.Fatalln(err)
	}

	var rs Message
	err = json.Unmarshal([]byte(msg.Message), &rs)

	if err != nil {
		log.Fatalln(err)
	}

	log.Println(rs)

}
