package controller

import (
	"database/sql"
	"net/http"
	"strconv"

	"github.com/hotstone-seo/hotstone-seo/server/repository"
	"github.com/hotstone-seo/hotstone-seo/server/service"
	"github.com/labstack/echo"
	"go.uber.org/dig"
)

// RoleTypeCntrl is controller to role_type entity
type RoleTypeCntrl struct {
	dig.In
	service.RoleTypeService
}

// Route to define API Route
func (c *RoleTypeCntrl) Route(e *echo.Group) {
	e.GET("/role_types", c.Find)
	e.GET("/role_types/:id", c.FindOne)
}

// Find all role_type
func (c *RoleTypeCntrl) Find(ctx echo.Context) (err error) {
	var roleTypes []*repository.RoleType
	ctx0 := ctx.Request().Context()

	validCols := []string{"id", "name", "updated_at", "created_at"}
	paginationParam := repository.BuildPaginationParam(ctx.QueryParams(), validCols)
	if roleTypes, err = c.RoleTypeService.Find(ctx0, paginationParam); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return ctx.JSON(http.StatusOK, roleTypes)
}

// FindOne role_type
func (c *RoleTypeCntrl) FindOne(ec echo.Context) (err error) {
	var (
		id       int64
		roleType *repository.RoleType
	)

	ctx := ec.Request().Context()
	if id, err = strconv.ParseInt(ec.Param("id"), 10, 64); err != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, "Invalid ID")
	}
	roleType, err = c.RoleTypeService.FindOne(ctx, id)
	if err == sql.ErrNoRows {
		return echo.NewHTTPError(http.StatusNotFound)
	}
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return ec.JSON(http.StatusOK, roleType)
}
