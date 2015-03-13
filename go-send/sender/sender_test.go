package sender

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	. "go-indexer/testUtils"
	"testing"
)

var queueName = "testQueue11" //todo: plus machine id

func TestSend(t *testing.T) {
	queue, err := GetCleanQueue(queueName)
	if err != nil {
		t.Error(err)
	}
	Send("mybucket/path /to /the /file", queue)
	msg := GetMessages(queue, 1)[0]

	var out map[string]interface{}
	json.Unmarshal([]byte(msg.Body), &out)

	assert.Equal(t, "mybucket", out["bucket"], "")
	assert.Equal(t, "path /to /the /file", out["path"], "")
}
