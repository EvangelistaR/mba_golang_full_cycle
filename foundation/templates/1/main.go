package main

import (
	"os"
	"text/template"
)

type Course struct {
	Name  string
	Hours int
}

func main() {
	course := Course{"Go", 40}

	tmp := template.New("TemplateCourse")
	tmp, _ = tmp.Parse("Course: {{.Name}} - Hours: {{.Hours}}\n")
	err := tmp.Execute(os.Stdout, course)

	if err != nil {
		panic(err)
	}
}
