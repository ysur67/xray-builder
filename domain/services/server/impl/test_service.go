package impl

import (
	"testing"
	"xraybuilder/internal"
	"xraybuilder/models"
)

func config() *models.ServerConfig {
	result := models.ServerConfig{}
	internal.ReadJson(internal.RootDir()+"/configs/server.template.json", &result)
	return &result
}

func service() *ServerServiceImpl {
	return &ServerServiceImpl{}
}

func TestEmptyAppend(t *testing.T) {
	cfg := config()
	svc := service()
	var clients []models.ClientDto
	svc.AppendClients(cfg, &clients, &cfg.Inbounds[0].StreamSettings)
	if len(cfg.Inbounds[0].Settings.Clients) != 0 {
		t.Error("Expected no clients found", len(cfg.Inbounds))
	}
}

func TestAppend(t *testing.T) {
	cfg := config()
	svc := service()
	clients := make([]models.ClientDto, 5)
	for ind := 0; ind < len(clients); ind++ {
		clients[ind] = models.ClientDto{
			Client: models.Client{
				ID:   "asdfadsfd",
				Flow: "asdfasdfdsa",
			},
			ShortId: "jopajopajopa",
		}
	}
	svc.AppendClients(cfg, &clients, &cfg.Inbounds[0].StreamSettings)
	if len(cfg.Inbounds[0].Settings.Clients) != len(clients) {
		t.Errorf("Expected %v clients found %v", len(clients), len(cfg.Inbounds[0].Settings.Clients))
	}
}

func TestSetPrivateKey(t *testing.T) {
	cfg := config()
	svc := service()
	keyPair := models.KeyPair{
		Pub:     "pub",
		Private: "private",
	}
	svc.SetPrivateKey(cfg, &keyPair)
	if cfg.Inbounds[0].StreamSettings.RealitySettings.PrivateKey != keyPair.Private {
		t.Error("Expected private key to be set, got invalid value instead")
	}
}

func TestSetDestinationAddress(t *testing.T) {
	cfg := config()
	addr := "https://rkn.gov.ru/"
	svc := service()
	svc.SetDestinationAddress(cfg, addr)
	if cfg.Inbounds[0].StreamSettings.RealitySettings.Dest != addr+":443" {
		t.Error("Expected dest to be set, got value invalid instead")
	}
}
