package main

import "os"

func NewMakefileWriter(filename string) (*MakefileWriter, error) {
	file, err := os.Create(filename)
	if err != nil {
		return nil, err
	}
	return &MakefileWriter{
		file: file,
	}, nil
}

type MakefileWriter struct {
	file *os.File
}

func (m *MakefileWriter) Write(w []byte) (int, error) {
	return m.file.Write(w)
}

func (m *MakefileWriter) Close() error {
	return m.file.Close()
}
