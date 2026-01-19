package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rkaturabara/flickly/internal/application/commons/handlers"
	"github.com/rkaturabara/flickly/internal/application/commons/helpers"
	"github.com/rkaturabara/flickly/internal/application/simplerule/viewmodel"
	"github.com/rkaturabara/flickly/internal/infra/crosscutting/utilities"
	"github.com/rkaturabara/flickly/internal/domain/simplerule/commandhandlers"
)

type SimpleRuleHandler struct {
	handlers.Handler
}

// NewSimpleRuleHandler cria uma nova instância de SimpleRuleHandler
func NewSimpleRuleHandler(serviceCollection utilities.IServiceCollection) *SimpleRuleHandler {
	return &SimpleRuleHandler{
		Handler: handlers.NewHandler(serviceCollection),	
	}
}

// PostExpense cria um novo expense
// @Summary Criar expense
// @Description Cria um novo expense com os dados fornecidos
// @Tags simplerule
// @Accept json
// @Produce json
// @Param expense body viewmodel.PostExpenseViewModel true "Dados do expense"
// @Success 200 {object} viewmodel.PostExpenseViewModelResponse
// @Failure 400 {object} object
// @Router /expense [post]
func (u *SimpleRuleHandler) PostExpense(c *gin.Context) {
	helpers.ViewHelperWithSuccessStatusCode[viewmodel.PostExpenseViewModel, commandhandlers.CreateExpenseCommandHandler, viewmodel.PostExpenseViewModelResponse](
		c, &u.Handler, http.StatusCreated)
}
