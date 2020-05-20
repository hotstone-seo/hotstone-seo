package service_test

import (
	"context"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/hotstone-seo/hotstone-seo/internal/api/repository"
	"github.com/hotstone-seo/hotstone-seo/internal/api/repository_mock"
	"github.com/hotstone-seo/hotstone-seo/internal/api/service"
	"github.com/hotstone-seo/hotstone-seo/internal/api/service_mock"
	"github.com/stretchr/testify/require"
)

func TestClientKeyService_Insert(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repoMock := repository_mock.NewMockClientKeyRepo(ctrl)
	auditSvcMock := service_mock.NewMockAuditTrailService(ctrl)
	ctx := context.Background()

	svc := service.ClientKeyServiceImpl{
		ClientKeyRepo:     repoMock,
		AuditTrailService: auditSvcMock,
	}

	newClientKey := repository.ClientKey{Name: "Foo", Prefix: "123", Key: "456"}
	repoMock.EXPECT().Insert(ctx, gomock.Any()).Return(newClientKey, nil)
	auditSvcMock.EXPECT().RecordChanges(ctx, "client_keys", gomock.Any(), repository.Insert, nil, gomock.Any()).Return(int64(1), nil)

	givenClientKey := repository.ClientKey{Name: "Foo"}
	data, err := svc.Insert(ctx, givenClientKey)
	require.NoError(t, err)
	require.Equal(t, "Foo", data.Name)
	require.Equal(t, "123", data.Prefix)
	require.NotEqual(t, "456", data.Key)
	require.Equal(t, 32, len(data.Key))

	time.Sleep(time.Millisecond) // waiting for submitting to audit trails
}

func Test_ExtractClientKey(t *testing.T) {
	tests := []struct {
		name          string
		clientKey     string
		wantPrefix    string
		wantKey       string
		wantKeyHashed string
		wantErr       bool
	}{
		{name: "Empty", clientKey: "", wantErr: true},
		{name: "Bad format", clientKey: "123.8888.34", wantErr: true},
		{name: "No prefix", clientKey: ".777", wantErr: true},
		{name: "No key", clientKey: "123.", wantErr: true},
		{name: "Prefix length != 7", clientKey: "12345.12345678901234567890123456789012", wantErr: true},
		{name: "Key length != 32", clientKey: "1234567.1234567890", wantErr: true},
		{name: "Valid", clientKey: "1234567.12345678901234567890123456789012",
			wantErr: false, wantPrefix: "1234567", wantKey: "12345678901234567890123456789012",
			wantKeyHashed: "e1b85b27d6bcb05846c18e6a48f118e89f0c0587140de9fb3359f8370d0dba08"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotPrefix, gotKey, gotKeyHashed, err := service.ExtractClientKey(tt.clientKey)
			require.Equal(t, tt.wantErr, err != nil)
			require.Equal(t, tt.wantPrefix, gotPrefix)
			require.Equal(t, tt.wantKey, gotKey)
			require.Equal(t, tt.wantKeyHashed, gotKeyHashed)
		})
	}
}
