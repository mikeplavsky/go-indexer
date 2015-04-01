package main

import (
	"fmt"
	"net/http"

	"gopkg.in/olivere/elastic.v1"
)

const (
	idx  = "jobs"
	idxT = "job"
)

type job struct {
	Customer string `form:"customer" binding:"required"`
	From     string `form:"from" binding:"required"`
	To       string `form:"to" binding:"required"`
}

func main() {

	c, _ := elastic.NewClient(
		http.DefaultClient,
		"http://localhost:8080")

	ex, _ := c.IndexExists(idx).Do()

	if !ex {

		fmt.Println("creating...")
		_, err := c.CreateIndex(idx).Do()

		if err != nil {
			fmt.Println(err)
		}

	}

	_, err := c.Index().
		Index(idx).
		Type(idxT).
		BodyJson(job{
		Customer: "BMW",
		From:     "A",
		To:       "B"}).
		Do()

	if err != nil {
		fmt.Println(err)
	}

}
