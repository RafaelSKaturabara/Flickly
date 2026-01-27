package mockedapi

import (
	"github.com/RafaelSKaturabara/Flickly/internal/infra/crosscutting/utilities"
	"github.com/RafaelSKaturabara/Flickly/internal/presentation/commons/middleware"
	"github.com/RafaelSKaturabara/Flickly/internal/presentation/mockedapi/handlers"
	"github.com/gin-gonic/gin"
)

// Configura e inicia o roteador Gin
func Startup(router *gin.Engine, serviceCollection utilities.IServiceCollection) {
	// Cria o middleware JWT
	jwtMiddleware := middleware.NewJWTMiddleware("config.JWTSecret")
	mockHandler := handlers.NewMockHandler(serviceCollection)

	router.PUT("/mock",
		jwtMiddleware.Auth(),
		jwtMiddleware.Role("user"),
		mockHandler.PutRequestMock)
}
