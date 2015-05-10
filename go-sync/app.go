package main

import (
	"go-indexer/go-sync/sqs"
	"log"
)

func main() {

	sqs := sqs.Sqs{}
	msg, err := sqs.GetMessage()

	if err != nil {
		log.Fatalln(err)
	}

	log.Println(msg)

}
