package users

import (
	"flickly/internal/api/users/controllers"
	"flickly/internal/infra/cross-cutting/utilities"
	"github.com/gin-gonic/gin"
)

// Configura e inicia o roteador Gin
func Startup(router *gin.Engine, serviceCollection utilities.IServiceCollection) {

	userController := controllers.NewUserController(serviceCollection)

	// Configurando rotas
	router.POST("/oauth/token", userController.PostOauthToken)
	router.POST("/user", userController.PostUser)
}
