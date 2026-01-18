package viewmodel

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTokenResponse_JSON(t *testing.T) {
	// Configuração
	token := TokenResponse{
		AccessToken: "test_token",
		TokenType:   "Bearer",
		ExpiresIn:   3600,
	}

	// Execução
	jsonData, err := json.Marshal(token)
	assert.NoError(t, err, "A serialização para JSON não deve gerar erro")

	// Verificações
	var parsedToken TokenResponse
	err = json.Unmarshal(jsonData, &parsedToken)
	assert.NoError(t, err, "A deserialização do JSON não deve gerar erro")

	assert.Equal(t, token.AccessToken, parsedToken.AccessToken, "O accessToken deve ser serializado corretamente")
	assert.Equal(t, token.TokenType, parsedToken.TokenType, "O tokenType deve ser serializado corretamente")
	assert.Equal(t, token.ExpiresIn, parsedToken.ExpiresIn, "O expiresIn deve ser serializado corretamente")

	// Verificar se os campos estão com os nomes corretos no JSON
	jsonString := string(jsonData)
	assert.Contains(t, jsonString, `"access_token":"test_token"`, "O campo access_token deve estar presente no JSON")
	assert.Contains(t, jsonString, `"token_type":"Bearer"`, "O campo token_type deve estar presente no JSON")
	assert.Contains(t, jsonString, `"expires_in":3600`, "O campo expires_in deve estar presente no JSON")
}
