package model

import (
	"time"

	"github.com/zitadel/zitadel/internal/domain"
)

type LabelPolicyView struct {
	AggregateID     string
	PrimaryColor    string
	BackgroundColor string
	WarnColor       string
	FontColor       string
	LogoURL         string
	IconURL         string

	PrimaryColorDark    string
	BackgroundColorDark string
	WarnColorDark       string
	FontColorDark       string
	LogoDarkURL         string
	IconDarkURL         string
	FontURL             string

	HideLoginNameSuffix bool
	ErrorMsgPopup       bool
	DisableWatermark    bool

	Default bool

	CreationDate time.Time
	ChangeDate   time.Time
	Sequence     uint64
}

type LabelPolicySearchRequest struct {
	Offset        uint64
	Limit         uint64
	SortingColumn LabelPolicySearchKey
	Asc           bool
	Queries       []*LabelPolicySearchQuery
}

type LabelPolicySearchKey int32

const (
	LabelPolicySearchKeyUnspecified LabelPolicySearchKey = iota
	LabelPolicySearchKeyAggregateID
	LabelPolicySearchKeyState
	LabelPolicySearchKeyInstanceID
	LabelPolicySearchKeyOwnerRemoved
)

type LabelPolicySearchQuery struct {
	Key    LabelPolicySearchKey
	Method domain.SearchMethod
	Value  interface{}
}

type LabelPolicySearchResponse struct {
	Offset      uint64
	Limit       uint64
	TotalResult uint64
	Result      []*LabelPolicyView
	Sequence    uint64
	Timestamp   time.Time
}
