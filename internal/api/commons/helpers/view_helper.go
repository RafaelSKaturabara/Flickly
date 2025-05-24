package helpers

import (
	"github.com/gin-gonic/gin"
	"github.com/rkaturabara/flickly/internal/api/commons/controllers"
)

func ViewHelper[VMRequest any, CCommand any, VMResponse any](ctx *gin.Context, controller *controllers.Controller, statusCode int) {
	var vmRequest VMRequest
	if err := ctx.ShouldBindJSON(&vmRequest); err != nil {

		controller.ErrorResponse(ctx, err)

		return
	}

	var command CCommand
	if err := controller.Mapper.Map(vmRequest, &command); err != nil {
		controller.ErrorResponse(ctx, err)
		return
	}

	response, err := controller.Mediator.Send(ctx, command)
	if err != nil {
		controller.ErrorResponse(ctx, err)
		return
	}

	var vmResponse VMResponse
	if err = controller.Mapper.Map(response, &vmResponse); err != nil {
		controller.ErrorResponse(ctx, err)
		return
	}
	controller.SuccessResponse(ctx, vmResponse, statusCode)
}
