package token

import "errors"

var (
	ErrTokenInvalid = errors.New("not valid token")
	ErrTokenExpired = errors.New("token expired")
)
