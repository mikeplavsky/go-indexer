package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"go-indexer/go-sync/sqs"
	"log"
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

func parseMessage(raw string) (
	string,
	error) {

	var msg Event
	err := json.Unmarshal([]byte(raw), &msg)

	if err != nil {
		return "", err
	}

	var rs Message
	err = json.Unmarshal([]byte(msg.Message), &rs)

	if err != nil {
		return "", err
	}

	if len(rs.Records) == 0 {
		return "", errors.New("wrong event")
	}

	r := rs.Records[0]

	l := fmt.Sprintf("%v\t%v/%v",
		r.S3.Object.Size,
		r.S3.Bucket.Name,
		r.S3.Object.Key)

	return l, nil

}

func run() {

	for {

		s := sqs.Sqs{}
		res, err := s.GetMessage()

		if err != nil {

			log.Println(err)
			return

		}

		obj, err := parseMessage(res.Body)

		if err != nil {
			log.Println(err)
		} else {
			fmt.Println(obj)
		}

		s.RemoveMessage(res)

	}

}

func main() {
	run()
}
