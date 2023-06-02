package serverclients

import (
	"strings"
	"xraybuilder/internal"
	"xraybuilder/models"
)

func CreateClients(amount int) (*[]models.ClientDto, error) {
	result := make([]models.ClientDto, amount)
	for ind := 0; ind < amount; ind++ {
		shortId, _, err := internal.Shellout("openssl rand -hex 8")
		if err != nil {
			return &result, err
		}
		shortId = strings.TrimSuffix(shortId, "\n")
		result[ind] = *models.NewClient(shortId)
	}
	return &result, nil
}
