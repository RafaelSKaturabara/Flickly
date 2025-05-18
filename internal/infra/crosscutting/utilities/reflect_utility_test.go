package utilities

import (
	"testing"
)

// Estruturas para testar a função GetStructName
type TestStruct struct{}
type AnotherTestStruct struct{}

func TestGetStructName(t *testing.T) {
	// Testando com uma instância direta da estrutura
	directInstance := TestStruct{}
	name := GetStructName(directInstance)
	if name != "TestStruct" {
		t.Fatalf("Esperado 'TestStruct', obtido '%s'", name)
	}

	// Testando com um ponteiro para a estrutura
	pointerInstance := &TestStruct{}
	name = GetStructName(pointerInstance)
	if name != "TestStruct" {
		t.Fatalf("Esperado 'TestStruct', obtido '%s'", name)
	}

	// Testando com outro tipo de estrutura
	anotherInstance := AnotherTestStruct{}
	name = GetStructName(anotherInstance)
	if name != "AnotherTestStruct" {
		t.Fatalf("Esperado 'AnotherTestStruct', obtido '%s'", name)
	}
} 