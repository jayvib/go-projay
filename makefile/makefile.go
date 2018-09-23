package makefile

import (
	"io"
	"os"
	"text/template"
	projay "go-projay"
)

func NewMakefileWriter(templatePath string, proj *projay.Project) (*MakefileWriter, error) {
	file, err := os.Create("Makefile")
	if err != nil {
		return nil, err
	}
	return &MakefileWriter{
		project:      proj,
		writerCloser: file,
		tmpl:         template.Must(template.ParseFiles(templatePath)),
	}, nil
}

type MakefileWriter struct {
	project      *projay.Project
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
