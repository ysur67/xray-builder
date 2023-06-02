package serverconfig

import (
	"encoding/json"
	"io/ioutil"
	"testing"
	"xraybuilder/internal"
	"xraybuilder/models"
)

func config() *models.ServerConfig {
	file, _ := ioutil.ReadFile(internal.RootDir() + "/configs/server.template.json")
	cfg := models.ServerConfig{}
	json.Unmarshal([]byte(file), &cfg)
	return &cfg
}

func TestEmptyAppend(t *testing.T) {
	config := config()
	var clients []models.ClientDto
	AppendClients(config, &clients, &config.Inbounds[0].StreamSettings)
	if len(config.Inbounds[0].Settings.Clients) != 0 {
		t.Error("Expected no clients found", len(config.Inbounds))
	}
}

func TestAppend(t *testing.T) {
	config := config()
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
	AppendClients(config, &clients, &config.Inbounds[0].StreamSettings)
	if len(config.Inbounds[0].Settings.Clients) != len(clients) {
		t.Errorf("Expected %v clients found %v", len(clients), len(config.Inbounds[0].Settings.Clients))
	}
}
