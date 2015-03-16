package main

import (
	"github.com/stretchr/testify/assert"
	//	. "go-indexer/testUtils"
	"testing"
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

	_, response := getEta()
	assert.Equal(t, "{\"files\":1,\"queue\":\""+queueName+"\",\"time\":\"5s\"}", response)
}
