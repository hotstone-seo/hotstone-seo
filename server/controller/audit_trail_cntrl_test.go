package controller_test

import (
	"errors"
	"net/http"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/hotstone-seo/hotstone-seo/server/controller"
	"github.com/hotstone-seo/hotstone-seo/server/repository"
	"github.com/hotstone-seo/hotstone-seo/server/service_mock"

	"github.com/stretchr/testify/require"
	"github.com/typical-go/typical-rest-server/pkg/echotest"
)

func TestAuditTrailController_Find(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	auditTrailSvcMock := service_mock.NewMockAuditTrailService(ctrl)
	auditTrailCntrl := controller.AuditTrailCntrl{
		AuditTrailService: auditTrailSvcMock,
	}

	t.Run("WHEN retrieved error", func(t *testing.T) {
		auditTrailSvcMock.EXPECT().Find(gomock.Any(), gomock.Any()).Return(nil, errors.New("retrieve error"))
		_, err := echotest.DoGET(auditTrailCntrl.Find, "/", nil)
		require.EqualError(t, err, "code=500, message=retrieve error")
	})

	t.Run("WHEN successful", func(t *testing.T) {
		auditTrailTime := time.Now()
		auditTrailTimeStr := auditTrailTime.Format(time.RFC3339Nano)
		auditTrailSvcMock.EXPECT().Find(gomock.Any(), gomock.Any()).Return(
			[]*repository.AuditTrail{
				&repository.AuditTrail{
					ID:         100,
					Time:       auditTrailTime,
					EntityName: "dssource1.com",
					EntityID:   10,
					Operation:  "insert",
					OldData:    []byte("{\"content\": \"some-string\"}"),
					NewData:    []byte("{\"content\": \"some-string\"}"),
				},
			},
			nil,
		)
		rr, err := echotest.DoGET(auditTrailCntrl.Find, "/", nil)
		require.NoError(t, err)
		require.Equal(t, http.StatusOK, rr.Code)
		require.Equal(t, "[{\"id\":100,\"time\":\""+auditTrailTimeStr+"\",\"entity_name\":\"dssource1.com\",\"entity_id\":10,\"operation\":\"insert\",\"old_data\":{\"content\":\"some-string\"},\"new_data\":{\"content\":\"some-string\"}}]\n", rr.Body.String())
	})
}
