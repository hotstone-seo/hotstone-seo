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
	Svc service.UserSvc
}

// Route to define API Route
func (u *UserCntrl) Route(e *echo.Group) {
	e.GET("/users", u.Find)
	e.POST("/users", u.Create)
	e.GET("/users/:id", u.FindOne)
	e.PUT("/users", u.Update)
	e.DELETE("/users/:id", u.Delete)
}

// Find all users
func (u *UserCntrl) Find(c echo.Context) (err error) {
	var users []*repository.User
	ctx := c.Request().Context()

	if users, err = u.Svc.Find(ctx); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, users)
}

// FindOne user
func (u *UserCntrl) FindOne(ec echo.Context) (err error) {
	var (
		id   int64
		user *repository.User
	)

	ctx := ec.Request().Context()
	if id, err = strconv.ParseInt(ec.Param("id"), 10, 64); err != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, "Invalid ID")
	}
	user, err = u.Svc.FindOne(ctx, id)
	if err == sql.ErrNoRows {
		return echo.NewHTTPError(http.StatusNotFound)
	}
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return ec.JSON(http.StatusOK, user)
}

// Create user
func (u *UserCntrl) Create(c echo.Context) (err error) {
	var user repository.User
	ctx := c.Request().Context()
	if err = c.Bind(&user); err != nil {
		return err
	}
	if err = user.Validate(); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	lastInsertID, err := u.Svc.Insert(ctx, user)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, err.Error())
	}
	user.ID = lastInsertID
	return c.JSON(http.StatusCreated, user)
}

// Delete user
func (u *UserCntrl) Delete(c echo.Context) (err error) {
	var id int64
	ctx := c.Request().Context()
	if id, err = strconv.ParseInt(c.Param("id"), 10, 64); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid ID")
	}
	if err = u.Svc.Delete(ctx, id); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, GeneralResponse{
		Message: fmt.Sprintf("Success delete user #%d", id),
	})
}

// Update user
func (u *UserCntrl) Update(c echo.Context) (err error) {
	var user repository.User
	ctx0 := c.Request().Context()
	if err = c.Bind(&user); err != nil {
		return err
	}
	if user.ID <= 0 {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid ID")
	}
	if err = user.Validate(); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err = u.Svc.Update(ctx0, user); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, GeneralResponse{
		Message: fmt.Sprintf("Success update user #%d", user.ID),
	})
}
