package main

import (
	"flickly/internal/api/flickly"
	"flickly/internal/api/users"
	inversionofcontrol "flickly/internal/infra/cross-cutting/inversion-of-control"
	"flickly/internal/infra/cross-cutting/utilities"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	serviceCollection := utilities.NewServiceCollection()

	inversionofcontrol.InjectServices(serviceCollection)
	inversionofcontrol.InjectMediatorHandlers(serviceCollection)

	users.Startup(router, serviceCollection)
	flickly.Startup(router)

	// Inicia o servidor na porta 8080
	err := router.Run(":8080")
	if err != nil {
		return
	}
}
