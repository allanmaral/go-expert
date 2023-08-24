package main

import (
	"net/http"
	"strings"
	"text/template"
)

type Course struct {
	Name     string
	Workload int
}

type Courses []Course

func ToUpper(s string) string {
	return strings.ToUpper(s)
}

func main() {
	templates := []string{
		"header.html",
		"content.html",
		"footer.html",
	}

	t := template.New("content.html")
	t.Funcs(template.FuncMap{"ToUpper": ToUpper})

	t = template.Must(t.ParseFiles(templates...))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		err := t.Execute(w, Courses{
			{"Go", 40},
			{"Java", 20},
			{"Python", 30},
		})
		if err != nil {
			panic(err)
		}
	})

	http.ListenAndServe(":8080", nil)

}
