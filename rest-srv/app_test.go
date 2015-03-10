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
		assert.Equal(t, http.StatusBadRequest, response.Code, testCase + ": customer, from, to params are required")

	}
}

func TestCustomers(t *testing.T) {
	r := &http.Request{}
	response := httptest.NewRecorder()
	listCustomers(response, r)
	assert.Equal(t, http.StatusOK, response.Code, "")
}
