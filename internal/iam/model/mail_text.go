package model

import (
	"github.com/zitadel/zitadel/internal/eventstore/v1/models"
)

type MailTexts struct {
	Texts   []*MailText
	Default bool
}
type MailText struct {
	models.ObjectRoot

	State        PolicyState
	Default      bool
	MailTextType string
	Language     string
	Title        string
	PreHeader    string
	Subject      string
	Greeting     string
	Text         string
	ButtonText   string
	FooterText   string
}

func (p *MailText) IsValid() bool {
	return p.ObjectRoot.AggregateID != ""
}
