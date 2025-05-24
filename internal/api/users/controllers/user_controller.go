package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rkaturabara/flickly/internal/api/commons/controllers"
	"github.com/rkaturabara/flickly/internal/api/commons/helpers"
	viewmodel "github.com/rkaturabara/flickly/internal/api/users/viewmodel"
	"github.com/rkaturabara/flickly/internal/domain/users/command_handlers"
	"github.com/rkaturabara/flickly/internal/infra/crosscutting/utilities"
)

type UserController struct {
	controllers.Controller
}

// NewUserController cria uma nova instância de UserController
func NewUserController(serviceCollection utilities.IServiceCollection) *UserController {
	return &UserController{
		Controller: controllers.NewController(serviceCollection),
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
func (u *UserController) PostUser(c *gin.Context) {
	helpers.ViewHelper[viewmodel.CreateUserRequest, command_handlers.CreateUserCommand, viewmodel.CreateUserResponse](c, &u.Controller, http.StatusCreated)
}
