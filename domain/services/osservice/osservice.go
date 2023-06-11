package osservice

import (
	"xraybuilder/models"
)

type OsService interface {
	GenerateKeyPair() (*models.KeyPair, error)
	DownloadAndInstallXray(version string) error
	GetServerAddr() (*string, error)
	GenerateShortId() (*string, error)
	WriteConfigs(serverConfig *models.ServerConfig, clientConfigs *[]models.ClientConfig, clientStartIndex int)
	SaveKeyPair(pair *models.KeyPair) error
	RestartXray() error
	SuppressLoginMessage() error
	ApplyIptablesRules() error
	EnableTcpBBR() error
	IsSuperUser() (bool, error)
}
