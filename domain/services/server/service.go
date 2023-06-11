package server

import "xraybuilder/models"

type ServerService interface {
	ReadConfig(path string) (*models.ServerConfig, error)
	ReadKeyPair(path string) (*models.KeyPair, error)
	AppendClients(cfg *models.ServerConfig, clients *[]models.ClientDto, streamSettings *models.StreamSettingsObject)
	SetPrivateKey(cfg *models.ServerConfig, keyPair *models.KeyPair)
	SetDestinationAddr(cfg *models.ServerConfig, addr string)
	InflateServerConfig(cfg *models.ServerConfig, clients *[]models.ClientDto, keyPair *models.KeyPair, destination string)
	CurrentUsers(cfg *models.ServerConfig) int
}
