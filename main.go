//
// Autor: Davi Araujo
// Data: 02/12/2021
//

package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"

	"api/controller"
	"api/db"
	"api/models"
)

var (
	Url      = os.Getenv("URL")
	Endpoint = os.Getenv("ENDPOINT")
)

func init() {
	db.CreateDB()
	db.CreateTableHosts()
	if Url != "" && Endpoint != "" {
		models.InsertEndpoint(Url, Endpoint, "default")
	}
}

func main() {

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "3000"
	}

	r := mux.NewRouter()

	r.HandleFunc("/api/v1/ping", controller.PingHandler)
	r.HandleFunc("/api/v1/send", controller.SenderHandler)
	r.HandleFunc("/api/v1/endpoint", controller.EndpointHandler)
	r.HandleFunc("/api/v1/endpoints", controller.EndpointsHandler)

	srv := &http.Server{
		Addr:    ":" + PORT,
		Handler: r,
	}
	log.Printf("Starting server on " + PORT)
	err := srv.ListenAndServe()
	log.Fatal(err)

}
