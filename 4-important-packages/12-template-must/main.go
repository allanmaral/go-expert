package main

import (
	"os"
	"text/template"
)

type Course struct {
	Name     string
	Workload int
}

func main() {
	course := Course{"Go", 40}
	t := template.Must(template.New("CourseTemplate").Parse("Course: {{.Name}} - Workload: {{.Workload}}"))
	err := t.Execute(os.Stdout, course)
	if err != nil {
		panic(err)
	}
}
