package impl

import (
	"xraybuilder/internal"
	"xraybuilder/models"
)

type ServerService struct{}

func (s *ServerService) ReadConfig(path string) (*models.ServerConfig, error) {
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

func (s *ServerService) AppendClient(
	cfg *models.ServerConfig,
	client *models.Client,
) {
	inbound := cfg.FirstInbound()
	inbound.Settings.Clients = append(inbound.Settings.Clients, *client)
	inbound.StreamSettings.RealitySettings.ShortIds = append(inbound.StreamSettings.RealitySettings.ShortIds, client.ShortId)
}

func (s *ServerService) SetPrivateKey(
	cfg *models.ServerConfig,
	keyPair *models.KeyPair,
) {
	cfg.FirstInbound().StreamSettings.RealitySettings.PrivateKey = keyPair.Private
}

func (s *ServerService) SetDestinationAddress(cfg *models.ServerConfig, addr string) {
	first := cfg.FirstInbound()
	first.StreamSettings.RealitySettings.Dest = addr + ":443"
	first.StreamSettings.RealitySettings.ServerNames = []string{addr}
}

func (s *ServerService) SetupServer(
	cfg *models.ServerConfig,
	keyPair *models.KeyPair,
	destination string,
) {
	s.SetPrivateKey(cfg, keyPair)
	s.SetDestinationAddress(cfg, destination)
}

func (s *ServerService) GetUsers(cfg *models.ServerConfig) *[]models.Client {
	return &cfg.FirstInbound().Settings.Clients
}

func (s *ServerService) GetUser(cfg *models.ServerConfig, userIdOrComment string) *models.Client {
	inbound := cfg.FirstInbound()
	for _, user := range inbound.Settings.Clients {
		if user.Comment == userIdOrComment || user.Id == userIdOrComment {
			return &user
		}
	}

	return nil
}

func (s *ServerService) RemoveUser(cfg *models.ServerConfig, userIdOrComment string) *models.Client {
	inbound := cfg.FirstInbound()
	for i, user := range inbound.Settings.Clients {
		if user.Comment == userIdOrComment || user.Id == userIdOrComment {
			inbound.Settings.Clients = internal.RemoveByIndex(inbound.Settings.Clients, i)

			inbound.StreamSettings.RealitySettings.ShortIds = internal.Remove(
				inbound.StreamSettings.RealitySettings.ShortIds,
				user.ShortId,
			)

			return &user
		}
	}

	return nil
}

func (s *ServerService) ReadKeyPair(path string) (*models.KeyPair, error) {
	var result models.KeyPair
	if err := internal.ReadJson(path, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

func New() *ServerService {
	return &ServerService{}
}
