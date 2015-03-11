package main

import (
	"bufio"
	"log"
	"os"
	"sync"

	"github.com/codegangsta/cli"

	"go-indexer/go-send/sender"
)

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

				_, err := sqs.CreateQueue(qn)

				if err != nil {
					log.Println(err)
					return
				}

				_, err = sqs.CreateQueue(qn + "-error")

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
