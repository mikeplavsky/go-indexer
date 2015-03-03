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

func parseParams(r *http.Request) (job, error) {
	params := r.URL.Query()
	log.Println(params)

	customer := params.Get("customer")
	from := params.Get("from")
	to := params.Get("to")

	if len(customer) == 0 || len(from) == 0 || len(to) == 0 {
		return job{},
			fmt.Errorf("customer, from, to fields are required")
	}

	return job{customer: customer, from: from, to: to}, nil
}

func listCustomers(w http.ResponseWriter,
	r *http.Request) string {
	list, err := getCustomers()

	if err != nil {
		showError(w, err)
	}

	JSON, _ := json.Marshal(map[string]interface{}{
		"result": list,
	})

	return string(JSON)
}

func getJob(w http.ResponseWriter,
	r *http.Request) string {

	job, err := parseParams(r)

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

func startJob(w http.ResponseWriter,
	r *http.Request) string {

	j, err := parseParams(r)

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

func main() {

	num := runtime.NumCPU()
	runtime.GOMAXPROCS(num)

	m := martini.Classic()
	m.Use(martini.Logger())

	m.Get("/job", getJob)
	m.Get("/eta", getEta)
	m.Get("/customers", listCustomers)
	m.Post("/job", startJob)
	//todo:remove this as I understand how to enable post in CUI
	m.Get("/job/create", startJob)
	m.Run()
}
