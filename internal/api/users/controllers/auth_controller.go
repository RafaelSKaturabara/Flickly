package controllers

import (
	"net/http"

	"github.com/rkaturabara/flickly/internal/api/commons/controllers"
	"github.com/rkaturabara/flickly/internal/api/commons/helpers"
	viewmodel "github.com/rkaturabara/flickly/internal/api/users/viewmodel"
	"github.com/rkaturabara/flickly/internal/domain/users/command_handlers"
	"github.com/rkaturabara/flickly/internal/domain/users/repositories"
	"github.com/rkaturabara/flickly/internal/infra/auth"
	"github.com/rkaturabara/flickly/internal/infra/crosscutting/utilities"

	"github.com/gin-gonic/gin"
)

// AuthController gerencia as operações de autenticação
// @Summary Controlador de autenticação
// @Description Gerencia operações de registro e autenticação de usuários
// @Tags auth
type AuthController struct {
	controllers.Controller
	authService auth.AuthServiceInterface
}

// NewAuthController cria uma nova instância do AuthController
// @Summary Cria um novo controlador de autenticação
// @Description Inicializa um novo controlador com as dependências necessárias
// @Tags auth
// @Param serviceCollection body utilities.IServiceCollection true "Coleção de serviços"
// @Return *AuthController
func NewAuthController(serviceCollection utilities.IServiceCollection) *AuthController {
	return &AuthController{
		Controller:  controllers.NewController(serviceCollection),
		authService: utilities.GetService[auth.AuthServiceInterface](serviceCollection),
	}
}

// Register lida com o registro de novos usuários
// @Summary Registra um novo usuário
// @Description Cria uma nova conta de usuário no sistema
// @Tags auth
// @Accept json
// @Produce json
// @Param request body auth.RegisterRequest true "Dados do usuário para registro"
// @Success 201 {object} auth.RegisterResponse "Usuário registrado com sucesso"
// @Failure 400 {object} map[string]string "Dados inválidos"
// @Failure 409 {object} map[string]string "Usuário já existe"
// @Failure 500 {object} map[string]string "Erro interno do servidor"
// @Router /auth/register [post]
func (c *AuthController) Register(ctx *gin.Context) {
	var req auth.RegisterRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	response, err := c.authService.Register(ctx.Request.Context(), req)
	if err != nil {
		if err == repositories.ErrUserAlreadyExists {
			ctx.JSON(http.StatusConflict, gin.H{"error": "Usuário já existe"})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao registrar usuário"})
		return
	}

	ctx.JSON(http.StatusCreated, response)
}

// Token lida com a autenticação OAuth2
// @Summary Autentica um usuário
// @Description Gera um token de acesso usando OAuth2
// @Tags auth
// @Accept json
// @Produce json
// @Param request body auth.TokenRequest true "Credenciais de autenticação"
// @Success 200 {object} auth.TokenResponse "Token gerado com sucesso"
// @Failure 400 {object} map[string]string "Dados inválidos"
// @Failure 401 {object} map[string]string "Credenciais inválidas"
// @Failure 500 {object} map[string]string "Erro interno do servidor"
// @Router /auth/token [post]
func (c *AuthController) Token(ctx *gin.Context) {

	helpers.ViewHelper[viewmodel.CreateUserRequest, command_handlers.CreateTokenCommand, viewmodel.CreateUserResponse](ctx, &c.Controller, http.StatusCreated)

	var req auth.TokenRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	response, err := c.authService.Token(ctx.Request.Context(), req)
	if err != nil {
		switch err {
		case auth.ErrInvalidClient:
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Client ID ou Client Secret inválidos"})
		case auth.ErrInvalidGrant:
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Grant type inválido"})
		case auth.ErrInvalidCredentials:
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Credenciais inválidas"})
		default:
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao gerar token"})
		}
		return
	}

	ctx.JSON(http.StatusOK, response)
}
