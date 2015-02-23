package main

import "github.com/go-martini/martini"

func main() {

	m := martini.Classic()
	m.Use(martini.Logger())

	m.Get("/job", func() string {
		return `{"count":12, "size": 220}`
	})

	m.Run()

}
