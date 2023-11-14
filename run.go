package gosh

import (
	"os"
	"os/exec"
	"strings"

	log "github.com/sirupsen/logrus"
)

// Run a command in shell
func Run(args ...string) error {
	log.Debugf("Running: %s", strings.Join(args, " "))
	cmd := exec.Command(args[0], args[1:]...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	return cmd.Run()
}

// Run a command in shell
func RunS(command string) error {
	args := strings.Split(command, " ")
	return Run(args...)
}

// Run a command in shell as root
func SuRun(args ...string) error {
	return Run(append([]string{"sudo"}, args...)...)
}

// Run a command in shell as root
func SuRunS(command string) error {
	args := strings.Split(command, " ")
	return SuRun(args...)
}

// Run a command in shell in a directory
func RunIn(path string, args ...string) error {
	path = Expand(path)
	if err := os.MkdirAll(path, 0755); err != nil {
		return err
	}

	cmd := exec.Command(args[0], args[1:]...)
	cmd.Dir = path
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	log.Debugf("Running: %s in %s", strings.Join(args, " "), path)
	return cmd.Run()
}

// Run a command in shell in a directory
func RunInS(path, command string) error {
	args := strings.Split(command, " ")
	return RunIn(path, args...)
}

// Run a command in shell in a directory as root
func SuRunIn(path string, args ...string) error {
	return RunIn(path, append([]string{"sudo"}, args...)...)
}

// Run a command in shell in a directory
func SuRunInS(path, command string) error {
	args := strings.Split(command, " ")
	return SuRunIn(path, args...)
}
