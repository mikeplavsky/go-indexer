package main

import (
	"encoding/json"
	"errors"
	"net/http"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEtaCalc(t *testing.T) {

	getFilesPerSecond = func() float64 {
		return 42.5
	}

	res := calcEta(42.5 * 33 * 60)
	assert.Equal(t, "33m0s", res)

}

func TestEtaGet(t *testing.T) {

	getFilesPerSecond = func() float64 {
		return 0.2
	}

	nums := []int{2, 3, 4}

	nQueues = func() int {
		return len(nums)
	}

	getQueueNum = func(i int) (int, error) {
		return nums[i], nil
	}

	os.Setenv("ES_QUEUE", "eta_test")

	_, b := getEta()

	var res = map[string]interface{}{}
	json.Unmarshal([]byte(b), &res)

	t.Log(res)

	assert.Equal(t, "45s", res["time"])
	assert.Equal(t, 9, res["files"])

}

func TestEtaError(t *testing.T) {

	getFilesPerSecond = func() float64 {
		return 0.2
	}

	nums := []int{2, 3, 4}

	nQueues = func() int {
		return len(nums)
	}

	getQueueNum = func(i int) (int, error) {
		return 0, errors.New("does not work")
	}

	code, _ := getEta()

	assert.Equal(t,
		http.StatusInternalServerError, code)

}
