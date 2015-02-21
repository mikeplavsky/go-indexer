package converter

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"strings"
	"testing"
)

func TestMain(m *testing.M) {

	os.Remove("/tmp/mage.json")
	os.Exit(m.Run())

}

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

	Convert("path", r, parse)

	f, _ := ioutil.ReadFile("/tmp/mage.json")
	res := strings.Split(string(f), "\n")

	table := []struct {
		idx  int
		line string
		num  float64
	}{
		{1, "one", 0},
		{3, "two", 1},
		{5, "three", 2},
	}

	for _, v := range table {

		line := res[v.idx]

		var val map[string]interface{}
		err := json.Unmarshal([]byte(line), &val)

		if err != nil {
			t.Error(line, err)
		}

		if val["path"] != "path" {
			t.Error(v.idx, v, val, "Wrong parsing")
		}

		if val["line"] != v.line {
			t.Error(v.idx, v, val, "Wrong parsing")
		}

		if val["num"].(float64) != v.num {
			t.Error(v.idx, v, val, "Wrong parsing")
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

	Convert("path", r, parse)

	f, _ := ioutil.ReadFile("/tmp/mage.json")
	res := strings.Split(string(f), "\n")

	for _, v := range []int{0, 2, 4} {

		line := res[v]

		var idx map[string]map[string]string
		err := json.Unmarshal([]byte(line), &idx)

		if err != nil {
			t.Error(line, err)
		}

		if idx["index"]["_type"] != "log" {
			t.Error(line, "wrong index type")
		}
	}

}

func TestOutput(t *testing.T) {

	var parse = func(
		path,
		line string,
		num int) ([]byte, error) {
		return []byte(""), nil
	}

	r := strings.NewReader("one")

	Convert("path", r, parse)
	_, err := ioutil.ReadFile("/tmp/mage.json")

	if err != nil {
		t.Error(err)
	}

}
