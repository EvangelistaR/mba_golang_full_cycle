package main

import (
	"io/ioutil"
	"net/http"
	"time"
)

func main() {
	// set a limit timer to request
	call := http.Client{Timeout: time.Second}

	response, err := call.Get("http://google.com")

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)

	if err != nil {
		panic(err)
	}

	println(string(body))
}
