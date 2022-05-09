package main

import (
	"gotraining.com/database"
	"gotraining.com/docs"
	"gotraining.com/party"
	"log"
	"net/http"
	"os"
)

const apiBasePath = "/api/v1"

// @title Party API
// @version 1.0
// @description Swagger documentation for Party API
// @BasePath /api/v1
// @securityDefinitions.basic BasicAuth
func main() {
	database.SetupDb(true)
	party.SetupRoutes(apiBasePath)
	docs.SetupRoutes()
	log.Fatal(http.ListenAndServe(os.Getenv("SERVER_PORT"), nil))

}
