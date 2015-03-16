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
			"1998-08-28T00:00:00Z"}, //legacy format: US digit order, Europerian dots, awkward
		{"s3://dmp-log-analysis/Quadra2901/UKTCEUCMIGPRDV1/UnifiedCalendarSync/MAgE_20150129_031000 (1).zip",
			"2015-01-29T03:10:00Z"},
		{"dmp-log-analysis/Ijuf/Bil/Calendar00-20150303152329.zip", "2015-03-03T15:23:29Z"},
	}

	for _, testCase := range cases {

		timestamp, err := parseTime(testCase.uri)

		assert.Equal(t,
			testCase.expectedTimestamp,
			timestamp,
			fmt.Sprintf("they should be equal. Error: %s", err))

	}
}

func TestParseTimeWithError(t *testing.T) {
	cases := []string{
		"s3://s3bucket/path/mage_20230232_122134.zip",  // 32nd of February, 31 is OK in Go:)
		"s3://s3bucket/path/mage_20230228_236000.zip",  // 23:60
		"s3://s3bucket/path/mage_20231123_1221340.zip", // extra zero
		"s3://s3bucket/path/mage_20231123122134.zip",   // ambiguous
	}
	for _, testCase := range cases {
		timeStamp, err := parseTime(testCase)
		assert.NotNil(t, err, "expected error, but received"+timeStamp)
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

func TestParseLineWithError(t *testing.T) {

	cases := []string{
		"1010760", //size only
		"buck-et/C-UST-OM-ER/F O L D/ E R/Log file_20150112_135921.zip", //uri only
		"1010760	Log file_20150112_135921.zip", //there is no bucket
		"1010760	buck-et/Log file_20150112_135921.zip", //there is no customer
		"1010760	s3://buck-et/Log file_20150112_135921.zip", //there is no customer
		"1010760	https://s3.amazonaws.com/buck-et/Log file_20150112_135921.zip", //there is no customer
		"1010760	s3://s3bucket/path/mage_20231123122134", //not a file
	}

	for _, testCase := range cases {
		res, err := parseLine(testCase)
		assert.NotNil(t, err, fmt.Sprintf("expected error, but received", res))
	}
}
