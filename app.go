package main

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strconv"
	"time"

	"go-indexer/go-sync/sqs"

	"github.com/goamz/goamz/aws"
	"github.com/goamz/goamz/s3"
)

var (
	S3_PATH,
	ES_QUEUE,
	ES_FS_PER_INDEX,
	ES_INDEX,
	ES_INDEXER string
)

var queueMaxWaitTimeSeconds = 10

type awsIdx interface {
	sqs.SqsA

	getLog(bucket string,
		path string) ([]byte, error)

	exec(string) ([]byte, error)
}

type idx struct {
	sqs.Sqs
}

func getAuth() (aws.Auth, error) {
	return aws.GetAuth("", "", "", time.Time{})
}

func (idx) exec(c string) ([]byte, error) {

	cvrt := exec.Command(c)
	return cvrt.CombinedOutput()

}

func (idx) getLog(bucket, path string) ([]byte, error) {

	auth, err := getAuth()

	if err != nil {
		return nil, err
	}

	s3 := s3.New(auth, aws.USEast)

	if err != nil {
		return nil, err
	}

	b := s3.Bucket(bucket)
	data, err := b.Get(path)

	if err != nil {
		return nil, err
	}

	return data, nil

}

func index(i awsIdx) error {

	res, err := i.GetMessage()

	if err != nil {
		return err
	}

	raw := res.Body

	var msg map[string]interface{}
	err = json.Unmarshal([]byte(raw), &msg)

	if err != nil {
		return err
	}

	bucket := msg["bucket"].(string)
	path := msg["path"].(string)

	data, err := i.getLog(bucket, path)

	if err != nil {
		return err
	}

	S3_PATH = fmt.Sprintf(
		"https://s3.amazonaws.com/%v/%v",
		bucket,
		path)

	os.Setenv("S3_PATH", S3_PATH)

	h := md5.New()
	io.WriteString(h, S3_PATH)

	S3_FILE_ID := fmt.Sprintf("%x", h.Sum(nil))

	os.Setenv("S3_FILE_ID", S3_FILE_ID)

	f, err := ioutil.TempFile(
		"",
		"indexer")

	if err != nil {
		return err
	}

	defer f.Close()
	defer os.Remove(f.Name())

	f.Write(data)
	f.Sync()

	os.Setenv("ES_FILE", f.Name())

	out, err := i.exec(ES_INDEXER)
	log.Println(string(out))

	if err != nil {
		return err
	}

	return i.RemoveMessage(res)
}

func setVars() {

	ES_QUEUE = os.Getenv("ES_QUEUE")
	ES_INDEXER = os.Getenv("ES_INDEXER")
	ES_INDEX = os.Getenv("ES_INDEX")
	ES_FS_PER_INDEX = os.Getenv("ES_FS_PER_INDEX")
}

func nextIdx() <-chan int {

	res := make(chan int)

	go func() {

		currIdx := 0
		perIdx, _ := strconv.Atoi(ES_FS_PER_INDEX)

		for i := 1; ; i++ {

			res <- currIdx

			if i%perIdx == 0 && i != 0 {
				currIdx += 1
			}

		}
	}()

	return res

}

func run() {

	setVars()
	i := nextIdx()

	currIdx := 0

	for {

		idxName := fmt.Sprintf(
			"test%v_%v",
			currIdx,
			ES_INDEX)

		os.Setenv("ES_INDEX", idxName)

		if err := index(idx{}); err != nil {

			log.Println(err)
			continue

		}

		currIdx = <-i

	}

}

func main() {
	run()
}
