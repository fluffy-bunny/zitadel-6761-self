package domain

import (
	"github.com/zitadel/zitadel/internal/crypto"
	"github.com/zitadel/zitadel/internal/eventstore/v1/models"
)

type APIApp struct {
	models.ObjectRoot

	AppID              string
	AppName            string
	ClientID           string
	ClientSecret       *crypto.CryptoValue
	ClientSecretString string
	AuthMethodType     APIAuthMethodType

	State AppState
}

func (a *APIApp) GetApplicationName() string {
	return a.AppName
}

func (a *APIApp) GetState() AppState {
	return a.State
}

type APIAuthMethodType int32

const (
	APIAuthMethodTypeBasic APIAuthMethodType = iota
	APIAuthMethodTypePrivateKeyJWT
)

func (a *APIApp) IsValid() bool {
	return a.AppName != ""
}

func (a *APIApp) setClientID(clientID string) {
	a.ClientID = clientID
}

func (a *APIApp) setClientSecret(clientSecret *crypto.CryptoValue) {
	a.ClientSecret = clientSecret
}

func (a *APIApp) requiresClientSecret() bool {
	return a.AuthMethodType == APIAuthMethodTypeBasic
}

func (a *APIApp) GenerateClientSecretIfNeeded(generator crypto.Generator) (secret string, err error) {
	if a.AuthMethodType == APIAuthMethodTypePrivateKeyJWT {
		return "", nil
	}
	a.ClientSecret, secret, err = NewClientSecret(generator)
	if err != nil {
		return "", err
	}
	return secret, nil
}
