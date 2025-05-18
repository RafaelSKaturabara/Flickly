package utilities

import (
	"reflect"
)

func GetStructName(i interface{}) string {
	// Usa o pacote reflect para obter o tipo da variável
	t := reflect.TypeOf(i)

	// Verifica se o tipo é um ponteiro, em caso de ponteiro, precisamos dereferenciá-lo
	if t.Kind() == reflect.Ptr {
		t = t.Elem() // Obter o tipo subjacente do ponteiro
	}

	return t.Name()
}
