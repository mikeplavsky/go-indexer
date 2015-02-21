package main

import (
	"encoding/json"
	"go-indexer/go-convert/converter"
	"io/ioutil"
	"strings"
	"testing"
)

func TestParse(t *testing.T) {

	in := []string{
		"2014-12-08",
		"00:15:57.3561",
		"PxF078",
		"Tx3D6",
		"A27",
		"C94",
		"M1865",
		"Trace",
		"Exec SQL: SET_FOLDER_PROCESSING_LOCKED"}

	r := strings.NewReader(strings.Join(in, "\t"))

	converter.Convert(
		"tesing",
		r,
		parse)

	f, _ := ioutil.ReadFile("/tmp/mage.json")
	res := strings.Split(string(f), "\n")

	line := res[1]

	var val map[string]interface{}
	err := json.Unmarshal([]byte(line), &val)

	t.Log(val)

	if err != nil {
		t.Error(err)
	}

	w := val["@timestamp"].(string)

	if !strings.Contains(w, "2014") {
		t.Error("wrong year")
	}

	fs := []struct {
		pos  int
		name string
	}{
		{2, "pid"},
		{3, "tid"},
		{4, "agent"},
		{5, "coll"},
		{6, "mbx"},
		{7, "msg_type"},
		{8, "msg"},
	}

	for _, f := range fs {

		if val[f.name] != in[f.pos] {

			t.Errorf("expected %v, got %v",
				in[f.pos],
				val[f.name])

		}

	}

}
