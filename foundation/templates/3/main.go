package main

import (
	"os"
	"text/template"
)

type Course struct {
	Name string
	Time int
}

type Courses []Course

func main() {
	courses := Courses{
		{"Golang", 80},
		{"Ruby", 180},
		{"JS", 30},
		{"Java", 280},
	}

	courseTemplate := template.Must(template.New("template.html").ParseFiles("template.html"))
	err := courseTemplate.Execute(os.Stdout, courses)

	if err != nil {
		panic(err)
	}
}
