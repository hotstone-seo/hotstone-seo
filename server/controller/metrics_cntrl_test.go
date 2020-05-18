package controller_test

import (
	"errors"
	"net/http"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"github.com/typical-go/typical-rest-server/pkg/echotest"

	"github.com/hotstone-seo/hotstone-seo/server/controller"
	"github.com/hotstone-seo/hotstone-seo/internal/analyt"
	"github.com/hotstone-seo/hotstone-seo/server/service_mock"
)

func TestMetricsController_ListMismatched(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	metricsSvcMock := service_mock.NewMockMetricService(ctrl)
	metricsCntrl := controller.MetricsCntrl{
		MetricService: metricsSvcMock,
	}

	firstSeen := time.Now()
	lastSeen := time.Now()

	t.Run("WHEN retrieved error", func(t *testing.T) {
		metricsSvcMock.EXPECT().MismatchReports(gomock.Any(), gomock.Any()).Return(nil, errors.New("retrieve error"))
		_, err := echotest.DoGET(metricsCntrl.ListMismatched, "/", nil)
		require.EqualError(t, err, "code=500, message=retrieve error")
	})

	t.Run("WHEN successful", func(t *testing.T) {
		metricsSvcMock.EXPECT().MismatchReports(gomock.Any(), gomock.Any()).Return(
			[]*analyt.MismatchReport{
				&analyt.MismatchReport{
					URL:       "test.com/test",
					FirstSeen: firstSeen,
					LastSeen:  lastSeen,
					Count:     1,
				},
			},
			nil,
		)

		firstSeenStr := firstSeen.Format(time.RFC3339Nano)
		lastSeenStr := lastSeen.Format(time.RFC3339Nano)

		rr, err := echotest.DoGET(metricsCntrl.ListMismatched, "/", nil)
		require.NoError(t, err)
		require.Equal(t, http.StatusOK, rr.Code)
		require.Equal(t, "[{\"url\":\"test.com/test\",\"count\":1,\"first_seen\":\""+firstSeenStr+"\",\"last_seen\":\""+lastSeenStr+"\"}]\n", rr.Body.String())
	})
}
