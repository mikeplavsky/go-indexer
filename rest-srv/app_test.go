package main

import (
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestGetJob(t *testing.T) {
	getJobStats = func(j job) (map[string]uint64, error) {
		return map[string]uint64{
			"count": uint64(9000),
			"size":  uint64(100500),
		}, nil
	}

	code, body := getJob(job{})
	assert.Equal(t, http.StatusOK, code)

	var out map[string]interface{}
	json.Unmarshal([]byte(body), &out)

	t.Log(out)
	assert.Equal(t, "100KB", out["size"])
	assert.Equal(t, "9,000", out["count"])
}

func TestGetJobWithError(t *testing.T) {
	getJobStats = func(j job) (map[string]uint64, error) {
		return nil, fmt.Errorf("error")
	}
	code, _ := getJob(job{})
	assert.Equal(t, http.StatusInternalServerError, code, "")
}

func TestCustomers(t *testing.T) {
	getCustomers = func() ([]string, error) {
		return []string{"foo", "bar"}, nil
	}
	code, body := listCustomers()
	assert.Equal(t, http.StatusOK, code, "")
	assert.Equal(t, "[\"foo\",\"bar\"]", body, "")
}

func TestCustomersWithError(t *testing.T) {
	getCustomers = func() ([]string, error) {
		return nil, fmt.Errorf("error")
	}

	code, _ := listCustomers()
	assert.Equal(t, http.StatusInternalServerError, code, "")
}
