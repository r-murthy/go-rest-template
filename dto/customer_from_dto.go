package model

import "golang-rest/db/entity"

func GetCustomerFromEntity(c entity.Customer) (custModel Customer) {
	custModel.CustomerID = c.CustomerID
	custModel.Name = TrimmedString(c.Name)
	custModel.Email = TrimmedString(c.Email)
	custModel.PhoneNumber = TrimmedString(c.PhoneNumber)
	return
}

func GetCustomerFromDto(c Customer) (custEntity entity.Customer) {
	custEntity.CustomerID = c.CustomerID
	custEntity.Name = string(c.Name)
	custEntity.Email = string(c.Email)
	custEntity.PhoneNumber = string(c.PhoneNumber)
	return
}

func GetCustomersFromEntity(c []entity.Customer) (customersDto []Customer) {
	custsModel := make([]Customer, len(c))
	for i, cust := range c {
		custsModel[i] = GetCustomerFromEntity(cust)
	}
	customersDto = custsModel[:]
	return
}
