package token

import "time"

type Maker interface {
	CreateToken(string, time.Duration) (string, error)
	ValidToken(s string) (Payload, error)
}
