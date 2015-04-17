package main

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func main() {

	path := os.Getenv("S3_PATH")

	h := md5.New()
	io.WriteString(h, path)

	id := fmt.Sprintf("%x", h.Sum(nil))

	url := "http://localhost:8080/_count?q=fileId:"
	d, err := http.Get(url + id)

	if err != nil {
		log.Fatal(err)
	}

	body, _ := ioutil.ReadAll(d.Body)
	log.Println(string(body))

	res := map[string]interface{}{}
	err = json.Unmarshal(body, &res)

	if err != nil {
		log.Fatal(err)
	}

	if res["count"].(float64) == 0 {
		log.Fatal(path + " not found.")
	}

}
