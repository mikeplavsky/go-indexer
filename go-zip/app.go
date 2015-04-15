package main

import (
	"archive/zip"
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {

	f := os.Getenv("ES_FILE")

	r, err := zip.OpenReader(f)
	if err != nil {
		log.Fatal(err)
	}

	defer r.Close()

	for _, f := range r.File {

		rc, err := f.Open()

		if err != nil {
			log.Fatal(err)
		}

		scanner := bufio.NewScanner(rc)
		for scanner.Scan() {

			fmt.Println(scanner.Text())
		}

		rc.Close()

	}
}
