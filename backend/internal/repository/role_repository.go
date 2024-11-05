package repository

import (
	"chainwave/backend/internal/models"
	"database/sql"
)

// AddCustomer adds a new customer to the database
func AddCustomer(db *sql.DB, customer models.Customer) error {
	_, err := db.Exec(`INSERT INTO customers (id, customer_name, contact_info, location, user_id) VALUES (uuid_generate_v4(), $1, $2, $3, $4)`,
		customer.CustomerName, customer.ContactInfo, customer.Location, customer.UserId)
	return err
}

// EditCustomer updates an existing customer in the database
func EditCustomer(db *sql.DB, customer models.Customer) error {
	_, err := db.Exec(`UPDATE customers SET customer_name = $1, contact_info = $2, location = $3 WHERE id = $4`,
		customer.CustomerName, customer.ContactInfo, customer.Location, customer.Id)
	return err
}

// AddBusinessAdmin adds a new business admin to the database
func AddBusinessAdmin(db *sql.DB, businessAdmin models.BusinessAdmin) error {
	_, err := db.Exec(`INSERT INTO business_admins (id, company_name, contact_info, location, user_id) VALUES (uuid_generate_v4(), $1, $2, $3, $4)`,
		businessAdmin.CompanyName, businessAdmin.ContactInfo, businessAdmin.Location, businessAdmin.UserId)
	return err
}

// EditBusinessAdmin updates an existing business admin in the database
func EditBusinessAdmin(db *sql.DB, businessAdmin models.BusinessAdmin) error {
	_, err := db.Exec(`UPDATE business_admins SET company_name = $1, contact_info = $2, location = $3 WHERE id = $4`,
		businessAdmin.CompanyName, businessAdmin.ContactInfo, businessAdmin.Location, businessAdmin.Id)
	return err
}

// AddTransporter adds a new transporter to the database
func AddTransporter(db *sql.DB, transporter models.Transporter) error {
	_, err := db.Exec(`INSERT INTO transporters (id, driver_name, vehicle_details, contact_info, location, user_id) VALUES (uuid_generate_v4(), $1, $2, $3, $4, $5)`,
		transporter.DriverName, transporter.VehicleDetails, transporter.ContactInfo, transporter.Location, transporter.UserId)
	return err
}

// EditTransporter updates an existing transporter in the database
func EditTransporter(db *sql.DB, transporter models.Transporter) error {
	_, err := db.Exec(`UPDATE transporters SET driver_name = $1, vehicle_details = $2, contact_info = $3, location = $4 WHERE id = $5`,
		transporter.DriverName, transporter.VehicleDetails, transporter.ContactInfo, transporter.Location, transporter.Id)
	return err
}

// AddSupplier adds a new supplier to the database
func AddSupplier(db *sql.DB, supplier models.Supplier) error {
	_, err := db.Exec(`INSERT INTO suppliers (id, supplier_name, contact_info, address, location, user_id) VALUES (uuid_generate_v4(), $1, $2, $3, $4, $5)`,
		supplier.SupplierName, supplier.ContactInfo, supplier.Address, supplier.Location, supplier.UserId)
	return err
}

// EditSupplier updates an existing supplier in the database
func EditSupplier(db *sql.DB, supplier models.Supplier) error {
	_, err := db.Exec(`UPDATE suppliers SET supplier_name = $1, contact_info = $2, address = $3, location = $4 WHERE id = $5`,
		supplier.SupplierName, supplier.ContactInfo, supplier.Address, supplier.Location, supplier.Id)
	return err
}
