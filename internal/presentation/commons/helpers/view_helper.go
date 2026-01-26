package helpers

import (
	"github.com/gin-gonic/gin"
	"github.com/RafaelSKaturabara/Flickly/internal/domain/core/mediator"
	"github.com/RafaelSKaturabara/Flickly/internal/presentation/commons/handlers"
)

func ViewHelperWithSuccessStatusCode[VMRequest any, CCommand mediator.Request, VMResponse any](ctx *gin.Context, handler *handlers.Handler, statusCode int) {
	var vmRequest VMRequest
	if err := ctx.ShouldBindJSON(&vmRequest); err != nil {
		handler.ErrorResponse(ctx, err)
		return
	}

	var command CCommand
	if err := handler.Mapper.Map(vmRequest, &command); err != nil {
		handler.ErrorResponse(ctx, err)
		return
	}

	sendToMediatorAndGenerateResponse[VMResponse](ctx, handler, statusCode, command)
}

func ViewHelperUrlEncodedWith[VMRequest any, CCommand mediator.Request, VMResponse any](ctx *gin.Context, handler *handlers.Handler) {
	var vmRequest VMRequest
	if err := ctx.ShouldBind(&vmRequest); err != nil {
		handler.ErrorResponse(ctx, err)
		return
	}

	var command CCommand
	if err := handler.Mapper.Map(vmRequest, &command); err != nil {
		handler.ErrorResponse(ctx, err)
		return
	}

	sendToMediatorAndGenerateResponse[VMResponse](ctx, handler, 0, command)
}

func ViewHelperWith[VMRequest any, CCommand mediator.Request, VMResponse any](ctx *gin.Context, handler *handlers.Handler) {
	ViewHelperWithSuccessStatusCode[VMRequest, CCommand, VMResponse](ctx, handler, 0)
}

func sendToMediatorAndGenerateResponse[VMResponse any](ctx *gin.Context, handler *handlers.Handler, statusCode int, command mediator.Request) {
	response, err := handler.Mediator.Send(ctx.Request.Context(), command)
	if err != nil {
		handler.ErrorResponse(ctx, err)
		return
	}

	var vmResponse VMResponse
	if err = handler.Mapper.Map(response, &vmResponse); err != nil {
		handler.ErrorResponse(ctx, err)
		return
	}
	handler.SuccessResponse(ctx, vmResponse, statusCode)
}
