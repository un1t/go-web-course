package main

import (
	"os"
	"text/template"
)

func main() {
	funcMap := map[string]any{
		"add": func(a int, b int) int {
			return a + b
		},
	}

	tmpl, err := template.New("").Funcs(funcMap).ParseFiles("some.txt")
	if err != nil {
		panic(err)
	}

	data := map[string]any{
		"text":   "some text",
		"number": 42,
	}

	tmpl.ExecuteTemplate(os.Stdout, "some.txt", data)
}
