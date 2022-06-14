package entity

// Customer is the model of a person/business entity.
type Customer struct {
	ID          int    `db:"id"`
	CustomerID  string `db:"customer_id"`
	Name        string `db:"name"`
	Email       string `db:"email"`
	PhoneNumber string `db:"phone_no"`
}
