package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
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

func ReadConfig(path string) (*models.Config, error) {
	if path == "" {
		path = "config.template.json"
	}
	file, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	config := models.Config{}
	err = json.Unmarshal([]byte(file), &config)
	if err != nil {
		return nil, err
	}
	return &config, nil
}

func main() {
	_, err := GenerateKeyPair()
	if err != nil {
		panic(err)
	}
	config, err := ReadConfig("")
	if err != nil {
		panic(err)
	}
	fmt.Println(config)
}
