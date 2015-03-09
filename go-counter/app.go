package main

import (
	"encoding/json"
	"errors"
	"go-indexer/go-send/sender"
	"log"
	"strconv"
)

func getCount() (uint64, error) {

	q, err := sender.GetQueue()

	if err != nil {
		return 0, err
	}

	ps := map[string]string{
		"WaitTimeSeconds":     "10",
		"MaxNumberOfMessages": "1"}

	res, err := q.ReceiveMessageWithParameters(ps)

	if err != nil {
		return 0, err
	}

	if len(res.Messages) == 0 {
		return 0, errors.New("No messages")
	}

	q.DeleteMessage(&res.Messages[0])

	raw := res.Messages[0].Body

	var msg map[string]interface{}
	err = json.Unmarshal([]byte(raw), &msg)

	if err != nil {
		return 0, err
	}

	r := msg["path"].(string)
	count, _ := strconv.Atoi(r)

	return uint64(count), nil

}

func main() {

	count := uint64(0)

	for i := 0; ; i++ {
		c, err := getCount()

		if err != nil {
			log.Println(err)
		}

		count += c
		log.Printf("%v\t%v\t%v",
			i,
			c,
			count)
	}

}
