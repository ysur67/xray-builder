package linux

import (
	"os/user"
	commands "xray-builder/domain/commands"
	"xray-builder/internal"
	"xray-builder/models"
)

type LinuxOsService struct {
	XrayConfigPath  string
	XrayKeypairPath string
	executor        commands.CmdExecutor
}

func (s *LinuxOsService) GenerateKeyPair() (*models.KeyPair, error) {
	return s.executor.GenerateKeyPair()
}

func (s *LinuxOsService) GetServerAddr() (string, error) {
	addr, err := s.executor.GetServerAddr()
	return internal.DerefString(addr), err
}

func (s *LinuxOsService) GenerateShortId() (*string, error) {
	return s.executor.GenerateShortId()
}

func (s *LinuxOsService) WriteServerConfig(serverConfig *models.ServerConfig) error {
	return internal.WriteToFile(s.XrayConfigPath, &serverConfig)
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
