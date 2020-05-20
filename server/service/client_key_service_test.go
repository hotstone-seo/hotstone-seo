package service_test

import (
	"testing"

	"github.com/hotstone-seo/hotstone-seo/server/service"
	"github.com/stretchr/testify/require"
)

// FIXME:
// func TestClientKeyService_Insert(t *testing.T) {
// 	ctrl := gomock.NewController(t)
// 	defer ctrl.Finish()

// 	repoMock := repository_mock.NewMockClientKeyRepo(ctrl)
// 	auditSvcMock := service_mock.NewMockAuditTrailService(ctrl)
// 	ctx := context.Background()

// 	svc := service.ClientKeyServiceImpl{
// 		ClientKeyRepo: repoMock,
// 		AuditTrailService: auditSvcMock,
// 	}

// 	newClientKey := repository.ClientKey{Name: "Foo", Prefix: "123", Key: "456"}
// 	repoMock.EXPECT().Insert(ctx, gomock.Any()).Return(newClientKey, nil)

// 	givenClientKey := repository.ClientKey{Name: "Foo"}
// 	data, err := svc.Insert(ctx, givenClientKey)
// 	require.NoError(t, err)
// 	require.Equal(t, "Foo", data.Name)
// 	require.Equal(t, "123", data.Prefix)
// 	require.NotEqual(t, "456" , data.Key)
// 	require.Equal(t, 32, len(data.Key))
// }

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
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotPrefix, gotKey, gotKeyHashed, err := service.ExtractClientKey(tt.clientKey)
			if !tt.wantErr {
				require.NoError(t, err)
			}
			// require.Equal(t, tt.wantErr, tt.err)
			require.Equal(t, tt.wantPrefix, gotPrefix)
			require.Equal(t, tt.wantKey, gotKey)
			require.Equal(t, tt.wantKeyHashed, gotKeyHashed)
		})
	}
}
