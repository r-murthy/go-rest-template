package db

import (
	"context"
	"database/sql"
	"log"

	"golang-rest/config"
	"golang-rest/db/entity"
)

// UpsertCustomer updates customer, if not found inserts a new record.
func (p *PostgresDb) UpsertCustomer(c *entity.Customer) (err error) {
	opts := sql.TxOptions{Isolation: sql.LevelRepeatableRead}
	tx, err := p.dbConn.BeginTx(context.Background(), &opts)
	if err != nil {
		return
	}

	defer tx.Rollback()

	if len(c.CustomerID) == 0 {
		err = assignCustomerID(c, tx)
		if err != nil {
			return
		}
	}

	query :=
		`
			INSERT INTO golang_template.customers (customer_id, name, email, phone_no) 
			VALUES ($1, $2, $3, $4)
			ON CONFLICT (customer_id)
			DO UPDATE SET
				name = EXCLUDED.name,
				email = EXCLUDED.email,
				phone_no = EXCLUDED.phone_no
		`

	_, err = tx.ExecContext(context.Background(), query, c.CustomerID, c.Name, c.Email, c.PhoneNumber)
	if err != nil {
		return
	}
	err = tx.Commit()
	return
}

func assignCustomerID(customer *entity.Customer, tx *sql.Tx) (err error) {
	latestCustomerID, err := getLatestCustomerID(tx)
	if err != nil {
		return
	}
	customer.CustomerID, err = GenerateID(latestCustomerID, "CUST", "%04d")
	return
}

func getLatestCustomerID(tx *sql.Tx) (customerID string, err error) {
	var nullableID sql.NullString
	query :=
		`
			SELECT MAX(customer_id)
			FROM golang_template.customers
		`
	err = tx.QueryRow(query).Scan(&nullableID)
	customerID = nullableID.String
	return
}

// GetCustomersCount gets the total number customers.
func (p *PostgresDb) GetCustomersCount() (count int, err error) {
	query :=
		`
			SELECT COUNT(*)
			FROM golang_template.customers
		`
	err = p.dbConn.Get(&count, query)
	if err != nil {
		log.Printf("error getting customers count: %v", err.Error())
	}
	return
}

// GetCustomers gets a paginated list of customers.
func (p *PostgresDb) GetCustomers(pageToken int) (customers []entity.Customer, err error) {
	// customers = []*entity.Customer{}
	query :=
		`
			SELECT customer_id, name, email, phone_no
			FROM golang_template.customers
			ORDER BY customer_id ASC
			LIMIT CASE WHEN $1 > 0 THEN $2::integer END
			OFFSET $3
		`
	offset := 0
	if pageToken > 0 {
		offset = (pageToken - 1) * config.APIConfig.CustomersPageSize
	}

	err = p.dbConn.Select(&customers, query, pageToken, config.APIConfig.CustomersPageSize, offset)
	if err != nil {
		log.Printf("error getting paginated list of customers: %v", err.Error())
	}
	return
}

// GetCustomer gets customer with specified customer id.
func (p *PostgresDb) GetCustomer(customerID string) (customer entity.Customer, err error) {
	query :=
		`
			SELECT customer_id, name, email, phone_no
			FROM golang_template.customers
			WHERE customer_id = $1
		`
	err = p.dbConn.Get(&customer, query, customerID)
	if err != nil {
		log.Printf("error getting customer with ID %v: %v", customerID, err.Error())
	}
	return
}
