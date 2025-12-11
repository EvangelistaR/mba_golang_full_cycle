package main

import (
	"log"
	"net/http"
)

func main() {
	// fileServer := http.FileServer(http.Dir("./public"))
	mux := http.NewServeMux()
	mux.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir("./public"))))

	mux.HandleFunc("/blog", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello from"))
	})
	log.Fatal(http.ListenAndServe(":8080", mux))
}
