package controller_test

import (
	"errors"
	"net/http"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/hotstone-seo/hotstone-seo/server/controller"
	"github.com/hotstone-seo/hotstone-seo/server/mock_service"
	"github.com/hotstone-seo/hotstone-seo/server/repository"
	"github.com/stretchr/testify/require"
	"github.com/typical-go/typical-rest-server/pkg/echotest"
)

func TestRuleController_Create(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	ruleSvcMock := mock_service.NewMockRuleService(ctrl)
	ruleCntrl := controller.RuleCntrl{
		RuleService: ruleSvcMock,
	}
	t.Run("WHEN invalid rule request", func(t *testing.T) {
		_, err := echotest.DoPOST(ruleCntrl.Create, "/", `{ "name": "", "url_pattern": ""}`, nil)
		require.EqualError(t, err, "code=400, message=Key: 'Rule.Name' Error:Field validation for 'Name' failed on the 'required' tag\nKey: 'Rule.UrlPattern' Error:Field validation for 'UrlPattern' failed on the 'required' tag")
	})
	t.Run("WHEN invalid json format", func(t *testing.T) {
		_, err := echotest.DoPOST(ruleCntrl.Create, "/", `invalid`, nil)
		require.EqualError(t, err, "code=400, message=Syntax error: offset=1, error=invalid character 'i' looking for beginning of value")
	})
	t.Run("WHEN insert error", func(t *testing.T) {
		ruleSvcMock.EXPECT().Insert(gomock.Any(), gomock.Any()).Return(int64(-1), errors.New("some-insert-error"))
		_, err := echotest.DoPOST(ruleCntrl.Create, "/", `{ "name": "some-name", "url_pattern": "http://some-pattern", "data_source_id":1}`, nil)
		require.EqualError(t, err, "code=422, message=some-insert-error")
	})
	t.Run("WHEN insert success", func(t *testing.T) {
		ruleSvcMock.EXPECT().Insert(gomock.Any(), gomock.Any()).Return(int64(999), nil)
		rr, err := echotest.DoPOST(ruleCntrl.Create, "/", `{ "name": "some-name", "url_pattern": "http://some-pattern", "data_source_id":1}`, nil)
		require.NoError(t, err)
		require.Equal(t, http.StatusCreated, rr.Code)
		require.Equal(t, "{\"id\":999,\"name\":\"some-name\",\"url_pattern\":\"http://some-pattern\",\"data_source_id\":1,\"updated_at\":\"0001-01-01T00:00:00Z\",\"created_at\":\"0001-01-01T00:00:00Z\",\"status\":\"\",\"change_status_at\":\"0001-01-01T00:00:00Z\"}\n", rr.Body.String())
	})
}

func TestRuleController_Find(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	ruleSvcMock := mock_service.NewMockRuleService(ctrl)
	ruleCntrl := controller.RuleCntrl{
		RuleService: ruleSvcMock,
	}
	t.Run("WHEN retrieved error", func(t *testing.T) {
		ruleSvcMock.EXPECT().Find(gomock.Any(), gomock.Any()).Return(nil, errors.New("retrieve error"))
		_, err := echotest.DoGET(ruleCntrl.Find, "/", nil)
		require.EqualError(t, err, "code=500, message=retrieve error")
	})
	t.Run("WHEN successful", func(t *testing.T) {
		ruleSvcMock.EXPECT().Find(gomock.Any(), gomock.Any()).Return(
			[]*repository.Rule{
				&repository.Rule{
					ID:         999,
					Name:       "Airport Rule",
					UrlPattern: "/airport",
				},
			},
			nil,
		)
		rr, err := echotest.DoGET(ruleCntrl.Find, "/", nil)
		require.NoError(t, err)
		require.Equal(t, http.StatusOK, rr.Code)
		require.Equal(t, "[{\"id\":999,\"name\":\"Airport Rule\",\"url_pattern\":\"/airport\",\"data_source_id\":null,\"updated_at\":\"0001-01-01T00:00:00Z\",\"created_at\":\"0001-01-01T00:00:00Z\",\"status\":\"\",\"change_status_at\":\"0001-01-01T00:00:00Z\"}]\n", rr.Body.String())
	})
}

