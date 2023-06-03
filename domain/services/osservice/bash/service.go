package bash

import (
	"fmt"
	"strings"
	"xraybuilder/domain/commands"
	"xraybuilder/domain/services/osservice"
	"xraybuilder/models"
)

type BashOsService struct {
	executor commands.CmdExecutor
}

func (s *BashOsService) GenerateKeyPair() (*models.KeyPair, error) {
	out, err := s.executor.Execute("xray x25519")
	if err != nil {
		return nil, err
	}
	keyPair, err := fromStdOut(*out)
	if err != nil {
		return nil, err
	}
	return keyPair, nil
}

func fromStdOut(value string) (*models.KeyPair, error) {
	values := strings.Split(strings.TrimSuffix(value, "\n"), "\n")
	if len(values) != 2 {
		return nil, &osservice.KeyPairServiceError{Type: osservice.InvalidResponse}
	}
	private, public := values[0], values[1]
	if private == "" || public == "" {
		return nil, &osservice.KeyPairServiceError{Type: osservice.InvalidResponse}
	}
	private = removePrefix(private, "Private key:")
	public = removePrefix(public, "Public key:")
	return models.NewKeyPair(public, private), nil
}

func removePrefix(target string, prefix string) string {
	result := strings.TrimPrefix(target, prefix)
	return strings.TrimSpace(result)
}

func (s *BashOsService) DownloadAndInstallXray(version string) error {
	cmd := fmt.Sprintf(
		`bash -c "$(curl -L https://github.com/XTLS/Xray-install/raw/main/install-release.sh)" @ install -u root --version %s`,
		version,
	)
	out, err := s.executor.Execute(cmd)
	if err != nil {
		return err
	}
	fmt.Println(out)
	return nil
}

func (s *BashOsService) GetServerAddr() (*string, error) {
	out, err := s.executor.Execute("hostname -I")
	if err != nil {
		return nil, err
	}
	result := strings.Split(*out, " ")[0]
	return &result, nil
}

func (s *BashOsService) GenerateShortId() (*string, error) {
	out, err := s.executor.Execute("openssl rand -hex 8")
	if err != nil {
		return nil, err
	}
	result := strings.TrimSuffix(*out, "\n")
	return &result, nil
}

func NewBashOsService(executor commands.CmdExecutor) *BashOsService {
	return &BashOsService{executor: executor}
}
