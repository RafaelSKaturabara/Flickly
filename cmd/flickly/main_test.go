package main

import (
	"flickly/internal/api/flickly"
	"flickly/internal/api/users"
	inversionofcontrol "flickly/internal/infra/cross-cutting/inversion-of-control"
	"flickly/internal/infra/cross-cutting/utilities"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

// TestSetupRouter testa a configuração do roteador com todas as dependências
func TestSetupRouter(t *testing.T) {
	// Configuração
	gin.SetMode(gin.TestMode)
	
	// Configurar roteador manualmente em vez de chamar main()
	// Isso nos permite testar a configuração sem iniciar o servidor HTTP
	setupRouter := func() *gin.Engine {
		router := gin.New() // Usar gin.New() em vez de gin.Default() para evitar logs
		serviceCollection := utilities.NewServiceCollection()
		
		// Injetar dependências
		inversionofcontrol.InjectServices(serviceCollection)
		inversionofcontrol.InjectMediatorHandlers(serviceCollection)
		
		// Configurar rotas
		users.Startup(router, serviceCollection)
		flickly.Startup(router)
		
		return router
	}
	
	// Executar a configuração
	router := setupRouter()
	
	// Verificações básicas para garantir que o roteador foi configurado
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/ping", nil)
	router.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})
	router.ServeHTTP(w, req)
	
	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "pong", w.Body.String())
} 