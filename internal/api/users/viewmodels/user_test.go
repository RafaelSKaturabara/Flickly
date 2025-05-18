package view_models

import (
	"encoding/json"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestCreateUserResponse_JSON(t *testing.T) {
	// Configuração
	id := uuid.New()
	now := time.Now()
	response := CreateUserResponse{
		ID:        id,
		CreatedAt: now,
		Name:      "Test Person",
		Email:     "test@example.com",
	}

	// Execução
	jsonData, err := json.Marshal(response)
	assert.NoError(t, err, "A serialização para JSON não deve gerar erro")

	// Verificações
	var parsedResponse CreateUserResponse
	err = json.Unmarshal(jsonData, &parsedResponse)
	assert.NoError(t, err, "A deserialização do JSON não deve gerar erro")

	assert.Equal(t, response.ID, parsedResponse.ID, "O ID deve ser serializado corretamente")
	assert.Equal(t, response.Name, parsedResponse.Name, "O Name deve ser serializado corretamente")
	assert.Equal(t, response.Email, parsedResponse.Email, "O Email deve ser serializado corretamente")

	// Verificar campos JSON
	jsonString := string(jsonData)
	assert.Contains(t, jsonString, `"id":"`+id.String()+`"`, "O campo id deve estar presente no JSON")
	assert.Contains(t, jsonString, `"name":"Test Person"`, "O campo name deve estar presente no JSON")
	assert.Contains(t, jsonString, `"email":"test@example.com"`, "O campo email deve estar presente no JSON")
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
