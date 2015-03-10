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

type Parse func(
	path,
	line string,
	num int) (res map[string]interface{}, err error)

func worker(
	path string,
	parse Parse,
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

			obj, err := parse(
				path,
				e.line,
				e.num)

			if err != nil {
				continue
			}

			idx := [2]string{}
			lineId := id + strconv.Itoa(e.num)

			idx[0] = fmt.Sprintf(
				`{"index": {"_type": "log","_id":"%v"}}`,
				lineId)

			res, _ := json.Marshal(obj)
			idx[1] = string(res)

			out <- strings.Join(idx[:2], "\n")

		}
	}

}

func Convert(
	path string,
	r io.Reader,
	parse Parse,
	out chan<- string) {

	num := runtime.GOMAXPROCS(-1)

	in := make(chan event)
	done := make(chan bool)

	for i := 0; i < num; i++ {
		go worker(
			path,
			parse,
			in,
			out,
			done)
	}

	scanner := bufio.NewScanner(r)

	l := 0

	for scanner.Scan() {

		in <- event{
			scanner.Text(),
			l,
		}

		l += 1

	}

	close(in)

	for i := 0; i < num; i++ {
		<-done
	}

	close(out)

}
