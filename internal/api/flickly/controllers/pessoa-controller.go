package controllers

import (
	view_models "flickly/internal/api/flickly/view-models"
	"flickly/internal/domain/NomeDominioExemploPacote/commands"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
	"time"
)

// Função para obter todas as entidades de exemplo
func GetPessoaExemplo(c *gin.Context) {
	timeNow := time.Now()
	c.JSON(http.StatusOK, view_models.NewGetPessoaResponse(uuid.New(), time.Now().Add(-100), &timeNow, "Nome Exemplo", 40))
}

// Função para criar uma nova entidade de exemplo
func CreatePessoaExemplo(c *gin.Context) {
	var criarPessoaCommand commands.CriarPessoaCommand
	if err := c.ShouldBindJSON(&criarPessoaCommand); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	timeNow := time.Now()

	c.JSON(http.StatusCreated, view_models.NewGetPessoaResponse(uuid.New(), time.Now().Add(-100), &timeNow, "Nome Exemplo", 40))
}
