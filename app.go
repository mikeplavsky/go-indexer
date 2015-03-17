package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strconv"
	"time"

	"github.com/goamz/goamz/aws"
	"github.com/goamz/goamz/s3"
	"github.com/goamz/goamz/sqs"
)

var (
	S3_PATH,
	ES_QUEUE,
	ES_FS_PER_INDEX,
	ES_INDEX,
	ES_INDEXER string
)

var queueMaxWaitTimeSeconds = 10

func index() error {

	auth, err := aws.GetAuth("", "", "", time.Time{})

	if err != nil {
		return err
	}

	sqs := sqs.New(auth, aws.USEast)
	q, err := sqs.GetQueue(ES_QUEUE)

	if err != nil {
		return err
	}

	s3 := s3.New(auth, aws.USEast)

	if err != nil {
		return err
	}

	ps := map[string]string{
		"WaitTimeSeconds":     strconv.Itoa(queueMaxWaitTimeSeconds),
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

	S3_PATH = fmt.Sprintf(
		"https://s3.amazonaws.com/%v/%v",
		bucket,
		path)

	os.Setenv("S3_PATH", S3_PATH)

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
	ES_INDEX = os.Getenv("ES_INDEX")
	ES_FS_PER_INDEX = os.Getenv("ES_FS_PER_INDEX")

	currIdx := 0
	perIdx, _ := strconv.Atoi(ES_FS_PER_INDEX)

	for i := 0; ; i++ {

		if i%perIdx == 0 && i != 0 {
			currIdx += 1
		}

		idxName := fmt.Sprintf(
			"test%v_%v",
			currIdx,
			ES_INDEX)

		os.Setenv("ES_INDEX", idxName)

		if err := index(); err != nil {
			log.Println(err)
		}

	}

}
