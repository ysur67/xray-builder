package keypair

import (
	"strings"
	"xraybuilder/models"
)

func FromStdOut(out string) (*models.KeyPair, error) {
	values := strings.Split(strings.TrimSuffix(out, "\n"), "\n")
	if len(values) != 2 {
		return nil, &KeyPairServiceError{Type: InvalidResponse}
	}
	private, public := values[0], values[1]
	if private == "" || public == "" {
		return nil, &KeyPairServiceError{Type: InvalidResponse}
	}
	private = removePrefix(private, "Private key:")
	public = removePrefix(public, "Public key:")
	return models.NewKeyPair(public, private), nil
}

func removePrefix(target string, prefix string) string {
	result := strings.TrimPrefix(target, prefix)
	return strings.TrimSpace(result)
}
