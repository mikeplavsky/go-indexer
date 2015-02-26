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

func getFilesPerSecond() (fPerSec float64) {

	cpu := runtime.NumCPU()
	fPerSec = 2.5 / 70.0 * float64(cpu)

	return
}

func getEta() (int, string) {

	q, err := sender.GetQueue()

	if err != nil {

		return http.StatusInternalServerError,
			err.Error()

	}

	attr, err := q.GetQueueAttributes(
		"ApproximateNumberOfMessages")

	if err != nil {

		return http.StatusInternalServerError,
			err.Error()

	}

	num, _ := strconv.Atoi(attr.Attributes[0].Value)

	res := map[string]interface{}{}

	res["files"] = num
	res["time"] = calcEta(float64(num))

	data, _ := json.Marshal(res)

	return http.StatusOK, string(data)

}
