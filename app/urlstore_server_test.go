package app

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/hotstone-seo/hotstone-server/app/repository"
	"github.com/hotstone-seo/hotstone-server/mock"
	"github.com/stretchr/testify/require"
)

func TestURLStoreServerImpl_Sync(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	urlStoreSyncSvcMock := mock.NewMockURLStoreSyncService(ctrl)
	urlStore := repository.InitURLStore()

	// urlStoreServer := NewURLStoreServer(urlStoreSyncSvcMock)
	urlStoreServer := &URLStoreServerImpl{
		URLStoreSyncService: urlStoreSyncSvcMock,
		urlStore:            urlStore,
		latestVersion:       -1,
	}

	list1And2URLStoreSync := []*repository.URLStoreSync{
		&repository.URLStoreSync{Version: 1, Operation: "INSERT", RuleID: 1, LatestURLPattern: "/url/1"},
		&repository.URLStoreSync{Version: 2, Operation: "UPDATE", RuleID: 1, LatestURLPattern: "/url/1update"},
	}
	// list3And4URLStoreSync := []*repository.URLStoreSync{
	// 	&repository.URLStoreSync{Version: 3, Operation: "INSERT", RuleID: 2, LatestURLPattern: "/url/b"},
	// 	&repository.URLStoreSync{Version: 4, Operation: "UPDATE", RuleID: 2, LatestURLPattern: "/url/bupdate"},
	// }

	t.Run("WHEN first sync (s.latestVersion < latestVersionSyncDB)", func(t *testing.T) {
		require.Equal(t, 0, urlStore.Count())

		ctx := context.Background()
		urlStoreSyncSvcMock.EXPECT().GetListDiff(ctx, gomock.Eq(int64(-1))).Return(list1And2URLStoreSync, nil)
		urlStoreSyncSvcMock.EXPECT().GetLatestVersion(ctx).Return(int64(len(list1And2URLStoreSync)), nil)

		err := urlStoreServer.Sync()
		require.NoError(t, err)

		require.Equal(t, 1, urlStore.Count())
	})

	// ruleSvcMock := mock.NewMockRuleService(ctrl)
	// ruleCntrl := controller.RuleCntrl{
	// 	RuleService: ruleSvcMock,
	// }
	// t.Run("WHEN invalid rule request", func(t *testing.T) {
	// 	_, err := echokit.DoPOST(ruleCntrl.Create, "/", `{ "name": "", "url_pattern": ""}`)
	// 	require.EqualError(t, err, "code=400, message=Key: 'Rule.Name' Error:Field validation for 'Name' failed on the 'required' tag\nKey: 'Rule.UrlPattern' Error:Field validation for 'UrlPattern' failed on the 'required' tag")
	// })
	// t.Run("WHEN invalid json format", func(t *testing.T) {
	// 	_, err := echokit.DoPOST(ruleCntrl.Create, "/", `invalid`)
	// 	require.EqualError(t, err, "code=400, message=Syntax error: offset=1, error=invalid character 'i' looking for beginning of value")
	// })
	// t.Run("WHEN insert error", func(t *testing.T) {
	// 	ruleSvcMock.EXPECT().Insert(gomock.Any(), gomock.Any()).Return(int64(-1), errors.New("some-insert-error"))
	// 	_, err := echokit.DoPOST(ruleCntrl.Create, "/", `{ "name": "some-name", "url_pattern": "some-pattern", "data_source_id":1}`)
	// 	require.EqualError(t, err, "code=422, message=some-insert-error")
	// })
	// t.Run("WHEN insert success", func(t *testing.T) {
	// 	ruleSvcMock.EXPECT().Insert(gomock.Any(), gomock.Any()).Return(int64(999), nil)
	// 	rr, err := echokit.DoPOST(ruleCntrl.Create, "/", `{ "name": "some-name", "url_pattern": "some-pattern", "data_source_id":1}`)
	// 	require.NoError(t, err)
	// 	require.Equal(t, http.StatusCreated, rr.Code)
	// 	require.Equal(t, "{\"message\":\"Success create new rule #999\"}\n", rr.Body.String())
	// })
}
