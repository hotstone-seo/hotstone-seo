package service_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/golang/mock/gomock"
	"github.com/hotstone-seo/hotstone-seo/server/repository"
	"github.com/hotstone-seo/hotstone-seo/server/repository_mock"
	"github.com/hotstone-seo/hotstone-seo/server/service"
	"github.com/hotstone-seo/hotstone-seo/server/service_mock"
)

func TestClientKeyService_Insert(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repoMock := repository_mock.NewMockClientKeyRepo(ctrl)
	auditSvcMock := service_mock.NewMockAuditTrailService(ctrl)
	ctx := context.Background()

	svc := service.ClientKeyServiceImpl{
		ClientKeyRepo: repoMock,
		AuditTrailService: auditSvcMock,
	}

	newClientKey := repository.ClientKey{Name: "Foo", Prefix: "123", Key: "456"}
	repoMock.EXPECT().Insert(ctx, gomock.Any()).Return(newClientKey, nil)

	givenClientKey := repository.ClientKey{Name: "Foo"}
	data, err := svc.Insert(ctx, givenClientKey)
	require.NoError(t, err)
	require.Equal(t, "Foo", data.Name)
	require.Equal(t, "123", data.Prefix)
	require.NotEqual(t, "456" , data.Key)
	require.Equal(t, 32, len(data.Key))
}
