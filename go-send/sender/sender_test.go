package sender

import (
	"encoding/json"
	"github.com/goamz/goamz/sqs"
	"github.com/stretchr/testify/assert"
	"testing"
)

var queueName = "testQueue12" //todo: plus machine guid
var queue (*sqs.Queue)

func checkErr(t *testing.T, err error) {
	if err != nil {
		t.Error(t)
	}
}

func TestMain(t *testing.T) {
	//create own ctor, not app one
	sqs := GetSqs()

	_, err := sqs.CreateQueue(queueName)
	checkErr(t, err)

	queue, err = sqs.GetQueue(queueName)
        checkErr(t, err)

	//Purge is available only one time in 60 seconds
	//_, err = queue.Purge()
	resp, err := queue.ReceiveMessage(10)
	checkErr(t, err)
	
	if(len(resp.Messages) > 0) {
		queue.DeleteMessageBatch(resp.Messages)
	}	
}

func TestInvalidJobParameters(t *testing.T) {
	Send("mybucket/path /to /the /file", queue)
	resp, _ := queue.ReceiveMessage(1)
	msg := resp.Messages[0]

	var out map[string]interface{}
	json.Unmarshal([]byte(msg.Body), &out)

	assert.Equal(t, "mybucket", out["bucket"], "")
	assert.Equal(t, "path /to /the /file", out["path"], "")
}
