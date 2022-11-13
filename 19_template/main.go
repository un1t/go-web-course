package main

import (
	"os"
	"text/template"
)

func main() {
	tmpl, err := template.ParseGlob("templates/*")
	if err != nil {
		panic(err)
	}

	type Data struct {
		SomeInt    int
		SomeString string
		SomeSlice  []string
		SomeMap    map[string]string
	}
	data := Data{
		SomeInt:    42,
		SomeString: "вкусные яблоки",
		SomeSlice:  []string{"яблоко", "груша", "виноград"},
		SomeMap:    map[string]string{"aa": "11", "bb": "22", "cc": "33"},
	}

	tmpl.ExecuteTemplate(os.Stdout, "1.txt", data)
}
