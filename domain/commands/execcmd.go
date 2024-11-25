package commands

import "xray-builder/models"

type CmdExecutor interface {
	GenerateKeyPair() (*models.KeyPair, error)
	GenerateShortId() (*string, error)
	GetServerAddr() (*string, error)
	RestartXray() error
}
