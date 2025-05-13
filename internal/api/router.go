package api

import (
	"flickly/internal/api/flickly/controllers"
	"github.com/gin-gonic/gin"
)

// Configura e inicia o roteador Gin
func Startup() {
	router := gin.Default()

	// Configurando rotas
	router.GET("/entidade-exemplo/{id}", controllers.GetPessoaExemplo)
	router.POST("/entidade-exemplo", controllers.CreatePessoaExemplo)

	// Inicia o servidor na porta 8080
	router.Run(":8080")
}
