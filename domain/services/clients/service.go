package impl

import (
	"xray-builder/domain/services/osservice"
	"xray-builder/internal"
	"xray-builder/models"
)

type ClientCfgServiceImpl struct {
	configsDirectory string
	svc              osservice.OsService
}

func (b *ClientCfgServiceImpl) CreateClient(comment string) (*models.Client, error) {
	shortId, err := b.svc.GenerateShortId()
	if err != nil {
		return nil, err
	}
	return models.NewClient(*shortId, comment), nil
}

func (b *ClientCfgServiceImpl) CreateClientConfig(serverName string, client *models.Client, keyPair *models.KeyPair) (*models.ClientConfig, error) {
	clientConfig := models.ClientConfig{}
	internal.ReadJson(b.configsDirectory+"/client.template.json", &clientConfig)
	serverAddr, _ := b.svc.GetServerAddr()
	first := clientConfig.FirstOutbound()
	vnext := make([]models.ClientVnext, 1)
	vnext[0] = models.ClientVnext{
		Address: serverAddr,
		Port:    443,
		Users: []models.ClientUser{
			{
				Id:         client.Id,
				Flow:       "xtls-rprx-vision",
				Encryption: "none",
				Comment:    client.Comment,
			},
		},
	}
	first.Settings.Vnext = vnext
	first.StreamSettings.RealitySettings.ShortID = client.ShortId
	first.StreamSettings.RealitySettings.ServerName = serverName
	first.StreamSettings.RealitySettings.PublicKey = keyPair.Pub
	return &clientConfig, nil
}

func New(svc osservice.OsService) *ClientCfgServiceImpl {
	return &ClientCfgServiceImpl{
		svc:              svc,
		configsDirectory: internal.ResolveConfigPath(),
	}
}
