package users

import (
	"github.com/gin-gonic/gin"
	"github.com/rkaturabara/flickly/internal/api/commons/middleware"
	"github.com/rkaturabara/flickly/internal/api/users/controllers"
	"github.com/rkaturabara/flickly/internal/infra/crosscutting/utilities"
)

// Configura e inicia o roteador Gin
func Startup(router *gin.Engine, serviceCollection utilities.IServiceCollection) {
	userController := controllers.NewUserController(serviceCollection)

	// Configurando rotas
	router.POST("/user", userController.PostUser)

	// Cria o controlador de autenticação
	authController := controllers.NewAuthController(serviceCollection)

	// Cria o middleware JWT
	jwtMiddleware := middleware.NewJWTMiddleware("config.JWTSecret")

	// Grupo de rotas de autenticação
	authGroup := router.Group("/auth")
	{
		// Rotas públicas
		authGroup.POST("/register", authController.Register)
		authGroup.POST("/token", authController.Token)

		// Rotas protegidas
		authGroup.GET("/me", jwtMiddleware.Auth(), func(c *gin.Context) {
			user, _ := c.Get("user")
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
