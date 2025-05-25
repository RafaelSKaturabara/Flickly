package helpers

import (
	"github.com/gin-gonic/gin"
	"github.com/rkaturabara/flickly/internal/application/commons/handlers"
	"github.com/rkaturabara/flickly/internal/domain/core/mediator"
)

func ViewHelperWithSuccessStatusCode[VMRequest any, CCommand mediator.Request, VMResponse any](ctx *gin.Context, controller *handlers.Handler, statusCode int) {
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

	sendToMediatorAndGenerateResponse[VMResponse](ctx, controller, statusCode, command)
}

func ViewHelperUrlEncodedWith[VMRequest any, CCommand mediator.Request, VMResponse any](ctx *gin.Context, controller *handlers.Handler) {
	var vmRequest VMRequest
	if err := ctx.ShouldBind(&vmRequest); err != nil {
		controller.ErrorResponse(ctx, err)
		return
	}

	var command CCommand
	if err := controller.Mapper.Map(vmRequest, &command); err != nil {
		controller.ErrorResponse(ctx, err)
		return
	}

	sendToMediatorAndGenerateResponse[VMResponse](ctx, controller, 0, command)
}

func ViewHelperWith[VMRequest any, CCommand mediator.Request, VMResponse any](ctx *gin.Context, controller *handlers.Handler) {
	ViewHelperWithSuccessStatusCode[VMRequest, CCommand, VMResponse](ctx, controller, 0)
}

func sendToMediatorAndGenerateResponse[VMResponse any](ctx *gin.Context, controller *handlers.Handler, statusCode int, command mediator.Request) {
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
