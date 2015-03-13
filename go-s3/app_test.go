package main

import (
	"errors"
	"testing"

	"github.com/mitchellh/goamz/s3"
	"github.com/stretchr/testify/assert"
)

func TestRetry(t *testing.T) {

	count := 0

	f := func() (T, error) {

		count++

		if count == 3 {
			return true, nil
		} else if count == 4 {
			t.Fatal("cycling forever")
		}

		return false, errors.New("")

	}

	res := retryCall(f).(bool)

	assert.True(t, res, "retry should be finished when count = 3")
}

func TestTruncatedWoMarker(t *testing.T) {

	get := func(name,
		folder,
		delim,
		marker string) *s3.ListResp {

		switch folder {
		case "":
			return &s3.ListResp{
				CommonPrefixes: []string{"folder"}}
		case "folder":

			if marker == "" {
				return &s3.ListResp{
					Prefix:      "folder",
					IsTruncated: true,
					Contents:    []s3.Key{{}, {Key: "next"}}}
			} else if marker == "next" {

				return &s3.ListResp{
					Contents: []s3.Key{{}, {}, {}}}
			}

		}

		return &s3.ListResp{}

	}

	res := make(chan bucketSize)
	go calcBucket("test", "", res, get)

	b := <-res

	if b.count != 5 {
		t.Error("wrong number of items")
	}

}

func TestTruncated(t *testing.T) {

	get := func(name,
		folder,
		delim,
		marker string) *s3.ListResp {

		switch folder {
		case "":
			return &s3.ListResp{
				CommonPrefixes: []string{"folder"}}
		case "folder":

			if marker == "" {
				return &s3.ListResp{
					Prefix:      "folder",
					NextMarker:  "next",
					IsTruncated: true,
					Contents:    []s3.Key{{}, {}}}
			} else if marker == "next" {

				return &s3.ListResp{
					Contents: []s3.Key{{}, {}}}
			}

		}

		return &s3.ListResp{}

	}

	res := make(chan bucketSize)
	go calcBucket("test", "", res, get)

	b := <-res

	if b.count != 4 {
		t.Error("wrong number of items")
	}

}

func TestSubFolders(t *testing.T) {

	get := func(name,
		folder,
		delim,
		marker string) *s3.ListResp {

		switch folder {
		case "":
			return &s3.ListResp{
				CommonPrefixes: []string{"folder"}}
		case "folder":
			return &s3.ListResp{
				Contents: []s3.Key{{}, {}}}
		}

		return &s3.ListResp{}

	}

	res := make(chan bucketSize)
	go calcBucket("test", "", res, get)

	b := <-res

	if b.count != 2 {
		t.Error("wrong number of items")
	}
}

func TestBucket(t *testing.T) {

	data := []struct {
		in  s3.ListResp
		out int64
	}{
		{s3.ListResp{}, 0},

		{s3.ListResp{
			Contents: []s3.Key{
				{},
				{}}}, 2},
	}

	for _, d := range data {

		get := func(name,
			folder,
			delim,
			marker string) *s3.ListResp {

			return &d.in

		}

		res := make(chan bucketSize)
		go calcBucket("test", "", res, get)

		b := <-res

		if b.count != d.out {
			t.Error("wrong number of items")
		}

	}

}
