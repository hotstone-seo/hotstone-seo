package controller_test

import (
	"errors"
	"testing"

	"github.com/hotstone-seo/hotstone-server/mock"
	"github.com/stretchr/testify/require"
	"github.com/typical-go/typical-rest-server/pkg/echokit"

	"github.com/golang/mock/gomock"
	"github.com/hotstone-seo/hotstone-server/app/controller"
	"github.com/hotstone-seo/hotstone-server/app/repository"
)

func TestProviderCntrl_MatchRule(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	svc := mock.NewMockProviderService(ctrl)
	cntrl := controller.ProviderCntrl{ProviderService: svc}
	t.Run("WHEN invalid json body", func(t *testing.T) {
		_, err := echokit.DoPOST(cntrl.MatchRule, "/", `{invalid`)
		require.EqualError(t, err, "code=400, message=Syntax error: offset=2, error=invalid character 'i' looking for beginning of object key string")
	})
	t.Run("WHEN error match rule", func(t *testing.T) {
		svc.EXPECT().MatchRule(gomock.Any()).Return(nil, errors.New("some-error"))
		_, err := echokit.DoPOST(cntrl.MatchRule, "/", `{"path":"some-path"}`)
		require.EqualError(t, err, "code=422, message=some-error")
	})
	t.Run("WHEN okay", func(t *testing.T) {
		svc.EXPECT().MatchRule(gomock.Any()).Return(&repository.Rule{ID: 12345}, nil)
		rec, err := echokit.DoPOST(cntrl.MatchRule, "/", `{"path":"some-path"}`)
		require.NoError(t, err)
		require.Equal(t, 200, rec.Code)
		require.Equal(t, "{\"id\":12345,\"name\":\"\",\"url_pattern\":\"\",\"data_source_id\":null,\"updated_at\":\"0001-01-01T00:00:00Z\",\"created_at\":\"0001-01-01T00:00:00Z\"}\n", rec.Body.String())
	})
}

func TestProviderCntrl_RetrieveData(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	svc := mock.NewMockProviderService(ctrl)
	cntrl := controller.ProviderCntrl{ProviderService: svc}
	t.Run("WHEN invalid json body", func(t *testing.T) {
		_, err := echokit.DoPOST(cntrl.RetrieveData, "/", `{invalid`)
		require.EqualError(t, err, "code=400, message=Syntax error: offset=2, error=invalid character 'i' looking for beginning of object key string")
	})
	t.Run("WHEN error match rule", func(t *testing.T) {
		svc.EXPECT().RetrieveData(gomock.Any()).Return(nil, errors.New("some-error"))
		_, err := echokit.DoPOST(cntrl.RetrieveData, "/", `{"rule_id": 99999}`)
		require.EqualError(t, err, "code=422, message=some-error")
	})
	t.Run("WHEN okay", func(t *testing.T) {
		svc.EXPECT().RetrieveData(gomock.Any()).Return(map[string]string{"name": "CGK", "province": "Banten"}, nil)
		rec, err := echokit.DoPOST(cntrl.RetrieveData, "/", `{"rule_id": 99999}`)
		require.NoError(t, err)
		require.Equal(t, 200, rec.Code)
		require.Equal(t, "{\"name\":\"CGK\",\"province\":\"Banten\"}\n", rec.Body.String())
	})
}
