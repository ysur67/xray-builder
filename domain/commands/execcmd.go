package commands

import "xraybuilder/models"

type CmdExecutor interface {
	GenerateKeyPair() (*models.KeyPair, error)
	DownloadAndInstallXray(version string) error
	GenerateShortId() (*string, error)
	GetServerAddr() (*string, error)
	RestartXray() error
	SuppressLoginMessage() error
	ApplyIptablesRules() error
	EnableTcpBBR() error
}
