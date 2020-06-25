package controller_test

import (
	"errors"
	"net/http"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/typical-go/typical-rest-server/pkg/echotest"
	"github.com/typical-go/typical-rest-server/pkg/errvalid"

	"github.com/hotstone-seo/hotstone-seo/internal/api/controller"
	"github.com/hotstone-seo/hotstone-seo/internal/api/service_mock"
)

type (
	tagCntrlFn func(*service_mock.MockTagService)

	testCase struct {
		testName   string
		tagCntrlFn tagCntrlFn
		echotest.TestCase
	}
)

func createTagCntrl(t *testing.T, fn tagCntrlFn) (*controller.TagCntrl, *gomock.Controller) {
	mock := gomock.NewController(t)
	mockSvc := service_mock.NewMockTagService(mock)
	if fn != nil {
		fn(mockSvc)
	}
	return &controller.TagCntrl{
		TagService: mockSvc,
	}, mock
}

func TestTagController_Create(t *testing.T) {
	testCases := []testCase{
		{
			testName: "Malformed request body",
			TestCase: echotest.TestCase{
				Request: echotest.Request{
					Method: http.MethodPost,
					Target: "/tags",
					Body:   "{ invalid }",
					Header: echotest.HeaderForJSON(),
				},
				ExpectedCode: http.StatusInternalServerError,
				ExpectedErr:  "",
			},
		},
		{
			testName: "Invalid tag",
			TestCase: echotest.TestCase{
				Request: echotest.Request{
					Method: http.MethodPost,
					Target: "/tags",
					Body:   `{ rule_id: 999, locale: "en_US" }`,
					Header: echotest.HeaderForJSON(),
				},
				ExpectedCode: http.StatusUnprocessableEntity,
				ExpectedErr:  "",
			},
			tagCntrlFn: func(svc *service_mock.MockTagService) {
				svc.EXPECT().Create(gomock.Any(), gomock.Any()).Return(int64(-1), errvalid.New("Validation error"))
			},
		},
		{
			testName: "Error when inserting",
			TestCase: echotest.TestCase{
				Request: echotest.Request{
					Method: http.MethodPost,
					Target: "/tags",
					Body:   `{ "rule_id": 999, "locale": "en_US", "type": "title", "value": "Page Title" }`,
					Header: echotest.HeaderForJSON(),
				},
				ExpectedCode: http.StatusInternalServerError,
				ExpectedErr:  "",
			},
			tagCntrlFn: func(svc *service_mock.MockTagService) {
				svc.EXPECT().Create(gomock.Any(), gomock.Any()).Return(int64(-1), errors.New("Unexpected error"))
			},
		},
		{
			testName: "Successfully insert",
			TestCase: echotest.TestCase{
				Request: echotest.Request{
					Method: http.MethodPost,
					Target: "/tags",
					Body:   `{ "rule_id": 999, "locale": "en_US", "type": "title", "value": "Page Title" }`,
					Header: echotest.HeaderForJSON(),
				},
				ExpectedCode: http.StatusCreated,
				ExpectedHeader: map[string]string{
					"Location": "/tags/999",
				},
			},
			tagCntrlFn: func(svc *service_mock.MockTagService) {
				svc.EXPECT().Create(gomock.Any(), gomock.Any()).Return(int64(999), nil)
			},
		},
	}
}

