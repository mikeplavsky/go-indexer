package main

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"time"

	_ "net/http/pprof"
)

var S3_PATH string

type event struct {
	line string
	num  int
}

func parse(path string,
	i string,
	num int) (string, error) {

	ts := strings.SplitN(i, "\t", 9)

	if len(ts) != 9 {
		return i, errors.New("can't parse")
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

	res, _ := json.Marshal(obj)

	return string(res), nil

}

func worker(
	path string,
	in <-chan event,
	quit <-chan bool,
	done chan<- string) {

	f, err := ioutil.TempFile(
		"",
		"json")

	if err != nil {
		panic(err)
	}

	defer f.Close()

	for {
		select {
		case e := <-in:

			res, err := parse(
				path,
				e.line,
				e.num)

			if err != nil {
				continue
			}
			f.WriteString(
				`{"index": {"_type": "log"}}` + "\n")
			f.WriteString(res + "\n")

		case <-quit:

			log.Println(f.Name())
			done <- f.Name()

			return
		}
	}

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

	go func() {
		log.Println(
			http.ListenAndServe(
				"0.0.0.0:6060",
				nil))
	}()

	in := make(chan event)

	quit := make(chan bool)
	done := make(chan string)

	for i := 0; i < num; i++ {
		go worker(
			S3_PATH,
			in,
			quit,
			done)
	}

	scanner := bufio.NewScanner(os.Stdin)

	l := 0

	for scanner.Scan() {

		in <- event{
			scanner.Text(),
			l,
		}

		l += 1

	}

	close(quit)

	res := []string{}

	for i := 0; i < num; i++ {
		f := <-done
		res = append(res, f)
	}

	cmd := "cat " + strings.Join(res, " ") + "> /tmp/mage.json"
	cat := exec.Command("bash", "-c", cmd)

	_, err := cat.Output()
	if err != nil {
		log.Println(err)
	}

	for _, n := range res {
		os.Remove(n)
	}

}