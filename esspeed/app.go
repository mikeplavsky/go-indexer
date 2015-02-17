package main

import (
	"encoding/json"
	"flag"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/dustin/go-humanize"
)

type Stat struct {
	All struct {
		Primaries struct {
			Docs struct {
				Count int64 `json:"count"`
			} `json:"docs"`
		} `json:"primaries"`
	} `json:"_all"`
}

func GetEsDocs(url string) int64 {

	res, err := http.Get(url)

	if err != nil {
		log.Println(err)
		return 0
	}

	body, _ := ioutil.ReadAll(res.Body)

	var stat Stat
	json.Unmarshal(body, &stat)

	return stat.All.Primaries.Docs.Count

}

type Res struct {
	Docs  int64
	Speed int64
}

func GetSpeed(url string, res chan<- Res) {

	num1 := GetEsDocs(url)
	t1 := time.Now()

	time.Sleep(time.Second * 30)
	num2 := GetEsDocs(url)

	s := (num2 - num1) / int64(time.Since(t1).Seconds())

	res <- Res{num2, s}
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

	}

	log.Printf("%v %v\n",
		humanize.Comma(s.Speed),
		humanize.Comma(s.Docs))
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
