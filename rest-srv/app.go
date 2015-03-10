package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"runtime"

	"github.com/dustin/go-humanize"
	"github.com/go-martini/martini"
)

func parseParams(params martini.Params) (job, error) {
	log.Println(params)

	customer := params["customer"]
	from := params["from"]
	to := params["to"]

	if len(customer) == 0 || len(from) == 0 || len(to) == 0 {
		return job{},
			fmt.Errorf("customer, from, to fields are required")
	}

	return job{customer: customer, from: from, to: to}, nil
}

func listCustomers(params martini.Params, w http.ResponseWriter) string {
	list, err := getCustomers()

	if err != nil {
		showError(w, err)
	}

	JSON, _ := json.Marshal(map[string]interface{}{
		"result": list,
	})

	return string(JSON)
}

func getJob(params martini.Params, w http.ResponseWriter) string {

	job, err := parseParams(params)

	if err != nil {
		return showBadRequest(w, err)
	}

	stats, err := getJobStats(job)

	if err != nil {
		showError(w, err)
	}

	res := map[string]interface{}{}
	res["count"] = humanize.Comma(int64(stats["count"]))
	res["size"] = humanize.Bytes(stats["size"])
	res["eta"] = calcEta(float64(stats["count"]))

	data, _ := json.Marshal(res)
	return string(data)
}

func startJob(params martini.Params, w http.ResponseWriter) string {

	j, err := parseParams(params)

	if err != nil {
		return showBadRequest(w, err)
	}

	go sendJob(j)

	return "started"
}

//todo: show error stacktrace in debug localhost, show empty 500 in production
func showError(w http.ResponseWriter, err error) string {
	http.Error(w,
		err.Error(),
		http.StatusInternalServerError)
	return ""
}

func showBadRequest(w http.ResponseWriter, err error) string {
	http.Error(w,
		err.Error(),
		http.StatusBadRequest)
	return ""
}

func newServer() *martini.ClassicMartini {
	m := martini.Classic()
	m.Use(martini.Logger())

	m.Get("/job", getJob)
	m.Get("/eta", getEta)
	m.Get("/customers", listCustomers)
	m.Post("/job", startJob)
	//todo:remove this as I understand how to enable post in CUI
	m.Get("/job/create", startJob)
	return m
}

func main() {
	num := runtime.NumCPU()
	runtime.GOMAXPROCS(num)

	m := newServer()
	m.Run()
}
