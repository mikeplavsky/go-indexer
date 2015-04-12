package main

import (
	"bufio"
	"log"
	"os"
	"runtime"
	"strconv"
	"sync"
	"time"

	"github.com/codegangsta/cli"

	"go-indexer/go-send/sender"
)

func createQueue(qn string) string {

	for {

		log.Println("Creating queue:", qn)

		sqs := sender.GetSqs()
		q, err := sqs.CreateQueue(qn)

		if err != nil {

			log.Println(err)
			time.Sleep(time.Second * 2)

			continue
		}

		res, _ := q.GetQueueAttributes("QueueArn")
		return res.Attributes[0].Value

	}

}

func main() {

	sqs := sender.GetSqs()

	app := cli.NewApp()
	app.Name = "go-send"

	qn := os.Getenv("ES_QUEUE")

	cmds := []cli.Command{
		{
			Name:      "create",
			ShortName: "c",
			Usage:     "creates the queue",
			Action: func(c *cli.Context) {

				n := runtime.NumCPU()

				for i := 0; i < n; i++ {
					createQueue(
						qn + strconv.Itoa(i))
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
