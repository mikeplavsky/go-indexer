package main

import (
	"log"
	"net/http"

	"github.com/go-martini/martini"
	es "github.com/mattbaird/elastigo/lib"
)

func getJob(r *http.Request) string {

	log.Println(r.URL.Query())

	c := es.NewConn()

	c.Domain = "localhost"
	c.Port = "8080"

	return `{"count":12, "size": 220}`

}

func main() {

	m := martini.Classic()
	m.Use(martini.Logger())

	m.Get("/job", getJob)

	m.Run()

}
