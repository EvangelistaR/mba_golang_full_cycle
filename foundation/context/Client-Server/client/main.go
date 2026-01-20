package main

import (
	"context"
	"io"
	"net/http"
	"os"
	"time"
)

func main() {
	// Client code would go here
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "GET", "http://localhost:8080", nil)

	checkError(err)

	res, err := http.DefaultClient.Do(req)
	checkError(err)

	defer res.Body.Close()
	io.Copy(os.Stdout, res.Body)
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
