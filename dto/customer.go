package model

// Customer is the model of a person/business entity.
type Customer struct {
	ID          int           `json:"-"`
	CustomerID  string        `json:"customerid"`
	Name        TrimmedString `json:"name"`
	Email       TrimmedString `json:"email"`
	PhoneNumber TrimmedString `json:"phone"`
}

// ListCustomersResponse is the payload sent in response to a list customers request.
type ListCustomersResponse struct {
	Customers     []Customer `json:"customers"`
	NextPageToken int        `json:"nextpagetoken"`
	TotalSize     int        `json:"totalsize"`
}
