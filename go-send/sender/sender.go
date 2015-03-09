package sender

import (
	"encoding/json"
	"log"
	"os"
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

func GetQueue() (*sqs.Queue, error) {

	return GetSqs().
		GetQueue(GetQueueName())

}

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
