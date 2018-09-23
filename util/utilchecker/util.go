package utilchecker

import "os/exec"

func CheckDep() {

}

func isCommandAvailable(name string) bool {
	cmd := exec.Command("/bin/sh", "-c", "command -v " + name)
	if err := cmd.Run(); err != nil {
		return false
	}
	return true
}

func installDep() error {
	cmd := exec.Command(
		"/bin/sh", "go get -u", "github.com/golang/dep/cmd/dep")
	return cmd.Run()
}