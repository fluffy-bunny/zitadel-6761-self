package authz

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"golang.org/x/text/language"
)

func Test_Instance(t *testing.T) {
	type args struct {
		ctx context.Context
	}
	type res struct {
		instanceID string
		projectID  string
		consoleID  string
	}
	tests := []struct {
		name string
		args args
		res  res
	}{
		{
			"empty context",
			args{
				context.Background(),
			},
			res{
				instanceID: "",
				projectID:  "",
				consoleID:  "",
			},
		},
		{
			"WithInstanceID",
			args{
				WithInstanceID(context.Background(), "id"),
			},
			res{
				instanceID: "id",
				projectID:  "",
				consoleID:  "",
			},
		},
		{
			"WithInstance",
			args{
				WithInstance(context.Background(), &mockInstance{}),
			},
			res{
				instanceID: "instanceID",
				projectID:  "projectID",
				consoleID:  "consoleID",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GetInstance(tt.args.ctx)
			assert.Equal(t, tt.res.instanceID, got.InstanceID())
			assert.Equal(t, tt.res.projectID, got.ProjectID())
			assert.Equal(t, tt.res.consoleID, got.ConsoleClientID())
		})
	}
}

type mockInstance struct{}

func (m *mockInstance) InstanceID() string {
	return "instanceID"
}

func (m *mockInstance) ProjectID() string {
	return "projectID"
}

func (m *mockInstance) ConsoleClientID() string {
	return "consoleID"
}

func (m *mockInstance) ConsoleApplicationID() string {
	return "appID"
}

func (m *mockInstance) DefaultLanguage() language.Tag {
	return language.English
}

func (m *mockInstance) DefaultOrganisationID() string {
	return "orgID"
}

func (m *mockInstance) RequestedDomain() string {
	return "zitadel.cloud"
}

func (m *mockInstance) RequestedHost() string {
	return "zitadel.cloud:443"
}

func (m *mockInstance) SecurityPolicyAllowedOrigins() []string {
	return nil
}
