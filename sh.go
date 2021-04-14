package sh

import (
	"bytes"
	"errors"
	"io"
	"os"
	"os/exec"
	"strings"
	"time"
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

// RunWait executes command, and waits for command to succeed.
func RunWait(command string, timeOut time.Duration) error {
	var stdBuffer bytes.Buffer
	mw := io.MultiWriter(os.Stdout, &stdBuffer)

	cmd := exec.Command("sh", "-c", command)
	cmd.Stdout = mw
	cmd.Stderr = mw

	start := time.Now()

	for {
		if err := cmd.Run(); err == nil {
			break
		}
		if time.Since(start) > timeOut {
			return errors.New("command timeout: " + command)
		}
		time.Sleep(time.Second * 1)
	}

	return nil
}
