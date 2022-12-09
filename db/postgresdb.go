package db

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq" // used to specify postgres driver
)

// PostgresDb stores a connection to a postgres db and implements Db interface.
type PostgresDb struct {
	dbConn *sqlx.DB
}

// InitDb creates a table in postgres using the configuration provided.
func InitDb(connStr string) (*PostgresDb, error) {
	dbConn, err := sqlx.Connect("postgres", connStr)
	if err != nil {
		return nil, err
	}

	p := &PostgresDb{
		dbConn: dbConn,
	}
	err = p.dbConn.Ping()
	if err != nil {
		return nil, err
	}
	err = p.createSchema()
	if err != nil {
		return nil, err
	}
	err = p.createTable()
	if err != nil {
		return nil, err
	}

	return p, nil
}

func (p *PostgresDb) createTable() (err error) {

	customers :=
		`
			CREATE TABLE IF NOT EXISTS golang_template.customers (
				id 				BIGSERIAL NOT NULL PRIMARY KEY,
				customer_id 	TEXT NOT NULL,
				name			TEXT NOT NULL,
				email	  		TEXT,
				phone_no 		TEXT,
				CONSTRAINT unique_customer_org UNIQUE (customer_id)
			)
		`

	_, err = p.dbConn.Exec(customers)

	return
}

func (p *PostgresDb) createSchema() (err error) {

	schema :=
		`
	 		CREATE SCHEMA IF NOT EXISTS golang_template
	 	`
	_, err = p.dbConn.Exec(schema)

	return
}
