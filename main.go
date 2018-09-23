package main

import (
	"path/filepath"
	"text/template"
)

var (
	appName                  = "go-projay"
	makeFileTemplateFilename = "makefile.txt"
)

type project struct {
	Name    string
	Version string
}

func main() {

	makeFile, err := NewMakefileWriter("Makefile")
	if err != nil {
		panic(err)
	}
	defer func() {
		err = makeFile.Close()
		if err != nil {
			panic(err)
		}
	}()

	t := template.Must(template.ParseFiles(filepath.Join("templates", makeFileTemplateFilename)))
	err = t.Execute(makeFile, &project{
		Name: appName,
	})
	if err != nil {
		panic(err)
	}
}
