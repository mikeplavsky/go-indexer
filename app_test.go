package main

import (
	"encoding/json"
	"github.com/goamz/goamz/sqs"
	"github.com/stretchr/testify/assert"
	"go-indexer/go-send/sender"
	. "go-indexer/testUtils"
	"os"
	"testing"
)

var queueName = "testQueue15" //todo: plus machine id

func SetUp() *sqs.Queue {
	queueMaxWaitTimeSeconds = 1
	queue, _ := GetCleanQueue(queueName)
	sender.NQueues = 1
	ES_QUEUE = queueName
	//todo:use real indexer
	ES_INDEXER = "echo"
	sender.Init()
	return queue
}

func TestIndex(t *testing.T) {
	queue := SetUp()

	JSON, _ := json.Marshal(map[string]string{
		"bucket": "dmp-log-analysis",
		"path":   "Fuji/Lib/Calendar00-20150302071320.zip",
	})
	queue.SendMessage(string(JSON))
	err := index()

	queueSize := WaitMessagesInQueue(queue, "0")

	assert.Nil(t, err)
	assert.Equal(t, "0", queueSize, "queue is not empty")
	assert.Equal(t, "https://s3.amazonaws.com/dmp-log-analysis/Fuji/Lib/Calendar00-20150302071320.zip", os.Getenv("S3_PATH"))
	assert.True(t, fileExists(os.Getenv("ES_FILE")))
}

func TestIndex_ShouldNotDeleteMessageIfFailed(t *testing.T) {
	queue := SetUp()

	JSON, _ := json.Marshal(map[string]string{
		"bucket": "dmp-log-analysis",
		"path":   "not/existing/path",
	})
	queue.SendMessage(string(JSON))

	err := index()

	queueSize := WaitMessagesInQueue(queue, "1")
	assert.NotNil(t, err)
	assert.Equal(t, "1", queueSize, "message should not be deleted from queue")
}

func WaitMessagesInQueue(queue *sqs.Queue, count string) string {
	f := func() (num interface{}, err error) {
		attr, err := queue.GetQueueAttributes("ApproximateNumberOfMessages")
		num = attr.Attributes[0].Value
		return num, err
	}

	res := Wait(f, count, 1000, 100).(string)
	return res
}

func fileExists(name string) bool {
	_, err := os.Stat(name)
	return !os.IsNotExist(err)
}
