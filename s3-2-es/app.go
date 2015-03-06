package main

import (
	"encoding/json"
	"fmt"
	"go-indexer/go-convert/converter"
	"os"
	"regexp"
	"runtime"
	"strings"
	"time"
)

var formats = map[string]*regexp.Regexp{

	"20060102_150405": regexp.MustCompile(`.*_(\d+_\d+).*\.zip`),
	"01.02.2006":      regexp.MustCompile(`.*\.(\d{2}\.\d{2}\.\d{4}).*\.zip`),
	"20060102150405": regexp.MustCompile(`.*-(\d{14}).*\.zip`),
}

func parseTime(uri string) (string, error) {

	for f, r := range formats {

		if r.MatchString(uri) {

			dateStr := r.ReplaceAllString(uri, "$1")
			date, e := time.Parse(f, dateStr)

			timestamp := date.Format("2006-01-02T15:04:05Z")

			return timestamp, e

		}

	}
	return "", fmt.Errorf("URI %s does match any regex", uri)

}

func parseLine(
	line string) (map[string]string, error) {

	fields := strings.Split(line, "\t")

	size := fields[0]
	uri := fields[1]

	ps := strings.Split(uri, "/")

	timestamp, err := parseTime(ps[len(ps)-1])

	if err != nil {
		return nil, err
	}

	dataContract := map[string]string{
		"uri":        "https://s3.amazonaws.com/" + uri,
		"size":       size,
		"customer":   ps[1],
		"@timestamp": timestamp,
	}

	return dataContract, nil
}

func parse(
	path,
	i string,
	num int) ([]byte, error) {

	obj, err := parseLine(i)

	if err != nil {
		return []byte(i), err
	}

	return json.Marshal(obj)

}

func main() {

	num := runtime.NumCPU()
	runtime.GOMAXPROCS(num)

	out := make(chan string)
	go converter.Convert(
		"",
		os.Stdin,
		parse,
		out)

	for v := range out {
		fmt.Println(v)
	}

}
