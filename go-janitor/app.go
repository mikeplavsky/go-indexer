package main

import (
	"fmt"
	"strings"
	"sync"
	"time"

	"github.com/goamz/goamz/aws"
	"github.com/goamz/goamz/ec2"
	"github.com/goamz/goamz/sqs"
	"github.com/golang/glog"

	"github.com/cenkalti/backoff"
	"github.com/rubyist/circuitbreaker"
)

type T interface{}

func retryCall(
	f func() (T, error)) T {

	for {

		var res T

		get := func() error {

			return backoff.Retry(func() error {

				r, err := f()
				res = r

				return err

			}, backoff.NewExponentialBackOff())

		}

		cb := circuit.NewBreaker()
		err := cb.Call(get, time.Second*120)

		if err != nil {

			glog.Info(err)
			continue

		}

		return res

	}

}

func deleteQueue(
	res chan<- bool,
	qUrl string,
	ec *ec2.EC2,
	sqs *sqs.SQS) {

	qn := strings.Split(qUrl, "/")[4]
	id := qn[:10]

	opts := ec2.DescribeInstanceStatusOptions{
		InstanceIds: []string{id}}

	_, err := ec.DescribeInstanceStatus(&opts,
		ec2.NewFilter())

	if err != nil {

		if !strings.Contains(
			err.Error(),
			"InvalidInstanceID.NotFound") {
			return
		}

		q, err := sqs.GetQueue(qn)

		if err != nil {
			return
		}

		retryCall(
			func() (T, error) {
				return q.Delete()
			})

	}

	res <- true
	return

}

func main() {

	auth, _ := aws.GetAuth("", "", "", time.Time{})

	sqs := sqs.New(auth, aws.USEast)
	res, _ := sqs.ListQueues("i-")

	ec := ec2.New(auth, aws.USEast)

	r := make(chan bool)
	var w sync.WaitGroup

	go func() {

		l := len(res.QueueUrl)
		fmt.Printf("Queues:%v\n", l)

		for i := 0; ; i++ {

			<-r
			fmt.Print(".")

			w.Done()
		}

	}()

	for _, i := range res.QueueUrl {

		w.Add(1)
		go deleteQueue(r, i, ec, sqs)

	}

	w.Wait()
	fmt.Println("\n")

}
