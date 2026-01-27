package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	commandhandlers "github.com/RafaelSKaturabara/Flickly/internal/application/mockedapi/handlers"
	"github.com/RafaelSKaturabara/Flickly/internal/infra/crosscutting/utilities"
	"github.com/RafaelSKaturabara/Flickly/internal/presentation/commons/handlers"
	"github.com/RafaelSKaturabara/Flickly/internal/presentation/commons/helpers"
	"github.com/RafaelSKaturabara/Flickly/internal/presentation/mockedapi/viewmodel"
)

type MockHandler struct {
	handlers.Handler
}

// NewMockHandler cria uma nova instância de MockHandler
func NewMockHandler(serviceCollection utilities.IServiceCollection) *MockHandler {
	return &MockHandler{
		Handler: handlers.NewHandler(serviceCollection),
	}
}

// PutRequestMock cria um novo mock
// @Summary Criar mock
// @Description Cria um novo mock com os dados fornecidos
// @Tags mockedapi
// @Accept json
// @Produce json
// @Param mock body viewmodel.RouteViewModel true "Dados do mock"
// @Success 200 {object} viewmodel.RouteViewModel
// @Failure 400 {object} object
// @Router /mock [put]
func (u *MockHandler) PutRequestMock(c *gin.Context) {
	helpers.ViewHelperWithSuccessStatusCode[viewmodel.RouteViewModel, commandhandlers.CreateRouteCommandHandler, viewmodel.RouteViewModel](
		c, &u.Handler, http.StatusCreated)
}
