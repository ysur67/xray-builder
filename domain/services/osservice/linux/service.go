package linux

import (
	"fmt"
	"os/user"
	commands "xraybuilder/domain/commands"
	"xraybuilder/internal"
	"xraybuilder/models"
)

type LinuxOsService struct {
	XrayConfigPath  string
	XrayKeypairPath string
	executor        commands.CmdExecutor
}

func (s *LinuxOsService) GenerateKeyPair() (*models.KeyPair, error) {
	return s.executor.GenerateKeyPair()
}

func (s *LinuxOsService) DownloadAndInstallXray(version string) error {
	return s.executor.DownloadAndInstallXray(version)
}

func derefString(s *string) string {
	if s != nil {
		return *s
	}

	return ""
}

func (s *LinuxOsService) GetServerAddr() (string, error) {
	addr, err := s.executor.GetServerAddr()
	return derefString(addr), err
}

func (s *LinuxOsService) GenerateShortId() (*string, error) {
	return s.executor.GenerateShortId()
}

func (s *LinuxOsService) WriteConfigs(
	serverConfig *models.ServerConfig,
	clientConfig *models.ClientConfig,
	configIndex int,
) error {
	err := internal.WriteToFile(s.XrayConfigPath, &serverConfig)
	if err != nil {
		return err
	}

	fname := fmt.Sprintf("client%v.json", configIndex)
	err = internal.WriteToFile(fname, clientConfig)
	if err != nil {
		return err
	}

	return nil
}

func (s *LinuxOsService) SaveKeyPair(pair *models.KeyPair) error {
	if err := internal.WriteToFile(s.XrayKeypairPath, &pair); err != nil {
		return err
	}
	return nil
}

func (s *LinuxOsService) RestartXray() error {
	return s.executor.RestartXray()
}

func (s *LinuxOsService) IsSuperUser() (bool, error) {
	currentUser, err := user.Current()
	if err != nil {
		return false, err
	}
	return currentUser.Username == "root", nil
}

func New(xrayConfigPath string, xrayKeypairPath string, executor commands.CmdExecutor) *LinuxOsService {
	return &LinuxOsService{
		XrayConfigPath:  xrayConfigPath,
		XrayKeypairPath: xrayKeypairPath,
		executor:        executor,
	}
}
