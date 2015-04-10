package converter

import (
	"encoding/json"
	"errors"
	. "go-indexer/testUtils"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func callConvert(
	in string,
	parser Parser) []string {
	r := strings.NewReader(in)

	out := make(chan string)
	go Convert("testing", r, parser, out)

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

	var parseStub = func(
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

	out := callConvert("one\ntwo\nthree", parseStub)

	// index line between parsed lines is inserted
	outLineNums := []int{
		1,
		3,
		5}

	inFile := map[float64]string{
		//line number -> line content
		1: "one",
		2: "two",
		3: "three",
	}

	isParsed := map[float64]bool{}

	fileIDs := []string{}
	for _, outLineNum := range outLineNums {
		outLine := out[outLineNum]
		t.Log(outLine)

		var out map[string]interface{}
		err := json.Unmarshal([]byte(outLine), &out)
		assert.Nil(t, err, "Unable to parse JSON: "+outLine)

		assert.Equal(t, "testing", out["path"], "Wrong parsing")

		n := out["num"].(float64)
		assert.Equal(t, inFile[n], out["line"], "Wrong parsing")

		fileIDs = append(fileIDs, out["fileId"].(string))
		isParsed[n] = true
	}

	for _, fileID := range fileIDs {
		assert.NotEmpty(t, fileID)
		assert.Equal(t, fileIDs[0], fileID,
			"lines from the same file should have the same fileID")
	}

	for lineNum := range inFile {
		if !isParsed[lineNum] {
			t.Errorf(
				"Line %v has not been parsed",
				lineNum)
		}
	}

}

func TestNextIndex(t *testing.T) {

	var parseDummy = func(
		path,
		line string,
		num int) (map[string]interface{}, error) {
		return nil, nil
	}

	out := callConvert("one\ntwo\nthree", parseDummy)

	lineIds := []string{}
	for _, outLineNum := range []int{0, 2, 4} {

		line := out[outLineNum]
		t.Log(line)

		var idx map[string]map[string]string
		err := json.Unmarshal([]byte(line), &idx)

		assert.Nil(t, err, "Unable to parse JSON: "+line)
		assert.Equal(t, "log", idx["index"]["_type"], "wrong index type")
		assert.NotEmpty(t, idx["index"]["_id"])
		lineIds = append(lineIds, idx["index"]["_id"])
	}

	AssertAllValuesAreUnique(t, lineIds)
}

func TestParsingError(t *testing.T) {

	parseResults := []struct {
		line map[string]interface{}
		err  error
	}{
		{nil, nil},
		//skipped
		{nil, errors.New("")},
		{nil, nil},
	}

	i := 0

	var parseStub = func(
		path,
		line string,
		num int) (map[string]interface{}, error) {

		res := parseResults[i].line
		err := parseResults[i].err

		i = i + 1
		return res, err

	}

	out := callConvert("one\ntwo\nthree", parseStub)

	// (3 total - 1 failed) * (1 index line + 1 content line)
	assert.Equal(t, 4, len(out), "Wrong length")
}
