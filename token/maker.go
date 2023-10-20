package token

import (
	"fmt"
	"time"

	"github.com/aead/chacha20poly1305"
	"github.com/o1egl/paseto"
)

// Doubt
// General token maker interface, so that we can
// type Make interface {
// }

type PasetoMaker struct {
	paseto *paseto.V2

	// use symmetric key, since, we want to use the token locally for the api, see #19
	symmetricKey []byte
}

func NewPasetoMaker(symmetricKey string) (*PasetoMaker, error) {
	if len(symmetricKey) != chacha20poly1305.KeySize {
		return nil, fmt.Errorf("invalid key must be exactly %d characters", chacha20poly1305.KeySize)
	}

	pasetoMaker := &PasetoMaker{
		paseto:       paseto.NewV2(),
		symmetricKey: []byte(symmetricKey),
	}

	return pasetoMaker, nil
}

// creates new token
func (pasetoMaker *PasetoMaker) CreateToken(username string, userId int64, duration time.Duration) (string, error) {
	payload, err := NewPayload(username, userId, duration)
	if err != nil {
		return "", err
	}

	// footer is nil, and is optional
	return pasetoMaker.paseto.Encrypt(pasetoMaker.symmetricKey, payload, nil)
}

// checks if the token is valid or not
func (pasetoMaker *PasetoMaker) VerifyToken(token string) (*Payload, error) {
	payload := &Payload{}

	err := pasetoMaker.paseto.Decrypt(token, pasetoMaker.symmetricKey, payload, nil)
	if err != nil {
		return nil, ErrInvalidToken
	}

	err = payload.Valid()
	if err != nil {
		return nil, err
	}

	return payload, nil
}
