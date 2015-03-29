package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/dustin/go-humanize"
	"github.com/goamz/goamz/aws"
	"github.com/goamz/goamz/cloudwatch"
)

type Stat struct {
	All struct {
		Primaries struct {
			Docs struct {
				Count int64 `json:"count"`
			} `json:"docs"`
			Store struct {
				Size int64 `json:"size_in_bytes"`
			} `json:"store"`
		} `json:"primaries"`
	} `json:"_all"`
}

func putCwMetrics(r Res) {

	auth, _ := aws.GetAuth("", "", "", time.Time{})

	cw, _ := cloudwatch.NewCloudWatch(auth,
		aws.USEast.CloudWatchServicepoint)

	dim := cloudwatch.Dimension{
		Name:  "InstanceId",
		Value: os.Getenv("ES_INSTANCE_ID")}

	d := []cloudwatch.MetricDatum{
		{
			Unit:       "Count/Second",
			Value:      float64(r.Speed),
			MetricName: "IxSpeed",
			Dimensions: []cloudwatch.Dimension{dim}},
		{
			Unit:       "Count",
			Value:      float64(r.Docs),
			MetricName: "IxDocs",
			Dimensions: []cloudwatch.Dimension{dim}},
		{
			Unit:       "Megabytes",
			Value:      float64(r.Size / 1024 / 1024),
			MetricName: "IxSize",
			Dimensions: []cloudwatch.Dimension{dim}}}

	_, err := cw.PutMetricDataNamespace(
		d, "LogManagement")

	if err != nil {
		fmt.Println(err)
	}

}

func GetEsDocs(url string) (int64, int64) {

	res, err := http.Get(url)

	if err != nil {
		log.Println(err)
		return 0, 0
	}

	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		log.Println(err)
		return 0, 0
	}

	var stat Stat
	json.Unmarshal(body, &stat)

	return stat.All.Primaries.Docs.Count,
		stat.All.Primaries.Store.Size

}

type Res struct {
	Docs  int64
	Speed int64
	Size  int64
}

func GetSpeed(url string, res chan<- Res) {

	num1, _ := GetEsDocs(url)
	t1 := time.Now()

	time.Sleep(time.Second * 30)
	num2, size := GetEsDocs(url)

	s := (num2 - num1) / int64(time.Since(t1).Seconds())

	res <- Res{num2, s, size}
}

func GetAvgSpeed(ip string, p, num int) {

	res := make(chan Res)

	for i := 0; i < num; i++ {

		go func(p int) {

			url := "http://" + ip + ":" + strconv.Itoa(p) + "/_stats"
			GetSpeed(url, res)

		}(p + i)

	}

	s := Res{}

	for i := 0; i < num; i++ {

		r := <-res

		s.Speed += r.Speed
		s.Docs += r.Docs
		s.Size += r.Size

	}

	putCwMetrics(s)

	log.Printf("%v %v %v\n",
		humanize.Comma(s.Speed),
		humanize.Comma(s.Docs),
		humanize.Bytes(uint64(s.Size)))
}

func main() {

	ip := flag.String(
		"ip",
		"127.0.0.1",
		"elasticsearch ip")

	port := flag.Int(
		"p",
		8080,
		"port to start from")

	num := flag.Int(
		"n",
		1,
		"number of ports to scan")

	flag.Parse()

	for {
		GetAvgSpeed(*ip, *port, *num)
	}

}
