package serverclients

import (
	"strings"
	"xraybuilder/internal"
	"xraybuilder/models"
)

func CreateClients(amount int) ([]*models.ClientDto, error) {
	result := make([]*models.ClientDto, amount)
	for i := 0; i < amount; i++ {
		shortId, _, err := internal.Shellout("openssl rand -hex 8")
		if err != nil {
			return result, err
		}
		shortId = strings.TrimSuffix(shortId, "\n")
		result[i] = models.NewClient(shortId)
	}
	return result, nil
}
