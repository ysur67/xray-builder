package commands

import (
	"bytes"
	"fmt"
	"os/exec"
)

const ShellToUse = "bash"

type BashCmdExecutor struct{}

func (b *BashCmdExecutor) Execute(cmd string) (*string, error) {
	out, stderr, err := shellout(cmd)
	if err != nil {
		return nil, err
	}
	fmt.Println(stderr)
	return &out, nil
}

func shellout(command string) (string, string, error) {
	var stdout bytes.Buffer
	var stderr bytes.Buffer
	cmd := exec.Command(ShellToUse, "-c", command)
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err := cmd.Run()
	return stdout.String(), stderr.String(), err
}

func NewBashExecutor() *BashCmdExecutor {
	return &BashCmdExecutor{}
}