/*
func TestTagController_Find(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	tagSvcMock := service_mock.NewMockTagService(ctrl)
	tagCntrl := controller.TagCntrl{
		TagService: tagSvcMock,
	}
	t.Run("WHEN rule_id is not an integer", func(t *testing.T) {
		_, err := echotest.DoGET(tagCntrl.Find, "/tags?rule_id=invalid", nil)
		require.EqualError(t, err, "code=400, message=Invalid Rule ID")
	})
	t.Run("WHEN received error", func(t *testing.T) {
		tagSvcMock.EXPECT().Find(gomock.Any(), gomock.Any()).Return(nil, errors.New("retrieve error"))
		_, err := echotest.DoGET(tagCntrl.Find, "/tags?rule_id=999&locale=en_US", nil)
		require.EqualError(t, err, "code=500, message=retrieve error")
	})
	t.Run("WHEN successful", func(t *testing.T) {
		tagSvcMock.EXPECT().Find(gomock.Any(), gomock.Any()).Return(
			[]*repository.Tag{
				&repository.Tag{
					ID:         999,
					RuleID:     999,
					Locale:     "en_US",
					Type:       "title",
					Value:      "Page Title",
					Attributes: map[string]string{},
				},
			},
			nil,
		)
		rr, err := echotest.DoGET(tagCntrl.Find, "/tags?rule_id=999&locale=en_US", nil)
		require.NoError(t, err)
		require.Equal(t, http.StatusOK, rr.Code)
		require.Equal(t, "[{\"id\":999,\"rule_id\":999,\"locale\":\"en_US\",\"type\":\"title\",\"attributes\":{},\"value\":\"Page Title\",\"updated_at\":\"0001-01-01T00:00:00Z\",\"created_at\":\"0001-01-01T00:00:00Z\"}]\n", rr.Body.String())
	})
}

func TestTagController_FindOne(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	tagSvcMock := service_mock.NewMockTagService(ctrl)
	tagCntrl := controller.TagCntrl{
		TagService: tagSvcMock,
	}
	t.Run("WHEN ID param is not an integer", func(t *testing.T) {
		_, err := echotest.DoGET(tagCntrl.FindOne, "/", map[string]string{"id": "invalid"})
		require.EqualError(t, err, "code=422, message=Invalid ID")
	})
	t.Run("WHEN received error", func(t *testing.T) {
		tagSvcMock.EXPECT().FindOne(gomock.Any(), int64(999)).Return(nil, errors.New("find error"))
		_, err := echotest.DoGET(tagCntrl.FindOne, "/", map[string]string{"id": "999"})
		require.EqualError(t, err, "code=500, message=find error")
	})
	t.Run("WHEN tag not found", func(t *testing.T) {
		tagSvcMock.EXPECT().FindOne(gomock.Any(), int64(999)).Return(nil, sql.ErrNoRows)
		_, err := echotest.DoGET(tagCntrl.FindOne, "/", map[string]string{"id": "999"})
		require.EqualError(t, err, "code=404, message=Not Found")
	})
	t.Run("WHEN successful", func(t *testing.T) {
		tagSvcMock.EXPECT().FindOne(gomock.Any(), int64(999)).Return(
			&repository.Tag{
				ID:         999,
				RuleID:     999,
				Locale:     "en_US",
				Type:       "title",
				Value:      "Page Title",
				Attributes: map[string]string{},
			},
			nil,
		)
		rr, err := echotest.DoGET(tagCntrl.FindOne, "/", map[string]string{"id": "999"})
		require.NoError(t, err)
		require.Equal(t, http.StatusOK, rr.Code)
		require.Equal(t, "{\"id\":999,\"rule_id\":999,\"locale\":\"en_US\",\"type\":\"title\",\"attributes\":{},\"value\":\"Page Title\",\"updated_at\":\"0001-01-01T00:00:00Z\",\"created_at\":\"0001-01-01T00:00:00Z\"}\n", rr.Body.String())
	})
}

func TestTagController_Delete(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	tagSvcMock := service_mock.NewMockTagService(ctrl)
	tagCntrl := controller.TagCntrl{
		TagService: tagSvcMock,
	}
	t.Run("WHEN id is not an integer", func(t *testing.T) {
		_, err := echotest.DoDELETE(tagCntrl.Delete, "/", map[string]string{"id": "invalid"})
		require.EqualError(t, err, "code=400, message=Invalid ID")
	})
	t.Run("WHEN received error", func(t *testing.T) {
		tagSvcMock.EXPECT().Delete(gomock.Any(), int64(999)).Return(errors.New("delete error"))
		_, err := echotest.DoDELETE(tagCntrl.Delete, "/", map[string]string{"id": "999"})
		require.EqualError(t, err, "code=500, message=delete error")
	})
	t.Run("WHEN successfully delete", func(t *testing.T) {
		tagSvcMock.EXPECT().Delete(gomock.Any(), int64(999)).Return(nil)
		rr, err := echotest.DoDELETE(tagCntrl.Delete, "/", map[string]string{"id": "999"})
		require.NoError(t, err)
		require.Equal(t, http.StatusOK, rr.Code)
		require.Equal(t, "{\"message\":\"Success delete tag #999\"}\n", rr.Body.String())
	})
}

func TestTagController_Update(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	tagSvcMock := service_mock.NewMockTagService(ctrl)
	tagCntrl := controller.TagCntrl{
		TagService: tagSvcMock,
	}
	t.Run("WHEN body is malformed", func(t *testing.T) {
		_, err := echotest.DoPUT(tagCntrl.Update, "/", `{ "rule_id", "type": "title" }`, nil)
		require.EqualError(t, err, "code=400, message=Syntax error: offset=12, error=invalid character ',' after object key")
	})
	t.Run("WHEN id is invalid", func(t *testing.T) {
		_, err := echotest.DoPUT(tagCntrl.Update, "/", `{ "id": -1, "rule_id": 999, "type": "title" }`, nil)
		require.EqualError(t, err, "code=400, message=Invalid ID")
	})
	t.Run("WHEN tag is invalid", func(t *testing.T) {
		_, err := echotest.DoPUT(tagCntrl.Update, "/", `{ "id": 999, "rule_id": 999, "locale": "en_US", "type": "title" }`, nil)
		require.EqualError(t, err, "code=400, message=Key: 'Tag.Value' Error:Field validation for 'Value' failed on the 'noempty' tag")
	})
	t.Run("WHEN received error after updating", func(t *testing.T) {
		tagSvcMock.EXPECT().Update(gomock.Any(), gomock.Any()).Return(errors.New("update error"))
		_, err := echotest.DoPUT(tagCntrl.Update, "/", `{ "id": 999, "rule_id": 999, "locale": "en_US", "type": "title", "value": "Page Title" }`, nil)
		require.EqualError(t, err, "code=500, message=update error")
	})
	t.Run("WHEN successfully update tag", func(t *testing.T) {
		tagSvcMock.EXPECT().Update(gomock.Any(), gomock.Any()).Return(nil)
		rr, err := echotest.DoPUT(tagCntrl.Update, "/", `{ "id": 999, "rule_id": 999, "locale": "en_US", "type": "title", "value": "Page Title" }`, nil)
		require.NoError(t, err)
		require.Equal(t, http.StatusOK, rr.Code)
		require.Equal(t, "{\"message\":\"Success update tag #999\"}\n", rr.Body.String())
	})
}
*/
