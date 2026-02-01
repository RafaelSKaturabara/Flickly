package users

import (
	"github.com/RafaelSKaturabara/Flickly/internal/application"
	"github.com/RafaelSKaturabara/Flickly/internal/infra/crosscutting/utilities"
	"github.com/RafaelSKaturabara/Flickly/internal/presentation/commons/middleware"
	"github.com/RafaelSKaturabara/Flickly/internal/presentation/identity/handlers"
	"github.com/gin-gonic/gin"
)

// Configura e inicia o roteador Gin
func Startup(router *gin.Engine, serviceCollection utilities.IServiceCollection) {
	userController := handlers.NewUserHandler(serviceCollection)

	// Configurando rotas
	router.POST("/user", userController.PostUser)

	// Cria o controlador de autenticação
	oauthController := handlers.NewOAuthHandler(serviceCollection)

	// Cria o middleware JWT
	jwtMiddleware := middleware.NewJWTMiddleware("config.JWTSecret")

	// Grupo de rotas de autenticação
	authGroup := router.Group("/oauth")
	{
		// Rotas públicas
		authGroup.POST("/register", oauthController.Register)
		authGroup.POST("/token", jwtMiddleware.RefreshToken(), oauthController.Token)

		// Rotas protegidas
		authGroup.GET("/me", jwtMiddleware.Auth(), func(c *gin.Context) {
			user, _ := c.Request.Context().Value(application.UserContextKey).(*application.UserAuth)
			c.JSON(200, gin.H{"user": user})
		})

		// Rota protegida que requer role admin
		authGroup.GET("/admin",
			jwtMiddleware.Auth(),
			jwtMiddleware.Role("admin"),
			func(c *gin.Context) {
				c.JSON(200, gin.H{"message": "Acesso permitido"})
			})
	}
}
