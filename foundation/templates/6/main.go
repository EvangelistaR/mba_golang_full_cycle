package main

import (
	"html/template"
	"os"
	"strings"
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

	t := template.New("content.html")
	t.Funcs(template.FuncMap{"ToUpper": strings.ToUpper})
	t = template.Must(t.ParseFiles(templates...))

	err := t.Execute(os.Stdout, Courses{
		{"Golang", 80},
		{"Ruby", 180},
		{"JS", 30},
		{"Java", 280},
	})

	if err != nil {
		panic(err)
	}
}
