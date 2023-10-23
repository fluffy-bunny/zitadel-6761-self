package model

import (
	"github.com/zitadel/zitadel/internal/domain"
	iam_model "github.com/zitadel/zitadel/internal/iam/model"
	"github.com/zitadel/zitadel/internal/view/repository"
)

type LabelPolicySearchRequest iam_model.LabelPolicySearchRequest
type LabelPolicySearchQuery iam_model.LabelPolicySearchQuery
type LabelPolicySearchKey iam_model.LabelPolicySearchKey

func (req LabelPolicySearchRequest) GetLimit() uint64 {
	return req.Limit
}

func (req LabelPolicySearchRequest) GetOffset() uint64 {
	return req.Offset
}

func (req LabelPolicySearchRequest) GetSortingColumn() repository.ColumnKey {
	if req.SortingColumn == iam_model.LabelPolicySearchKeyUnspecified {
		return nil
	}
	return LabelPolicySearchKey(req.SortingColumn)
}

func (req LabelPolicySearchRequest) GetAsc() bool {
	return req.Asc
}

func (req LabelPolicySearchRequest) GetQueries() []repository.SearchQuery {
	result := make([]repository.SearchQuery, len(req.Queries))
	for i, q := range req.Queries {
		result[i] = LabelPolicySearchQuery{Key: q.Key, Value: q.Value, Method: q.Method}
	}
	return result
}

func (req LabelPolicySearchQuery) GetKey() repository.ColumnKey {
	return LabelPolicySearchKey(req.Key)
}

func (req LabelPolicySearchQuery) GetMethod() domain.SearchMethod {
	return req.Method
}

func (req LabelPolicySearchQuery) GetValue() interface{} {
	return req.Value
}

func (key LabelPolicySearchKey) ToColumnName() string {
	switch iam_model.LabelPolicySearchKey(key) {
	case iam_model.LabelPolicySearchKeyAggregateID:
		return LabelPolicyKeyAggregateID
	case iam_model.LabelPolicySearchKeyState:
		return LabelPolicyKeyState
	case iam_model.LabelPolicySearchKeyInstanceID:
		return LabelPolicyKeyInstanceID
	case iam_model.LabelPolicySearchKeyOwnerRemoved:
		return LabelPolicyKeyOwnerRemoved

	default:
		return ""
	}
}
