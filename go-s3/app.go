package main

import (
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof"
	"os"
	"runtime"
	"strings"
	"time"

	"github.com/codegangsta/cli"
	"github.com/dustin/go-humanize"
)

func l(o ...interface{}) {
	log.Println(o...)
}

type bucketSize struct {
	count, size int64
}

func listBucket(name,
	parent string,
	res chan bucketSize,
	get getFolderFunc) {

	items := walkBucket(name, parent, get)

	for {

		select {

		case k, ok := <-items:

			if !ok {
				close(res)
				return
			}

			fmt.Println(name + "/" + k.Key)
		}

	}
}

func calcBucket(name,
	parent string,
	res chan bucketSize,
	get getFolderFunc) {

	items := walkBucket(name, parent, get)

	size := int64(0)
	count := int64(0)

	p := func() {

		l("Items:", humanize.Comma(count))
		l("Size:", humanize.Bytes(uint64(size)))

	}

	start := time.Now()
	timeout := time.Duration(5)

	for {

		select {

		case k, ok := <-items:

			if !ok {

				p()

				res <- bucketSize{count, size}
				return

			}

			size += k.Size
			count++

		case <-time.After(time.Second * timeout):

			p()

		}

		if time.Since(start).Seconds() > float64(timeout) {

			p()
			start = time.Now()

		}
	}
}

func main() {

	num := runtime.NumCPU()
	runtime.GOMAXPROCS(num)

	go func() {
		l(http.ListenAndServe("0.0.0.0:6060", nil))
	}()

	app := cli.NewApp()
	app.Name = "go-s3"

	cmdBucket := func(
		c *cli.Context,
		cmd func(name,
			parent string,
			res chan bucketSize,
			get getFolderFunc)) {

		path := c.Args().First()
		ps := strings.SplitN(path, "/", 2)

		res := make(chan bucketSize)
		go cmd(ps[0], ps[1], res, getFolder)

		<-res
	}

	cmds := []cli.Command{
		{
			Name:      "calc",
			ShortName: "c",
			Usage:     "calculates size and number of items",
			Action: func(c *cli.Context) {
				cmdBucket(c, calcBucket)
			},
		},
		{
			Name:      "list",
			ShortName: "l",
			Usage:     "lists folder items",
			Action: func(c *cli.Context) {
				cmdBucket(c, listBucket)
			},
		},
	}

	app.Commands = cmds

	app.Action = func(c *cli.Context) {
		cli.ShowAppHelp(c)
	}

	app.Run(os.Args)

}
