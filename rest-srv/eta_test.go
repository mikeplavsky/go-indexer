package main

import (
	"go-indexer/go-send/sender"
	"net/http"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalcEta(t *testing.T) {
	getFilesPerSecond = func() float64 {
		return 42.5
	}

	res := calcEta(42.5 * 33 * 60)
	assert.Equal(t, "33m0s", res)
}

func TestGetEta(t *testing.T) {
	getFilesPerSecond = func() float64 {
		return 0.2
	}

	queue := SetUp()

	queue.SendMessage("dummy")

	code, body := getEta()
	assert.Equal(t, http.StatusOK, code)
	assert.Equal(t, "{\"files\":1,\"queue\":\""+queueName+"\",\"time\":\"5s\"}", body)
}

func TestGetEtaWithNonExistingQueue(t *testing.T) {
	os.Setenv("ES_QUEUE", "NOTEXISTING")
	sender.Init()

	code, _ := getEta()
	assert.Equal(t, http.StatusInternalServerError, code)
}
