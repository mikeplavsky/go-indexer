package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"regexp"
	"strings"
	"time"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func ParseTimestamp(uri string) (timestamp string, err error) {
	formats := map[string]string{
		"20060102_150405": `.*_(\d+_\d+).*\.zip`,
		"01.02.2006":      `.*\.(\d{2}\.\d{2}\.\d{4}).*\.zip`,
	}

	for format, pattern := range formats {
		dateRegex := regexp.MustCompile(pattern)
		if dateRegex.MatchString(uri) {
			dateStr := dateRegex.ReplaceAllString(uri, "$1")
			date, e := time.Parse(format, dateStr)
			timestamp = date.Format("2006-01-02T15:04:05Z")
			return timestamp, e
		}
	}
	return "", fmt.Errorf("URI %s does match any regex", uri)
}

func ParseLine(line string) (dataContract map[string]string, err error) {
	fields := strings.Split(line, "\t")
	size := fields[0]
	uri := fields[1]
	timestamp, err := ParseTimestamp(uri)
	customerID := regexp.MustCompile(`[^\/]*\/([^\/]*).*`).ReplaceAllString(uri, "$1")
	if err != nil {
		return nil, err
	}
	dataContract = map[string]string{
		"uri":        uri,
		"size":       size,
		"customer":   customerID,
		"@timestamp": timestamp,
	}
	return dataContract, nil
}

func main() {
	infpath := flag.String("in", "", "input file")
	index := flag.String("index", "s3-all", "index name for ES")
	outfpath := flag.String("out", "", "output file")
	flag.Parse()

	fin, err := os.Open(string(*infpath))
	check(err)
	defer fin.Close()

	fout, err := os.Create(*outfpath)
	check(err)
	defer fout.Close()
	scanner := bufio.NewScanner(fin)

	id := 1
	for scanner.Scan() {
		line := scanner.Text()
		dataObj, err := ParseLine(line)
		if err != nil {
			if strings.HasSuffix(line, ".zip") {
				fmt.Printf("unable to parse %s. Error: %s\n", line, err)
			}
			continue
		}

		documentJSON, _ := json.Marshal(map[string]interface{}{
			"index": map[string]interface{}{
				"_id":    id,
				"_index": *index,
				"_type":  "items",
			},
		})
		fout.WriteString(string(documentJSON) + "\n")

		documentJSON, _ = json.Marshal(dataObj)
		fout.WriteString(string(documentJSON) + "\n")

		id++
	}
}
