package controller

import (
	"golang-url-shortener/internal/dto"
	"golang-url-shortener/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Controller struct {
	shortService service.IShortService
}

func NewController(
	shortService service.IShortService,
) *Controller {
	return &Controller{
		shortService: shortService,
	}
}

func (c *Controller) Short(ctx *gin.Context) {
	var data dto.ShortRequest
	if err := ctx.ShouldBindJSON(&data); err != nil {
		ctx.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Code:    400,
			Message: err.Error(),
		})
		return
	}
	result, err := c.shortService.Short(ctx, data)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Code:    400,
			Message: err.Error(),
		})
	} else {
		ctx.JSON(http.StatusOK, result)
	}
}

func (c *Controller) Find(ctx *gin.Context) {
	link := ctx.Param("link")
	result, err := c.shortService.Find(ctx, link)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Code:    400,
			Message: err.Error(),
		})
	} else {
		ctx.Redirect(http.StatusFound, result)
	}
}
