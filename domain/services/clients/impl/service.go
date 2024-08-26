package impl

import (
	"xraybuilder/domain/services/osservice"
	"xraybuilder/internal"
	"xraybuilder/models"
)

type ClientCfgServiceImpl struct {
	svc osservice.OsService
}

func (b *ClientCfgServiceImpl) CreateClient(comment string) (*models.ClientDto, error) {
	shortId, err := b.svc.GenerateShortId()
	if err != nil {
		return nil, err
	}
	return models.NewClient(*shortId, comment), nil
}

func (b *ClientCfgServiceImpl) CreateClientConfig(serverName string, client *models.ClientDto, keyPair *models.KeyPair) (*models.ClientConfig, error) {
	clientConfig := models.ClientConfig{}
	internal.ReadJson("configs/client.template.json", &clientConfig)
	serverAddr, _ := b.svc.GetServerAddr()
	first := clientConfig.FirstOutbound()
	vnext := make([]models.ClientVnext, 1)
	vnext[0] = models.ClientVnext{
		Address: serverAddr,
		Port:    443,
		Users: []models.ClientUser{
			{
				Id:         client.Client.Id,
				Flow:       "xtls-rprx-vision",
				Encryption: "none",
				Comment:    client.Client.Comment,
			},
		},
	}
	first.Settings.Vnext = vnext
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

func New(svc osservice.OsService) *ClientCfgServiceImpl {
	return &ClientCfgServiceImpl{svc}
}
