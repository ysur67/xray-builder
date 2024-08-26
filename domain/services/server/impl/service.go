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

func (s *ServerServiceImpl) AppendClient(
	cfg *models.ServerConfig,
	client *models.ClientDto,
) {
	inbound := cfg.FirstInbound()
	inbound.Settings.Clients = append(inbound.Settings.Clients, client.Client)
	inbound.StreamSettings.RealitySettings.ShortIds = append(inbound.StreamSettings.RealitySettings.ShortIds, client.ShortId)
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

func (s *ServerServiceImpl) SetupServer(
	cfg *models.ServerConfig,
	keyPair *models.KeyPair,
	destination string,
) {
	s.SetPrivateKey(cfg, keyPair)
	s.SetDestinationAddress(cfg, destination)
}

func (s *ServerServiceImpl) GetUsers(cfg *models.ServerConfig) *[]models.Client {
	return &cfg.FirstInbound().Settings.Clients
}

func (s *ServerServiceImpl) RemoveUser(cfg *models.ServerConfig, userIdOrComment string) *models.Client {
	inbound := cfg.FirstInbound()
	for i, user := range inbound.Settings.Clients {
		if user.Comment == userIdOrComment || user.Id == userIdOrComment {
			inbound.Settings.Clients = append(inbound.Settings.Clients[:i], inbound.Settings.Clients[i+1:]...)
			return &user
		}
	}

	return nil
}

func (s *ServerServiceImpl) ReadKeyPair(path string) (*models.KeyPair, error) {
	var result models.KeyPair
	if err := internal.ReadJson(path, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

func New() *ServerServiceImpl {
	return &ServerServiceImpl{}
}
