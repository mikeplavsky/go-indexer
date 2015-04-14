package main

import (
	"fmt"
	"os/exec"
)

func main() {

	cvrt := exec.Command(
		"./test.sh")

	res, err := cvrt.CombinedOutput()

	fmt.Println(string(res))

	if err != nil {
		fmt.Println(err)
	}

}
