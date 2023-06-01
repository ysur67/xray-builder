package main

import (
	"bytes"
	"encoding/json"
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

func ReadServerConfig(path string) (*models.ServerConfig, error) {
	if path == "" {
		path = "configs/server.template.json"
	}
	file, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	config := models.ServerConfig{}
	err = json.Unmarshal([]byte(file), &config)
	if err != nil {
		return nil, err
	}
	return &config, nil
}

func ReadArgs() *models.CliArgs {
	return &models.CliArgs{}
}

func main() {
	_, err := GenerateKeyPair()
	if err != nil {
		panic(err)
	}
	_, err = ReadServerConfig("")
	if err != nil {
		panic(err)
	}
	args := ReadArgs()
	// args.
}
