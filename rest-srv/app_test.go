package main

import (
	"github.com/stretchr/testify/assert"
	"gopkg.in/olivere/elastic.v1"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestInvalidJobParameters(t *testing.T) {
	cases := []string{
		"http://example.com/job",
		"http://example.com/job?customer=contoso&from=2000", //missing params should not cause 500 status
		"http://example.com/job?from=2000&from=2002",
		"http://example.com/job?employer=contoso&from=2000&to=2002",
	}

	for _, testCase := range cases {
		r, _ := http.NewRequest("GET", testCase, nil)
		response := httptest.NewRecorder()
		getJob(response, r)
		assert.Equal(t, http.StatusBadRequest, response.Code, testCase+": customer, from, to params are required")
	}
}

func TestJobInfo(t *testing.T) {
	getJobStats = func(j job) (map[string]uint64, error) {
		return map[string]uint64{
			"count": uint64(9000),
			"size":  uint64(100500),
		}, nil
	}

	r, _ := http.NewRequest("GET", "http://example.com/job?customer=contoso&from=2000&to=2002", nil)
	response := httptest.NewRecorder()
	getJob(response, r)
	assert.Equal(t, http.StatusOK, response.Code)
}

func TestCustomers(t *testing.T) {
	getCustomers = func() ([]string, error) {
		return []string{"foo", "bar"}, nil
	}
	r, _ := http.NewRequest("GET", "", nil)
	response := httptest.NewRecorder()
	ret := listCustomers(response, r)
	assert.Equal(t, http.StatusOK, response.Code, "")
	assert.Equal(t, "[\"foo\",\"bar\"]", ret, "")
}

func TestStartJob(t *testing.T) {
	getFiles = func(j job, skip int, take int) (hits *elastic.SearchHits, err error) {
		return nil, nil
	}

	r, _ := http.NewRequest("GET", "http://example.com/job?customer=contoso&from=2000&to=2002", nil)
	response := httptest.NewRecorder()
	startJob(response, r)
	assert.Equal(t, http.StatusOK, response.Code)
}
