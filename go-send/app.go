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

	"github.com/codegangsta/cli"
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

	app := cli.NewApp()
	app.Name = "go-send"

	qn := os.Getenv("ES_QUEUE")

	cmds := []cli.Command{
		{
			Name:      "create",
			ShortName: "c",
			Usage:     "creates the queue",
			Action: func(c *cli.Context) {

				_, err := sqs.CreateQueue(qn)

				if err != nil {
					log.Println(err)
					return
				}
			},
		},
		{
			Name:      "send",
			ShortName: "s",
			Usage:     "sends messages to the queue",
			Action: func(c *cli.Context) {

				q, err := sqs.GetQueue(qn)

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
			},
		},
	}

	app.Commands = cmds

	app.Action = func(c *cli.Context) {
		cli.ShowAppHelp(c)
	}

	app.Run(os.Args)

}
