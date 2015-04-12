package converter

import (
	"bufio"
	"crypto/md5"
	"encoding/json"
	"fmt"
	"io"
	"runtime"
	"strconv"
	"strings"
)

type event struct {
	line string
	num  int
}

type Parser interface {
	Next(line string) bool

	Parse(path,
		line string,
		num int) (
		res map[string]interface{},
		err error)
}

func worker(
	path string,
	parser Parser,
	in <-chan event,
	out chan<- string,
	done chan<- bool) {

	h := md5.New()
	io.WriteString(h, path)

	id := fmt.Sprintf("%x", h.Sum(nil))

	for {
		select {
		case e, ok := <-in:

			if !ok {
				done <- true
				return
			}

			obj, err := parser.Parse(
				path,
				e.line,
				e.num)

			if err != nil {
				continue
			}

			idx := [2]string{}
			lineID := id + strconv.Itoa(e.num)

			idx[0] = fmt.Sprintf(
				`{"index": {"_type": "log","_id":"%v"}}`,
				lineID)

			if obj != nil {
				obj["fileId"] = id
			}

			res, _ := json.Marshal(obj)
			idx[1] = string(res)

			out <- strings.Join(idx[:2], "\n")

		}
	}

}

// Convert generates bulk request for ES in parallel
func Convert(
	path string,
	r io.Reader,
	parser Parser,
	out chan<- string) {

	num := runtime.GOMAXPROCS(-1)

	in := make(chan event)
	done := make(chan bool)

	for i := 0; i < num; i++ {
		go worker(
			path,
			parser,
			in,
			out,
			done)
	}

	scanner := bufio.NewScanner(r)

	l := 1
	msgL := 1

	msg := ""

	for scanner.Scan() {

		line := scanner.Text()

		if len(msg) == 0 {

			msg = line
			l++
			continue
		}

		if !parser.Next(line) {

			msg += "\n"
			msg += line

		} else {

			in <- event{
				msg,
				msgL}

			msg = line
			msgL = l
		}
		l++

	}

	in <- event{
		msg,
		msgL}

	close(in)

	for i := 0; i < num; i++ {
		<-done
	}

	close(out)

}
