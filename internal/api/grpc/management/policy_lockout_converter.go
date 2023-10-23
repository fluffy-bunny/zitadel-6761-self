package management

import (
	"github.com/zitadel/zitadel/internal/domain"
	mgmt "github.com/zitadel/zitadel/pkg/grpc/management"
)

func AddLockoutPolicyToDomain(p *mgmt.AddCustomLockoutPolicyRequest) *domain.LockoutPolicy {
	return &domain.LockoutPolicy{
		MaxPasswordAttempts: uint64(p.MaxPasswordAttempts),
	}
}

func UpdateLockoutPolicyToDomain(p *mgmt.UpdateCustomLockoutPolicyRequest) *domain.LockoutPolicy {
	return &domain.LockoutPolicy{
		MaxPasswordAttempts: uint64(p.MaxPasswordAttempts),
	}
}
