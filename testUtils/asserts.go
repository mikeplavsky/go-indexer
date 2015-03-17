package sender

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// AssertSetsAreEqual verifies that arrays with distinct values have the same
// values, excluding order. It does not modify both collections
func AssertSetsAreEqual(t *testing.T, expected []string, actual []string) {
	assert.Equal(t, len(expected), len(actual), "length are not equal")

	var actualSet = map[string]bool{}
	for _, v := range actual {
		actualSet[v] = true
	}

	for _, v := range expected {
		if !actualSet[v] {
			t.Errorf("%s is not found in %s", v, actual)
		}
	}
}

// AssertAllValuesAreUnique check that there is no repeating item
func AssertAllValuesAreUnique(t *testing.T, values []string) {
	var set = map[string]bool{}
	for _, v := range values {
		set[v] = true
	}
	if len(set) != len(values) {
		t.Errorf("%v contain non-distinct values", values)
	}
}
