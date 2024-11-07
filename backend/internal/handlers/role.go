package handlers

import (
	"chainwave/backend/internal/models"
	"chainwave/backend/internal/repository"
	"database/sql"
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// AddCustomerHandler handles adding a new customer
func AddCustomerHandler(db *sql.DB, c *gin.Context) {
	var request struct {
		Customer models.Customer `json:"customer"`
		Location models.Location `json:"location"`
	}

	// Bind the incoming JSON to the request struct
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Retrieve the user ID from the context
	userId, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User ID not found in context"})
		return
	}

	uid, err := uuid.Parse(userId.(string))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	// Add the customer and location
	customerID, locationID, err := repository.AddCustomer(db, uid, request.Customer, request.Location)
	if err != nil {
		if customerID != uuid.Nil && locationID == uuid.Nil {
			c.JSON(http.StatusConflict, gin.H{"error": "Customer already exists"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Update the customer ID, user ID, and location ID in the response
	request.Customer.Id = customerID
	request.Customer.UserId = uid
	request.Customer.LocationId = locationID

	// Respond with the created customer object
	c.JSON(http.StatusCreated, request.Customer)
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
	var request struct {
		BusinessAdmin models.BusinessAdmin `json:"businessAdmin"`
		Location models.Location `json:"location"`
	}
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	userId, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User ID not found in context"})
		return
	}
	uid, err := uuid.Parse(userId.(string))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}
	businessAdminID, locationID, err := repository.AddBusinessAdmin(db, uid, request.BusinessAdmin, request.Location)
	if err != nil {
		if businessAdminID != uuid.Nil && locationID == uuid.Nil {
			c.JSON(http.StatusConflict, gin.H{"error": "Business Admin already exists"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Update the business admin ID, user ID, and location ID in the response
	request.BusinessAdmin.Id = businessAdminID
	request.BusinessAdmin.UserId = uid
	request.BusinessAdmin.LocationId = locationID

	c.JSON(http.StatusCreated, request.BusinessAdmin)
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
	var request struct {
		Transporter models.Transporter `json:"transporter"`
		Location    models.Location    `json:"location"`
		Vehicle     models.Vehicle     `json:"vehicle"`
	}
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	userId, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User ID not found in context"})
		return
	}
	uid, err := uuid.Parse(userId.(string))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}
	transporterID, locationID, vehicleID, err := repository.AddTransporter(db, uid, request.Transporter, request.Location, request.Vehicle)
	if err != nil {
		if transporterID != uuid.Nil && locationID == uuid.Nil && vehicleID == uuid.Nil {
			c.JSON(http.StatusConflict, gin.H{"error": "Transporter already exists"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Update the transporter ID, user ID, location ID, and vehicle ID in the response
	request.Transporter.Id = transporterID
	request.Transporter.UserId = uid
	request.Transporter.LocationId = locationID
	request.Transporter.VehicleId = vehicleID

	c.JSON(http.StatusCreated, request.Transporter)
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
	var request struct {
		Supplier models.Supplier `json:"supplier"`
		Location models.Location `json:"location"`
	}
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	userId, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User ID not found in context"})
		return
	}
	uid, err := uuid.Parse(userId.(string))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}
	supplierID, locationID, err := repository.AddSupplier(db, uid, request.Supplier, request.Location)
	if err != nil {
		if supplierID != uuid.Nil && locationID == uuid.Nil {
			c.JSON(http.StatusConflict, gin.H{"error": "Supplier already exists"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Update the supplier ID, user ID, and location ID in the response
	request.Supplier.Id = supplierID
	request.Supplier.UserId = uid
	request.Supplier.LocationId = locationID

	c.JSON(http.StatusCreated, request.Supplier)
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

// GetRolesHandler handles fetching roles for a given user ID
func GetRolesHandler(db *sql.DB, c *gin.Context) {
	userId, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User ID not found in context"})
		return
	}
	uid, err := uuid.Parse(userId.(string))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}
	roles, err := repository.GetRolesByUserId(db, uid)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, roles)
}
