package controller

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
	"github.com/typical-go/typical-rest-server/pkg/dbkit"

	"github.com/hotstone-seo/hotstone-seo/server/repository"
	"github.com/hotstone-seo/hotstone-seo/server/service"
	"go.uber.org/dig"
	"gopkg.in/go-playground/validator.v9"
)

// ClientKeyCntrl is controller to client_key entity
type ClientKeyCntrl struct {
	dig.In
	service.ClientKeyService
}

// Route to define API Route
func (c *ClientKeyCntrl) Route(e *echo.Group) {
	e.GET("/client-keys", c.Find)
	e.POST("/client-keys", c.Create)
	e.GET("/client-keys/:id", c.FindOne)
	e.PUT("/client-keys", c.Update)
	e.DELETE("/client-keys/:id", c.Delete)
}

// Create client_key
func (c *ClientKeyCntrl) Create(ctx echo.Context) (err error) {
	var clientKey repository.ClientKey
	ctx0 := ctx.Request().Context()
	if err = ctx.Bind(&clientKey); err != nil {
		return err
	}
	if err = validator.New().Struct(clientKey); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if clientKey, err = c.ClientKeyService.Insert(ctx0, clientKey); err != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, err.Error())
	}
	return ctx.JSON(http.StatusCreated, clientKey)
}

// Find of client_key
func (c *ClientKeyCntrl) Find(ctx echo.Context) (err error) {
	var clientKey []*repository.ClientKey
	ctx0 := ctx.Request().Context()
	if clientKey, err = c.ClientKeyService.Find(ctx0); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return ctx.JSON(http.StatusOK, clientKey)
}

// FindOne client_key
func (c *ClientKeyCntrl) FindOne(ctx echo.Context) (err error) {
	var id int64
	var clientKey *repository.ClientKey
	ctx0 := ctx.Request().Context()
	if id, err = strconv.ParseInt(ctx.Param("id"), 10, 64); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid ID")
	}
	if clientKey, err = c.ClientKeyService.FindOne(ctx0, dbkit.Equal("id", id)); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	if clientKey == nil {
		return echo.NewHTTPError(http.StatusNotFound, fmt.Sprintf("ClientKey#%d not found", id))
	}
	return ctx.JSON(http.StatusOK, clientKey)
}

// Delete client_key
func (c *ClientKeyCntrl) Delete(ctx echo.Context) (err error) {
	var id int64
	ctx0 := ctx.Request().Context()
	if id, err = strconv.ParseInt(ctx.Param("id"), 10, 64); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid ID")
	}
	if err = c.ClientKeyService.Delete(ctx0, id); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return ctx.JSON(http.StatusOK, GeneralResponse{
		Message: fmt.Sprintf("Success delete API Key #%d", id),
	})
}

// Update client key
func (c *ClientKeyCntrl) Update(ctx echo.Context) (err error) {
	var tag repository.ClientKey
	ctx0 := ctx.Request().Context()
	if err = ctx.Bind(&tag); err != nil {
		return err
	}
	if tag.ID <= 0 {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid ID")
	}
	if err = tag.Validate(); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err = c.ClientKeyService.Update(ctx0, tag); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return ctx.JSON(http.StatusOK, GeneralResponse{
		Message: fmt.Sprintf("Success update tag #%d", tag.ID),
	})
}
