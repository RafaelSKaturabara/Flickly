package handlers

import (
	"errors"
	"net/http"

	"github.com/rkaturabara/flickly/internal/application/commons/view_model"
	"github.com/rkaturabara/flickly/internal/domain/core"
	"github.com/rkaturabara/flickly/internal/domain/core/mediator"
	"github.com/rkaturabara/flickly/internal/infra/crosscutting/utilities"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	Mapper   utilities.Mapper
	Mediator mediator.Mediator
}

func NewHandler(collection utilities.IServiceCollection) Handler {
	return Handler{
		Mapper:   utilities.GetService[utilities.Mapper](collection),
		Mediator: utilities.GetService[mediator.Mediator](collection),
	}
}

func (c *Handler) SuccessResponse(ctx *gin.Context, successResponse any, successStatusCode int) {
	statusCode := http.StatusOK
	if successStatusCode > 0 {
		statusCode = successStatusCode
	}

	ctx.JSON(statusCode, successResponse)
}

func (c *Handler) ErrorResponse(ctx *gin.Context, err error) {
	var domainError *core.DomainError
	if errors.As(err, &domainError) {
		var errorResponse view_model.ErrorResponse
		if errMap := c.Mapper.Map(domainError, &errorResponse); errMap != nil {
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
