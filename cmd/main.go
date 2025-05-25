// @title Flickly API
// @version 1.0
// @description API do projeto Flickly
// @license.name MIT
// @host localhost:8080
// @BasePath /
// @schemes http https
package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/rkaturabara/flickly/docs"
	"github.com/rkaturabara/flickly/internal/application/flickly"
	"github.com/rkaturabara/flickly/internal/application/users"
	"github.com/rkaturabara/flickly/internal/infra/crosscutting/ioc"
	swaggerConfig "github.com/rkaturabara/flickly/internal/infra/crosscutting/swagger"
	"github.com/rkaturabara/flickly/internal/infra/crosscutting/utilities"
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
