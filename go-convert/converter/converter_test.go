package converter

import (
	"encoding/json"
	"errors"
	"io"
	"strings"
	"testing"
)

func callConvert(
	r io.Reader,
	parse Parse) []string {

	out := make(chan string)
	go Convert("path", r, parse, out)

	res := []string{}
	for v := range out {

		ls := strings.Split(v, "\n")

		for _, l := range ls {
			res = append(res, l)
		}

	}

	return res

}

func TestValue(t *testing.T) {

	var parse = func(
		path,
		line string,
		num int) (map[string]interface{}, error) {

		res := map[string]interface{}{}

		res["path"] = path
		res["line"] = line
		res["num"] = num

		return res, nil

	}

	r := strings.NewReader("one\ntwo\nthree")
	res := callConvert(r, parse)

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
		t.Log(line)

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
		num int) (map[string]interface{}, error) {
		return nil, nil
	}

	r := strings.NewReader("one\ntwo\nthree")
	res := callConvert(r, parse)

	for _, v := range []int{0, 2, 4} {

		line := res[v]
		t.Log(line)

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
		line map[string]interface{}
		err  error
	}{
		{nil, nil},
		{nil, errors.New("")},
		{nil, nil},
	}

	i := 0

	var parse = func(
		path,
		line string,
		num int) (map[string]interface{}, error) {

		res := table[i].line
		err := table[i].err

		i = i + 1
		return res, err

	}

	r := strings.NewReader("one\ntwo\nthree")
	res := callConvert(r, parse)

	if len(res) != 4 {
		t.Error("Wrong length", len(res), res)
	}

}
