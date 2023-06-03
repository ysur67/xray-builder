package linux

import (
	"fmt"
	"strings"
	commands "xraybuilder/domain/commands/bash"
	"xraybuilder/domain/services/osservice"
	"xraybuilder/internal"
	"xraybuilder/models"
)

type LinuxOsService struct {
	executor commands.BashCmdExecutor
}

func (s *LinuxOsService) GenerateKeyPair() (*models.KeyPair, error) {
	out, err := s.executor.GenerateKeyPair()
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

func (s *LinuxOsService) DownloadAndInstallXray(version string) error {
	return s.executor.DownloadAndInstallXray(version)
}

func (s *LinuxOsService) GetServerAddr() (*string, error) {
	return s.executor.GetServerAddr()
}

func (s *LinuxOsService) GenerateShortId() (*string, error) {
	return s.executor.GenerateShortId()
}

func (s *LinuxOsService) WriteConfigs(serverConfig *models.ServerConfig, clientConfigs *[]models.ClientConfig) {
	sysPath := "/usr/local/etc/xray/config.json"
	internal.WriteToFile(sysPath, &serverConfig)
	for ind, elem := range *clientConfigs {
		internal.WriteToFile(fmt.Sprintf("client%v.json", ind), &elem)
	}
}

func (s *LinuxOsService) RestartXray() error {
	return s.executor.RestartXray()
}

func (s *LinuxOsService) SuppressLoginMessage() error {
	return s.executor.SuppressLoginMessage()
}
func (s *LinuxOsService) ApplyIptablesRules() error {
	return s.executor.ApplyIptablesRules()
}
func (s *LinuxOsService) EnableTcpBBR() error {
	return s.executor.EnableTcpBBR()
}

func NewLinuxOsService(executor *commands.BashCmdExecutor) *LinuxOsService {
	return &LinuxOsService{executor: *executor}
}
