package controller

import (
	"database/sql"
	"fmt"
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
	e.POST("/role_types", c.Create)
	e.PUT("/role_types", c.Update)
	e.DELETE("/role_types/:id", c.Delete)
}

// Create role_type
func (c *RoleTypeCntrl) Create(ctx echo.Context) (err error) {
	var roleType repository.RoleType
	var lastInsertID int64
	ctx0 := ctx.Request().Context()
	fmt.Print(ctx0)
	if err = ctx.Bind(&roleType); err != nil {
		return err
	}
	if err = roleType.Validate(); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if lastInsertID, err = c.RoleTypeService.Insert(ctx0, roleType); err != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, err.Error())
	}
	roleType.ID = lastInsertID
	return ctx.JSON(http.StatusCreated, roleType)
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

// Delete role_type
func (c *RoleTypeCntrl) Delete(ctx echo.Context) (err error) {
	var id int64
	ctx0 := ctx.Request().Context()
	if id, err = strconv.ParseInt(ctx.Param("id"), 10, 64); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid ID")
	}
	if err = c.RoleTypeService.Delete(ctx0, id); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return ctx.JSON(http.StatusOK, GeneralResponse{
		Message: fmt.Sprintf("Success delete role type #%d", id),
	})
}

// Update role_type
func (c *RoleTypeCntrl) Update(ctx echo.Context) (err error) {
	var roleType repository.RoleType
	ctx0 := ctx.Request().Context()
	if err = ctx.Bind(&roleType); err != nil {
		return err
	}
	if roleType.ID <= 0 {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid ID")
	}
	if err = roleType.Validate(); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err = c.RoleTypeService.Update(ctx0, roleType); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return ctx.JSON(http.StatusOK, GeneralResponse{
		Message: fmt.Sprintf("Success update role type #%d", roleType.ID),
	})
}
