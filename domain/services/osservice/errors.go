package osservice

import (
	"fmt"
)

type KeyPairServiceErrorType int

const (
	InvalidResponse KeyPairServiceErrorType = iota
)

type KeyPairServiceError struct {
	Type KeyPairServiceErrorType
}

func (e *KeyPairServiceError) Error() string {
	return fmt.Sprintf("ErrorType: %v", e.Type)
}
