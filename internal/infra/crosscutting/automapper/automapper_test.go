package automapper

import (
	"testing"
	"time"
)

// Source struct para teste
type Source struct {
	ID        int
	Email     string
	CreatedAt time.Time
	Age       int
	Address   string
	Name      string
	Message   string
}

func (s Source) Error() string {
	return s.Message
}

// Destination struct para teste
type Destination struct {
	ID              int
	Email           string
	CreatedAt       time.Time
	Age             int
	Address         string
	FullName        string
	InternalMessage string
}

func TestMap_WithErrorMethod(t *testing.T) {
	// Arrange
	source := Source{
		ID:        1,
		Email:     "test@test.com",
		CreatedAt: time.Now(),
		Age:       25,
		Address:   "Test Address",
		Name:      "Test Name",
		Message:   "Test Error Message",
	}
	dest := &Destination{}

	// Act
	err := Map(source, dest)

	// Assert
	if err != nil {
		t.Errorf("Map() error = %v", err)
	}

	if dest.ID != source.ID {
		t.Errorf("ID = %v, want %v", dest.ID, source.ID)
	}

	if dest.Email != source.Email {
		t.Errorf("Email = %v, want %v", dest.Email, source.Email)
	}

	if dest.Age != source.Age {
		t.Errorf("Age = %v, want %v", dest.Age, source.Age)
	}

	if dest.Address != source.Address {
		t.Errorf("Address = %v, want %v", dest.Address, source.Address)
	}

	if dest.FullName != source.Name {
		t.Errorf("FullName = %v, want %v", dest.FullName, source.Name)
	}

	if dest.InternalMessage != source.Message {
		t.Errorf("InternalMessage = %v, want %v", dest.InternalMessage, source.Message)
	}
}

func TestMap_WithoutErrorMethod(t *testing.T) {
	// Arrange
	type SimpleSource struct {
		ID    int
		Email string
	}
	type SimpleDest struct {
		ID              int
		Email           string
		InternalMessage string
	}

	source := SimpleSource{
		ID:    1,
		Email: "test@test.com",
	}
	dest := &SimpleDest{}

	// Act
	err := Map(source, dest)

	// Assert
	if err != nil {
		t.Errorf("Map() error = %v", err)
	}

	if dest.ID != source.ID {
		t.Errorf("ID = %v, want %v", dest.ID, source.ID)
	}

	if dest.Email != source.Email {
		t.Errorf("Email = %v, want %v", dest.Email, source.Email)
	}

	if dest.InternalMessage != "" {
		t.Errorf("InternalMessage = %v, want empty string", dest.InternalMessage)
	}
}
