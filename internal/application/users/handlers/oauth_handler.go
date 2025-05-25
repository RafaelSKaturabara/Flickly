package handlers

import (
	"net/http"

	"github.com/rkaturabara/flickly/internal/application/commons/handlers"
	"github.com/rkaturabara/flickly/internal/application/commons/helpers"
	viewmodel "github.com/rkaturabara/flickly/internal/application/users/viewmodel"
	"github.com/rkaturabara/flickly/internal/domain/users/command_handlers"
	"github.com/rkaturabara/flickly/internal/infra/crosscutting/utilities"

	"github.com/gin-gonic/gin"
)

// AuthController gerencia as operações de autenticação
// @Summary Controlador de autenticação
// @Description Gerencia operações de registro e autenticação de usuários
// @Tags oauth
type OAuthHandler struct {
	handlers.Handler
}

// NewAuthController cria uma nova instância do AuthController
// @Summary Cria um novo controlador de autenticação
// @Description Inicializa um novo controlador com as dependências necessárias
// @Tags oauth
// @Param serviceCollection body utilities.IServiceCollection true "Coleção de serviços"
// @Return *AuthController
func NewOAuthHandler(serviceCollection utilities.IServiceCollection) *OAuthHandler {
	return &OAuthHandler{
		Handler: handlers.NewHandler(serviceCollection),
	}
}

// Register lida com o registro de novos usuários
// @Summary Registra um novo usuário
// @Description Cria uma nova conta de usuário no sistema
// @Tags oauth
// @Accept json
// @Produce json
// @Param request body viewmodel.RegisterRequest true "Dados do usuário para registro"
// @Success 201 {object} viewmodel.RegisterResponse "Usuário registrado com sucesso"
// @Failure 400 {object} map[string]string "Dados inválidos"
// @Failure 409 {object} map[string]string "Usuário já existe"
// @Failure 500 {object} map[string]string "Erro interno do servidor"
// @Router /oauth/register [post]
func (c *OAuthHandler) Register(ctx *gin.Context) {
	var req viewmodel.RegisterRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//response, err := c.authService.Register(ctx.Request.Context(), req)
	//if err != nil {
	//	if err == repositories.ErrUserAlreadyExists {
	//		ctx.JSON(http.StatusConflict, gin.H{"error": "Usuário já existe"})
	//		return
	//	}
	//	ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao registrar usuário"})
	//	return
	//}

	ctx.JSON(http.StatusCreated, gin.H{"error": "err.Error()"})
}

// Token lida com a autenticação OAuth2
// @Summary Autentica um usuário
// @Description Gera um token de acesso usando OAuth2
// @Tags oauth
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param grant_type formData string true "Tipo de concessão (password)"
// @Param client_id formData string true "ID do cliente"
// @Param client_secret formData string true "Segredo do cliente"
// @Param username formData string true "Nome de usuário (email)"
// @Param password formData string true "Senha do usuário"
// @Param scope formData string false "Escopo do token"
// @Success 200 {object} viewmodel.TokenResponse "Token gerado com sucesso"
// @Failure 400 {object} map[string]string "Dados inválidos"
// @Failure 401 {object} map[string]string "Credenciais inválidas"
// @Failure 500 {object} map[string]string "Erro interno do servidor"
// @Router /oauth/token [post]
func (c *OAuthHandler) Token(ctx *gin.Context) {
	helpers.ViewHelperUrlEncodedWith[viewmodel.TokenRequest, command_handlers.CreateTokenCommand, viewmodel.TokenResponse](ctx, &c.Handler)
}

// RefreshToken lida com a renovação do token de acesso
// @Summary Renova o token de acesso
// @Description Gera um novo token de acesso usando o refresh token
// @Tags oauth
// @Accept json
// @Produce json
// @Param request body viewmodel.RefreshTokenRequest true "Refresh token"
// @Success 200 {object} viewmodel.TokenResponse "Token renovado com sucesso"
// @Failure 400 {object} map[string]string "Dados inválidos"
// @Failure 401 {object} map[string]string "Token inválido"
// @Failure 500 {object} map[string]string "Erro interno do servidor"
// @Router /oauth/refresh [post]
func (c *OAuthHandler) RefreshToken(ctx *gin.Context) {
	helpers.ViewHelperWith[viewmodel.RefreshTokenRequest, command_handlers.RefreshTokenCommand, viewmodel.TokenResponse](ctx, &c.Handler)
}
