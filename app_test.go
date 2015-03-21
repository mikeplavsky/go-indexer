package main

import (
	"errors"
	"io/ioutil"
	"os"
	"testing"

	"github.com/goamz/goamz/sqs"
	"github.com/stretchr/testify/assert"
)

type getLogPtr func(string, string) ([]byte, error)
type getMessagePtr func() (*sqs.Message, error)
type removeMessagePtr func(*sqs.Message) error
type execPtr func(string) ([]byte, error)

type ti struct {
	gl getLogPtr
	gm getMessagePtr
	rm removeMessagePtr
	e  execPtr
}

func (i ti) exec(c string) ([]byte, error) {

	if i.e != nil {
		return i.e(c)
	}

	return nil, errors.New("exec not implemented")
}

func (i ti) getLog(b string,
	p string) ([]byte, error) {

	if i.gl != nil {
		return i.gl(b, p)
	}

	return nil, errors.New("getLog not implemented")
}

func (i ti) getMessage() (*sqs.Message, error) {

	if i.gm != nil {
		return i.gm()
	}

	return nil, errors.New("getMessage not implemented")
}

func (i ti) removeMessage(m *sqs.Message) error {

	if i.rm != nil {
		return i.rm(m)
	}

	return errors.New("removeMessage not implemented")
}

func TestEnv(t *testing.T) {

	os.Setenv("ES_INDEXER", "indexer1")
	os.Setenv("ES_INDEX", "index1")
	os.Setenv("ES_QUEUE", "queue1")
	os.Setenv("ES_FS_PER_INDEX", "perIndex")

	setVars()

	assert.Equal(t, "index1", ES_INDEX)
	assert.Equal(t, "indexer1", ES_INDEXER)
	assert.Equal(t, "queue1", ES_QUEUE)
	assert.Equal(t, "perIndex", ES_FS_PER_INDEX)

}

func TestIndexMessageError(t *testing.T) {

	err := index(ti{})

	assert.EqualError(t,
		err,
		"getMessage not implemented")
}

func TestIndexJsonError(t *testing.T) {

	getMessage := func() (*sqs.Message, error) {
		return &sqs.Message{}, nil
	}

	err := index(ti{gm: getMessage})

	assert.EqualError(t,
		err,
		"unexpected end of JSON input")
}

func TestIndexLogError(t *testing.T) {

	getMessage := func() (*sqs.Message, error) {
		return &sqs.Message{
				Body: `{"bucket":"a", "path": "b"}`},
			nil
	}

	err := index(ti{gm: getMessage})

	assert.EqualError(t,
		err,
		"getLog not implemented")
}

func TestIndexSuccess(t *testing.T) {

	getMessage := func() (*sqs.Message, error) {
		return &sqs.Message{
				Body: `{"bucket":"a", "path": "b"}`},
			nil
	}

	getLog := func(string, string) ([]byte, error) {
		return []byte("Bytes"), nil
	}

	removeMessage := func(*sqs.Message) error {
		return nil
	}

	exec := func(c string) ([]byte, error) {

		assert.Equal(t, "test_indexer", c)

		data, err := ioutil.ReadFile(os.Getenv("ES_FILE"))

		assert.NoError(t, err)
		assert.Equal(t, "Bytes", string(data))

		return []byte("Combined Output"), nil

	}

	ES_INDEXER = "test_indexer"

	os.Unsetenv("S3_PATH")
	os.Unsetenv("ES_FILE")
	os.Unsetenv("ES_FILE_CONTENT")

	err := index(ti{
		gm: getMessage,
		gl: getLog,
		rm: removeMessage,
		e:  exec})

	assert.NoError(t, err)

	assert.Equal(t,
		"https://s3.amazonaws.com/a/b",
		os.Getenv("S3_PATH"))

}

func TestIndexError(t *testing.T) {

	getMessage := func() (*sqs.Message, error) {
		return &sqs.Message{
				Body: `{"bucket":"a", "path": "b"}`},
			nil
	}

	getLog := func(string, string) ([]byte, error) {
		return []byte("Bytes"), nil
	}

	err := index(ti{
		gm: getMessage,
		gl: getLog})

	assert.EqualError(t,
		err,
		"exec not implemented")

}
