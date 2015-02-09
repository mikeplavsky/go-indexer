package main

import (
	"log"

	"github.com/goamz/goamz/aws"
	"github.com/goamz/goamz/sqs"
)

func main() {

	auth, _ := aws.EnvAuth()
	sqs := sqs.New(auth, aws.USEast)

	q, err := sqs.GetQueue("lm-test")

	if err != nil {
		log.Println(err)
	}

	log.Println(q)

	res, err := q.ReceiveMessage(1)

	if err != nil {
		log.Println(err)
	}

	log.Println(res)

}
