package view_models

import (
	"encoding/json"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestCreatePessoaResponse_JSON(t *testing.T) {
	// Configuração
	id := uuid.New()
	now := time.Now()
	response := CreatePessoaResponse{
		ID:        id,
		CreatedAt: now,
		Nome:      "Test Person",
		Idade:     30,
	}
	
	// Execução
	jsonData, err := json.Marshal(response)
	assert.NoError(t, err, "A serialização para JSON não deve gerar erro")
	
	// Verificações
	var parsedResponse CreatePessoaResponse
	err = json.Unmarshal(jsonData, &parsedResponse)
	assert.NoError(t, err, "A deserialização do JSON não deve gerar erro")
	
	assert.Equal(t, response.ID, parsedResponse.ID, "O ID deve ser serializado corretamente")
	assert.Equal(t, response.Nome, parsedResponse.Nome, "O Nome deve ser serializado corretamente")
	assert.Equal(t, response.Idade, parsedResponse.Idade, "A Idade deve ser serializada corretamente")
	
	// Verificar campos JSON
	jsonString := string(jsonData)
	assert.Contains(t, jsonString, `"id":"` + id.String() + `"`, "O campo id deve estar presente no JSON")
	assert.Contains(t, jsonString, `"nome":"Test Person"`, "O campo nome deve estar presente no JSON")
	assert.Contains(t, jsonString, `"idade":30`, "O campo idade deve estar presente no JSON")
}

func TestCreateUserRequest_JSON(t *testing.T) {
	// Configuração
	request := CreateUserRequest{
		Name:  "Test User",
		Email: "test@example.com",
	}
	
	// Execução
	jsonData, err := json.Marshal(request)
	assert.NoError(t, err, "A serialização para JSON não deve gerar erro")
	
	// Verificações
	var parsedRequest CreateUserRequest
	err = json.Unmarshal(jsonData, &parsedRequest)
	assert.NoError(t, err, "A deserialização do JSON não deve gerar erro")
	
	assert.Equal(t, request.Name, parsedRequest.Name, "O Name deve ser serializado corretamente")
	assert.Equal(t, request.Email, parsedRequest.Email, "O Email deve ser serializado corretamente")
	
	// Verificar campos JSON
	jsonString := string(jsonData)
	assert.Contains(t, jsonString, `"name":"Test User"`, "O campo name deve estar presente no JSON")
	assert.Contains(t, jsonString, `"email":"test@example.com"`, "O campo email deve estar presente no JSON")
} 