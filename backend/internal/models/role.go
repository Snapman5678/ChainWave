package models

import "github.com/google/uuid"

// Customer struct
type Customer struct {
	Id          uuid.UUID `json:"id"`
	CustomerName string   `json:"customer_name"`
	ContactInfo string   `json:"contact_info"`
	Location    string   `json:"location"`
	UserId      uuid.UUID `json:"user_id"`
}

// BusinessAdmin struct
type BusinessAdmin struct {
	Id             uuid.UUID `json:"id"`
	CompanyName   string    `json:"company_name"`
	ContactInfo   string    `json:"contact_info"`
	Location      string    `json:"location"`
	UserId        uuid.UUID `json:"user_id"`
}

// Transporter struct
type Transporter struct {
	Id           uuid.UUID `json:"id"`
	DriverName   string    `json:"driver_name"`
	VehicleDetails string   `json:"vehicle_details"`
	ContactInfo  string    `json:"contact_info"`
	Location     string    `json:"location"`
	UserId       uuid.UUID `json:"user_id"`
}

// Supplier struct
type Supplier struct {
	Id           uuid.UUID `json:"id"`
	SupplierName string    `json:"supplier_name"`
	ContactInfo  string    `json:"contact_info"`
	Address      string    `json:"address"`
	Location     string    `json:"location"`
	UserId       uuid.UUID `json:"user_id"`
}