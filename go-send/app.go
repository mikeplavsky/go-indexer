package main

import (
	"bufio"
	"encoding/json"
	"log"
	"os"
	"runtime"
	"strconv"
	"sync"
	"time"

	"github.com/codegangsta/cli"

	"go-indexer/go-send/sender"
)

var ES_QUEUE string

func createQueue(qn string,
	attrs map[string]string) string {

	for {

		log.Println("Creating queue:", qn)

		sqs := sender.GetSqs()
		q, err := sqs.CreateQueueWithAttributes(
			qn, attrs)

		if err != nil {

			log.Println(err)
			time.Sleep(time.Second * 2)

			continue
		}

		res, _ := q.GetQueueAttributes("QueueArn")
		arn := res.Attributes[0].Value

		log.Println(arn)

		return arn

	}

}

func createQueues() {

	arn := createQueue(ES_QUEUE+"_dl",
		map[string]string{})

	n := runtime.NumCPU()

	rd := map[string]string{

		"maxReceiveCount":     "5",
		"deadLetterTargetArn": arn,
	}

	res, _ := json.Marshal(rd)

	attrs := map[string]string{

		"VisibilityTimeout": "30",
		"RedrivePolicy":     string(res),
	}

	for i := 0; i < n; i++ {
		createQueue(
			ES_QUEUE+strconv.Itoa(i),
			attrs)
	}

}

func main() {

	ES_QUEUE = os.Getenv("ES_QUEUE")
	sqs := sender.GetSqs()

	app := cli.NewApp()
	app.Name = "go-send"

	cmds := []cli.Command{
		{
			Name:      "create",
			ShortName: "c",
			Usage:     "creates the queue",
			Action: func(c *cli.Context) {
				createQueues()
			},
		},
		{
			Name:      "send",
			ShortName: "s",
			Usage:     "sends messages to the queue",
			Action: func(c *cli.Context) {

				q, err := sqs.GetQueue(ES_QUEUE)

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

						sender.Send(s, q)
						w.Done()

					}(s3path)

				}
				w.Wait()
			},
		},
	}

	app.Commands = cmds

	app.Action = func(c *cli.Context) {
		cli.ShowAppHelp(c)
	}

	app.Run(os.Args)

}
