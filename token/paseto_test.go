package token

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func createPastoMaker(t *testing.T) Maker {
	secret := RandString(32)
	maker, err := NewPastoMaker(secret)
	require.NoError(t, err)
	require.NotEmpty(t, maker)
	return maker
}

func TestPasetoToken(t *testing.T) {
	maker := createPastoMaker(t)
	username := RandString(5)
	duration, _ := time.ParseDuration("24h")
	token, err := maker.CreateToken(username, duration)
	require.NoError(t, err)
	require.NotZero(t, token)

	payload, err := maker.ValidToken(token)
	require.NoError(t, err)
	require.NotEmpty(t, payload)
	require.Equal(t, payload.UserName, username)
	fmt.Println(payload)
}

func TestExpiredToken(t *testing.T) {
	maker := createPastoMaker(t)
	username := RandString(5)
	duration, _ := time.ParseDuration("-24h")
	token, err := maker.CreateToken(username, duration)
	require.NoError(t, err)

	payload, err := maker.ValidToken(token)
	require.NoError(t, err)
	err = payload.Valid()
	require.Error(t, err)
	require.EqualError(t, err, ErrTokenExpired.Error())
}
