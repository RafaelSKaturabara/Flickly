package flickly

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestStartup(t *testing.T) {
	// Configuração
	gin.SetMode(gin.TestMode)
	router := gin.New()

	// Execução
	Startup(router)

	// Verificações para endpoint /health
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/health", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code, "O código de status para /health deve ser 200 OK")

	var healthResponse map[string]string
	err := json.Unmarshal(w.Body.Bytes(), &healthResponse)
	assert.NoError(t, err, "Não deve ocorrer erro ao desserializar a resposta JSON")
	assert.Equal(t, "ok", healthResponse["status"], "O status deve ser 'ok'")
	assert.Equal(t, "flickly", healthResponse["service"], "O serviço deve ser 'flickly'")

	// Verificações para endpoint /api/flickly/version
	w = httptest.NewRecorder()
	req, _ = http.NewRequest(http.MethodGet, "/api/flickly/version", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code, "O código de status para /api/flickly/version deve ser 200 OK")

	var versionResponse map[string]string
	err = json.Unmarshal(w.Body.Bytes(), &versionResponse)
	assert.NoError(t, err, "Não deve ocorrer erro ao desserializar a resposta JSON")
	assert.Equal(t, "1.0.0", versionResponse["version"], "A versão deve ser '1.0.0'")
	assert.Equal(t, "flickly", versionResponse["api"], "A API deve ser 'flickly'")
}

func TestGetHealth(t *testing.T) {
	// Configuração
	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	// Execução
	getHealth(c)

	// Verificações
	assert.Equal(t, http.StatusOK, w.Code, "O código de status deve ser 200 OK")

	var response map[string]string
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err, "Não deve ocorrer erro ao desserializar a resposta JSON")
	assert.Equal(t, "ok", response["status"], "O status deve ser 'ok'")
	assert.Equal(t, "flickly", response["service"], "O serviço deve ser 'flickly'")
}

func TestGetVersion(t *testing.T) {
	// Configuração
	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	// Execução
	getVersion(c)

	// Verificações
	assert.Equal(t, http.StatusOK, w.Code, "O código de status deve ser 200 OK")

	var response map[string]string
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err, "Não deve ocorrer erro ao desserializar a resposta JSON")
	assert.Equal(t, "1.0.0", response["version"], "A versão deve ser '1.0.0'")
	assert.Equal(t, "flickly", response["api"], "A API deve ser 'flickly'")
}
