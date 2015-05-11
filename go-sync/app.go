package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"go-indexer/go-sync/sqs"
	"go-indexer/s3-2-es/parser"
	"log"
	"net/http"
)

type Object struct {
	Key  string
	Size int64
}

type Bucket struct {
	Name string
}

type S3 struct {
	Bucket Bucket `json:"bucket"`
	Object Object `json:"object"`
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

func createEsDoc(obj map[string]interface{}) {

	data, err := json.Marshal(obj)

	if err != nil {

		log.Println(err)
		return

	}

	path := fmt.Sprintf(
		"http://localhost:8080/s3data/log/%v/_create",
		obj["_id"])

	_, err = http.Post(path,
		"application/json",
		bytes.NewBuffer(data))

	if err != nil {
		log.Println(err)
	}

}

func parseMessage(raw string) (
	map[string]interface{},
	error) {

	var msg Event
	err := json.Unmarshal([]byte(raw), &msg)

	if err != nil {
		return nil, err
	}

	var rs Message
	err = json.Unmarshal([]byte(msg.Message), &rs)

	if err != nil {
		return nil, err
	}

	r := rs.Records[0]

	l := fmt.Sprintf("%v\t%v/%v",
		r.S3.Object.Size,
		r.S3.Bucket.Name,
		r.S3.Object.Key)

	return parser.ParseLine(l)

}

func run() {

	for {

		sqs := sqs.Sqs{}
		res, err := sqs.GetMessage()

		if err != nil {

			log.Println(err)
			continue

		}

		obj, err := parseMessage(res.Body)

		if err != nil {
			log.Println(err)
		} else {
			createEsDoc(obj)
		}

		sqs.RemoveMessage(res)

	}

}

func main() {
	run()
}
