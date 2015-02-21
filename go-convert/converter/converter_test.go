package convert

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

	r := strings.NewReader("one")

	Convert("path", r, parse)

	f, _ := ioutil.ReadFile("/tmp/mage.json")
	res := strings.Split(string(f), "\n")

	var idx map[string]interface{}

	err := json.Unmarshal([]byte(res[1]), &idx)

	if err != nil {
		t.Error(res[1], err)
	}

}

func TestIndex(t *testing.T) {

	var parse = func(
		path,
		line string,
		num int) ([]byte, error) {
		return []byte(""), nil
	}

	r := strings.NewReader("one")

	Convert("path", r, parse)

	f, _ := ioutil.ReadFile("/tmp/mage.json")
	res := strings.Split(string(f), "\n")

	var idx map[string]map[string]string

	err := json.Unmarshal([]byte(res[0]), &idx)

	if err != nil {
		t.Error(res[0], err)
	}

	if idx["index"]["_type"] != "log" {
		t.Error(res[0], "wrong index type")
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
