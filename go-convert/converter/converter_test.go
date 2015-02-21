package convert

import (
	"io/ioutil"
	"os"
	"strings"
	"testing"
)

func TestMain(m *testing.M) { 

	os.Remove("/tmp/mage.json")
	os.Exit(m.Run()) 

}

func TestOutput(t *testing.T) {

	var parse = func(
		path,
		line string,
		num int) (string, error) {
		return "", nil
	}

	r := strings.NewReader("one")

	Convert("path", r, parse)
	_, err := ioutil.ReadFile("/tmp/mage.json")

	if err != nil {
		t.Error(err)
	}

}
