package converter

import (
	"encoding/json"
	"errors"
	"github.com/stretchr/testify/assert"
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

		res := map[string]interface{}{
			"path": path,
			"line": line,
			"num":  num,
		}

		return res, nil

	}

	r := strings.NewReader("one\ntwo\nthree")
	res := callConvert(r, parse)

	// index line between parsed lines has been inserted
	outLineNums := []int{
		1,
		3,
		5}

	inFile := map[float64]string{
		//line number -> line content
		0: "one",
		1: "two",
		2: "three",
	}

	isParsed := map[float64]bool{}

	for _, outLineNum := range outLineNums {

		outLine := res[outLineNum]
		t.Log(outLine)

		var out map[string]interface{}
		err := json.Unmarshal([]byte(outLine), &out)
		assert.Nil(t, err, "Unable to parse JSON: "+outLine)

		assert.NotEmpty(t, out["fileId"])
		assert.Equal(t, "path", out["path"], "Wrong parsing")

		n := out["num"].(float64)
		assert.Equal(t, inFile[n], out["line"], "Wrong parsing")

		isParsed[n] = true
	}

	for lineNum, _ := range inFile {
		if !isParsed[lineNum] {
			t.Errorf(
				"Line %v has not been parsed",
				lineNum)
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

	assert.Equal(t, 4, len(res), "Wrong length")
}
