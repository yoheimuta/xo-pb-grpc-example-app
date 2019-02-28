package authtoken

import "fmt"

// Parser represents a parser for an auth token string.
type Parser struct {
	secret []byte
}

// NewParser creates a new Parser.
func NewParser(
	secret []byte,
) (Parser, error) {
	if len(secret) == 0 {
		return Parser{}, fmt.Errorf("secret should not be empty")
	}
	return Parser{
		secret: secret,
	}, nil
}

// Parse parses an auth token string.
func (p Parser) Parse(
	tokenString string,
) (
	*Token,
	error,
) {
	// TODO: Return a genuine token.
	return nil, nil
}
