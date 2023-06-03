package clients

import "xraybuilder/models"

type ClientCfgsService interface {
	CreateClients(count int) (*[]models.ClientDto, error)
	CreateClientConfig(serverName string, client *models.ClientDto, keyPair *models.KeyPair) (*models.ClientConfig, error)
	CreateMultipleConfigs(serverName string, clients *[]models.ClientDto, keyPair *models.KeyPair) (*[]models.ClientConfig, error)
}
