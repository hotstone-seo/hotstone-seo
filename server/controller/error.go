package controller

import (
	"database/sql"
	"net/http"

	"github.com/hotstone-seo/hotstone-seo/pkg/cachekit"
	"github.com/labstack/echo"
	"github.com/typical-go/typical-rest-server/pkg/errvalid"
)

func httpError(err error) *echo.HTTPError {
	if err == sql.ErrNoRows {
		return echo.NewHTTPError(http.StatusNotFound)
	}

	if errvalid.Check(err) {
		return echo.NewHTTPError(
			http.StatusUnprocessableEntity,
			errvalid.Message(err),
		)
	}

	if cachekit.NotModifiedError(err) {
		return echo.NewHTTPError(http.StatusNotModified)
	}

	return echo.NewHTTPError(
		http.StatusInternalServerError,
		err.Error(),
	)
}
