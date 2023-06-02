package serverconfig

import (
	"xraybuilder/models"
)

func AppendClients(
	serverConfig *models.ServerConfig,
	clients *[]models.ClientDto,
	streamSettings *models.StreamSettingsObject,
) {
	serverClients := make([]models.Client, len(*clients))
	for ind, elem := range *clients {
		serverClients[ind] = elem.Client
	}
	first := &serverConfig.Inbounds[0]
	first.Settings.Clients = serverClients
}

func SetPrivateKey(
	serverConfig *models.ServerConfig,
	keyPair *models.KeyPair,
) {
	serverConfig.Inbounds[0].StreamSettings.RealitySettings.PrivateKey = keyPair.Private
}
