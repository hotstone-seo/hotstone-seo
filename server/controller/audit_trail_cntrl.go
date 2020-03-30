package controller

import (
	"net/http"

	"github.com/hotstone-seo/hotstone-seo/server/repository"
	"github.com/hotstone-seo/hotstone-seo/server/service"
	"github.com/labstack/echo"
	"go.uber.org/dig"
)

// AuditTrailCntrl is controller to rule entity
type AuditTrailCntrl struct {
	dig.In
	service.AuditTrailService
}

// Route to define API Route
func (c *AuditTrailCntrl) Route(e *echo.Group) {
	e.GET("/audit-trail", c.Find)
}

// Find all rule
func (c *AuditTrailCntrl) Find(ctx echo.Context) (err error) {
	var listAuditTrail []*repository.AuditTrail
	ctx0 := ctx.Request().Context()

	validCols := []string{"id", "time", "entity_name", "entity_id", "operation", "username", "old_data", "new_data"}
	paginationParam := repository.BuildPaginationParam(ctx.QueryParams(), validCols)
	if listAuditTrail, err = c.AuditTrailService.Find(ctx0, paginationParam); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return ctx.JSON(http.StatusOK, listAuditTrail)
}
