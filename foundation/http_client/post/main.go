package main

import (
	"bytes"
	"io"
	"net/http"
	"os"
)

func main() {
	// set a limit timer to request
	call := http.Client{}
	jsonNames := bytes.NewBuffer([]byte(`{"name": "John Doe"}`))
	response, err := call.Post("http://google.com", "application/json", jsonNames)

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	io.CopyBuffer(os.Stdout, response.Body, nil) // copying data from response.Body and printing it
}

// using Buffer is basically to write and read data
