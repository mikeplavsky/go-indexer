package main

import (
	"bytes"
	"io/ioutil"
	"os"
	"strconv"
	"text/template"

	"github.com/codegangsta/cli"
)

func add(a, b int) int { return a + b }

func parse(cnt int, in, out string) {

	inStr, _ := ioutil.ReadFile(in)

	funcMap := template.FuncMap{"add": add}
	t := template.New("T").Funcs(funcMap)

	t, err := t.Parse(string(inStr))

	if err != nil {
		panic(err)
	}

	c := []int{}
	for i := 0; i < cnt; i++ {
		c = append(c, i)
	}

	buff := bytes.NewBufferString("")
	t.Execute(buff, c)

	ioutil.WriteFile(out, buff.Bytes(), 0644)

}

func main() {

	app := cli.NewApp()

	app.Name = "repeater"
	app.Usage = "repeater count input output"
	app.Version = "v0.1"

	app.Action = func(c *cli.Context) {

		cnt, _ := strconv.Atoi(c.Args()[0])
		in := c.Args()[1]
		out := c.Args()[2]

		parse(cnt, in, out)

	}

	app.Run(os.Args)

}
