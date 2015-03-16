package main

import (
	"encoding/json"
	"fmt"
	"go-indexer/go-send/sender"
	"net/http"
	"runtime"
	"strconv"
	"time"
)

func calcEta(files float64) (eta string) {

	secs := files / getFilesPerSecond()
	eta = fmt.Sprintf("%v",
		time.Second*time.Duration(secs))

	return

}

var getFilesPerSecond = func() (fPerSec float64) {

	cpu := runtime.NumCPU()
	fPerSec = 2.5 / 70.0 * float64(cpu)

	return
}

func getQueueNum(i int) (int, error) {

	q, err := sender.GetQueue(i)

	if err != nil {
		return 0, err
	}

	attr, err := q.GetQueueAttributes(
		"ApproximateNumberOfMessages")

	if err != nil {
		return 0, err
	}

	num, _ := strconv.Atoi(attr.Attributes[0].Value)
	return num, nil

}

func getEta() (int, string) {

	num := 0

	for i := 0; i < sender.NQueues; i++ {

		n, err := getQueueNum(i)

		if err != nil {
			return http.StatusInternalServerError,
				err.Error()
		}

		num += n
	}

	res := map[string]interface{}{}

	res["files"] = num
	res["time"] = calcEta(float64(num))
	res["queue"] = sender.GetQueueName()

	data, _ := json.Marshal(res)

	return http.StatusOK, string(data)

}
