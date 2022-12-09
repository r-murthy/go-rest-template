package main

import (
	"context"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	firebase "firebase.google.com/go"
	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
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

	httpApi := r.PathPrefix("/api").Subrouter()
	wsApi := r.PathPrefix("/ws").Subrouter()

	httpApi.Use(loggingMiddleware)
	httpApi.Use(rest.AuthMiddleware(authClient))

	wsApi.Use(loggingMiddleware)
	wsApi.Use(rest.AuthMiddleware(authClient))

	httpApi.Handle("/customer", customer.UpsertCustomer(db)).Methods("PUT", "OPTIONS")
	httpApi.Handle("/customers", customer.ListCustomers(db)).Methods("GET", "OPTIONS")
	httpApi.Handle("/customers/{customerid}", customer.GetCustomer(db)).Methods("GET", "OPTIONS")

	httpApi.Handle("/products", product.UpsertProduct()).Methods("POST", "OPTIONS")

	httpApi.Handle("/ping", pingHandler()).Methods("GET", "OPTIONS")

	wsApi.Handle("/greet", greetHandler())
	log.Println("Starting server on port " + config.APIConfig.Port)
	err = http.ListenAndServe(":"+config.APIConfig.Port, r)
	if err != nil {
		log.Fatalln("error starting server:", err)
	}
}

func greetHandler() http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		upgrader.CheckOrigin = func(req *http.Request) bool {
			log.Printf("Here is the %v", req.Header)
			return true
		}

		ws, err := upgrader.Upgrade(rw, req, nil)
		if err != nil {
			log.Println("error")
		}
		log.Printf("Client Successfully connected... ")

		reader(ws)
	})
}

func reader(conn *websocket.Conn) {
	messageType, p, err := conn.ReadMessage()
	if err != nil {
		log.Printf("error %v", err)
		return
	}
	log.Printf("Incoming message %v", string(p))

	for i := 1; i < 5; i++ {
		time.Sleep(5 * time.Second)
		sendBack := []byte("Message from the server")

		if err := conn.WriteMessage(messageType, sendBack); err != nil {
			log.Printf("error %v", err)
			return
		}
	}
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.RequestURI)
		next.ServeHTTP(w, r)
	})
}

func pingHandler() http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		log.Printf("here is the header %v", req.Header.Get(http.CanonicalHeaderKey("Origin")))
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
