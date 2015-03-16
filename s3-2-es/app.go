package main

import (
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
	"20060102150405":  regexp.MustCompile(`.*-(\d{14}).*\.zip`),
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
	line string) (map[string]interface{}, error) {

	fields := strings.Split(line, "\t")

	if len(fields) < 2 {
		return nil, fmt.Errorf("line should contain size and uri")
	}

	size := fields[0]
	uri := fields[1]
	uri = strings.TrimPrefix(uri, "https://s3.amazonaws.com/")
	uri = strings.TrimPrefix(uri, "s3://")

	ps := strings.Split(uri, "/")

	if len(ps) < 3 {
		return nil, fmt.Errorf("uri should contain a bucket and a customer")
	}

	timestamp, err := parseTime(ps[len(ps)-1])

	if err != nil {
		return nil, err
	}

	dataContract := map[string]interface{}{
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
	num int) (map[string]interface{}, error) {

	return parseLine(i)

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
