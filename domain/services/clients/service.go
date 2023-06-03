package clients

import "xraybuilder/models"

type ClientCfgsService interface {
	CreateClients(count int) (*[]models.ClientDto, error)
	CreateClientsConfig(serverName string, client *models.ClientDto, keyPair *models.KeyPair) (*models.ClientConfig, error)
}
