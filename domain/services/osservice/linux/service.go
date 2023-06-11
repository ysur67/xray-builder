package linux

import (
	"fmt"
	"os/user"
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

func (s *LinuxOsService) WriteConfigs(
	serverConfig *models.ServerConfig,
	clientConfigs *[]models.ClientConfig,
	clientStartIndex int,
) {
	internal.WriteToFile(internal.LinuxConfigPath, &serverConfig)
	for ind, elem := range *clientConfigs {
		configIndex := clientStartIndex + ind
		internal.WriteToFile(fmt.Sprintf("client%v.json", configIndex), &elem)
	}
}

func (s *LinuxOsService) SaveKeyPair(pair *models.KeyPair) error {
	if err := internal.WriteToFile(internal.LinuxKeyPairPath, &pair); err != nil {
		return err
	}
	return nil
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

func (s *LinuxOsService) IsSuperUser() (bool, error) {
	currentUser, err := user.Current()
	if err != nil {
		return false, err
	}
	return currentUser.Username == "root", nil
}

func NewLinuxOsService(executor commands.CmdExecutor) *LinuxOsService {
	return &LinuxOsService{executor: executor}
}
