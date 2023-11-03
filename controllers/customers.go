package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"opus127.com.br/models"
)

type CreateCustomerInput struct {
	Name    string `json:"name" binding:"required"`
	Address string `json:"address" binding:"required"`
}

type UpdateCustomerInput struct {
	Name    string `json:"name"`
	Address string `json:"address"`
}

// @Summary Get all the customers
// @ID get-all-customers
// @Produce json
// @Success 200 {object} models.Customer
// @Router /customers [get]
func FindAllCustomers(c *gin.Context) {
	var customers []models.Customer
	models.DB.Find(&customers)

	c.JSON(http.StatusOK, gin.H{"data": customers})
}

// @Summary Create a new customer
// @ID create-customer
// @Produce json
// @Param data body models.Customer true "customer data"
// @Success 200 {object} models.Customer
// @Router /customers [post]
func CreateCustomer(c *gin.Context) {
	// Validate input
	var input CreateCustomerInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create Customer
	customer := models.Customer{Name: input.Name, Address: input.Address}
	models.DB.Create(&customer)

	c.JSON(http.StatusOK, gin.H{"data": customer})
}

// @Summary Find a customer
// @ID find-customer
// @Produce json
// @Param id path string true "Customer ID"
// @Success 200 {object} models.Customer
// @Router /customers/{id} [get]
func FindCustomer(c *gin.Context) { // Get model if exist
	var customer models.Customer

	if err := models.DB.Where("id = ?", c.Param("id")).First(&customer).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": customer})
}

// @Summary Update a customer
// @ID update-customer
// @Produce json
// @Param id path string true "Customer ID"
// @Param data body models.Customer true "customer data"
// @Success 200 {object} models.Customer
// @Router /customers/{id} [patch]
func UpdateCustomer(c *gin.Context) {
	// Get model if exist
	var customer models.Customer
	if err := models.DB.Where("id = ?", c.Param("id")).First(&customer).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	// Validate input
	var input UpdateCustomerInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	models.DB.Model(&customer).Updates(input)

	c.JSON(http.StatusOK, gin.H{"data": customer})
}

// @Summary Delete a customer
// @ID delete-customer
// @Produce json
// @Param id path string true "Customer ID"
// @Success 200 {object} models.Customer
// @Router /customers/{id} [delete]
func DeleteCustomer(c *gin.Context) {
	// Get model if exist
	var customer models.Customer
	if err := models.DB.Where("id = ?", c.Param("id")).First(&customer).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	models.DB.Delete(&customer)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
