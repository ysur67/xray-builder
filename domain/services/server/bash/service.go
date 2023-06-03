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

func NewBashServerService() *BashServerService {
	return &BashServerService{}
}
