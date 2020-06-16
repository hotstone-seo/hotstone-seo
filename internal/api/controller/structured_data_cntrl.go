package controller

import (
	"database/sql"
	"fmt"
	"net/http"
	"strconv"

	"github.com/typical-go/typical-rest-server/pkg/dbkit"

	"github.com/hotstone-seo/hotstone-seo/internal/api/repository"
	"github.com/hotstone-seo/hotstone-seo/internal/api/service"
	"github.com/hotstone-seo/hotstone-seo/internal/app/infra"
	"github.com/labstack/echo"
	"go.uber.org/dig"
)

type StructuredDataCntrl struct {
	dig.In
	service.StructuredDataService
	*infra.App
}

func (c *StructuredDataCntrl) Route(e *echo.Group) {
	e.GET("/structured-data", c.Find)
	e.POST("/structured-data", c.Create)
	e.GET("/structured-data/:id", c.FindOne)
	e.PUT("/structured-data/:id", c.Update)
	e.DELETE("/structured-data/:id", c.Delete)
}

// Create new structured data
func (c *StructuredDataCntrl) Create(ctx echo.Context) (err error) {
	var structData repository.StructuredData
	var lastInsertID int64
	ctx0 := ctx.Request().Context()
	if err = ctx.Bind(&structData); err != nil {
		return err
	}
	if err = structData.Validate(); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if lastInsertID, err = c.StructuredDataService.Insert(ctx0, structData); err != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, err.Error())
	}
	structData.ID = lastInsertID
	return ctx.JSON(http.StatusCreated, structData)
}

// Find all structured data
func (c *StructuredDataCntrl) Find(ce echo.Context) (err error) {
	var (
		structDatas []*repository.StructuredData
		opts        []dbkit.SelectOption
		ctx         = ce.Request().Context()
	)
	if ruleID := ce.QueryParam("rule_id"); ruleID != "" {
		if _, err := strconv.Atoi(ruleID); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, "Invalid Rule ID")
		}
		opts = append(opts, dbkit.Equal("rule_id", ruleID))
	} else {
		// TODO: return validation error
	}
	if structDatas, err = c.StructuredDataService.Find(ctx, opts...); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return ce.JSON(http.StatusOK, structDatas)
}

// FindOne structured data
func (c *StructuredDataCntrl) FindOne(ec echo.Context) (err error) {
	var (
		id         int64
		structData *repository.StructuredData
	)
	ctx := ec.Request().Context()
	if id, err = strconv.ParseInt(ec.Param("id"), 10, 64); err != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, "Invalid ID")
	}
	structData, err = c.StructuredDataService.FindOne(ctx, id)
	if err == sql.ErrNoRows {
		return echo.NewHTTPError(http.StatusNotFound)
	}
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return ec.JSON(http.StatusOK, structData)
}

// Delete structured data
func (c *StructuredDataCntrl) Delete(ctx echo.Context) (err error) {
	var id int64
	ctx0 := ctx.Request().Context()
	if id, err = strconv.ParseInt(ctx.Param("id"), 10, 64); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid ID")
	}
	if err = c.StructuredDataService.Delete(ctx0, id); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return ctx.JSON(http.StatusOK, GeneralResponse{
		Message: fmt.Sprintf("Success delete structured data #%d", id),
	})
}

// Update structured data
func (c *StructuredDataCntrl) Update(ctx echo.Context) (err error) {
	var structData repository.StructuredData
	ctx0 := ctx.Request().Context()
	if err = ctx.Bind(&structData); err != nil {
		return err
	}
	if structData.ID <= 0 {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid ID")
	}
	if err = structData.Validate(); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err = c.StructuredDataService.Update(ctx0, structData); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return ctx.JSON(http.StatusOK, GeneralResponse{
		Message: fmt.Sprintf("Success update structured data #%d", structData.ID),
	})
}
