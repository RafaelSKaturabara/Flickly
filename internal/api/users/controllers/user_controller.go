package controllers

import (
	viewmodels "flickly/internal/api/users/viewmodels"
	"flickly/internal/domain/core/mediator"
	"flickly/internal/domain/users/commands"
	"flickly/internal/domain/users/repositories"
	"flickly/internal/infra/crosscutting/utilities"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserController struct {
	mediator       mediator.Mediator
	userRepository repositories.IUserRepository
}

// NewUserController cria uma nova instância de UserController
func NewUserController(collection utilities.IServiceCollection) *UserController {
	return &UserController{
		mediator:       utilities.GetService[mediator.Mediator](collection),
		userRepository: utilities.GetService[repositories.IUserRepository](collection),
	}
}

// PostUser cria um novo usuário
// @Summary Criar um novo usuário
// @Description Cria um novo usuário com base nos dados fornecidos
// @Tags users
// @Accept json
// @Produce json
// @Param user body commands.CreateUserCommand true "Dados do usuário"
// @Success 200 {object} interface{}
// @Failure 400 {object} interface{}
// @Router /user [post]
func (u *UserController) PostUser(c *gin.Context) {
	var createUserCommand commands.CreateUserCommand

	if err := c.ShouldBindJSON(&createUserCommand); err != nil {
		// Retorna um erro se não puder vincular
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	response, err := u.mediator.Send(c, createUserCommand)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
	}

	// Caso as credenciais sejam inválidas
	c.JSON(http.StatusOK, response)
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
	grantType := c.PostForm("grant_type")
	clientID := c.PostForm("client_id")
	clientSecret := c.PostForm("client_secret")
	username := c.PostForm("username")
	password := c.PostForm("password")

	// Aqui você deve implementar a lógica para verificar as credenciais
	if grantType == "password" && clientID == "my_client_id" && clientSecret == "my_client_secret" {

		user, err := u.userRepository.GetUserByEmail(username)

		if err != nil {
			c.JSON(http.StatusUnauthorized, err)
		}

		if user != nil && password != "" {
			// Se as credenciais estão corretas, retornar um token de exemplo
			response := viewmodels.TokenResponse{
				AccessToken: "some_generated_token",
				TokenType:   "Bearer " + user.Name,
				ExpiresIn:   3600, // Expires in 1 hour
			}
			c.JSON(http.StatusOK, response)
			return
		}
	}

	// Caso as credenciais sejam inválidas
	c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid_grant"})
}
