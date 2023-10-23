package model

import (
	"github.com/zitadel/zitadel/internal/crypto"
	es_models "github.com/zitadel/zitadel/internal/eventstore/v1/models"
)

type IDPConfig struct {
	es_models.ObjectRoot
	IDPConfigID  string
	Type         IdpConfigType
	Name         string
	StylingType  IDPStylingType
	State        IDPConfigState
	OIDCConfig   *OIDCIDPConfig
	JWTIDPConfig *JWTIDPConfig
}

type OIDCIDPConfig struct {
	es_models.ObjectRoot
	IDPConfigID           string
	ClientID              string
	ClientSecret          *crypto.CryptoValue
	ClientSecretString    string
	Issuer                string
	Scopes                []string
	IDPDisplayNameMapping OIDCMappingField
	UsernameMapping       OIDCMappingField
}

type JWTIDPConfig struct {
	es_models.ObjectRoot
	IDPConfigID  string
	JWTEndpoint  string
	Issuer       string
	KeysEndpoint string
}

type IdpConfigType int32

const (
	IDPConfigTypeOIDC IdpConfigType = iota
	IDPConfigTypeSAML
	IDPConfigTypeJWT
)

type IDPConfigState int32

const (
	IDPConfigStateActive IDPConfigState = iota
	IDPConfigStateInactive
	IDPConfigStateRemoved
)

type IDPStylingType int32

const (
	IDPStylingTypeUnspecified IDPStylingType = iota
	IDPStylingTypeGoogle
)

type OIDCMappingField int32

const (
	OIDCMappingFieldUnspecified OIDCMappingField = iota
	OIDCMappingFieldPreferredLoginName
	OIDCMappingFieldEmail
)

func NewIDPConfig(iamID, idpID string) *IDPConfig {
	return &IDPConfig{ObjectRoot: es_models.ObjectRoot{AggregateID: iamID}, IDPConfigID: idpID}
}

func (idp *IDPConfig) IsValid(includeConfig bool) bool {
	if idp.Name == "" || idp.AggregateID == "" {
		return false
	}
	if !includeConfig {
		return true
	}
	if idp.Type == IDPConfigTypeOIDC && !idp.OIDCConfig.IsValid(true) {
		return false
	}
	return true
}

func (oi *OIDCIDPConfig) IsValid(withSecret bool) bool {
	if withSecret {
		return oi.ClientID != "" && oi.Issuer != "" && oi.ClientSecretString != ""
	}
	return oi.ClientID != "" && oi.Issuer != ""
}

func (oi *OIDCIDPConfig) CryptSecret(crypt crypto.Crypto) error {
	cryptedSecret, err := crypto.Crypt([]byte(oi.ClientSecretString), crypt)
	if err != nil {
		return err
	}
	oi.ClientSecret = cryptedSecret
	return nil
}

func (st IDPStylingType) GetCSSClass() string {
	switch st {
	case IDPStylingTypeGoogle:
		return "google"
	default:
		return ""
	}
}
