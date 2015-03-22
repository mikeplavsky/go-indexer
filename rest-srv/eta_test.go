package main

import (
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
