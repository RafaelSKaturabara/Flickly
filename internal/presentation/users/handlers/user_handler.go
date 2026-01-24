package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	command_handlers "github.com/RafaelSKaturabara/Flickly/internal/application/users/handlers"
	"github.com/RafaelSKaturabara/Flickly/internal/infra/crosscutting/utilities"
	"github.com/RafaelSKaturabara/Flickly/internal/presentation/commons/handlers"
	"github.com/RafaelSKaturabara/Flickly/internal/presentation/commons/helpers"
	viewmodel "github.com/RafaelSKaturabara/Flickly/internal/presentation/users/viewmodel"
)

type UserHandler struct {
	handlers.Handler
}

// NewUserController cria uma nova instância de UserController
func NewUserHandler(serviceCollection utilities.IServiceCollection) *UserHandler {
	return &UserHandler{
		Handler: handlers.NewHandler(serviceCollection),
	}
}

// PostUser cria um novo usuário
// @Summary Criar usuário
// @Description Cria um novo usuário com os dados fornecidos
// @Tags users
// @Accept json
// @Produce json
// @Param user body viewmodel.CreateUserRequest true "Dados do usuário"
// @Success 200 {object} viewmodel.CreateUserResponse
// @Failure 400 {object} object
// @Router /user [post]
func (u *UserHandler) PostUser(c *gin.Context) {
	helpers.ViewHelperWithSuccessStatusCode[viewmodel.CreateUserRequest, command_handlers.CreateUserCommand, viewmodel.CreateUserResponse](
		c, &u.Handler, http.StatusCreated)
}
