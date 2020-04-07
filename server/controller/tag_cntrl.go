package controller

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/typical-go/typical-rest-server/pkg/dbkit"

	"github.com/hotstone-seo/hotstone-seo/server/config"
	"github.com/hotstone-seo/hotstone-seo/server/repository"
	"github.com/hotstone-seo/hotstone-seo/server/service"
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
	ctx0 := ctx.Request().Context()
	if err = ctx.Bind(&tag); err != nil {
		return err
	}
	if err = tag.Validate(); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if lastInsertID, err = c.TagService.Insert(ctx0, tag); err != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, err.Error())
	}
	tag.ID = lastInsertID
	return ctx.JSON(http.StatusCreated, tag)
}

// Find all tag
func (c *TagCntrl) Find(ce echo.Context) (err error) {
	var (
		tags []*repository.Tag
		opts []dbkit.FindOption
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
func (c *TagCntrl) FindOne(ctx echo.Context) (err error) {
	var id int64
	var tag *repository.Tag
	ctx0 := ctx.Request().Context()
	if id, err = strconv.ParseInt(ctx.Param("id"), 10, 64); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid ID")
	}
	if tag, err = c.TagService.FindOne(ctx0, id); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	if tag == nil {
		return echo.NewHTTPError(http.StatusNotFound, fmt.Sprintf("Tag#%d not found", id))
	}
	return ctx.JSON(http.StatusOK, tag)
}

// Delete tag
func (c *TagCntrl) Delete(ctx echo.Context) (err error) {
	var id int64
	ctx0 := ctx.Request().Context()
	if id, err = strconv.ParseInt(ctx.Param("id"), 10, 64); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid ID")
	}
	if err = c.TagService.Delete(ctx0, id); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return ctx.JSON(http.StatusOK, GeneralResponse{
		Message: fmt.Sprintf("Success delete tag #%d", id),
	})
}

// Update tag
func (c *TagCntrl) Update(ctx echo.Context) (err error) {
	var tag repository.Tag
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
	if err = c.TagService.Update(ctx0, tag); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return ctx.JSON(http.StatusOK, GeneralResponse{
		Message: fmt.Sprintf("Success update tag #%d", tag.ID),
	})
}
