package converter

import (
	"bufio"
	"io"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
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
			f.WriteString(string(res) + "\n")

		case <-quit:

			done <- f.Name()
			return
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

	quit := make(chan bool)
	done := make(chan string)

	for i := 0; i < num; i++ {
		go worker(
			path,
			parse,
			in,
			quit,
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

	close(quit)

	res := []string{}

	for i := 0; i < num; i++ {
		f := <-done
		res = append(res, f)
	}

	cmd := "cat " + strings.Join(res, " ")
	cat := exec.Command("bash", "-c", cmd)

	catout, err := cat.StdoutPipe()

	if err != nil {
		log.Println(err)
	}

	if err := cat.Start(); err != nil {
		log.Println(err)
	}

	ret := bufio.NewScanner(catout)

	for ret.Scan() {
		out <- ret.Text()
	}
	defer close(out)

	for _, n := range res {
		os.Remove(n)
	}

}
