package token

import (
	"errors"
	"time"

	"github.com/o1egl/paseto"
	"golang.org/x/crypto/chacha20poly1305"
)

type PastoMaker struct {
	paseto       *paseto.V2
	symmetricKey string
}

func NewPastoMaker(secret string) (Maker, error) {
	if len(secret) != chacha20poly1305.KeySize {
		return PastoMaker{}, errors.New("symmetricKey should be fulfil chacha20")
	}

	return PastoMaker{
		paseto:       paseto.NewV2(),
		symmetricKey: secret,
	}, nil
}

func (p PastoMaker) CreateToken(s string, duration time.Duration) (string, error) {
	payload, err := NewPayload(s, duration)
	if err != nil {
		return "", err
	}

	return p.paseto.Encrypt([]byte(p.symmetricKey), &payload, nil)

}

func (p PastoMaker) ValidToken(token string) (Payload, error) {
	payload := Payload{}
	err := p.paseto.Decrypt(token, []byte(p.symmetricKey), &payload, nil)
	if err != nil {
		return payload, errors.New("Valid failed")
	}

	return payload, nil
}
