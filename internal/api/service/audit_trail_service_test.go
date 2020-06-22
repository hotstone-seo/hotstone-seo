package service_test

import (
	"context"
	"testing"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/golang/mock/gomock"

	"github.com/hotstone-seo/hotstone-seo/internal/api/repository"
	"github.com/hotstone-seo/hotstone-seo/internal/api/repository_mock"
	"github.com/hotstone-seo/hotstone-seo/internal/api/service"
)

type auditTrailSvcFn func(repo *repository_mock.MockAuditTrailRepo)

func createAuditTrailSvc(t *testing.T, fn auditTrailSvcFn) (service.AuditTrailSvc, *gomock.Controller) {
	mock := gomock.NewController(t)
	repo := repository_mock.NewMockAuditTrailRepo(mock)
	if fn != nil {
		fn(repo)
	}
	return &service.AuditTrailSvcImpl{
		AuditTrailRepo: repo,
	}, mock
}

func TestAuditTrailService_RecordInsert(t *testing.T) {
	testcases := []struct {
		testName        string
		username        string
		entity          string
		id              int64
		obj             interface{}
		onAuditTrailSvc func(mock *repository_mock.MockAuditTrailRepo)
	}{
		{
			entity:   "rules",
			id:       9999,
			obj:      struct{ ID int }{ID: 9999},
			username: "bejo@email.com",
			onAuditTrailSvc: func(mock *repository_mock.MockAuditTrailRepo) {
				mock.EXPECT().Insert(gomock.Any(), repository.AuditTrail{
					EntityName: "rules",
					EntityID:   9999,
					Operation:  "INSERT",
					Username:   "bejo@email.com",
					OldData:    []byte(`{}`),
					NewData:    []byte(`{"ID":9999}`),
				})
			},
		},
	}
	for _, tt := range testcases {
		t.Run(tt.testName, func(t *testing.T) {
			svc, mock := createAuditTrailSvc(t, tt.onAuditTrailSvc)
			defer mock.Finish()

			ctx := context.WithValue(
				context.Background(),
				service.TokenCtxKey,
				&jwt.Token{
					Claims: jwt.MapClaims{
						"email": tt.username,
					},
				},
			)
			svc.RecordInsert(ctx, tt.entity, tt.id, tt.obj)
			time.Sleep(5 * time.Millisecond)
		})
	}
}

func TestAuditTrailService_RecordDelete(t *testing.T) {
	testcases := []struct {
		testName        string
		username        string
		entity          string
		id              int64
		obj             interface{}
		onAuditTrailSvc func(mock *repository_mock.MockAuditTrailRepo)
	}{
		{
			entity:   "rules",
			id:       9999,
			obj:      struct{ ID int }{ID: 9999},
			username: "bejo@email.com",
			onAuditTrailSvc: func(mock *repository_mock.MockAuditTrailRepo) {
				mock.EXPECT().Insert(gomock.Any(), repository.AuditTrail{
					EntityName: "rules",
					EntityID:   9999,
					Operation:  "DELETE",
					Username:   "bejo@email.com",
					OldData:    []byte(`{"ID":9999}`),
					NewData:    []byte(`{}`),
				})
			},
		},
	}
	for _, tt := range testcases {
		t.Run(tt.testName, func(t *testing.T) {
			svc, mock := createAuditTrailSvc(t, tt.onAuditTrailSvc)
			defer mock.Finish()

			ctx := context.WithValue(
				context.Background(),
				service.TokenCtxKey,
				&jwt.Token{
					Claims: jwt.MapClaims{
						"email": tt.username,
					},
				},
			)
			svc.RecordDelete(ctx, tt.entity, tt.id, tt.obj)
			time.Sleep(5 * time.Millisecond)
		})
	}
}

func TestAuditTrailService_RecordUpdate(t *testing.T) {
	testcases := []struct {
		testName        string
		username        string
		entity          string
		id              int64
		oldObj          interface{}
		newObj          interface{}
		onAuditTrailSvc func(mock *repository_mock.MockAuditTrailRepo)
	}{
		{
			entity:   "rules",
			id:       9999,
			oldObj:   struct{ ID int }{ID: 9999},
			newObj:   struct{ ID int }{ID: 8888},
			username: "bejo@email.com",
			onAuditTrailSvc: func(mock *repository_mock.MockAuditTrailRepo) {
				mock.EXPECT().Insert(gomock.Any(), repository.AuditTrail{
					EntityName: "rules",
					EntityID:   9999,
					Operation:  "UPDATE",
					Username:   "bejo@email.com",
					OldData:    []byte(`{"ID":9999}`),
					NewData:    []byte(`{"ID":8888}`),
				})
			},
		},
	}
	for _, tt := range testcases {
		t.Run(tt.testName, func(t *testing.T) {
			svc, mock := createAuditTrailSvc(t, tt.onAuditTrailSvc)
			defer mock.Finish()

			ctx := context.WithValue(
				context.Background(),
				service.TokenCtxKey,
				&jwt.Token{
					Claims: jwt.MapClaims{
						"email": tt.username,
					},
				},
			)
			svc.RecordUpdate(ctx, tt.entity, tt.id, tt.oldObj, tt.newObj)
			time.Sleep(5 * time.Millisecond)
		})
	}
}
