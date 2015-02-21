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

func parseTime(uri string) (string, error) {

	formats := map[string]string{

		"20060102_150405": `.*_(\d+_\d+).*\.zip`,
		"01.02.2006":      `.*\.(\d{2}\.\d{2}\.\d{4}).*\.zip`,

	}

	for format, pattern := range formats {

		dateRegex := regexp.MustCompile(pattern)

		if dateRegex.MatchString(uri) {

			dateStr := dateRegex.ReplaceAllString(uri, "$1")
			date, e := time.Parse(format, dateStr)

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

	timestamp, err := parseTime(uri)

	if err != nil {
		return nil, err
	}

	customerID := regexp.MustCompile(
		`[^\/]*\/([^\/]*).*`).ReplaceAllString(uri, "$1")


	dataContract := map[string]string{
		"uri":        uri,
		"size":       size,
		"customer":   customerID,
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

	converter.Convert(
		"",
		os.Stdin,
		parse)

}
