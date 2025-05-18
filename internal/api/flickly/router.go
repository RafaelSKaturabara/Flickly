package flickly

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// Configura e inicia o roteador Gin
func Startup(router *gin.Engine) {
	// Configurando rota de saúde/verificação
	router.GET("/health", getHealth)
	
	// Configurando grupo de rotas para a API do flickly
	flicklyGroup := router.Group("/api/flickly")
	{
		flicklyGroup.GET("/version", getVersion)
	}
}

// Handler para endpoint de saúde
func getHealth(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
		"service": "flickly",
	})
}

// Handler para endpoint de versão
func getVersion(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"version": "1.0.0",
		"api": "flickly",
	})
}
