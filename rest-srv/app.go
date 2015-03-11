package main

import (
	"encoding/json"
	"net/http"
	"runtime"

	"github.com/dustin/go-humanize"
	"github.com/go-martini/martini"
	"github.com/martini-contrib/binding"
)

func listCustomers(w http.ResponseWriter) string {
	list, err := getCustomers()

	if err != nil {
		return showError(w, err)
	}

	JSON, _ := json.Marshal(list)

	return string(JSON)
}

func getJob(job job, w http.ResponseWriter) string {
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

func startJob(job job) string {
	go sendJob(job)
	return "started"
}

//todo: show error stacktrace in debug localhost, show empty 500 in production
func showError(w http.ResponseWriter, err error) string {
	http.Error(w,
		err.Error(),
		http.StatusInternalServerError)
	return ""
}

func newServer() *martini.ClassicMartini {
	m := martini.Classic()
	m.Use(martini.Logger())

	m.Get("/job", binding.Bind(job{}), getJob)
	m.Get("/eta", getEta)
	m.Get("/customers", listCustomers)
	m.Post("/job", binding.Bind(job{}), startJob)
	//todo:remove this as I understand how to enable post in CUI
	m.Get("/job/create", binding.Bind(job{}), startJob)
	return m
}

func main() {
	num := runtime.NumCPU()
	runtime.GOMAXPROCS(num)

	m := newServer()
	m.Run()
}
