package main

import (
	"encoding/json"
	"go-indexer/go-send/sender"
	. "go-indexer/testUtils"
	"os"
	"testing"

	"github.com/goamz/goamz/sqs"
	"github.com/stretchr/testify/assert"
	"gopkg.in/olivere/elastic.v1"
)

var queueName = "testQueue11" //todo: plus machine id

func SetUp() *sqs.Queue {
	queue, _ := GetCleanQueue(queueName + "0")
	sender.NQueues = 1
	os.Setenv("ES_QUEUE", queueName)
	sender.Init()
	return queue
}

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

	filesJSONPage1 = []byte(`{
		"total": 2,
        "hits": [
            {
                "_source": {
                    "uri": "https://s3.amazonaws.com/mybucket/path/page1.zip"
                }
            }
        ]
    }`)

	filesJSONPage2 = []byte(`{
		"total": 2,
        "hits": [
            {
                "_source": {
                    "uri": "https://s3.amazonaws.com/mybucket/path/page2.zip"
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

	queue := SetUp()
	sendJob(job{})

	// there is no set datatype in stdlib
	expectedPaths := []string{"path/1.zip", "path/2.zip"}

	messages := GetMessages(queue, len(expectedPaths))

	assert.Equal(t, len(expectedPaths), len(messages))

	paths := []string{}
	for _, msg := range messages {
		var out map[string]interface{}
		json.Unmarshal([]byte(msg.Body), &out)
		assert.Equal(t, "mybucket", out["bucket"], "")
		paths = append(paths, out["path"].(string))
	}
	AssertSetsAreEqual(t, expectedPaths, paths)
}

// Sends two different messages in single queue
// and checks that all messages has been delivered
func TestStartJob_Paging(t *testing.T) {
	PageSize = 1
	getFiles = func(job job, skip int, take int) (h *elastic.SearchHits, err error) {
		var hits elastic.SearchHits

		switch skip {
		case 0:
			err = json.Unmarshal(filesJSONPage1, &hits)
			break
		case 1:
			err = json.Unmarshal(filesJSONPage2, &hits)
			break
		}
		return &hits, err
	}

	queue := SetUp()
	sendJob(job{})

	// there is no set datatype in stdlib
	expectedPaths := []string{"path/page1.zip", "path/page2.zip"}

	messages := GetMessages(queue, len(expectedPaths))

	assert.Equal(t, len(expectedPaths), len(messages))

	paths := []string{}
	for _, msg := range messages {
		var out map[string]interface{}
		json.Unmarshal([]byte(msg.Body), &out)
		assert.Equal(t, "mybucket", out["bucket"], "")
		paths = append(paths, out["path"].(string))
	}
	AssertSetsAreEqual(t, expectedPaths, paths)
}
