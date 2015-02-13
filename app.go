package main

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"os"
	"os/exec"

	"github.com/goamz/goamz/aws"
	"github.com/goamz/goamz/s3"
	"github.com/goamz/goamz/sqs"
)

var (
	ES_QUEUE,
	ES_INDEXER string
)

func index(q *sqs.Queue, s3 *s3.S3) error {

	ps := map[string]string{
		"WaitTimeSeconds":     "10",
		"MaxNumberOfMessages": "1"}

	res, err := q.ReceiveMessageWithParameters(ps)

	if err != nil {
		return err
	}

	if len(res.Messages) == 0 {
		return errors.New("No messages")
	}

	raw := res.Messages[0].Body

	var msg map[string]interface{}
	err = json.Unmarshal([]byte(raw), &msg)

	if err != nil {
		return err
	}

	bucket := msg["bucket"].(string)
	path := msg["path"].(string)

	b := s3.Bucket(bucket)
	data, err := b.Get(path)

	if err != nil {
		return err
	}

	f, err := ioutil.TempFile(
		"",
		"indexer")

	if err != nil {
		return err
	}

	defer f.Close()

	f.Write(data)
	f.Sync()

	os.Setenv("ES_FILE", f.Name())

	cvrt := exec.Command(ES_INDEXER)
	out, err := cvrt.CombinedOutput()

	log.Println(string(out))

	os.Remove(f.Name())

	if err != nil {
		return err
	}

	q.DeleteMessage(&res.Messages[0])
	return nil

}

func main() {

	ES_QUEUE = os.Getenv("ES_QUEUE")
	ES_INDEXER = os.Getenv("ES_INDEXER")

	auth, _ := aws.EnvAuth()
	sqs := sqs.New(auth, aws.USEast)

	q, err := sqs.GetQueue(ES_QUEUE)

	if err != nil {
		log.Println(err)
	}

	log.Println(q)

	s3 := s3.New(auth, aws.USEast)

	if err != nil {
		log.Println(err)
	}

	for {
		if err := index(q, s3); err != nil {
			log.Println(err)
		}
	}

}
