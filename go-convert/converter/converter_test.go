package converter

import (
	"encoding/json"
	"errors"
	"strings"
	"testing"
)

func TestValue(t *testing.T) {

	var parse = func(
		path,
		line string,
		num int) ([]byte, error) {

		res := map[string]interface{}{}

		res["path"] = path
		res["line"] = line
		res["num"] = num

		return json.Marshal(res)

	}

	r := strings.NewReader("one\ntwo\nthree")

	out := make(chan string)
	go Convert("path", r, parse, out)

	res := []string{}
	for v := range out {
		res = append(res, v)
	}

	lines := []int{
		1,
		3,
		5}

	table := map[float64]string{
		0: "one",
		1: "two",
		2: "three",
	}

	check := map[float64]bool{}

	for _, v := range lines {

		line := res[v]

		var val map[string]interface{}
		err := json.Unmarshal([]byte(line), &val)

		if err != nil {
			t.Error(line, err)
		}

		if val["path"] != "path" {
			t.Error(v, v, val, "Wrong parsing")
		}

		n := val["num"].(float64)

		if val["line"] != table[n] {
			t.Error(v, v, val, "Wrong parsing")
		}

		check[n] = true

	}

	for _, i := range []float64{0, 1, 2} {
		if !check[i] {
			t.Errorf(
				"Line %v has not been parsed",
				i)
		}
	}

}

func TestNextIndex(t *testing.T) {

	var parse = func(
		path,
		line string,
		num int) ([]byte, error) {
		return []byte(""), nil
	}

	r := strings.NewReader("one\ntwo\nthree")

	out := make(chan string)
	go Convert("path", r, parse, out)

	res := []string{}
	for v := range out {
		res = append(res, v)
	}

	for _, v := range []int{0, 2, 4} {

		line := res[v]

		var idx map[string]map[string]string
		err := json.Unmarshal([]byte(line), &idx)

		if err != nil {
			t.Error(line, err)
		}

		if idx["index"]["_type"] != "log" {
			t.Error(line, "wrong index")
		}
	}

}

func TestParsingError(t *testing.T) {

	table := []struct {
		line string
		err  error
	}{
		{"a", nil},
		{"b", errors.New("")},
		{"c", nil},
	}

	i := 0

	var parse = func(
		path,
		line string,
		num int) (res []byte, err error) {

		res = []byte(table[i].line)
		err = table[i].err

		i = i + 1
		return

	}

	r := strings.NewReader("one\ntwo\nthree")

	out := make(chan string)
	go Convert("path", r, parse, out)

	res := []string{}
	for v := range out {
		res = append(res, v)
	}

	if len(res) != 4 {
		t.Error("Wrong length", len(res), res)
	}

}