func TestRuleController_FindOne(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	ruleSvcMock := mock_service.NewMockRuleService(ctrl)
	ruleCntrl := controller.RuleCntrl{
		RuleService: ruleSvcMock,
	}
	t.Run("WHEN id is not an integer", func(t *testing.T) {
		_, err := echotest.DoGET(ruleCntrl.FindOne, "/", map[string]string{"id": "invalid"})
		require.EqualError(t, err, "code=400, message=Invalid ID")
	})
	t.Run("WHEN received error", func(t *testing.T) {
		ruleSvcMock.EXPECT().FindOne(gomock.Any(), int64(999)).Return(nil, errors.New("find error"))
		_, err := echotest.DoGET(ruleCntrl.FindOne, "/", map[string]string{"id": "999"})
		require.EqualError(t, err, "code=500, message=find error")
	})
	t.Run("WHEN rule is not found", func(t *testing.T) {
		ruleSvcMock.EXPECT().FindOne(gomock.Any(), int64(999)).Return(nil, nil)
		_, err := echotest.DoGET(ruleCntrl.FindOne, "/", map[string]string{"id": "999"})
		require.EqualError(t, err, "code=404, message=Rule #999 not found")
	})
	t.Run("WHEN successful", func(t *testing.T) {
		ruleSvcMock.EXPECT().FindOne(gomock.Any(), int64(999)).Return(
			&repository.Rule{
				ID:         999,
				Name:       "Airport Rule",
				UrlPattern: "/airport",
			},
			nil,
		)
		rr, err := echotest.DoGET(ruleCntrl.FindOne, "/", map[string]string{"id": "999"})
		require.NoError(t, err)
		require.Equal(t, http.StatusOK, rr.Code)
		require.Equal(t, "{\"id\":999,\"name\":\"Airport Rule\",\"url_pattern\":\"/airport\",\"data_source_id\":null,\"updated_at\":\"0001-01-01T00:00:00Z\",\"created_at\":\"0001-01-01T00:00:00Z\"}\n", rr.Body.String())
	})
}

func TestRuleController_Delete(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	ruleSvcMock := mock_service.NewMockRuleService(ctrl)
	ruleCntrl := controller.RuleCntrl{
		RuleService: ruleSvcMock,
	}
	t.Run("WHEN id is not an integer", func(t *testing.T) {
		_, err := echotest.DoDELETE(ruleCntrl.Delete, "/", map[string]string{"id": "invalid"})
		require.EqualError(t, err, "code=400, message=Invalid ID")
	})
	t.Run("WHEN received error", func(t *testing.T) {
		ruleSvcMock.EXPECT().Delete(gomock.Any(), int64(999)).Return(errors.New("find error"))
		_, err := echotest.DoDELETE(ruleCntrl.Delete, "/", map[string]string{"id": "999"})
		require.EqualError(t, err, "code=500, message=find error")
	})
	t.Run("WHEN successful", func(t *testing.T) {
		ruleSvcMock.EXPECT().Delete(gomock.Any(), int64(999)).Return(nil)
		rr, err := echotest.DoDELETE(ruleCntrl.Delete, "/", map[string]string{"id": "999"})
		require.NoError(t, err)
		require.Equal(t, http.StatusOK, rr.Code)
		require.Equal(t, "{\"message\":\"Success delete rule #999\"}\n", rr.Body.String())
	})
}

func TestRuleController_Update(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	ruleSvcMock := mock_service.NewMockRuleService(ctrl)
	ruleCntrl := controller.RuleCntrl{
		RuleService: ruleSvcMock,
	}
	t.Run("WHEN body is malformed", func(t *testing.T) {
		_, err := echotest.DoPUT(ruleCntrl.Update, "/", "invalid", nil)
		require.EqualError(t, err, "code=400, message=Syntax error: offset=1, error=invalid character 'i' looking for beginning of value")
	})
	t.Run("WHEN id is invalid", func(t *testing.T) {
		_, err := echotest.DoPUT(ruleCntrl.Update, "/", `{ "id": -1, "name": "Airport Rule", "url_pattern": "/airport"}`, nil)
		require.EqualError(t, err, "code=400, message=Invalid ID")
	})
	t.Run("WHEN rule is invalid", func(t *testing.T) {
		_, err := echotest.DoPUT(ruleCntrl.Update, "/", `{ "id": 999, "name": "Airport Rule" }`, nil)
		require.EqualError(t, err, "code=400, message=Key: 'Rule.UrlPattern' Error:Field validation for 'UrlPattern' failed on the 'required' tag")
	})
	t.Run("WHEN received error after updating", func(t *testing.T) {
		ruleSvcMock.EXPECT().Update(gomock.Any(), gomock.Any()).Return(errors.New("update error"))
		_, err := echotest.DoPUT(ruleCntrl.Update, "/", `{ "id": 999, "name": "Airport Rule", "url_pattern": "/airport" }`, nil)
		require.EqualError(t, err, "code=500, message=update error")
	})
	t.Run("WHEN successfully update rule", func(t *testing.T) {
		ruleSvcMock.EXPECT().Update(gomock.Any(), gomock.Any()).Return(nil)
		rr, err := echotest.DoPUT(ruleCntrl.Update, "/", `{ "id": 999, "name": "Airport Rule", "url_pattern": "/airport" }`, nil)
		require.NoError(t, err)
		require.Equal(t, http.StatusOK, rr.Code)
		require.Equal(t, "{\"message\":\"Success update rule #999\"}\n", rr.Body.String())
	})
}
