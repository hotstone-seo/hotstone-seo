package controller

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo"

	"github.com/hotstone-seo/hotstone-seo/server/repository"
	"github.com/hotstone-seo/hotstone-seo/server/service"
	"go.uber.org/dig"
	"gopkg.in/go-playground/validator.v9"
)

// APIKeyCntrl is controller to data_source entity
type APIKeyCntrl struct {
	dig.In
	service.APIKeyService
}

// Route to define API Route
func (c *APIKeyCntrl) Route(e *echo.Group) {
	e.GET("/api_keys", c.Find)
	e.POST("/api_keys", c.Create)
	e.GET("/api_keys/:id", c.FindOne)
	e.DELETE("/api_keys/:id", c.Delete)
}

// Create data_source
func (c *APIKeyCntrl) Create(ctx echo.Context) (err error) {
	var apiKey repository.APIKey
	var lastInsertID int64
	ctx0 := ctx.Request().Context()
	if err = ctx.Bind(&apiKey); err != nil {
		return err
	}
	if err = validator.New().Struct(apiKey); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if lastInsertID, err = c.APIKeyService.Insert(ctx0, apiKey); err != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, err.Error())
	}
	apiKey.ID = lastInsertID
	return ctx.JSON(http.StatusCreated, apiKey)
}

// Find of data_source
func (c *APIKeyCntrl) Find(ctx echo.Context) (err error) {
	var apiKey []*repository.APIKey
	ctx0 := ctx.Request().Context()
	if apiKey, err = c.APIKeyService.Find(ctx0); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return ctx.JSON(http.StatusOK, apiKey)
}

// FindOne data_source
func (c *APIKeyCntrl) FindOne(ctx echo.Context) (err error) {
	var id int64
	var apiKey *repository.APIKey
	ctx0 := ctx.Request().Context()
	if id, err = strconv.ParseInt(ctx.Param("id"), 10, 64); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid ID")
	}
	if apiKey, err = c.APIKeyService.FindOne(ctx0, id); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	if apiKey == nil {
		return echo.NewHTTPError(http.StatusNotFound, fmt.Sprintf("APIKey#%d not found", id))
	}
	return ctx.JSON(http.StatusOK, apiKey)
}

// Delete data_source
func (c *APIKeyCntrl) Delete(ctx echo.Context) (err error) {
	var id int64
	ctx0 := ctx.Request().Context()
	if id, err = strconv.ParseInt(ctx.Param("id"), 10, 64); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid ID")
	}
	if err = c.APIKeyService.Delete(ctx0, id); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return ctx.JSON(http.StatusOK, GeneralResponse{
		Message: fmt.Sprintf("Success delete API Key #%d", id),
	})
}