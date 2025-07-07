package mockedapi

import (
	"github.com/gin-gonic/gin"
	"github.com/rkaturabara/flickly/internal/application/commons/middleware"
	"github.com/rkaturabara/flickly/internal/application/mockedapi/handlers"
	"github.com/rkaturabara/flickly/internal/infra/crosscutting/utilities"
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
