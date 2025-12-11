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

	courseTemplate := template.Must(template.New("TemplateCourse").Parse("Course: {{.Name}} - Hours: {{.Hours}}\n"))

	err := courseTemplate.Execute(os.Stdout, course)

	if err != nil {
		panic(err)
	}
}
