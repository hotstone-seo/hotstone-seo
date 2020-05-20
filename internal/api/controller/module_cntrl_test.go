package controller_test

import (
	"database/sql"
	"errors"
	"fmt"
	"net/http"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/hotstone-seo/hotstone-seo/internal/api/controller"
	"github.com/hotstone-seo/hotstone-seo/internal/api/repository"
	"github.com/hotstone-seo/hotstone-seo/internal/api/service_mock"
	"github.com/typical-go/typical-rest-server/pkg/echotest"
	"github.com/typical-go/typical-rest-server/pkg/errvalid"
)

type (
	testCase struct {
		testName string
		echotest.TestCase
		moduleCntrlBuilder
	}

	moduleCntrlBuilder struct {
		moduleSvcFn func(*service_mock.MockModuleService)
	}
)

func (c *moduleCntrlBuilder) build(mock *gomock.Controller) *controller.ModuleCntrl {
	mockSvc := service_mock.NewMockModuleService(mock)
	if c.moduleSvcFn != nil {
		c.moduleSvcFn(mockSvc)
	}
	return &controller.ModuleCntrl{
		ModuleService: mockSvc,
	}
}

func TestModuleController_Find(t *testing.T) {
	testcases := []testCase{
		{
			TestCase: echotest.TestCase{
				Request: echotest.Request{
					Method: http.MethodGet,
					Target: "/",
				},
				ExpectedCode: http.StatusOK,
				ExpectedBody: "[{\"id\":1,\"name\":\"rule\",\"path\":\"rules\",\"pattern\":\"rules*\",\"label\":\"Rules\",\"updated_at\":\"0001-01-01T00:00:00Z\",\"created_at\":\"0001-01-01T00:00:00Z\"}]\n",
			},
			moduleCntrlBuilder: moduleCntrlBuilder{
				moduleSvcFn: func(svc *service_mock.MockModuleService) {
					svc.EXPECT().Find(gomock.Any(), gomock.Any()).Return([]*repository.Module{
						&repository.Module{ID: 1, Name: "rule", Path: "rules", Pattern: "rules*", Label: "Rules"},
					}, nil)
				},
			},
		},
		{
			TestCase: echotest.TestCase{
				Request: echotest.Request{
					Method: http.MethodGet,
					Target: "/",
				},
				ExpectedErr: "code=500, message=some-error",
			},
			moduleCntrlBuilder: moduleCntrlBuilder{
				moduleSvcFn: func(svc *service_mock.MockModuleService) {
					svc.EXPECT().Find(gomock.Any(), gomock.Any()).Return(nil, fmt.Errorf("some-error"))
				},
			},
		},
	}

	for _, tt := range testcases {
		t.Run(tt.testName, func(t *testing.T) {
			mock := gomock.NewController(t)
			defer mock.Finish()
			tt.Execute(t, tt.build(mock).Find)
		})
	}
}

func TestModuleController_FindOne(t *testing.T) {
	testcases := []testCase{
		{
			testName: "valid ID",
			TestCase: echotest.TestCase{
				Request: echotest.Request{
					Method:    http.MethodGet,
					Target:    "/",
					URLParams: map[string]string{"id": "1"},
				},
				ExpectedCode: http.StatusOK,
				ExpectedBody: "{\"id\":1,\"name\":\"rule\",\"path\":\"rules\",\"pattern\":\"rules*\",\"label\":\"Rules\",\"updated_at\":\"0001-01-01T00:00:00Z\",\"created_at\":\"0001-01-01T00:00:00Z\"}\n",
			},
			moduleCntrlBuilder: moduleCntrlBuilder{
				moduleSvcFn: func(svc *service_mock.MockModuleService) {
					svc.EXPECT().FindOne(gomock.Any(), int64(1)).Return(&repository.Module{ID: 1, Name: "rule", Path: "rules", Pattern: "rules*", Label: "Rules"}, nil)
				},
			},
		},
		{
			testName: "entity not found",
			TestCase: echotest.TestCase{
				Request: echotest.Request{
					Method:    http.MethodGet,
					Target:    "/",
					URLParams: map[string]string{"id": "5"},
				},
				ExpectedErr: "code=404, message=Not Found",
			},
			moduleCntrlBuilder: moduleCntrlBuilder{
				moduleSvcFn: func(svc *service_mock.MockModuleService) {
					svc.EXPECT().FindOne(gomock.Any(), int64(5)).Return(nil, sql.ErrNoRows)
				},
			},
		},
		{
			testName: "validation error",
			TestCase: echotest.TestCase{
				Request: echotest.Request{
					Method:    http.MethodGet,
					Target:    "/",
					URLParams: map[string]string{"id": "2"},
				},
				ExpectedErr: "code=500, message=Validation: some-validation",
			},
			moduleCntrlBuilder: moduleCntrlBuilder{
				moduleSvcFn: func(svc *service_mock.MockModuleService) {
					svc.EXPECT().FindOne(gomock.Any(), int64(2)).Return(nil, errvalid.New("some-validation"))
				},
			},
		},
		{
			testName: "error",
			TestCase: echotest.TestCase{
				Request: echotest.Request{
					Method:    http.MethodGet,
					Target:    "/",
					URLParams: map[string]string{"id": "2"},
				},
				ExpectedErr: "code=500, message=some-error",
			},
			moduleCntrlBuilder: moduleCntrlBuilder{
				moduleSvcFn: func(svc *service_mock.MockModuleService) {
					svc.EXPECT().FindOne(gomock.Any(), int64(2)).Return(nil, errors.New("some-error"))
				},
			},
		},
	}

	for _, tt := range testcases {
		t.Run(tt.testName, func(t *testing.T) {
			mock := gomock.NewController(t)
			defer mock.Finish()
			tt.Execute(t, tt.build(mock).FindOne)
		})
	}
}

