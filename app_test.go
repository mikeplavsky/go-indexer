package main

import (
	"encoding/json"
	"github.com/goamz/goamz/sqs"
	"github.com/stretchr/testify/assert"
	"go-indexer/go-send/sender"
	. "go-indexer/testUtils"
	"testing"
)

var queueName = "testQueue15" //todo: plus machine id

func SetUp() *sqs.Queue {
	queue, _ := GetCleanQueue(queueName)
	sender.NQueues = 1
	ES_QUEUE = queueName
	//todo:use real indexer
	ES_INDEXER = "echo"
	sender.Init()
	return queue
}

func TestIndex(t *testing.T) {
	queueMaxWaitTimeSeconds = 1
	queue := SetUp()

	JSON, _ := json.Marshal(map[string]string{
		"bucket": "dmp-log-analysis",
		"path":   "Fuji/Lib/Calendar00-20150302071320.zip",
	})
	queue.SendMessage(string(JSON))
	err := index()

	attr, _ := queue.GetQueueAttributes("ApproximateNumberOfMessages")
	num := attr.Attributes[0].Value

	assert.Nil(t, err)
	assert.Equal(t, "0", num, "queue is not empty")
}
