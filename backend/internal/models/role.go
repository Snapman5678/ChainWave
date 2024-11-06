package models

import "github.com/google/uuid"

// Customer struct
type Customer struct {
	Id           uuid.UUID `json:"id"`
	CustomerName string    `json:"customer_name"`
	ContactInfo  string    `json:"contact_info"`
	LocationId   uuid.UUID `json:"location_id"`
	UserId       uuid.UUID `json:"user_id"`
}

// BusinessAdmin struct
type BusinessAdmin struct {
	Id           uuid.UUID `json:"id"`
	CompanyName  string    `json:"company_name"`
	ContactInfo  string    `json:"contact_info"`
	LocationId   uuid.UUID `json:"location_id"`
	UserId       uuid.UUID `json:"user_id"`
}

// Transporter struct
type Transporter struct {
	Id           uuid.UUID `json:"id"`
	DriverName   string    `json:"driver_name"`
	VehicleId    uuid.UUID `json:"vehicle_id"`
	ContactInfo  string    `json:"contact_info"`
	LocationId   uuid.UUID `json:"location_id"`
	UserId       uuid.UUID `json:"user_id"`
}

// Supplier struct
type Supplier struct {
	Id           uuid.UUID `json:"id"`
	SupplierName string    `json:"supplier_name"`
	ContactInfo  string    `json:"contact_info"`
	Address      string    `json:"address"`
	LocationId   uuid.UUID `json:"location_id"`
	UserId       uuid.UUID `json:"user_id"`
}

// Role struct
type Role struct {
	UserId          uuid.UUID  `json:"user_id"`
	CustomerId      *uuid.UUID `json:"customer_id,omitempty"`
	BusinessAdminId *uuid.UUID `json:"business_admin_id,omitempty"`
	TransporterId   *uuid.UUID `json:"transporter_id,omitempty"`
	SupplierId      *uuid.UUID `json:"supplier_id,omitempty"`
}