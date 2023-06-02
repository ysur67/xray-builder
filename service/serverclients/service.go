package serverclients

import (
	"strings"
	"xraybuilder/internal"
	"xraybuilder/models"
)

func CreateClients(amount int) (*[]models.ClientDto, error) {
	result := make([]models.ClientDto, amount)
	for ind := 0; ind < amount; ind++ {
		shortId, _, err := internal.Shellout("openssl rand -hex 8")
		if err != nil {
			return &result, err
		}
		shortId = strings.TrimSuffix(shortId, "\n")
		result[ind] = *models.NewClient(shortId)
	}
	return &result, nil
}

func CreateClientConfig(
	cfg *models.ServerConfig,
	client *models.ClientDto,
	keyPair *models.KeyPair,
) *models.ClientConfig {
	clientConfig := models.ClientConfig{}
	internal.ReadJson("client.template.json", &clientConfig)
	serverAddr, _ := internal.GetServerAddr()
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
	first.StreamSettings.RealitySettings.ServerName = cfg.FirstInbound().StreamSettings.RealitySettings.ServerNames[0]
	first.StreamSettings.RealitySettings.PublicKey = keyPair.Pub
	return &clientConfig
}
