package bash

import (
	"xraybuilder/internal"
	"xraybuilder/models"
)

type BashServerService struct{}

func (s *BashServerService) ReadConfig(path string) (*models.ServerConfig, error) {
	if path == "" {
		path = "server.template.json"
	}
	config := models.ServerConfig{}
	err := internal.ReadJson(path, &config)
	if err != nil {
		return nil, err
	}
	return &config, nil
}

func (s *BashServerService) AppendClients(
	cfg *models.ServerConfig,
	clients *[]models.ClientDto,
	streamSettings *models.StreamSettingsObject,
) {
	serverClients := make([]models.Client, len(*clients))
	shortIds := make([]string, len(*clients))
	for ind, elem := range *clients {
		serverClients[ind] = elem.Client
		shortIds[ind] = elem.ShortId
	}
	first := cfg.FirstInbound()
	first.Settings.Clients = append(first.Settings.Clients, serverClients...)
	first.StreamSettings.RealitySettings.ShortIds = append(first.StreamSettings.RealitySettings.ShortIds, shortIds...)
}

func (s *BashServerService) SetPrivateKey(
	cfg *models.ServerConfig,
	keyPair *models.KeyPair,
) {
	cfg.FirstInbound().StreamSettings.RealitySettings.PrivateKey = keyPair.Private
}

func (s *BashServerService) SetDestinationAddress(cfg *models.ServerConfig, addr string) {
	first := cfg.FirstInbound()
	first.StreamSettings.RealitySettings.Dest = addr + ":443"
	first.StreamSettings.RealitySettings.ServerNames = []string{addr}
}

func (s *BashServerService) InflateServerConfig(cfg *models.ServerConfig, clients *[]models.ClientDto, keyPair *models.KeyPair, destination string) {
	s.AppendClients(
		cfg,
		clients,
		&cfg.FirstInbound().StreamSettings,
	)
	s.SetPrivateKey(cfg, keyPair)
	s.SetDestinationAddress(cfg, destination)
}

func NewBashServerService() *BashServerService {
	return &BashServerService{}
}
