package bash

import (
	"fmt"
	"strings"
	"xraybuilder/domain/commands"
	"xraybuilder/models"
	"xraybuilder/service/keypair"
)

type BashOsService struct {
	executor commands.CmdExecutor
}

func (s *BashOsService) GenerateKeyPair() (*models.KeyPair, error) {
	out, err := s.executor.Execute("xray x25519")
	if err != nil {
		return nil, err
	}
	keyPair, err := keypair.FromStdOut(*out)
	if err != nil {
		return nil, err
	}
	return keyPair, nil
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

func NewBashOsService(executor commands.CmdExecutor) *BashOsService {
	return &BashOsService{executor: executor}
}
