package main

import (
	"encoding/json"
	"net/http"
	"runtime"

	"github.com/dustin/go-humanize"
	"github.com/go-martini/martini"
	"github.com/martini-contrib/binding"
)

func listCustomers() (int, string) {
	list, err := getCustomers()

	if err != nil {
		return showError(err)
	}

	return outputJSON(list)
}

func getJob(job job) (int, string) {
	stats, err := getJobStats(job)

	if err != nil {
		return showError(err)
	}

	res := map[string]interface{}{}
	res["count"] = humanize.Comma(int64(stats["count"]))
	res["size"] = humanize.Bytes(stats["size"])
	res["eta"] = calcEta(float64(stats["count"]))

	return outputJSON(res)
}

func startJob(job job) (int, string) {

	go sendJob(job)
	saveJob(job)

	return http.StatusOK, "started"

}

//todo: show error stacktrace in debug localhost, show empty 500 in production
func showError(err error) (int, string) {
	return http.StatusInternalServerError, err.Error()
}

func outputJSON(v interface{}) (int, string) {
	JSON, err := json.Marshal(v)
	if err != nil {
		return showError(err)
	}
	return http.StatusOK, string(JSON)
}

func newServer() *martini.ClassicMartini {
	m := martini.Classic()
	m.Use(martini.Logger())

	m.Get("/job", binding.Bind(job{}), getJob)
	m.Get("/jobs", getJobs)
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

	//sender.Init()

	m := newServer()
	m.Run()

}
