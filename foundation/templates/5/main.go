package main

import (
	"net/http"
	"text/template"
)

type Course struct {
	Name string
	Time int
}

type Courses []Course

func main() {
	templates := []string{
		"header.html",
		"content.html",
		"footer.html",
	}
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		courses := Courses{
			{"Golang", 80},
			{"Ruby", 180},
			{"JS", 30},
			{"Java", 280},
		}
		courseTemplate := template.Must(template.New("content.html").ParseFiles(templates...))
		err := courseTemplate.Execute(w, courses)

		if err != nil {
			panic(err)
		}
	})

	http.ListenAndServe(":8282", nil)
}
