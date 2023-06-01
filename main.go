package main

import (
	"bytes"
	"fmt"
	"os/exec"
	"xraybuilder/models"
	"xraybuilder/service/keypair"
)

const ShellToUse = "bash"

func Shellout(command string) (string, string, error) {
	var stdout bytes.Buffer
	var stderr bytes.Buffer
	cmd := exec.Command(ShellToUse, "-c", command)
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err := cmd.Run()
	return stdout.String(), stderr.String(), err
}

func GenerateKeyPair() (*models.KeyPair, error) {
	out, _, err := Shellout("xray x25519")
	if err != nil {
		return nil, err
	}
	keyPair, err := keypair.FromStdOut(out)
	if err != nil {
		return nil, err
	}
	return keyPair, nil
}

func main() {
	pair, err := GenerateKeyPair()
	if err != nil {
		panic(err)
	}
	fmt.Println(pair)
}
