package internal

import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"
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

func DownloadAndInstallXray(config *models.InstallArgs) error {
	cmd := fmt.Sprintf(`bash -c "$(curl -L https://github.com/XTLS/Xray-install/raw/main/install-release.sh)" @ install -u root --version %s`, config.XrayVersion)
	out, _, err := Shellout(cmd)
	if err != nil {
		return err
	}
	fmt.Println(out)
	return nil
}

func GetServerAddr() (*string, error) {
	out, _, err := Shellout("hostname -I")
	if err != nil {
		return nil, err
	}
	out = strings.Split(out, " ")[0]
	return &out, nil
}
