package token

import "time"

// maker is an interface for managing(create and verrify) tokens
type Maker interface {
	//CreateToken creats a new token for a specific username and duration
	CreateToken(username string, duration time.Duration) (string, error)

	//verifytoken checks if the token is valid or not
	VerifyToken(token string) (*Payload, error)
}
