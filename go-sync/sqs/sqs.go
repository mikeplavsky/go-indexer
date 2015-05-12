package sqs

import (
	"errors"
	"os"
	"time"

	"github.com/goamz/goamz/aws"
	"github.com/goamz/goamz/sqs"
)

type SqsA interface {
	GetMessage() (*sqs.Message, error)
	RemoveMessage(*sqs.Message) error
}

type Sqs struct{}

func getAuth() (aws.Auth, error) {
	return aws.GetAuth("", "", "", time.Time{})
}

func getQueue() (*sqs.Queue, error) {

	auth, err := getAuth()

	if err != nil {
		return nil, err
	}

	sqs := sqs.New(auth, aws.USEast)

	ES_QUEUE := os.Getenv("ES_QUEUE")
	return sqs.GetQueue(ES_QUEUE)

}

func (Sqs) GetMessage() (*sqs.Message, error) {

	q, err := getQueue()

	if err != nil {
		return nil, err
	}

	ps := map[string]string{
		"WaitTimeSeconds":     "20",
		"MaxNumberOfMessages": "1"}

	res, err := q.ReceiveMessageWithParameters(ps)

	if err != nil {
		return nil, err
	}

	if len(res.Messages) == 0 {
		return nil, errors.New("No messages")
	}

	return &res.Messages[0], nil

}

func (Sqs) RemoveMessage(msg *sqs.Message) error {

	q, err := getQueue()

	if err != nil {
		return err
	}

	_, err = q.DeleteMessage(msg)
	return err

}
