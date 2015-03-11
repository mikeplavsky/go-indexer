package sender

import (
	"encoding/json"
	"log"
	"os"
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

var queues = map[int]*sqs.Queue{}

func GetQueue(i int) (*sqs.Queue, error) {

	var err error = nil

	if _, ok := queues[i]; !ok {

		qn := GetQueueName() + strconv.Itoa(i)
		log.Printf("Getting queue %v", qn)

		queues[i], err = GetSqs().GetQueue(qn)

	}

	return queues[i], err

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
