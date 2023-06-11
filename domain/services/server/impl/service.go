package impl

import (
	"xraybuilder/internal"
	"xraybuilder/models"
)

type ServerServiceImpl struct{}

func (s *ServerServiceImpl) ReadConfig(path string) (*models.ServerConfig, error) {
	if path == "" {
		path = "configs/server.template.json"
	}
	config := models.ServerConfig{}
	err := internal.ReadJson(path, &config)
	if err != nil {
		return nil, err
	}
	return &config, nil
}

func (s *ServerServiceImpl) AppendClients(
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

func (s *ServerServiceImpl) SetPrivateKey(
	cfg *models.ServerConfig,
	keyPair *models.KeyPair,
) {
	cfg.FirstInbound().StreamSettings.RealitySettings.PrivateKey = keyPair.Private
}

func (s *ServerServiceImpl) SetDestinationAddress(cfg *models.ServerConfig, addr string) {
	first := cfg.FirstInbound()
	first.StreamSettings.RealitySettings.Dest = addr + ":443"
	first.StreamSettings.RealitySettings.ServerNames = []string{addr}
}

func (s *ServerServiceImpl) InflateServerConfig(cfg *models.ServerConfig, clients *[]models.ClientDto, keyPair *models.KeyPair, destination string) {
	s.AppendClients(
		cfg,
		clients,
		&cfg.FirstInbound().StreamSettings,
	)
	s.SetPrivateKey(cfg, keyPair)
	s.SetDestinationAddress(cfg, destination)
}

func (s *ServerServiceImpl) CurrentUsers(cfg *models.ServerConfig) int {
	return len(cfg.FirstInbound().Settings.Clients)
}

func (s *ServerServiceImpl) ReadKeyPair(path string) (*models.KeyPair, error) {
	var result models.KeyPair
	if err := internal.ReadJson(path, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

func NewServerServiceImpl() *ServerServiceImpl {
	return &ServerServiceImpl{}
}
