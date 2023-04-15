package rest

import (
	"net/http"
	"strconv"

	"golang-rest/client"
	"golang-rest/config"
	model "golang-rest/dto"
)

// ListCustomers is a handler to get all customers of an organisation.
func ListCustomers(db model.Db) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		var pageToken int
		if page := r.URL.Query().Get("page"); len(page) > 0 {
			pageToken, _ = strconv.Atoi(page)
		}
		listCustomers(pageToken, db, w)
	})
}

func listCustomers(page int, db model.Db, w http.ResponseWriter) {
	count, err := db.GetCustomersCount()
	if err != nil {
		client.RespondError(w, http.StatusInternalServerError, "")
		return
	}

	if isValid := validatePageToken(page, count); !isValid {
		client.RespondError(w, http.StatusBadRequest, "page does not exist")
		return
	}

	customers, err := db.GetCustomers(page)
	if err != nil {
		client.RespondError(w, http.StatusInternalServerError, "")
		return
	}

	var customersModel []model.Customer
	if len(customers) > 0 {
		customersModel = model.GetCustomersFromEntity(customers)
	} else {
		customersModel = make([]model.Customer, 0)
	}

	nextPage := 0
	if page*config.APIConfig.CustomersPageSize < count {
		nextPage = page + 1
	}

	res := model.ListCustomersResponse{Customers: customersModel, NextPageToken: nextPage, TotalSize: count}
	r := client.Response{HTTPStatus: http.StatusOK, Payload: res}
	client.Respond(w, r)
}

func validatePageToken(pageToken int, customersCount int) bool {
	return customersCount > (pageToken-1)*config.APIConfig.CustomersPageSize ||
		(customersCount == 0 && pageToken == 1) ||
		pageToken == 0
}
