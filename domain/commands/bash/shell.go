package commands

import (
	"strings"
	"xray-builder/domain/services/osservice"
	"xray-builder/models"
)

func fromStdOut(value string) (*models.KeyPair, error) {
	values := strings.Split(strings.TrimSuffix(value, "\n"), "\n")
	if len(values) != 2 {
		return nil, &osservice.KeyPairServiceError{Type: osservice.InvalidResponse}
	}
	private, public := values[0], values[1]
	if private == "" || public == "" {
		return nil, &osservice.KeyPairServiceError{Type: osservice.InvalidResponse}
	}
	private = removePrefix(private, "Private key:")
	public = removePrefix(public, "Public key:")
	return models.NewKeyPair(public, private), nil
}

func removePrefix(target string, prefix string) string {
	result := strings.TrimPrefix(target, prefix)
	return strings.TrimSpace(result)
}
