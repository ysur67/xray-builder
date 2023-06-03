package osservice

import (
	"fmt"
)

type KeyPairServiceErrorType int

const (
	EmptyResponse KeyPairServiceErrorType = iota
	InvalidResponse
)

type KeyPairServiceError struct {
	Type KeyPairServiceErrorType
}

func (e *KeyPairServiceError) Error() string {
	return fmt.Sprintf("ErrorType: %v", e.Type)
}
