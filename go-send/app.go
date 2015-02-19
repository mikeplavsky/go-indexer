package main

import (
	"bufio"
	"encoding/json"
	"log"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/goamz/goamz/aws"
	"github.com/goamz/goamz/sqs"
)

func send(s3path string, q *sqs.Queue) {

	ps := strings.SplitN(s3path, "/", 2)

	msg := map[string]string{
		"bucket": ps[0],
		"path":   ps[1],
	}

	res, _ := json.Marshal(msg)

	for {

		_, err := q.SendMessage(string(res))

		if err != nil {
			time.Sleep(time.Second * 2)
			continue
		}

		return

	}

}

func main() {

	auth, _ := aws.GetAuth("", "", "", time.Time{})
	sqs := sqs.New(auth, aws.USEast)

	q, err := sqs.GetQueue(os.Getenv("ES_QUEUE"))

	if err != nil {
		log.Println(err)
		return
	}

	scanner := bufio.NewScanner(os.Stdin)
	var w sync.WaitGroup

	for scanner.Scan() {

		s3path := scanner.Text()
		w.Add(1)

		go func(s string) {

			send(s, q)
			w.Done()

		}(s3path)

	}
	w.Wait()
}
