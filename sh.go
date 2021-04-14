package sh

import (
	"bytes"
	"io"
	"os"
	"os/exec"
	"strings"
)

func Run(command string) error {
	var stdBuffer bytes.Buffer
	mw := io.MultiWriter(os.Stdout, &stdBuffer)

	cmd := exec.Command("sh", "-c", command)
	cmd.Stdout = mw
	cmd.Stderr = mw

	if err := cmd.Run(); err != nil {
		return err
	}
	return nil
}

func RunR(command string) (string, error) {
	var stdBuffer bytes.Buffer
	mw := io.MultiWriter(os.Stdout, &stdBuffer)

	cmd := exec.Command("sh", "-c", command)
	cmd.Stderr = mw

	out, err := cmd.Output()
	if err != nil {
		return "", err
	}
	return strings.TrimRight(string(out), "\n"), nil
}
