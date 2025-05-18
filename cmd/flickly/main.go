// @title Flickly API
// @version 1.0
// @description API do projeto Flickly
// @license.name MIT
// @host localhost:8080
// @BasePath /
// @schemes http https
package main

import (
	"flickly/docs"
	"flickly/internal/api/flickly"
	"flickly/internal/api/users"
	"flickly/internal/infra/crosscutting/ioc"
	swaggerConfig "flickly/internal/infra/crosscutting/swagger"
	"flickly/internal/infra/crosscutting/utilities"
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Servidor iniciando em http://localhost:8080")

	// Configurações básicas do Swagger (apenas para permitir sua geração)
	docs.SwaggerInfo.Title = "Flickly API"
	docs.SwaggerInfo.BasePath = "/"

	router := gin.Default()
	serviceCollection := utilities.NewServiceCollection()

	ioc.InitAutomapper(serviceCollection)
	ioc.InjectServices(serviceCollection)
	ioc.InjectMediatorHandlers(serviceCollection)

	users.Startup(router, serviceCollection)
	flickly.Startup(router)

	// Configuração do Swagger usando o novo pacote
	swaggerConfig.SetupSwagger(router)

	// Inicia o servidor na porta 8080
	err := router.Run(":8080")
	if err != nil {
		return
	}
}
