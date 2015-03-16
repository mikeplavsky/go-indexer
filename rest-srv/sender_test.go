package main

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"go-indexer/go-send/sender"
	. "go-indexer/testUtils"
	"gopkg.in/olivere/elastic.v1"
	"os"
	"testing"
)

var queueName = "testQueue11" //todo: plus machine id

var (
	filesJSON = []byte(`{
        "hits": [
            {
                "_source": {
                    "uri": "mybucket/path/1.zip"
                }
            },
            {
                "_source": {
                    "uri": "https://s3.amazonaws.com/mybucket/path/2.zip"
                }
            }
        ]
    }`)
)

// Sends two different messages in single queue
// and checks that all messages has been delivered
func TestStartJob_DifferentMessagesUpload(t *testing.T) {
	getFiles = func(job job, skip int, take int) (h *elastic.SearchHits, err error) {
		var hits elastic.SearchHits
		err = json.Unmarshal(filesJSON, &hits)
		return &hits, err
	}

	queue, _ := GetCleanQueue(queueName + "0")
	sender.NQueues = 1
	os.Setenv("ES_QUEUE", queueName)
	sender.Init()
	sendJob(job{})

	messages := GetMessages(queue, 2)

	// there is no set datatype in stdlib
	expectedPaths := []string{"path/1.zip", "path/2.zip"}

	assert.Equal(t, 2, len(messages))

	paths := []string{}
	for index := range expectedPaths {
		msg := messages[index]
		var out map[string]interface{}
		json.Unmarshal([]byte(msg.Body), &out)
		assert.Equal(t, "mybucket", out["bucket"], "")
		paths = append(paths, out["path"].(string))
	}
	assertSetsAreEqual(t, expectedPaths, paths)
}

// test utils
// verifies that arrays with distinct values have the same values, excluding order
// it should not modify order
func assertSetsAreEqual(t *testing.T, expected []string, actual []string) {
	assert.Equal(t, len(expected), len(actual), "length are not equal")

	var actualSet = map[string]bool{}
	for _, v := range actual {
		actualSet[v] = true
	}

	for _, v := range expected {
		if !actualSet[v] {
			t.Errorf("%s is not found in %s", v, actual)
		}
	}
}
