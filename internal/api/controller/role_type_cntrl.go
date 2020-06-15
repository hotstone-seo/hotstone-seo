package controller

import (
	"database/sql"
	"fmt"
	"net/http"
	"strconv"

	"github.com/hotstone-seo/hotstone-seo/internal/api/repository"
	"github.com/hotstone-seo/hotstone-seo/internal/api/service"
	"github.com/labstack/echo"
	"go.uber.org/dig"
)

// RoleTypeCntrl is controller to role_type entity
type RoleTypeCntrl struct {
	dig.In
	Svc service.UserRoleSvc
}

// Route to define API Route
func (r *RoleTypeCntrl) Route(e *echo.Group) {
	e.GET("/role_types", r.Find)
	e.GET("/role_types/:id", r.FindOne)
	e.POST("/role_types", r.Create)
	e.PUT("/role_types", r.Update)
	e.DELETE("/role_types/:id", r.Delete)
}

// Find all role_type
func (r *RoleTypeCntrl) Find(ctx echo.Context) (err error) {
	var roleTypes []*repository.UserRole
	ctx0 := ctx.Request().Context()

	validCols := []string{"id", "name", "modules", "updated_at", "created_at"}
	paginationParam := repository.BuildPaginationParam(ctx.QueryParams(), validCols)
	if roleTypes, err = r.Svc.Find(ctx0, paginationParam); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return ctx.JSON(http.StatusOK, roleTypes)
}

// FindOne role_type
func (r *RoleTypeCntrl) FindOne(c echo.Context) (err error) {

	ctx := c.Request().Context()
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, "Invalid ID")
	}
	roleType, err := r.Svc.FindOne(ctx, id)
	if err == sql.ErrNoRows {
		return echo.NewHTTPError(http.StatusNotFound)
	}
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, roleType)
}

// Create role_type
func (r *RoleTypeCntrl) Create(c echo.Context) (err error) {
	var (
		req          service.UserRoleRequest
		roleType     repository.UserRole
		lastInsertID int64
	)
	ctx0 := c.Request().Context()
	if err = c.Bind(&req); err != nil {
		return
	}
	if lastInsertID, err = r.Svc.Insert(ctx0, req); err != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, err.Error())
	}
	roleType.ID = lastInsertID
	roleType.Name = req.Name
	return c.JSON(http.StatusCreated, roleType)
}

// Update role_type
func (r *RoleTypeCntrl) Update(c echo.Context) (err error) {
	var (
		req service.UserRoleRequest
	)
	ctx0 := c.Request().Context()
	if err = c.Bind(&req); err != nil {
		return err
	}
	if req.ID <= 0 {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid ID")
	}
	if err = r.Svc.Update(ctx0, req); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, GeneralResponse{
		Message: fmt.Sprintf("Success update role type #%d", req.ID),
	})

}

// Delete role_type
func (r *RoleTypeCntrl) Delete(c echo.Context) (err error) {
	var id int64
	ctx0 := c.Request().Context()
	if id, err = strconv.ParseInt(c.Param("id"), 10, 64); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid ID")
	}
	if err = r.Svc.Delete(ctx0, id); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, GeneralResponse{
		Message: fmt.Sprintf("Success delete role type #%d", id),
	})
}
