package main

import (
	"fmt"
	"go-indexer/go-convert/converter"
	"os"
	"runtime"

	"go-indexer/s3-2-es/parser"
)

type s3l struct{}

func (s3l) Next(string) bool {
	return true
}

func (s3l) Parse(
	path,
	i string,
	num int) (map[string]interface{}, error) {

	return parser.ParseLine(i)

}

func main() {

	num := runtime.NumCPU()
	runtime.GOMAXPROCS(num)

	out := make(chan string)
	go converter.Convert(
		"",
		os.Stdin,
		s3l{},
		out)

	for v := range out {
		fmt.Println(v)
	}

}
