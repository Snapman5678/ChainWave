package handlers

import (
	"chainwave/backend/internal/models"
	"chainwave/backend/internal/repository"
	"database/sql"
	"net/http"
	"github.com/gin-gonic/gin"
)


// AddCustomerHandler handles adding a new customer
func AddCustomerHandler(db *sql.DB, c *gin.Context) {
	var customer models.Customer
	if err := c.ShouldBindJSON(&customer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := repository.AddCustomer(db, customer); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, customer)
}

// EditCustomerHandler handles editing an existing customer
func EditCustomerHandler(db *sql.DB, c *gin.Context) {
	var customer models.Customer
	if err := c.ShouldBindJSON(&customer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := repository.EditCustomer(db, customer); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, customer)
}

// AddBusinessAdminHandler handles adding a new business admin
func AddBusinessAdminHandler(db *sql.DB, c *gin.Context) {
	var businessAdmin models.BusinessAdmin
	if err := c.ShouldBindJSON(&businessAdmin); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := repository.AddBusinessAdmin(db, businessAdmin); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, businessAdmin)
}

// EditBusinessAdminHandler handles editing an existing business admin
func EditBusinessAdminHandler(db *sql.DB, c *gin.Context) {
	var businessAdmin models.BusinessAdmin
	if err := c.ShouldBindJSON(&businessAdmin); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := repository.EditBusinessAdmin(db, businessAdmin); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, businessAdmin)
}

// AddTransporterHandler handles adding a new transporter
func AddTransporterHandler(db *sql.DB, c *gin.Context) {
	var transporter models.Transporter
	if err := c.ShouldBindJSON(&transporter); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := repository.AddTransporter(db, transporter); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, transporter)
}

// EditTransporterHandler handles editing an existing transporter
func EditTransporterHandler(db *sql.DB, c *gin.Context) {
	var transporter models.Transporter
	if err := c.ShouldBindJSON(&transporter); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := repository.EditTransporter(db, transporter); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, transporter)
}

// AddSupplierHandler handles adding a new supplier
func AddSupplierHandler(db *sql.DB, c *gin.Context) {
	var supplier models.Supplier
	if err := c.ShouldBindJSON(&supplier); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := repository.AddSupplier(db, supplier); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, supplier)
}

// EditSupplierHandler handles editing an existing supplier
func EditSupplierHandler(db *sql.DB, c *gin.Context) {
	var supplier models.Supplier
	if err := c.ShouldBindJSON(&supplier); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := repository.EditSupplier(db, supplier); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, supplier)
}
