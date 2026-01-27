// @title Flickly API
// @version 1.0
// @description API do projeto Flickly
// @license.name MIT
// @host localhost:8090
// @BasePath /
// @schemes http https
package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/RafaelSKaturabara/Flickly/docs"
	"github.com/RafaelSKaturabara/Flickly/internal/infra/crosscutting/ioc"
	swaggerConfig "github.com/RafaelSKaturabara/Flickly/internal/infra/crosscutting/swagger"
	"github.com/RafaelSKaturabara/Flickly/internal/infra/crosscutting/utilities"
	"github.com/RafaelSKaturabara/Flickly/internal/presentation/flickly"
	"github.com/RafaelSKaturabara/Flickly/internal/presentation/simplerule"
	"github.com/RafaelSKaturabara/Flickly/internal/presentation/users"
)

func main() {
	fmt.Println("Servidor iniciando em http://localhost:8090")

	// Configurações básicas do Swagger (apenas para permitir sua geração)
	docs.SwaggerInfo.Title = "Flickly API"
	docs.SwaggerInfo.BasePath = "/"

	router := gin.Default()
	serviceCollection := utilities.NewServiceCollection()

	ioc.InitAutomapper(serviceCollection)
	ioc.InjectServices(serviceCollection)
	ioc.InjectMediatorHandlers(serviceCollection)

	users.Startup(router, serviceCollection)
	simplerule.Startup(router, serviceCollection)
	flickly.Startup(router)

	// Configuração do Swagger usando o novo pacote
	swaggerConfig.SetupSwagger(router)

	// Inicia o servidor na porta 8090

	err := router.Run(":8090")
	if err != nil {
		return
	}
}
