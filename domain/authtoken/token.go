package authtoken

import "time"

// Token represents a content of the parsed token string.
type Token struct {
	UserID string
	Iat    time.Time
	Exp    time.Time
}
