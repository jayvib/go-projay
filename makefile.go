package main

import (
	"io"
	"os"
	"path/filepath"
	"text/template"
)

func NewMakefileWriter(templateName string, proj *Project) (*MakefileWriter, error) {
	file, err := os.Create("Makefile")
	if err != nil {
		return nil, err
	}
	return &MakefileWriter{
		project:      proj,
		writerCloser: file,
		tmpl:         template.Must(template.ParseFiles(filepath.Join("templates", templateName))),
	}, nil
}

type MakefileWriter struct {
	project      *Project
	writerCloser io.WriteCloser
	tmpl         *template.Template
}

func (m *MakefileWriter) Write(w []byte) (int, error) {
	return m.writerCloser.Write(w)
}

func (m *MakefileWriter) Close() error {
	return m.writerCloser.Close()
}

func (m *MakefileWriter) Init() error {
	defer m.Close()
	return m.tmpl.Execute(m.writerCloser, m.project)
}
