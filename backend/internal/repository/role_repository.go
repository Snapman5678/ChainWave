package repository

import (
	"chainwave/backend/internal/models"
	"database/sql"
	"errors"
	"github.com/google/uuid"
)

var ErrRoleAlreadyExists = errors.New("role already exists")

// Check if customer already exists
func CustomerExists(db *sql.DB, userId uuid.UUID) (bool, error) {
	var existingCustomerID uuid.UUID
	err := db.QueryRow(`SELECT id FROM customers WHERE user_id = $1`, userId).Scan(&existingCustomerID)
	if err == sql.ErrNoRows {
		return false, nil
	}
	if err != nil {
		return false, err
	}
	return true, nil
}

// AddCustomer adds a new customer to the database
func AddCustomer(db *sql.DB, userId uuid.UUID, customer models.Customer, location models.Location) (uuid.UUID, uuid.UUID, error) {
	tx, err := db.Begin()
	if err != nil {
		return uuid.Nil, uuid.Nil, err
	}

	var locationID uuid.UUID
	err = tx.QueryRow(`INSERT INTO locations (id, address, city, state, country, postal_code, latitude, longitude) VALUES (uuid_generate_v4(), $1, $2, $3, $4, $5, $6, $7) RETURNING id`,
		location.Address, location.City, location.State, location.Country, location.PostalCode, location.Latitude, location.Longitude).Scan(&locationID)
	if err != nil {
		tx.Rollback()
		return uuid.Nil, uuid.Nil, err
	}

	var customerID uuid.UUID
	err = tx.QueryRow(`INSERT INTO customers (id, customer_name, contact_info, location_id, user_id) VALUES (uuid_generate_v4(), $1, $2, $3, $4) RETURNING id`,
		customer.CustomerName, customer.ContactInfo, locationID, userId).Scan(&customerID)
	if err != nil {
		tx.Rollback()
		return uuid.Nil, uuid.Nil, err
	}

	_, err = tx.Exec(`SELECT upsert_user_role($1, $2, NULL, NULL, NULL)`,
		userId, customerID)
	if err != nil {
		tx.Rollback()
		return uuid.Nil, uuid.Nil, err
	}

	return customerID, locationID, tx.Commit()
}

// EditCustomer updates an existing customer in the database
func EditCustomer(db *sql.DB, customer models.Customer) error {
	_, err := db.Exec(`UPDATE customers SET customer_name = $1, contact_info = $2, location_id = $3 WHERE id = $4`,
		customer.CustomerName, customer.ContactInfo, customer.LocationId, customer.Id)
	return err
}

// Check if business admin already exists
func BusinessAdminExists(db *sql.DB, userId uuid.UUID) (bool, error) {
	var existingBusinessAdminID uuid.UUID
	err := db.QueryRow(`SELECT id FROM business_admins WHERE user_id = $1`, userId).Scan(&existingBusinessAdminID)
	if err == sql.ErrNoRows {
		return false, nil
	}
	if err != nil {
		return false, err
	}
	return true, nil
}

// AddBusinessAdmin adds a new business admin to the database
func AddBusinessAdmin(db *sql.DB, userId uuid.UUID, businessAdmin models.BusinessAdmin, location models.Location) (uuid.UUID, uuid.UUID, error) {
	tx, err := db.Begin()
	if err != nil {
		return uuid.Nil, uuid.Nil, err
	}

	var locationID uuid.UUID
	err = tx.QueryRow(`INSERT INTO locations (id, address, city, state, country, postal_code, latitude, longitude) VALUES (uuid_generate_v4(), $1, $2, $3, $4, $5, $6, $7) RETURNING id`,
		location.Address, location.City, location.State, location.Country, location.PostalCode, location.Latitude, location.Longitude).Scan(&locationID)
	if err != nil {
		tx.Rollback()
		return uuid.Nil, uuid.Nil, err
	}

	var businessAdminID uuid.UUID
	err = tx.QueryRow(`INSERT INTO business_admins (id, company_name, contact_info, location_id, user_id) VALUES (uuid_generate_v4(), $1, $2, $3, $4) RETURNING id`,
		businessAdmin.CompanyName, businessAdmin.ContactInfo, locationID, userId).Scan(&businessAdminID)
	if err != nil {
		tx.Rollback()
		return uuid.Nil, uuid.Nil, err
	}

	_, err = tx.Exec(`SELECT upsert_user_role($1, NULL, $2, NULL, NULL)`,
		userId, businessAdminID)
	if err != nil {
		tx.Rollback()
		return uuid.Nil, uuid.Nil, err
	}

	return businessAdminID, locationID, tx.Commit()
}

// EditBusinessAdmin updates an existing business admin in the database
func EditBusinessAdmin(db *sql.DB, businessAdmin models.BusinessAdmin) error {
	_, err := db.Exec(`UPDATE business_admins SET company_name = $1, contact_info = $2, location_id = $3 WHERE id = $4`,
		businessAdmin.CompanyName, businessAdmin.ContactInfo, businessAdmin.LocationId, businessAdmin.Id)
	return err
}

// Check if transporter already exists
func TransporterExists(db *sql.DB, userId uuid.UUID) (bool, error) {
	var existingTransporterID uuid.UUID
	err := db.QueryRow(`SELECT id FROM transporters WHERE user_id = $1`, userId).Scan(&existingTransporterID)
	if err == sql.ErrNoRows {
		return false, nil
	}
	if err != nil {
		return false, err
	}
	return true, nil
}

