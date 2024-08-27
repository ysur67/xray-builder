package osservice

import (
	"xraybuilder/models"
)

type OsService interface {
	GenerateKeyPair() (*models.KeyPair, error)
	DownloadAndInstallXray(version string) error
	GetServerAddr() (string, error)
	GenerateShortId() (*string, error)
	WriteServerConfig(serverConfig *models.ServerConfig) error
	SaveKeyPair(pair *models.KeyPair) error
	RestartXray() error
	IsSuperUser() (bool, error)
}
