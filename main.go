package main

var (
	project = &Project{
		Name: "go-projay",
	}
)

func main() {
	makeFileWriter, err := NewMakefileWriter("makefile.txt", project)
	if err != nil {
		panic(err)
	}
	err = makeFileWriter.Init()
	if err != nil {
		panic(err)
	}
}
