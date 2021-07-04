package token

import (
	"time"

	"github.com/google/uuid"
)

type Payload struct {
	ID         uuid.UUID
	UserName   string
	StartTime  time.Time
	ExpireTime time.Time
}

func NewPayload(username string, duration time.Duration) (Payload, error) {
	uid, err := uuid.NewRandom()
	if err != nil {
		return Payload{}, err
	}
	now := time.Now()
	expire := now.Add(duration)

	return Payload{
		ID:         uid,
		UserName:   username,
		StartTime:  now,
		ExpireTime: expire,
	}, nil
}

func (p Payload) Valid() error {
	if time.Now().After(p.ExpireTime) {
		return ErrTokenExpired
	}

	return nil
}
