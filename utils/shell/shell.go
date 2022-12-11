package shell

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
)

func ExecuteShellCommand(command string) (string, string, error) {
	var stdout bytes.Buffer
	var stderr bytes.Buffer
	cmd := exec.Command("bash", "-c", command)
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err := cmd.Run()
	return stdout.String(), stderr.String(), err
}

func CloneRepository(url string, dirName string) error {
	_, _, err := ExecuteShellCommand(fmt.Sprintf("git clone %s %s", url, dirName))
	if err != nil {
		return err
	}
	return nil
}

func ChangeDir(dirName string) error {
	return os.Chdir(dirName)
}
