package osservice

import (
	"xraybuilder/models"
)

type OsService interface {
	GenerateKeyPair() (*models.KeyPair, error)
	DownloadAndInstallXray(version string) error
	GetServerAddr() (string, error)
	GenerateShortId() (*string, error)
	WriteConfigs(
		serverConfig *models.ServerConfig,
		clientConfig *models.ClientConfig,
		clientIndex int,
	) error
	SaveKeyPair(pair *models.KeyPair) error
	RestartXray() error
	IsSuperUser() (bool, error)
}
