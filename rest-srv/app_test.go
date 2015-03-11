package main

import (
	"github.com/stretchr/testify/assert"
	"gopkg.in/olivere/elastic.v1"
	"net/http"
	"net/http/httptest"
	"testing"
	"encoding/json"
	"fmt"
)

func TestJobInfo(t *testing.T) {
	getJobStats = func(j job) (map[string]uint64, error) {
		return map[string]uint64{
			"count": uint64(9000),
			"size":  uint64(100500),
		}, nil
	}

	r := job{Customer: "constoso", From: "200", To: "2001"}
	response := httptest.NewRecorder()
	ret := getJob(r, response)
	assert.Equal(t, http.StatusOK, response.Code)
	
	var out map[string]interface{}
	json.Unmarshal([]byte(ret), &out)

	fmt.Println(out)
	assert.Equal(t, "100KB", out["size"])
	assert.Equal(t, "9,000", out["count"])
}

func TestCustomers(t *testing.T) {
	getCustomers = func() ([]string, error) {
		return []string{"foo", "bar"}, nil
	}
	response := httptest.NewRecorder()
	ret := listCustomers(response)
	assert.Equal(t, http.StatusOK, response.Code, "")
	assert.Equal(t, "[\"foo\",\"bar\"]", ret, "")
}

func TestStartJob(t *testing.T) {
	getFiles = func(j job, skip int, take int) (hits *elastic.SearchHits, err error) {
		return nil, nil
	}

	r := job{Customer: "constoso", From: "200", To: "2001"}
	startJob(r)
}
