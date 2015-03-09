package converter

import (
	"bufio"
	"crypto/md5"
	"io"
	"runtime"
	"strings"
)

type event struct {
	line string
	num  int
}

type Parse func(
	path,
	line string,
	num int) (res []byte, err error)

func worker(
	path string,
	parse Parse,
	in <-chan event,
	out chan<- string,
	done chan<- bool) {

	h := md5.New()
	io.WriteString(h, path)

	id := h.Sum(nil)

	for {
		select {
		case e, ok := <-in:

			if !ok {
				done <- true
				return
			}

			res, err := parse(
				path,
				e.line,
				e.num)

			if err != nil {
				continue
			}

			idx := [2]string{}

			idx[0] = `{"index": {"_type": "log"}}`
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
