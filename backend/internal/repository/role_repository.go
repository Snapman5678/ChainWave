package repository

import (
	"chainwave/backend/internal/models"
	"database/sql"
)

// AddCustomer adds a new customer to the database
func AddCustomer(db *sql.DB, customer models.Customer) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}

	_, err = tx.Exec(`INSERT INTO customers (id, customer_name, contact_info, location, user_id) VALUES (uuid_generate_v4(), $1, $2, $3, $4)`,
		customer.CustomerName, customer.ContactInfo, customer.Location, customer.UserId)
	if err != nil {
		tx.Rollback()
		return err
	}

	_, err = tx.Exec(`INSERT INTO user_roles (user_id, customer_id) VALUES ($1, (SELECT id FROM customers WHERE user_id = $1))`,
		customer.UserId)
	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit()
}

// EditCustomer updates an existing customer in the database
func EditCustomer(db *sql.DB, customer models.Customer) error {
	_, err := db.Exec(`UPDATE customers SET customer_name = $1, contact_info = $2, location = $3 WHERE id = $4`,
		customer.CustomerName, customer.ContactInfo, customer.Location, customer.Id)
	return err
}

// AddBusinessAdmin adds a new business admin to the database
func AddBusinessAdmin(db *sql.DB, businessAdmin models.BusinessAdmin) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}

	_, err = tx.Exec(`INSERT INTO business_admins (id, company_name, contact_info, location, user_id) VALUES (uuid_generate_v4(), $1, $2, $3, $4)`,
		businessAdmin.CompanyName, businessAdmin.ContactInfo, businessAdmin.Location, businessAdmin.UserId)
	if err != nil {
		tx.Rollback()
		return err
	}

	_, err = tx.Exec(`INSERT INTO user_roles (user_id, business_admin_id) VALUES ($1, (SELECT id FROM business_admins WHERE user_id = $1))`,
		businessAdmin.UserId)
	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit()
}

// EditBusinessAdmin updates an existing business admin in the database
func EditBusinessAdmin(db *sql.DB, businessAdmin models.BusinessAdmin) error {
	_, err := db.Exec(`UPDATE business_admins SET company_name = $1, contact_info = $2, location = $3 WHERE id = $4`,
		businessAdmin.CompanyName, businessAdmin.ContactInfo, businessAdmin.Location, businessAdmin.Id)
	return err
}

// AddTransporter adds a new transporter to the database
func AddTransporter(db *sql.DB, transporter models.Transporter) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}

	_, err = tx.Exec(`INSERT INTO transporters (id, driver_name, vehicle_details, contact_info, location, user_id) VALUES (uuid_generate_v4(), $1, $2, $3, $4, $5)`,
		transporter.DriverName, transporter.VehicleDetails, transporter.ContactInfo, transporter.Location, transporter.UserId)
	if err != nil {
		tx.Rollback()
		return err
	}

	_, err = tx.Exec(`INSERT INTO user_roles (user_id, transporter_id) VALUES ($1, (SELECT id FROM transporters WHERE user_id = $1))`,
		transporter.UserId)
	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit()
}

// EditTransporter updates an existing transporter in the database
func EditTransporter(db *sql.DB, transporter models.Transporter) error {
	_, err := db.Exec(`UPDATE transporters SET driver_name = $1, vehicle_details = $2, contact_info = $3, location = $4 WHERE id = $5`,
		transporter.DriverName, transporter.VehicleDetails, transporter.ContactInfo, transporter.Location, transporter.Id)
	return err
}

// AddSupplier adds a new supplier to the database
func AddSupplier(db *sql.DB, supplier models.Supplier) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}

	_, err = tx.Exec(`INSERT INTO suppliers (id, supplier_name, contact_info, address, location, user_id) VALUES (uuid_generate_v4(), $1, $2, $3, $4, $5)`,
		supplier.SupplierName, supplier.ContactInfo, supplier.Address, supplier.Location, supplier.UserId)
	if err != nil {
		tx.Rollback()
		return err
	}

	_, err = tx.Exec(`INSERT INTO user_roles (user_id, supplier_id) VALUES ($1, (SELECT id FROM suppliers WHERE user_id = $1))`,
		supplier.UserId)
	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit()
}

// EditSupplier updates an existing supplier in the database
func EditSupplier(db *sql.DB, supplier models.Supplier) error {
	_, err := db.Exec(`UPDATE suppliers SET supplier_name = $1, contact_info = $2, address = $3, location = $4 WHERE id = $5`,
		supplier.SupplierName, supplier.ContactInfo, supplier.Address, supplier.Location, supplier.Id)
	return err
}


