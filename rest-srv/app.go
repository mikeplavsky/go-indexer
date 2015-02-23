package main

import (
	"log"
	"net/http"

	"github.com/go-martini/martini"
)

func getJob(r *http.Request) string {

	log.Println(r.URL.Query())
	return `{"count":12, "size": 220}`

}

func main() {

	m := martini.Classic()
	m.Use(martini.Logger())

	m.Get("/job", getJob)

	m.Run()

}
