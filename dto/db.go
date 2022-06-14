package model

import "golang-rest/db/entity"

// Db specifies methods to manage resources.
type Db interface {
	UpsertCustomer(*entity.Customer) error
	GetCustomers(int) ([]entity.Customer, error)
	GetCustomer(string) (entity.Customer, error)
	GetCustomersCount() (int, error)
}
