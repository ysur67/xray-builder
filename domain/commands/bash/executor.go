package commands

import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"
)

const ShellToUse = "bash"

type BashCmdExecutor struct{}

func (b *BashCmdExecutor) GenerateKeyPair() (*string, error) {
	out, _, err := shellout("xray x25519")
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (b *BashCmdExecutor) DownloadAndInstallXray(version string) error {
	cmd := fmt.Sprintf(
		`bash -c "$(curl -L https://github.com/XTLS/Xray-install/raw/main/install-release.sh)" @ install -u root --version %s`,
		version,
	)
	out, _, err := shellout(cmd)
	if err != nil {
		return err
	}
	fmt.Println(out)
	return nil
}

func (b *BashCmdExecutor) GenerateShortId() (*string, error) {
	out, _, err := shellout("openssl rand -hex 8")
	if err != nil {
		return nil, err
	}
	result := strings.TrimSuffix(out, "\n")
	return &result, nil
}

func (b *BashCmdExecutor) GetServerAddr() (*string, error) {
	out, _, err := shellout("hostname -I")
	if err != nil {
		return nil, err
	}
	result := strings.Split(out, " ")[0]
	return &result, nil
}

func (b *BashCmdExecutor) RestartXray() error {
	_, _, err := shellout("systemctl restart xray")
	return err
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
