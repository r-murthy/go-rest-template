package rest

import (
	"encoding/json"
	"log"
	"net/http"

	"golang-rest/client"
	model "golang-rest/dto"
)

// UpsertCustomer processes PUT requests to upsert customers to db.
func UpsertCustomer(db model.Db) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if r.Method == http.MethodOptions {
			response := client.Response{HTTPStatus: http.StatusOK}
			client.Respond(w, response)
		}
		var customer model.Customer
		err := json.NewDecoder(r.Body).Decode(&customer)
		if err != nil {
			log.Printf("error unmarshalling customer: %v", err)
			client.RespondError(w, http.StatusBadRequest, "error reading customer from request payload")
			return
		}

		var isNewCustomer bool
		if len(customer.CustomerID) == 0 {
			isNewCustomer = true
		}

		custEntity := model.GetCustomerFromDto(customer)
		err = db.UpsertCustomer(&custEntity)
		if err != nil {
			log.Printf("error inserting customer: %v\n", err.Error())
			client.RespondError(w, http.StatusInternalServerError, "please try again!")
			return
		}

		c := model.GetCustomerFromEntity(custEntity)

		var status int
		if isNewCustomer {
			status = http.StatusCreated
		} else {
			status = http.StatusOK
		}
		res := client.Response{HTTPStatus: status, Payload: c}
		client.Respond(w, res)
	})
}
