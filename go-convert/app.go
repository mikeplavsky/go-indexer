package main

import (
	"fmt"
	"log"
	"os"
	"runtime"

	"go-indexer/go-convert/converter"
)

var S3_PATH string

func main() {

	S3_PATH := os.Getenv("S3_PATH")

	if S3_PATH == "" {
		log.Println("S3_PATH is not set")
		return
	}

	num := runtime.NumCPU()
	runtime.GOMAXPROCS(num)

	log.Println("Num CPUs:", num)

	out := make(chan string)
	go converter.Convert(
		S3_PATH,
		os.Stdin,
		l{},
		out)

	for v := range out {
		fmt.Println(v)
	}

}
