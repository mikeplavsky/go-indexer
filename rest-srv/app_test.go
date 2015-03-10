package main

import (
	"github.com/go-martini/martini"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestInvalidJobParameters(t *testing.T) {
	cases := []martini.Params{
		martini.Params{},
		martini.Params{"dummy": "dummy"},
		martini.Params{"from": "2000", "to": "2010"},
		martini.Params{"customer": "val1", "to": "2010"},
	}

	for _, testCase := range cases {
		response := httptest.NewRecorder()
		getJob(testCase, response)
		assert.Equal(t,
			response.Code, http.StatusBadRequest, "customer, from, to params are required")
	}
}
