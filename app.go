package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os/exec"

	"github.com/goamz/goamz/aws"
	"github.com/goamz/goamz/s3"
	"github.com/goamz/goamz/sqs"
)

func main() {

	auth, _ := aws.EnvAuth()
	sqs := sqs.New(auth, aws.USEast)

	q, err := sqs.GetQueue("lm-test")

	if err != nil {
		log.Println(err)
	}

	ps := map[string]string{
		"WaitTimeSeconds":     "10",
		"MaxNumberOfMessages": "1"}

	res, err := q.ReceiveMessageWithParameters(ps)

	if err != nil {
		log.Println(err)
	}

	raw := res.Messages[0].Body

	var msg map[string]interface{}
	err = json.Unmarshal([]byte(raw), &msg)

	if err != nil {
		log.Println(err)
	}

	bucket := msg["bucket"].(string)
	path := msg["path"].(string)

	log.Println(path)

	s3 := s3.New(auth, aws.USEast)
	b := s3.Bucket(bucket)

	data, err := b.Get(path)

	if err != nil {
		log.Println(err)
	}

	f, _ := ioutil.TempFile(
		"",
		"indexer")

	defer f.Close()

	f.Write(data)
	f.Sync()

	log.Println(f.Name())

	exec := func(cmd string) {

		cvrt := exec.Command("bash", "-c", cmd)

		_, err = cvrt.Output()
		if err != nil {
			log.Println()
		}

	}

	exec("unzip -p  " + f.Name() + " | ./go-convert")

	q.DeleteMessage(&res.Messages[0])

}
