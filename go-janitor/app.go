package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/goamz/goamz/aws"
	"github.com/goamz/goamz/ec2"
	"github.com/goamz/goamz/sqs"
)

func main() {

	auth, _ := aws.GetAuth("", "", "", time.Time{})

	sqs := sqs.New(auth, aws.USEast)
	res, _ := sqs.ListQueues("i-")

	ec := ec2.New(auth, aws.USEast)

	for _, i := range res.QueueUrl {

		id := strings.Split(i, "/")[4]

		opts := ec2.DescribeInstanceStatusOptions{
			InstanceIds: []string{id}}

		res, err := ec.DescribeInstanceStatus(&opts,
			ec2.NewFilter())

		if err != nil {
			fmt.Println(err)
		}

		fmt.Println(res)

	}

}
