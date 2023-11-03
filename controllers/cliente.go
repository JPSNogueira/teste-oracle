package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"opus127.com.br/models"
)

// @Summary Busca todos os clientes
// @ID get-all-clientes
// @Produce json
// @Success 200 {object} models.Cliente
// @Router /clientes [get]
func FindAllClientes(c *gin.Context) {
	var customers []models.Cliente
	models.DB.Find(&customers)

	c.JSON(http.StatusOK, gin.H{"data": customers})
}

// @Summary Encontra um cliente
// @ID find-customer
// @Produce json
// @Param id path string true "Cliente ID"
// @Success 200 {object} models.Cliente
// @Router /clientes/{id} [get]
func FindCliente(c *gin.Context) { // Get model if exist
	var customer models.Cliente

	if err := models.DB.Where("idCliente = ?", c.Param("id")).First(&customer).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": customer})
}


