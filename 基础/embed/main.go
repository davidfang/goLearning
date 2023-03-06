package main

import (
	"embed"
	"html/template"
	"os"
)

//go:embed  *.tmpl
var tmp embed.FS

func main() {
	name := "en.tmpl"
	if len(os.Args) > 1 {
		name =  os.Args[1] + ".tmpl"
	}
	arg := "word"
	if len(os.Args) > 2 {
		arg = os.Args[2]
	}

	t, err := template.ParseFS(tmp, "*")

	if err != nil {
		panic(err)
	}

	if err = t.ExecuteTemplate(os.Stdout, name, arg); err != nil {
		panic(err)
	}

}
