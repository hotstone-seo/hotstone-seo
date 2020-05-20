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

// UserCntrl is controller to user entity
type UserCntrl struct {
	dig.In
	service.UserService
}

// Route to define API Route
func (c *UserCntrl) Route(e *echo.Group) {
	e.GET("/users", c.Find)
	e.POST("/users", c.Create)
	e.GET("/users/:id", c.FindOne)
	e.PUT("/users", c.Update)
	e.DELETE("/users/:id", c.Delete)
	e.POST("/users_is_exists", c.FindOneByEmail)
}

// Find all users
func (c *UserCntrl) Find(ctx echo.Context) (err error) {
	var users []*repository.User
	ctx0 := ctx.Request().Context()

	validCols := []string{"id", "email", "role_type_id", "updated_at", "created_at"}
	paginationParam := repository.BuildPaginationParam(ctx.QueryParams(), validCols)
	if users, err = c.UserService.Find(ctx0, paginationParam); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return ctx.JSON(http.StatusOK, users)
}

// FindOne user
func (c *UserCntrl) FindOne(ec echo.Context) (err error) {
	var (
		id   int64
		user *repository.User
	)

	ctx := ec.Request().Context()
	if id, err = strconv.ParseInt(ec.Param("id"), 10, 64); err != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, "Invalid ID")
	}
	user, err = c.UserService.FindOne(ctx, id)
	if err == sql.ErrNoRows {
		return echo.NewHTTPError(http.StatusNotFound)
	}
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return ec.JSON(http.StatusOK, user)
}

// Create user
func (c *UserCntrl) Create(ctx echo.Context) (err error) {
	var user repository.User
	var lastInsertID int64
	ctx0 := ctx.Request().Context()
	if err = ctx.Bind(&user); err != nil {
		return err
	}
	if err = user.Validate(); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if lastInsertID, err = c.UserService.Insert(ctx0, user); err != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, err.Error())
	}
	user.ID = lastInsertID
	return ctx.JSON(http.StatusCreated, user)
}

// Delete user
func (c *UserCntrl) Delete(ctx echo.Context) (err error) {
	var id int64
	ctx0 := ctx.Request().Context()
	if id, err = strconv.ParseInt(ctx.Param("id"), 10, 64); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid ID")
	}
	if err = c.UserService.Delete(ctx0, id); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return ctx.JSON(http.StatusOK, GeneralResponse{
		Message: fmt.Sprintf("Success delete user #%d", id),
	})
}

// Update user
func (c *UserCntrl) Update(ctx echo.Context) (err error) {
	var user repository.User
	ctx0 := ctx.Request().Context()
	if err = ctx.Bind(&user); err != nil {
		return err
	}
	if user.ID <= 0 {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid ID")
	}
	if err = user.Validate(); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err = c.UserService.Update(ctx0, user); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return ctx.JSON(http.StatusOK, GeneralResponse{
		Message: fmt.Sprintf("Success update user #%d", user.ID),
	})
}

// FindOneByEmail user
func (c *UserCntrl) FindOneByEmail(ctx echo.Context) (err error) {
	var user *repository.User
	ctx0 := ctx.Request().Context()
	if err = ctx.Bind(&user); err != nil {
		return err
	}
	if err = user.Validate(); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	user, err = c.UserService.FindOneByEmail(ctx0, user.Email)
	if err == sql.ErrNoRows {
		user = nil
	}

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return ctx.JSON(http.StatusOK, user)
}
