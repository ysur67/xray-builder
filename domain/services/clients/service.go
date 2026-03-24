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

func (b *ClientCfgServiceImpl) CreateClient(comment string, network string) (*models.Client, error) {
	shortId, err := b.svc.GenerateShortId()
	if err != nil {
		return nil, err
	}
	return models.NewClient(*shortId, comment, network), nil
}

func (b *ClientCfgServiceImpl) CreateClientConfig(serverName string, client *models.Client, keyPair *models.KeyPair, serverStreamSettings *models.StreamSettingsObject) (*models.ClientConfig, error) {
	clientConfig := models.ClientConfig{}
	internal.ReadJson(b.configsDirectory+"/client.template.json", &clientConfig)
	serverAddr, _ := b.svc.GetServerAddr()
	first := clientConfig.FirstOutbound()

	flow := "xtls-rprx-vision"
	if serverStreamSettings.Network == "xhttp" {
		flow = ""
	}

	vnext := make([]models.ClientVnext, 1)
	vnext[0] = models.ClientVnext{
		Address: serverAddr,
		Port:    443,
		Users: []models.ClientUser{
			{
				Id:         client.Id,
				Flow:       flow,
				Encryption: "none",
				Comment:    client.Comment,
			},
		},
	}
	first.Settings.Vnext = vnext
	first.StreamSettings.Network = serverStreamSettings.Network
	first.StreamSettings.RealitySettings.ShortID = client.ShortId
	first.StreamSettings.RealitySettings.ServerName = serverName
	first.StreamSettings.RealitySettings.PublicKey = keyPair.Pub

	if serverStreamSettings.XhttpSettings != nil {
		first.StreamSettings.XhttpSettings = &struct {
			Path string `json:"path"`
		}{Path: serverStreamSettings.XhttpSettings.Path}
	}

	return &clientConfig, nil
}

func New(svc osservice.OsService) *ClientCfgServiceImpl {
	return &ClientCfgServiceImpl{
		svc:              svc,
		configsDirectory: internal.ResolveConfigPath(),
	}
}
