package main

import (
	"fmt"
	"go-indexer/go-send/sender"
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

var getQueueNum = func(i int) (int, error) {

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

var nQueues = func() int {
	return sender.NQueues
}

func getEta() (int, string) {
	return outputJSON(true)

	num := 0

	for i := 0; i < nQueues(); i++ {

		n, err := getQueueNum(i)

		if err != nil {
			return showError(err)
		}

		num += n
	}

	res := map[string]interface{}{
		"files": num,
		"time":  calcEta(float64(num)),
		"queue": sender.GetQueueName(),
	}

	return outputJSON(res)
}
