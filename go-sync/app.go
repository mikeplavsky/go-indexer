package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"go-indexer/go-sync/sqs"
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

func createEsDoc(obj map[string]interface{}) error {

	data, err := json.Marshal(obj)

	if err != nil {
		return err
	}

	path := fmt.Sprintf(
		"http://localhost:8080/s3data/log/%v/_create",
		obj["_id"])

	res, err := http.Post(path,
		"application/json",
		bytes.NewBuffer(data))

	if err != nil {
		return err
	}

	log.Println(res)

	return nil

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

	buf := []string{}

	for {

		s := sqs.Sqs{}
		res, err := s.GetMessage()

		if _, ok := err.(sqs.ErrNoMessages); ok {

			if len(buf) == 0 {
				continue
			}

			buf = []string{}
			continue

		}

		if err != nil {

			log.Println(err)
			continue

		}

		obj, err := parseMessage(res.Body)

		if err != nil {
			log.Println(err)
		} else {
			buf = append(buf, obj)
		}

		s.RemoveMessage(res)

	}

}

func main() {

	go sqs.AuthGen()
	go run()

	w := make(<-chan bool)
	<-w

}
