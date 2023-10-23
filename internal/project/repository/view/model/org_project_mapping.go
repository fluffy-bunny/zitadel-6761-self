package model

const (
	OrgProjectMappingKeyProjectID      = "project_id"
	OrgProjectMappingKeyOrgID          = "org_id"
	OrgProjectMappingKeyProjectGrantID = "project_grant_id"
	OrgProjectMappingKeyInstanceID     = "instance_id"
	OrgProjectMappingOwnerRemoved      = "owner_removed"
)

type OrgProjectMapping struct {
	ProjectID      string `json:"-" gorm:"column:project_id;primary_key"`
	OrgID          string `json:"-" gorm:"column:org_id;primary_key"`
	ProjectGrantID string `json:"-" gorm:"column:project_grant_id"`
	InstanceID     string `json:"instanceID" gorm:"column:instance_id"`
}
