package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"runtime"
	"strings"
	"time"

	"go-indexer/go-convert/converter"
)

var S3_PATH string

type l struct{}

func (l) Next(string) bool {
	return true
}

func (l) Parse(
	path,
	i string,
	num int) (map[string]interface{}, error) {

	ts := strings.SplitN(i, "\t", 9)

	if len(ts) != 9 {
		return nil, errors.New("can't parse")
	}

	d := ts[0] + "\t" + ts[1]
	t, _ := time.Parse("2006-01-02\t15:04:05.0000", d)

	obj := map[string]interface{}{}

	obj["@timestamp"] = t
	obj["path"] = path + fmt.Sprintf("#%v", num)
	obj["pid"] = ts[2]
	obj["tid"] = ts[3]
	obj["agent"] = ts[4]
	obj["coll"] = ts[5]
	obj["mbx"] = ts[6]
	obj["msg_type"] = ts[7]
	obj["msg"] = ts[8]

	return obj, nil

}

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
