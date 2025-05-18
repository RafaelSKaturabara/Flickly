package controllers

import (
	"errors"
	"flickly/internal/api/commons/view_model"
	"flickly/internal/domain/core"
	"flickly/internal/infra/crosscutting/utilities"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Controller struct {
	mapper utilities.Mapper
}

func NewController(collection utilities.IServiceCollection) Controller {
	return Controller{
		mapper: utilities.GetService[utilities.Mapper](collection),
	}
}

func (c *Controller) SuccessOrErrorResponse(ctx *gin.Context, execute func(ct *gin.Context) (interface{}, error), statusCode int) {
	response, err := execute(ctx)
	if err != nil {
		c.errorResponse(ctx, err)
		return
	}
	c.successResponse(ctx, response, statusCode)
}

func (c *Controller) successResponse(ctx *gin.Context, successResponse interface{}, successStatusCode int) {
	statusCode := http.StatusOK
	if successStatusCode > 0 {
		statusCode = successStatusCode
	}

	ctx.JSON(statusCode, successResponse)
}

func (c *Controller) errorResponse(ctx *gin.Context, err error) {
	var domainError *core.DomainError
	if errors.As(err, &domainError) {
		var errorResponse view_model.ErrorResponse
		if errMap := c.mapper.Map(domainError, &errorResponse); errMap != nil {
			ctx.JSON(http.StatusTeapot, gin.H{
				"message": errMap.Error(),
			})
			return
		}

		ctx.JSON(domainError.StatusCode, errorResponse)
		return
	}

	ctx.JSON(http.StatusTeapot, gin.H{
		"message": err.Error(),
	})
}