// AddTransporter adds a new transporter to the database
func AddTransporter(db *sql.DB, userId uuid.UUID, transporter models.Transporter, location models.Location, vehicle models.Vehicle) (uuid.UUID, uuid.UUID, uuid.UUID, error) {
	tx, err := db.Begin()
	if err != nil {
		return uuid.Nil, uuid.Nil, uuid.Nil, err
	}

	var locationID uuid.UUID
	err = tx.QueryRow(`INSERT INTO locations (id, address, city, state, country, postal_code, latitude, longitude) VALUES (uuid_generate_v4(), $1, $2, $3, $4, $5, $6, $7) RETURNING id`,
		location.Address, location.City, location.State, location.Country, location.PostalCode, location.Latitude, location.Longitude).Scan(&locationID)
	if err != nil {
		tx.Rollback()
		return uuid.Nil, uuid.Nil, uuid.Nil, err
	}

	var vehicleID uuid.UUID
	err = tx.QueryRow(`INSERT INTO vehicles (id, make, model, year, latitude, longitude, max_distance, max_capacity, current_capacity) VALUES (uuid_generate_v4(), $1, $2, $3, $4, $5, $6, $7, $8) RETURNING id`,
		vehicle.Make, vehicle.Model, vehicle.Year, vehicle.Latitude, vehicle.Longitude, vehicle.MaxDistance, vehicle.MaxCapacity, vehicle.CurrentCapacity).Scan(&vehicleID)
	if err != nil {
		tx.Rollback()
		return uuid.Nil, uuid.Nil, uuid.Nil, err
	}

	var transporterID uuid.UUID
	err = tx.QueryRow(`INSERT INTO transporters (id, driver_name, vehicle_id, contact_info, location_id, user_id) VALUES (uuid_generate_v4(), $1, $2, $3, $4, $5) RETURNING id`,
		transporter.DriverName, vehicleID, transporter.ContactInfo, locationID, userId).Scan(&transporterID)
	if err != nil {
		tx.Rollback()
		return uuid.Nil, uuid.Nil, uuid.Nil, err
	}

	_, err = tx.Exec(`SELECT upsert_user_role($1, NULL, NULL, $2, NULL)`,
		userId, transporterID)
	if err != nil {
		tx.Rollback()
		return uuid.Nil, uuid.Nil, uuid.Nil, err
	}

	return transporterID, locationID, vehicleID, tx.Commit()
}

// EditTransporter updates an existing transporter in the database
func EditTransporter(db *sql.DB, transporter models.Transporter) error {
	_, err := db.Exec(`UPDATE transporters SET driver_name = $1, vehicle_id = $2, contact_info = $3, location_id = $4 WHERE id = $5`,
		transporter.DriverName, transporter.VehicleId, transporter.ContactInfo, transporter.LocationId, transporter.Id)
	return err
}

// Check if supplier already exists
func SupplierExists(db *sql.DB, userId uuid.UUID) (bool, error) {
	var existingSupplierID uuid.UUID
	err := db.QueryRow(`SELECT id FROM suppliers WHERE user_id = $1`, userId).Scan(&existingSupplierID)
	if err == sql.ErrNoRows {
		return false, nil
	}
	if err != nil {
		return false, err
	}
	return true, nil
}

// AddSupplier adds a new supplier to the database
func AddSupplier(db *sql.DB, userId uuid.UUID, supplier models.Supplier, location models.Location) (uuid.UUID, uuid.UUID, error) {
	tx, err := db.Begin()
	if err != nil {
		return uuid.Nil, uuid.Nil, err
	}

	var locationID uuid.UUID
	err = tx.QueryRow(`INSERT INTO locations (id, address, city, state, country, postal_code, latitude, longitude) VALUES (uuid_generate_v4(), $1, $2, $3, $4, $5, $6, $7) RETURNING id`,
		location.Address, location.City, location.State, location.Country, location.PostalCode, location.Latitude, location.Longitude).Scan(&locationID)
	if err != nil {
		tx.Rollback()
		return uuid.Nil, uuid.Nil, err
	}

	var supplierID uuid.UUID
	err = tx.QueryRow(`INSERT INTO suppliers (id, supplier_name, contact_info, location_id, user_id) VALUES (uuid_generate_v4(), $1, $2, $3, $4) RETURNING id`,
		supplier.SupplierName, supplier.ContactInfo, locationID, userId).Scan(&supplierID)
	if err != nil {
		tx.Rollback()
		return uuid.Nil, uuid.Nil, err
	}

	_, err = tx.Exec(`SELECT upsert_user_role($1, NULL, NULL, NULL, $2)`,
		userId, supplierID)
	if err != nil {
		tx.Rollback()
		return uuid.Nil, uuid.Nil, err
	}

	return supplierID, locationID, tx.Commit()
}

// EditSupplier updates an existing supplier in the database
func EditSupplier(db *sql.DB, supplier models.Supplier) error {
	_, err := db.Exec(`UPDATE suppliers SET supplier_name = $1, contact_info = $2, address = $3, location_id = $4 WHERE id = $5`,
		supplier.SupplierName, supplier.ContactInfo, supplier.Address, supplier.LocationId, supplier.Id)
	return err
}

// GetRolesByUserId fetches roles for a given user ID
func GetRolesByUserId(db *sql.DB, userId uuid.UUID) ([]models.Role, error) {
	rows, err := db.Query(`SELECT user_id, customer_id, business_admin_id, transporter_id, supplier_id FROM user_roles WHERE user_id = $1`, userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var roles []models.Role
	for rows.Next() {
		var role models.Role
		if err := rows.Scan(&role.UserId, &role.CustomerId, &role.BusinessAdminId, &role.TransporterId, &role.SupplierId); err != nil {
			return nil, err
		}
		roles = append(roles, role)
	}

	return roles, nil
}





