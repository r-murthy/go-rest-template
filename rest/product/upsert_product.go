package rest

import (
	"encoding/json"
	"log"
	"net/http"

	"golang-rest/client"
	"golang-rest/dto"
)

// UpsertCustomer processes PUT requests to upsert customers to db.
func UpsertProduct() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var prod model.Product

		err := json.NewDecoder(r.Body).Decode(&prod)
		if err != nil {
			log.Printf("error unmarshalling product: %v", err)
			// For a graphql endpoint, an appropriate response needs to be sent
			client.RespondError(w, http.StatusBadRequest, "error reading product from request payload")
			return
		}

		log.Printf("the product %v", prod)

		status := http.StatusOK
		res := client.Response{HTTPStatus: status, Payload: prod}
		client.Respond(w, res)
	})
}
