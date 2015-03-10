package main

import (
	//"github.com/go-martini/martini"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"net/url"
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
		url, _ := url.Parse(testCase)
		r := &http.Request{Method: "GET", URL: url}
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

	testCase := "http://example.com/job?customer=contoso&from=2000&to=2002"

	url, _ := url.Parse(testCase)
	r := &http.Request{Method: "GET", URL: url}
	response := httptest.NewRecorder()
	getJob(response, r)
	assert.Equal(t, http.StatusOK, response.Code)
}

func TestCustomers(t *testing.T) {
	getCustomers = func() ([]string, error) {
		return []string{"foo", "bar"}, nil
	}
	r := &http.Request{}
	response := httptest.NewRecorder()
	ret := listCustomers(response, r)
	assert.Equal(t, http.StatusOK, response.Code, "")
	assert.Equal(t, "[\"foo\",\"bar\"]", ret, "")
}
