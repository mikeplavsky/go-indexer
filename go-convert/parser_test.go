package main

import (
	"encoding/json"
	"go-indexer/go-convert/converter"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestError(t *testing.T) {

	_, err := parse("testing", "", 0)
	assert.NotNil(t, err, "parsing does not return error")

}

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

	res := callConvert(strings.Join(in, "\t"))
	line := res[1]

	var out map[string]interface{}
	err := json.Unmarshal([]byte(line), &out)

	t.Log(out)

	assert.Nil(t, err)

	assert.Contains(t, out["@timestamp"], "2014-12-08", "wrong timestamp")
	assert.Equal(t, "testing#1", out["path"])

	filedsToVerify := []struct {
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

	for _, f := range filedsToVerify {
		assert.Equal(t, in[f.pos], out[f.name])
	}
}

func callConvert(in string) []string {
	r := strings.NewReader(in)

	out := make(chan string)
	go converter.Convert("testing", r, parse, out)

	res := []string{}
	for v := range out {

		ls := strings.Split(v, "\n")

		for _, l := range ls {
			res = append(res, l)
		}

	}

	return res
}
