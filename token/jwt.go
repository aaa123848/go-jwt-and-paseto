package token

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

type JwtMaker struct {
	secretKey string
}

func NewJwtMaker(secretKey string) Maker {
	return JwtMaker{
		secretKey: secretKey,
	}
}

func (j JwtMaker) CreateToken(username string, duration time.Duration) (string, error) {
	payload, err := NewPayload(username, duration)
	if err != nil {
		return "", err
	}

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, payload) // HS256 = HMAC with SHA-256
	res, err := jwtToken.SignedString([]byte(j.secretKey))
	if err != nil {
		return "", err
	}
	return res, nil
}

func (j JwtMaker) ValidToken(token string) (Payload, error) {
	keyFunc := func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, ErrTokenInvalid
		}
		return []byte(j.secretKey), nil
	}
	jwtToken, err := jwt.ParseWithClaims(token, &Payload{}, keyFunc)
	if err != nil {
		return Payload{}, err
	}
	payload, ok := jwtToken.Claims.(*Payload)
	if !ok {
		return Payload{}, ErrTokenInvalid
	}
	return *payload, nil
}
