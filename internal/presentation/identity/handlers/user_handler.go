package handlers

import (
	"net/http"
	"time"

	"github.com/RafaelSKaturabara/Flickly/internal/application/identity/commands"
	"github.com/RafaelSKaturabara/Flickly/internal/domain/core/mediator"
	"github.com/RafaelSKaturabara/Flickly/internal/domain/identity/entities"
	"github.com/RafaelSKaturabara/Flickly/internal/infra/crosscutting/utilities"
	"github.com/RafaelSKaturabara/Flickly/internal/presentation/commons/handlers"
	"github.com/RafaelSKaturabara/Flickly/internal/presentation/commons/helpers"
	viewmodel "github.com/RafaelSKaturabara/Flickly/internal/presentation/identity/viewmodel"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type UserHandler struct {
	handlers.Handler
	mediator mediator.Mediator
}

// NewUserController cria uma nova instância de UserController
func NewUserHandler(serviceCollection utilities.IServiceCollection) *UserHandler {
	return &UserHandler{
		Handler: handlers.NewHandler(serviceCollection),
		mediator: utilities.GetService[mediator.Mediator](serviceCollection),
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
	helpers.ViewHelperWithSuccessStatusCode[viewmodel.CreateUserRequest, commands.CreateUserCommand, viewmodel.CreateUserResponse](
		c, &u.Handler, http.StatusCreated)
}

// GetAuthorize cria um novo usuário
// @Summary Criar usuário
// @Description Cria um novo usuário com os dados fornecidos
// @Tags users
// @Accept json
// @Produce json
// @Param user body viewmodel.CreateUserRequest true "Dados do usuário"
// @Success 200 {object} viewmodel.CreateUserResponse
// @Failure 400 {object} object
// @Router /authorize [get]
func (u *UserHandler) GetAuthorize(c *gin.Context) {

	// Get ClientID from header
	clientIDStr := c.Query("client_id")
		
	clientID, _ := uuid.Parse(clientIDStr)

	command := commands.GetAuthorizeCommand{
		Challenge: c.Query("challenge"),
		ClientID: clientID,
	}

	authCodeResponse, _ := u.mediator.Send(c.Request.Context(), command)
	authCode := authCodeResponse.(entities.AuthCode)

	// ver se client já vem preenchido. 
	response := viewmodel.GetAuthorizeResponse{
		Code:       authCode.Code,
		ClientName: authCode.Client.Name,
		Expires_in: int(authCode.ExpiresAt.Sub(time.Now()).Seconds()),
	}

	u.Handler.SuccessResponse(c, response, http.StatusOK)
}

// PostLogin cria um login para um usuário
// @Summary Cria um login para um usuário
// @Description Cria um novo login com os dados fornecidos
// @Tags users
// @Accept json
// @Produce json
// @Param user body viewmodel.CreateLoginRequest true "Dados do usuário"
// @Success 200 {object} viewmodel.CreateLoginResponse
// @Failure 400 {object} object
// @Router /login [post]
func (u *UserHandler) PostLogin(c *gin.Context) {
	helpers.ViewHelperWithSuccessStatusCode[viewmodel.CreateUserRequest, commands.CreateLoginCommand, viewmodel.CreateUserResponse](
		c, &u.Handler, http.StatusCreated)
}
