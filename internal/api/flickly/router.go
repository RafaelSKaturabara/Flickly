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

// getHealth verifica o estado do servidor
// @Summary Verificar saúde do servidor
// @Description Retorna o status de saúde do servidor
// @Tags system
// @Produce json
// @Success 200 {object} interface{}
// @Router /health [get]
func getHealth(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":  "ok",
		"service": "flickly",
	})
}

// getVersion retorna a versão da API
// @Summary Verificar versão da API
// @Description Retorna informações sobre a versão da API
// @Tags system
// @Produce json
// @Success 200 {object} interface{}
// @Router /api/flickly/version [get]
func getVersion(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"version": "1.0.0",
		"api":     "flickly",
	})
}
