package linux

import (
	"fmt"
	commands "xraybuilder/domain/commands"
	"xraybuilder/internal"
	"xraybuilder/models"
)

type LinuxOsService struct {
	executor commands.CmdExecutor
}

func (s *LinuxOsService) GenerateKeyPair() (*models.KeyPair, error) {
	return s.executor.GenerateKeyPair()
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
