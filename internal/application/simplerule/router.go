package simplerule

import (
	"github.com/gin-gonic/gin"
	"github.com/rkaturabara/flickly/internal/application/commons/middleware"
	"github.com/rkaturabara/flickly/internal/application/simplerule/handlers"
	"github.com/rkaturabara/flickly/internal/domain/users/entities"
	"github.com/rkaturabara/flickly/internal/infra/crosscutting/utilities"
)

// Configura e inicia o roteador Gin
func Startup(router *gin.Engine, serviceCollection utilities.IServiceCollection) {
	simpleRuleHandler := handlers.NewSimpleRuleHandler(serviceCollection)

	// Cria o middleware JWT
	jwtMiddleware := middleware.NewJWTMiddleware("config.JWTSecret")

	// Grupo de rotas de autenticação
	authGroup := router.Group("/simplerule")
	{
		// Rotas públicas
		authGroup.POST("/expense", jwtMiddleware.Auth(), simpleRuleHandler.PostExpense)

		// Rotas protegidas
		authGroup.GET("/me", jwtMiddleware.Auth(), func(c *gin.Context) {
			user, _ := c.Request.Context().Value(middleware.UserContextKey).(*entities.User)
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
