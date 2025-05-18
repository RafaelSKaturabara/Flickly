package controllers

import (
	"flickly/internal/api/commons/controllers"
	viewmodels "flickly/internal/api/users/viewmodels"
	"flickly/internal/domain/core"
	"flickly/internal/domain/core/mediator"
	"flickly/internal/domain/users/commands"
	"flickly/internal/domain/users/repositories"
	"flickly/internal/infra/crosscutting/utilities"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserController struct {
	controllers.Controller
	mediator       mediator.Mediator
	userRepository repositories.IUserRepository
	mapper         utilities.Mapper
}

// NewUserController cria uma nova instância de UserController
func NewUserController(collection utilities.IServiceCollection) *UserController {
	return &UserController{
		Controller:     controllers.NewController(collection),
		mediator:       utilities.GetService[mediator.Mediator](collection),
		userRepository: utilities.GetService[repositories.IUserRepository](collection),
		mapper:         utilities.GetService[utilities.Mapper](collection),
	}
}

// PostUser cria um novo usuário
// @Summary Criar usuário
// @Description Cria um novo usuário com os dados fornecidos
// @Tags users
// @Accept json
// @Produce json
// @Param user body viewmodels.CreateUserRequest true "Dados do usuário"
// @Success 200 {object} viewmodels.CreateUserResponse
// @Failure 400 {object} object
// @Router /user [post]
func (u *UserController) PostUser(c *gin.Context) {
	u.SuccessOrErrorResponse(c, func(ct *gin.Context) (interface{}, error) {

		var createUserRequest viewmodels.CreateUserRequest
		if err := c.ShouldBindJSON(&createUserRequest); err != nil {
			return nil, err
		}

		var createUserCommand commands.CreateUserCommand
		if err := u.mapper.Map(createUserRequest, &createUserCommand); err != nil {
			return nil, err
		}

		response, err := u.mediator.Send(c, createUserCommand)
		if err != nil {
			return nil, err
		}

		var userResponse viewmodels.CreateUserResponse
		if err = u.mapper.Map(response, &userResponse); err != nil {
			return nil, err
		}
		return userResponse, nil
	}, http.StatusCreated)
}

// PostOauthToken autentica um usuário e gera um token
// @Summary Gerar token de autenticação
// @Description Autentica um usuário e retorna um token de acesso
// @Tags auth
// @Accept x-www-form-urlencoded
// @Produce json
// @Param grant_type formData string true "Tipo de concessão (password)"
// @Param client_id formData string true "ID do cliente"
// @Param client_secret formData string true "Segredo do cliente"
// @Param username formData string true "E-mail do usuário"
// @Param password formData string true "Senha do usuário"
// @Success 200 {object} viewmodels.TokenResponse
// @Failure 401 {object} object
// @Router /oauth/token [post]
func (u *UserController) PostOauthToken(c *gin.Context) {
	u.SuccessOrErrorResponse(c, func(ct *gin.Context) (interface{}, error) {
		grantType := c.PostForm("grant_type")
		clientID := c.PostForm("client_id")
		clientSecret := c.PostForm("client_secret")
		username := c.PostForm("username")
		password := c.PostForm("password")

		// Aqui você deve implementar a lógica para verificar as credenciais
		if grantType == "password" && clientID == "my_client_id" && clientSecret == "my_client_secret" {

			user, err := u.userRepository.GetUserByEmail(username)

			if err != nil {
				return nil, err
			}

			if user != nil && password != "" {
				// Se as credenciais estão corretas, retornar um token de exemplo
				response := viewmodels.TokenResponse{
					AccessToken: "some_generated_token",
					TokenType:   "Bearer " + user.Name,
					ExpiresIn:   3600, // Expires in 1 hour
				}
				return response, nil
			}
		}

		// Caso as credenciais sejam inválidas
		return nil, core.NewDomainErrorBuilder(nil).WithStatusCode(http.StatusUnauthorized).WithMessage("Invalid credentials").WithErrorCode(2).Build()
	}, http.StatusOK)
}
