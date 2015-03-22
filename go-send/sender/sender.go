package sender

import (
	"encoding/json"
	"log"
	"os"
	"runtime"
	"strconv"
	"strings"
	"time"

	"github.com/goamz/goamz/aws"
	"github.com/goamz/goamz/sqs"
)

func GetSqs() *sqs.SQS {

	auth, _ := aws.GetAuth("", "", "", time.Time{})
	return sqs.New(auth, aws.USEast)

}

func GetQueueName() (qn string) {

	qn = os.Getenv("ES_QUEUE")
	return

}

type queue struct {
	q   *sqs.Queue
	err error
}

var queues = map[int]queue{}

// NQueues is the number of target queues
var NQueues = runtime.NumCPU()

func init() {
	//Init()
}

// Init creates queues
// todo: init queuename other way, without reading env vars on start
func Init() {
	for i := 0; i < NQueues; i++ {

		qn := GetQueueName() + strconv.Itoa(i)
		log.Printf("Getting queue %v", qn)

		q, err := GetSqs().GetQueue(qn)
		queues[i] = queue{q, err}

	}
}

// GetQueue returns Queue with specific ID
func GetQueue(i int) (*sqs.Queue, error) {
	return queues[i].q, queues[i].err
}

// Send sends given path on S3 (bucket/path) to specified queue
func Send(s3path string, q *sqs.Queue) {

	ps := strings.SplitN(s3path, "/", 2)

	msg := map[string]string{
		"bucket": ps[0],
		"path":   ps[1],
	}

	res, _ := json.Marshal(msg)

	for {

		_, err := q.SendMessage(string(res))

		if err != nil {

			log.Println(err)

			time.Sleep(time.Second * 2)
			continue
		}

		return

	}

}