func TestModuleController_Delete(t *testing.T) {
	testcases := []testCase{
		{
			TestCase: echotest.TestCase{
				Request: echotest.Request{
					Method:    http.MethodDelete,
					Target:    "/",
					URLParams: map[string]string{"id": "1"},
				},
				ExpectedCode: 200,
				ExpectedBody: "{\"message\":\"Success delete module #1\"}\n",
			},
			moduleCntrlBuilder: moduleCntrlBuilder{
				moduleSvcFn: func(svc *service_mock.MockModuleService) {
					svc.EXPECT().Delete(gomock.Any(), int64(1)).Return(nil)
				},
			},
		},
		{
			TestCase: echotest.TestCase{
				Request: echotest.Request{
					Method:    http.MethodDelete,
					Target:    "/",
					URLParams: map[string]string{"id": "1"},
				},
				ExpectedErr: "code=500, message=some-error",
			},
			moduleCntrlBuilder: moduleCntrlBuilder{
				moduleSvcFn: func(svc *service_mock.MockModuleService) {
					svc.EXPECT().Delete(gomock.Any(), int64(1)).Return(errors.New("some-error"))
				},
			},
		},
		{
			TestCase: echotest.TestCase{
				Request: echotest.Request{
					Method:    http.MethodDelete,
					Target:    "/",
					URLParams: map[string]string{"id": "1"},
				},
				ExpectedErr: "code=500, message=Validation: some-validation",
			},
			moduleCntrlBuilder: moduleCntrlBuilder{
				moduleSvcFn: func(svc *service_mock.MockModuleService) {
					svc.EXPECT().Delete(gomock.Any(), int64(1)).Return(errvalid.New("some-validation"))
				},
			},
		},
	}

	for _, tt := range testcases {
		t.Run(tt.testName, func(t *testing.T) {
			mock := gomock.NewController(t)
			defer mock.Finish()
			tt.Execute(t, tt.build(mock).Delete)
		})
	}
}

func TestModuleController_Create(t *testing.T) {
	testcases := []testCase{
		{
			TestCase: echotest.TestCase{
				Request: echotest.Request{
					Method: http.MethodPost,
					Target: "/",
					Body:   `invalid}`,
					Header: echotest.HeaderForJSON(),
				},
				ExpectedErr: `code=400, message=Syntax error: offset=1, error=invalid character 'i' looking for beginning of value`,
			},
		},
		{
			TestCase: echotest.TestCase{
				Request: echotest.Request{
					Method: http.MethodPost,
					Target: "/",
					Body:   `{"name":"some-name", "path":"some-path"}`,
					Header: echotest.HeaderForJSON(),
				},
				ExpectedErr: "code=422, message=some-error",
			},
			moduleCntrlBuilder: moduleCntrlBuilder{
				moduleSvcFn: func(svc *service_mock.MockModuleService) {
					svc.EXPECT().Insert(gomock.Any(), gomock.Any()).Return(int64(-1), fmt.Errorf("some-error"))
				},
			},
		},
		{
			TestCase: echotest.TestCase{
				Request: echotest.Request{
					Method: http.MethodPost,
					Target: "/",
					Body:   `{"name":"some-name", "path":"some-path"}`,
					Header: echotest.HeaderForJSON(),
				},
				ExpectedCode: http.StatusCreated,
				ExpectedBody: "{\"id\":100,\"name\":\"some-name\",\"path\":\"some-path\",\"pattern\":\"\",\"label\":\"\",\"updated_at\":\"0001-01-01T00:00:00Z\",\"created_at\":\"0001-01-01T00:00:00Z\"}\n",
			},
			moduleCntrlBuilder: moduleCntrlBuilder{
				moduleSvcFn: func(svc *service_mock.MockModuleService) {
					svc.EXPECT().Insert(gomock.Any(), gomock.Any()).Return(int64(100), nil)
				},
			},
		},
	}

	for _, tt := range testcases {
		t.Run(tt.testName, func(t *testing.T) {
			mock := gomock.NewController(t)
			defer mock.Finish()
			tt.Execute(t, tt.build(mock).Create)
		})
	}
}

func TestModuleController_Update(t *testing.T) {
	testcases := []testCase{
		{
			TestCase: echotest.TestCase{
				Request: echotest.Request{
					Method: http.MethodPost,
					Target: "/",
					Body:   `invalid}`,
					Header: echotest.HeaderForJSON(),
				},
				ExpectedErr: `code=400, message=Syntax error: offset=1, error=invalid character 'i' looking for beginning of value`,
			},
		},
		{
			TestCase: echotest.TestCase{
				Request: echotest.Request{
					Method: http.MethodPost,
					Target: "/",
					Body:   `{"id" : 1, "name":"some-name", "path":"some-path"}`,
					Header: echotest.HeaderForJSON(),
				},
				ExpectedCode: http.StatusOK,
				ExpectedBody: "{\"message\":\"Success update module #1\"}\n",
			},
			moduleCntrlBuilder: moduleCntrlBuilder{
				moduleSvcFn: func(svc *service_mock.MockModuleService) {
					svc.EXPECT().Update(gomock.Any(), gomock.Any()).Return(nil)
				},
			},
		},
	}

	for _, tt := range testcases {
		t.Run(tt.testName, func(t *testing.T) {
			mock := gomock.NewController(t)
			defer mock.Finish()
			tt.Execute(t, tt.build(mock).Update)
		})
	}
}
