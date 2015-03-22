package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"gopkg.in/olivere/elastic.v1"
)

type tq struct {
	s func(int, string)
	n int
}

func (t tq) send(i int, s string) {
	if t.s != nil {
		t.s(i, s)
	}
}

func (t tq) qNum() int {
	return t.n
}

func TestSendJobError(t *testing.T) {

	getFiles = func(job,
		int,
		int) (*elastic.SearchHits, error) {
		return nil, errors.New("elastic error")
	}

	err := sendJobImpl(job{}, tq{n: 1})

	assert.Error(t, err)
}

func TestSendJobZero(t *testing.T) {

	getFiles = func(job,
		int,
		int) (*elastic.SearchHits, error) {

		hits := elastic.SearchHits{}
		return &hits, nil
	}

	send := func(int, string) {
		assert.False(t, true)
	}

	err := sendJobImpl(job{}, tq{s: send, n: 1})

	assert.NoError(t, err)
}

func TestSendQueues(t *testing.T) {

	mNum := 9
	qNum := 4

	getFiles = func(job,
		int,
		int) (*elastic.SearchHits, error) {

		hits := elastic.SearchHits{}

		for i := 0; i < mNum; i++ {

			msg := fmt.Sprintf(
				`{"uri":"https://s3.amazonaws.com/%v"}`,
				i)

			raw := json.RawMessage(msg)

			hits.Hits = append(
				hits.Hits,
				&elastic.SearchHit{
					Source: &raw})

		}

		return &hits, nil
	}

	type r struct {
		q int
		m string
	}

	res := make(chan r, mNum)

	send := func(q int, m string) {
		res <- r{q, m}
	}

	tq := tq{s: send, n: qNum}

	err := sendJobImpl(job{}, tq)
	close(res)

	assert.NoError(t, err)
	f := map[int]map[string]bool{}

	for r := range res {

		t.Log(r)

		if _, ok := f[r.q]; !ok {
			f[r.q] = map[string]bool{}
		}

		f[r.q][r.m] = true
	}

	t.Log(f)

	assert.Equal(t, 3, len(f[0]))

	for q := 1; q < qNum; q++ {

		t.Log(f[q])
		assert.Equal(t, 2, len(f[q]))

	}

	assert.True(t, f[0]["4"])
	assert.True(t, f[1]["1"])
	assert.True(t, f[3]["7"])

}

func TestSendPaging(t *testing.T) {

	total := 8
	mNum := 4
	qNum := 4

	PageSize = 2

	nextId := make(chan int)
	go func() {
		for i := 0; ; i++ {
			nextId <- i
		}
	}()

	getFiles = func(job job,
		skip int,
		take int) (*elastic.SearchHits, error) {

		hits := elastic.SearchHits{
			TotalHits: int64(total)}

		for i := 0; i < mNum; i++ {

			msg := fmt.Sprintf(
				`{"uri":"https://s3.amazonaws.com/%v/%v/%v"}`,
                                skip,
                                take,
				<-nextId)

			raw := json.RawMessage(msg)

			hits.Hits = append(
				hits.Hits,
				&elastic.SearchHit{
					Source: &raw})

		}

		return &hits, nil
	}

	type r struct {
		q int
		m string
	}

	res := make(chan r, total/PageSize*mNum)

	send := func(q int, m string) {
		res <- r{q, m}
	}

	tq := tq{s: send, n: qNum}

	err := sendJobImpl(job{}, tq)
	close(res)

	assert.NoError(t, err)
	f := map[int]map[string]bool{}

	for r := range res {

		t.Log(r)

		if _, ok := f[r.q]; !ok {
			f[r.q] = map[string]bool{}
		}

		f[r.q][r.m] = true
	}

	for q := 0; q < qNum; q++ {

		t.Log("Queue:", q)
		t.Log(f[q])

		assert.Equal(t, 4, len(f[q]))

	}

	assert.True(t, f[0]["2/2/4"])
	assert.True(t, f[1]["0/2/1"])
	assert.True(t, f[3]["2/2/7"])

}
