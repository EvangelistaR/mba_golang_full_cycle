package main

import (
	"io"
	"net/http"
)

func main() {
	request, err := http.Get("https://google.com")

	if err != nil {
		panic(err)
	}

	//closing connection
	defer request.Body.Close()

	res, err := io.ReadAll(request.Body)

	if err != nil {
		panic(err)
	}

	println(string(res))

}
