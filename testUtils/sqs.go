package sender

import (
	"github.com/goamz/goamz/aws"
	"github.com/goamz/goamz/sqs"
	"log"
	"time"
)

func GetCleanQueue(queueName string) (*sqs.Queue, error) {
	//create own ctor, not app one
	auth, _ := aws.GetAuth("", "", "", time.Time{})
	sqs := sqs.New(auth, aws.USEast)

	_, err := sqs.CreateQueueWithTimeout(queueName, 1)
	if err != nil {
		return nil, err
	}

	queue, err := sqs.GetQueue(queueName)
	if err != nil {
		return nil, err
	}

	//Purge is available only one time in 60 seconds
	//_, err = queue.Purge()
	resp, err := queue.ReceiveMessage(10)
	if err != nil {
		return nil, err
	}

	for len(resp.Messages) > 0 {
		log.Printf("deleting %d messages", len(resp.Messages))
		_, err = queue.DeleteMessageBatch(resp.Messages)
		if err != nil {
			return nil, err
		}
		resp, err = queue.ReceiveMessage(10)
		if err != nil {
			return nil, err
		}
	}
	return queue, nil
}

func GetMessages(queue *sqs.Queue, count int) []sqs.Message {
	ret := make([]sqs.Message, count)

	i := 0
	for i < count {
		resp, _ := queue.ReceiveMessage(1)
		if len(resp.Messages) > 0 {
			ret[i] = resp.Messages[0]
			i++
		}
	}

	return ret
}
