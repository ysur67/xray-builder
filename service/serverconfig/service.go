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
	shortIds := make([]string, len(*clients))
	for ind, elem := range *clients {
		serverClients[ind] = elem.Client
		shortIds[ind] = elem.ShortId
	}
	first := serverConfig.FirstInbound()
	first.Settings.Clients = append(first.Settings.Clients, serverClients...)
	first.StreamSettings.RealitySettings.ShortIds = append(first.StreamSettings.RealitySettings.ShortIds, shortIds...)
}

func SetPrivateKey(
	serverConfig *models.ServerConfig,
	keyPair *models.KeyPair,
) {
	serverConfig.FirstInbound().StreamSettings.RealitySettings.PrivateKey = keyPair.Private
}

func SetDestinationAddress(serverConfig *models.ServerConfig, addr string) {
	first := serverConfig.FirstInbound()
	first.StreamSettings.RealitySettings.Dest = addr + ":443"
	first.StreamSettings.RealitySettings.ServerNames = []string{addr}
}
