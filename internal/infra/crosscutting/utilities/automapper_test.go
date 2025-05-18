package utilities

import (
	"reflect"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

// Structs de exemplo para teste
type SourceUser struct {
	ID        int64
	Name      string
	Email     string
	CreatedAt time.Time
	Age       int
	Address   string
}

type DestinationUser struct {
	ID        int64
	Name      string
	Email     string
	CreatedAt time.Time
	Age       int
	Address   string
}

type DifferentUser struct {
	ID        int64
	FullName  string // Nome diferente
	Email     string
	CreatedAt time.Time
	Age       int
	Address   string
}

func TestAutoMapper_Map(t *testing.T) {
	// Criar instância do AutoMapper
	mapper := NewAutoMapper()

	// Criar dados de origem
	source := SourceUser{
		ID:        1,
		Name:      "John Doe",
		Email:     "john@example.com",
		CreatedAt: time.Now(),
		Age:       30,
		Address:   "123 Main St",
	}

	// Teste 1: Mapeamento básico
	t.Run("Mapeamento básico", func(t *testing.T) {
		var dest DestinationUser
		err := mapper.Map(source, &dest)

		assert.NoError(t, err)
		assert.Equal(t, source.ID, dest.ID)
		assert.Equal(t, source.Name, dest.Name)
		assert.Equal(t, source.Email, dest.Email)
		assert.Equal(t, source.CreatedAt, dest.CreatedAt)
		assert.Equal(t, source.Age, dest.Age)
		assert.Equal(t, source.Address, dest.Address)
	})

	// Teste 2: Mapeamento com campos diferentes
	t.Run("Mapeamento com campos diferentes", func(t *testing.T) {
		var dest DifferentUser
		err := mapper.Map(source, &dest)

		assert.NoError(t, err)
		assert.Equal(t, source.ID, dest.ID)
		assert.Equal(t, source.Email, dest.Email)
		assert.Equal(t, source.CreatedAt, dest.CreatedAt)
		assert.Equal(t, source.Age, dest.Age)
		assert.Equal(t, source.Address, dest.Address)
		// FullName não é mapeado porque o nome do campo é diferente
		assert.Empty(t, dest.FullName)
	})

	// Teste 3: Mapeamento personalizado
	t.Run("Mapeamento personalizado", func(t *testing.T) {
		// Adicionar mapeamento personalizado
		mapper.AddMapping(
			reflect.TypeOf(SourceUser{}),
			reflect.TypeOf(DifferentUser{}),
			func(source, dest reflect.Value) error {
				// Mapear campos comuns
				dest.FieldByName("ID").Set(source.FieldByName("ID"))
				dest.FieldByName("Email").Set(source.FieldByName("Email"))
				dest.FieldByName("CreatedAt").Set(source.FieldByName("CreatedAt"))
				dest.FieldByName("Age").Set(source.FieldByName("Age"))
				dest.FieldByName("Address").Set(source.FieldByName("Address"))
				// Mapear Name para FullName
				dest.FieldByName("FullName").Set(source.FieldByName("Name"))
				return nil
			},
		)

		var dest DifferentUser
		err := mapper.Map(source, &dest)

		assert.NoError(t, err)
		assert.Equal(t, source.ID, dest.ID)
		assert.Equal(t, source.Name, dest.FullName) // Agora o nome é mapeado corretamente
		assert.Equal(t, source.Email, dest.Email)
		assert.Equal(t, source.CreatedAt, dest.CreatedAt)
		assert.Equal(t, source.Age, dest.Age)
		assert.Equal(t, source.Address, dest.Address)
	})
}

func TestAutoMapper_MapSlice(t *testing.T) {
	mapper := NewAutoMapper()

	// Criar slice de origem
	source := []SourceUser{
		{
			ID:        1,
			Name:      "John Doe",
			Email:     "john@example.com",
			CreatedAt: time.Now(),
			Age:       30,
			Address:   "123 Main St",
		},
		{
			ID:        2,
			Name:      "Jane Doe",
			Email:     "jane@example.com",
			CreatedAt: time.Now(),
			Age:       28,
			Address:   "456 Oak St",
		},
	}

	// Teste de mapeamento de slice
	t.Run("Mapeamento de slice", func(t *testing.T) {
		var dest []DestinationUser
		err := mapper.MapSlice(source, &dest)

		assert.NoError(t, err)
		assert.Len(t, dest, 2)

		// Verificar primeiro item
		assert.Equal(t, source[0].ID, dest[0].ID)
		assert.Equal(t, source[0].Name, dest[0].Name)
		assert.Equal(t, source[0].Email, dest[0].Email)
		assert.Equal(t, source[0].CreatedAt, dest[0].CreatedAt)
		assert.Equal(t, source[0].Age, dest[0].Age)
		assert.Equal(t, source[0].Address, dest[0].Address)

		// Verificar segundo item
		assert.Equal(t, source[1].ID, dest[1].ID)
		assert.Equal(t, source[1].Name, dest[1].Name)
		assert.Equal(t, source[1].Email, dest[1].Email)
		assert.Equal(t, source[1].CreatedAt, dest[1].CreatedAt)
		assert.Equal(t, source[1].Age, dest[1].Age)
		assert.Equal(t, source[1].Address, dest[1].Address)
	})
}
