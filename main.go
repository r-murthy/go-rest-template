package main

import (
	"context"
	"io/ioutil"
	"log"
	"net/http"

	firebase "firebase.google.com/go"
	"github.com/gorilla/mux"
	"google.golang.org/api/option"
	"gopkg.in/yaml.v2"

	"golang-rest/client"
	"golang-rest/config"
	"golang-rest/db"
	"golang-rest/rest"
	customer "golang-rest/rest/customer"
	product "golang-rest/rest/product"
)

func main() {
	configFile, err := ioutil.ReadFile("config.yaml")
	if err != nil {
		log.Fatalln("error reading configuration file:", err)
	}

	err = yaml.Unmarshal([]byte(configFile), &config.APIConfig)
	if err != nil {
		log.Fatalln("error parsing configuration data:", err)
	}

	opt := option.WithCredentialsFile(config.APIConfig.FBServiceFile)
	fbApp, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		log.Fatalf("error initializing app: %v\n", err)
	}

	authClient, err := fbApp.Auth(context.Background())
	if err != nil {
		log.Fatalf("error getting Auth client: %v\n", err)
	}

	db, err := db.InitDb(config.APIConfig.PGConn)
	if err != nil {
		log.Fatalln("error InitDb:", err)
	}

	r := mux.NewRouter()

	api := r.PathPrefix("/api").Subrouter()
	api.Use(loggingMiddleware)
	api.Use(rest.AuthMiddleware(authClient))

	api.Handle("/customer", customer.UpsertCustomer(db)).Methods("PUT", "OPTIONS")
	api.Handle("/customers", customer.ListCustomers(db)).Methods("GET", "OPTIONS")
	api.Handle("/customers/{customerid}", customer.GetCustomer(db)).Methods("GET", "OPTIONS")

	api.Handle("/products", product.UpsertProduct()).Methods("POST", "OPTIONS")

	api.Handle("/ping", pingHandler()).Methods("GET", "OPTIONS")

	log.Println("Starting server on port " + config.APIConfig.Port)
	err = http.ListenAndServe(":"+config.APIConfig.Port, r)
	if err != nil {
		log.Fatalln("error starting server:", err)
	}
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.RequestURI)
		next.ServeHTTP(w, r)
	})
}

func pingHandler() http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		var response client.Response
		response = client.Response{HTTPStatus: http.StatusOK}
		client.Respond(res, response)
	})
}

func handleOptions() http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		response := client.Response{HTTPStatus: http.StatusOK}
		client.Respond(res, response)
	})
}
