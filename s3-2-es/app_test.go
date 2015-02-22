package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseTime(t *testing.T) {

	cases := []struct {
		uri,
		expectedTimestamp string
	}{
		{"s3://s3bucket/path/mage_20231123_122134.zip",
			"2023-11-23T12:21:34Z"},
		{"s3://s3bucket/path/mage.08.28.1998.zip",
			"1998-08-28T00:00:00Z"}, //legacy format: US digitd order, Europerian dots, akward
		{"s3://dmp-log-analysis/Quadra2901/UKTCEUCMIGPRDV1/UnifiedCalendarSync/MAgE_20150129_031000 (1).zip",
			"2015-01-29T03:10:00Z"},
		{"", ""},
	}

	for _, testCase := range cases {

		timestamp, err := parseTime(testCase.uri)

		assert.Equal(t,
			testCase.expectedTimestamp,
			timestamp,
			fmt.Sprintf("they should be equal. Error: %s", err))

	}
}

func TestParseLine(t *testing.T) {

	res, err := parseLine("1010760	buck-et/C-UST-OM-ER/F O L D/ E R/Log file_20150112_135921.zip")

	assert.Equal(t, nil, err, fmt.Sprintf("they should be equal. Error: %s", err))
	assert.Equal(t, "1010760", res["size"])
	assert.Equal(t, "C-UST-OM-ER", res["customer"])
	assert.Equal(t, "2015-01-12T13:59:21Z", res["@timestamp"])
	assert.Equal(t, "https://s3.amazonaws.com/buck-et/C-UST-OM-ER/F O L D/ E R/Log file_20150112_135921.zip", res["uri"])

}
