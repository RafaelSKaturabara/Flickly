package integration_tests

import (
	"os"
	"testing"
)

// TestMain é o ponto de entrada para todos os testes de integração
func TestMain(m *testing.M) {
	// Configuração adicional antes dos testes, como:
	// - Configurar bancos de dados de teste
	// - Inicializar qualquer estado necessário
	// - Configurar ambiente
	
	setup()
	
	// Executar todos os testes
	code := m.Run()
	
	// Limpar após os testes
	teardown()
	
	// Definir o código de saída do processo
	os.Exit(code)
}

// setup configura o ambiente para os testes de integração
func setup() {
	// Configurar para ambiente de teste
	// Por exemplo, mudar para usar um banco de dados em memória
	// ou configurar mocks para serviços externos
}

// teardown limpa após a execução dos testes
func teardown() {
	// Limpar recursos, conexões de banco de dados, etc.
} 