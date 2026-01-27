package simplerule

import (
	"github.com/RafaelSKaturabara/Flickly/internal/domain/users/entities"
	"github.com/RafaelSKaturabara/Flickly/internal/infra/crosscutting/utilities"
	"github.com/RafaelSKaturabara/Flickly/internal/presentation/commons/middleware"
	"github.com/RafaelSKaturabara/Flickly/internal/presentation/simplerule/handlers"
	"github.com/gin-gonic/gin"
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
		authGroup.GET("/subcategories", jwtMiddleware.Auth(), simpleRuleHandler.ListAllSubcategories)

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
