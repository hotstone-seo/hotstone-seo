package controller

import (
	"database/sql"
	"fmt"
	"net/http"
	"strconv"

	"github.com/typical-go/typical-rest-server/pkg/dbkit"

	"github.com/hotstone-seo/hotstone-seo/internal/api/repository"
	"github.com/hotstone-seo/hotstone-seo/internal/api/service"
	"github.com/hotstone-seo/hotstone-seo/internal/app/config"
	"github.com/labstack/echo"
	"go.uber.org/dig"
)

// TagCntrl is controller to tag entity
type TagCntrl struct {
	dig.In
	service.TagService
	*config.Config
}

// Route to define API Route
func (c *TagCntrl) Route(e *echo.Group) {
	e.GET("/tags", c.Find)
	e.POST("/tags", c.Create)
	e.GET("/tags/:id", c.FindOne)
	e.PUT("/tags", c.Update)
	e.DELETE("/tags/:id", c.Delete)
}

// Create tag
func (c *TagCntrl) Create(ctx echo.Context) (err error) {
	var tag repository.Tag
	var lastInsertID int64
	reqCtx := ctx.Request().Context()
	if err = ctx.Bind(&tag); err != nil {
		return err
	}
	if err = tag.Validate(); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if lastInsertID, err = c.TagService.Insert(reqCtx, tag); err != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, err.Error())
	}
	tag.ID = lastInsertID
	return ctx.JSON(http.StatusCreated, tag)
}

// Find all tag
func (c *TagCntrl) Find(ce echo.Context) (err error) {
	var (
		tags []*repository.Tag
		opts []dbkit.SelectOption
		ctx  = ce.Request().Context()
	)

	if ruleID := ce.QueryParam("rule_id"); ruleID != "" {
		if _, err := strconv.Atoi(ruleID); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, "Invalid Rule ID")
		}
		opts = append(opts, dbkit.Equal("rule_id", ruleID))
	} else {
		// TODO: return validation error
	}

	if locale := ce.QueryParam("locale"); locale != "" {
		opts = append(opts, dbkit.Equal("locale", locale))
	} else {
		// TODO: return validation error
	}

	if tags, err = c.TagService.Find(ctx, opts...); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return ce.JSON(http.StatusOK, tags)
}

// FindOne tag
func (c *TagCntrl) FindOne(ec echo.Context) (err error) {
	var (
		id  int64
		tag *repository.Tag
	)
	ctx := ec.Request().Context()
	if id, err = strconv.ParseInt(ec.Param("id"), 10, 64); err != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, "Invalid ID")
	}

	tag, err = c.TagService.FindOne(ctx, id)
	if err == sql.ErrNoRows {
		return echo.NewHTTPError(http.StatusNotFound)
	}

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return ec.JSON(http.StatusOK, tag)
}

// Delete tag
func (c *TagCntrl) Delete(ctx echo.Context) (err error) {
	var id int64
	reqCtx := ctx.Request().Context()
	if id, err = strconv.ParseInt(ctx.Param("id"), 10, 64); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid ID")
	}
	if err = c.TagService.Delete(reqCtx, id); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return ctx.JSON(http.StatusOK, GeneralResponse{
		Message: fmt.Sprintf("Success delete tag #%d", id),
	})
}

// Update tag
func (c *TagCntrl) Update(ctx echo.Context) (err error) {
	var tag repository.Tag
	reqCtx := ctx.Request().Context()
	if err = ctx.Bind(&tag); err != nil {
		return err
	}
	if tag.ID <= 0 {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid ID")
	}
	if err = tag.Validate(); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err = c.TagService.Update(reqCtx, tag); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return ctx.JSON(http.StatusOK, GeneralResponse{
		Message: fmt.Sprintf("Success update tag #%d", tag.ID),
	})
}
