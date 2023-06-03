package impl

import (
	"xraybuilder/domain/services/osservice"
	"xraybuilder/internal"
	"xraybuilder/models"
)

type ClientCfgServiceImpl struct {
	svc osservice.OsService
}

func (b *ClientCfgServiceImpl) CreateClients(count int) (*[]models.ClientDto, error) {
	result := make([]models.ClientDto, count)
	for ind := 0; ind < count; ind++ {
		shortId, err := b.svc.GenerateShortId()
		if err != nil {
			return &result, err
		}
		result[ind] = *models.NewClient(*shortId)
	}
	return &result, nil
}

func (b *ClientCfgServiceImpl) CreateClientConfig(serverName string, client *models.ClientDto, keyPair *models.KeyPair) (*models.ClientConfig, error) {
	clientConfig := models.ClientConfig{}
	internal.ReadJson("configs/client.template.json", &clientConfig)
	serverAddr, _ := b.svc.GetServerAddr()
	first := clientConfig.FirstOutbound()
	first.Settings.Vnext = models.ClientVnext{
		Address: *serverAddr,
		Port:    443,
		Users: []models.ClientUser{
			{
				ID:         client.Client.ID,
				Flow:       "xtls-rprx-vision",
				Encryption: "none",
			},
		},
	}
	first.StreamSettings.RealitySettings.ShortID = client.ShortId
	first.StreamSettings.RealitySettings.ServerName = serverName
	first.StreamSettings.RealitySettings.PublicKey = keyPair.Pub
	return &clientConfig, nil
}

func (b *ClientCfgServiceImpl) CreateMultipleConfigs(serverName string, clients *[]models.ClientDto, keyPair *models.KeyPair) (*[]models.ClientConfig, error) {
	result := make([]models.ClientConfig, len(*clients))
	for ind, elem := range *clients {
		config, err := b.CreateClientConfig(serverName, &elem, keyPair)
		if err != nil {
			return nil, err
		}
		result[ind] = *config
	}
	return &result, nil
}

func NewClientCfgServiceImpl(svc osservice.OsService) *ClientCfgServiceImpl {
	return &ClientCfgServiceImpl{svc: svc}
}
