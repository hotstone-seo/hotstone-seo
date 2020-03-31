package oauth2google_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/require"

	"github.com/hotstone-seo/hotstone-seo/pkg/oauth2google"
	"github.com/labstack/echo"
)

func TestMiddleware(t *testing.T) {
	t.Run("GIVEN bad config", func(t *testing.T) {
		cntrl := &oauth2google.AuthCntrl{
			Config: &oauth2google.Config{},
		}

		var debugger strings.Builder
		logrus.SetOutput(&debugger)
		logrus.SetFormatter(&testformatter{})

		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		cntrl.Middleware()(func(ce echo.Context) error {
			return nil
		})(c)

		require.Equal(t, "[warning] JWT Error: code=400, message=missing or malformed jwt", debugger.String())
	})
}

// simple formatter without timestamp
type testformatter struct {
}

func (f *testformatter) Format(entry *logrus.Entry) ([]byte, error) {
	return []byte(fmt.Sprintf("[%s] %s", entry.Level, entry.Message)), nil
}
