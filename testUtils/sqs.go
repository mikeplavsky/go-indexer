package sender

import (
	"github.com/goamz/goamz/aws"
	"github.com/goamz/goamz/sqs"
	"log"
	"time"
)

var (
	maxGetCleanQueueAttempts = 100
	maxGetMessagesAttempts   = 100
	defaultRequestTimeout    = time.Duration(100) * time.Millisecond
)

// GetCleanQueue creates new queue if it does not exist, purges the messages
func GetCleanQueue(queueName string) (*sqs.Queue, error) {
	//create own ctor, not app one
	auth, _ := aws.GetAuth("", "", "", time.Time{})
	sqs := sqs.New(auth, aws.USEast)

	//wait one second to wait all messages in flight switched to active state
	time.Sleep(time.Duration(1) * time.Second)
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

	//todo:parallel
	attempts := 0
	for len(resp.Messages) > 0 && attempts < maxGetCleanQueueAttempts {
		attempts++
		log.Printf("deleting %d messages", len(resp.Messages))
		_, err = queue.DeleteMessageBatch(resp.Messages)
		if err != nil {
			return nil, err
		}
		resp, err = queue.ReceiveMessage(10)
		if len(resp.Messages) == 0 {
			time.Sleep(defaultRequestTimeout)
		}
		if err != nil {
			return nil, err
		}
	}
	return queue, nil
}

// GetMessages returns top messages with specified count
func GetMessages(queue *sqs.Queue, count int) []sqs.Message {
	ret := make([]sqs.Message, count)

	i := 0
	attempts := 0
	for i < count && attempts < maxGetMessagesAttempts {
		attempts++
		resp, _ := queue.ReceiveMessage(1)
		if len(resp.Messages) > 0 {
			ret[i] = resp.Messages[0]
			i++
		} else {
			time.Sleep(defaultRequestTimeout)
		}
	}

	return ret
}

func Wait(
	f func() (interface{}, error),
	expected interface{},
	timeout time.Duration,
	maxAttempts int) interface{} {

	attempts := 0

	for attempts < maxAttempts {

		attempts++

		objs, err := f()

		if err == nil && objs == expected {
			return objs
		}

		log.Println(err, objs)
		time.Sleep(timeout)

	}

	return nil
}
