package server

import "xraybuilder/models"

type ServerService interface {
	ReadConfig(path string) (*models.ServerConfig, error)
}
