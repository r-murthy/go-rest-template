package rest

import (
	"net/http"

	"golang-rest/client"
	dto "golang-rest/dto"

	"github.com/gorilla/mux"
)

// GetCustomer process requests to get customer details of a single customer.
func GetCustomer(db dto.Db) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		params := mux.Vars(r)
		customerID := params["customerid"]

		customer, err := db.GetCustomer(customerID)
		if err != nil {
			client.RespondError(w, http.StatusInternalServerError, "")
			return
		}
		customerDto := dto.GetCustomerFromEntity(customer)

		res := client.Response{HTTPStatus: http.StatusOK, Payload: customerDto}
		client.Respond(w, res)
	})
}
