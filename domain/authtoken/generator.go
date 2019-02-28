package authtoken

import (
	"fmt"
	"time"
)

// Generator represents a generator of an auth token string.
type Generator struct {
	secret   []byte
	lifetime time.Duration
}

// NewGenerator creates a new Generator.
func NewGenerator(
	secret []byte,
	lifetime time.Duration,
) (*Generator, error) {
	if len(secret) == 0 {
		return nil, fmt.Errorf("secret should not be empty")
	}

	if lifetime == 0 {
		return nil, fmt.Errorf("lifetime should not be zero")
	}

	return &Generator{
		secret:   secret,
		lifetime: lifetime,
	}, nil
}

// Generate generates a new auth token string.
func (g Generator) Generate(
	userID string,
	now time.Time,
) (string, error) {
	if now.IsZero() {
		return "", fmt.Errorf("now should not be zero")
	}

	// TODO: Return a genuine token string.
	return "TODO", nil
}
