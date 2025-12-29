package main

import (
	"io/ioutil"
	"net/http"
)

func main() {
	// set a limit timer to request
	client := http.Client{}
	request, err := http.NewRequest("GET", "http://google.com", nil)

	if err != nil {
		panic(err)
	}

	request.Header.Set("Accept", "application/json")
	response, err := client.Do(request)

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)

	if err != nil {
		panic(err)
	}

	println(string(body))
	// or just doing
	// io.CopyBuffer(os.Stdout, response.Body, nil)

}
