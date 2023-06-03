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

func (b *BashCmdExecutor) SuppressLoginMessage() error {
	_, _, err := shellout("touch .hushlogin")
	return err
}
func (b *BashCmdExecutor) ApplyIptablesRules() error {
	_, _, err := shellout(`
	iptables -A INPUT -i lo -j ACCEPT;
	iptables -A INPUT -m conntrack --ctstate ESTABLISHED,RELATED -j ACCEPT;
	iptables -A INPUT -p icmp -j ACCEPT;
	iptables -A INPUT -p tcp -m tcp --dport 22 -j ACCEPT;
	iptables -A OUTPUT -p tcp --sport 22 -m state --state ESTABLISHED -j ACCEPT;
	iptables -A INPUT -p tcp --dport 80 -j ACCEPT;
	iptables -A INPUT -p tcp --dport 443 -j ACCEPT;
	iptables -P INPUT DROP;
	ip6tables -A INPUT -i lo -j ACCEPT;
	ip6tables -A INPUT -m state --state RELATED,ESTABLISHED -j ACCEPT;
	ip6tables -A INPUT -p ipv6-icmp -j ACCEPT;
	ip6tables -P INPUT DROP;
	apt install iptables-persistent;
	`)

	return err
}
func (b *BashCmdExecutor) EnableTcpBBR() error {
	_, _, err := shellout(`
	echo "net.core.default_qdisc=fq" >> /etc/sysctl.conf;
	echo "net.ipv4.tcp_congestion_control=bbr" >> /etc/sysctl.conf;
	sysctl -p;
	`)

	return err
}

func NewBashExecutor() *BashCmdExecutor {
	return &BashCmdExecutor{}
}
