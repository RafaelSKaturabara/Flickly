package handlers

import (
	"fmt"
	"net/http"

	"github.com/RafaelSKaturabara/Flickly/internal/application/simplerule/commands"
	"github.com/RafaelSKaturabara/Flickly/internal/application/simplerule/queries"
	"github.com/RafaelSKaturabara/Flickly/internal/domain/core"
	"github.com/RafaelSKaturabara/Flickly/internal/domain/simplerule/entities"
	usersEntities "github.com/RafaelSKaturabara/Flickly/internal/domain/users/entities"
	"github.com/RafaelSKaturabara/Flickly/internal/infra/crosscutting/utilities"
	"github.com/RafaelSKaturabara/Flickly/internal/presentation/commons/handlers"
	"github.com/RafaelSKaturabara/Flickly/internal/presentation/commons/helpers"
	"github.com/RafaelSKaturabara/Flickly/internal/presentation/commons/middleware"
	"github.com/RafaelSKaturabara/Flickly/internal/presentation/simplerule/viewmodel"
	"github.com/gin-gonic/gin"
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
// @Param expense body viewmodel.PostExpenseRequest true "Dados do expense"
// @Success 200 {object} viewmodel.PostExpenseResponse
// @Failure 400 {object} object
// @Router /simplerule/expense [post]
func (u *SimpleRuleHandler) PostExpense(c *gin.Context) {
	helpers.ViewHelperWithSuccessStatusCode[viewmodel.PostExpenseRequest, commands.CreateExpenseCommand, viewmodel.PostExpenseResponse](
		c, &u.Handler, http.StatusCreated)
}

// ListAllSubcategories lista todas as subcategorias
// @Summary Listar subcategorias
// @Description Lista todas as subcategorias
// @Tags simplerule
// @Accept json
// @Produce json
// @Success 200 {object} viewmodel.ListAllSubcategoriesViewModelResponse
// @Failure 400 {object} object
// @Router /simplerule/subcategories [get]
func (u *SimpleRuleHandler) ListAllSubcategories(c *gin.Context) {
	user, _ := c.Request.Context().Value(middleware.UserContextKey).(*usersEntities.User)

	query := queries.ListAllSubcategoriesQuery{
		UserSimpleRuleID: user.ID,
	}

	response, err := u.Mediator.Send(c.Request.Context(), query)
	if err != nil {
		u.ErrorResponse(c, err)
		return
	}

	// The query returns a slice of Subcategory entities.
	// Map each entity to the corresponding view model response item.
	// The query returns a slice of Subcategory pointers.
	var vmResponse viewmodel.ListAllSubcategoriesViewModelResponse
	subcategories, ok := response.([]*entities.Subcategory)
	if !ok {
		u.ErrorResponse(c, core.DoNotUseThisGenericError(fmt.Errorf("unexpected response type")))
		return
	}
	for _, subPtr := range subcategories {
		var subVM viewmodel.SubcategoryResponse
		if err = u.Mapper.Map(*subPtr, &subVM); err != nil {
			u.ErrorResponse(c, err)
			return
		}
		vmResponse.Subcategories = append(vmResponse.Subcategories, subVM)
	}
	u.SuccessResponse(c, vmResponse, http.StatusOK)
}
