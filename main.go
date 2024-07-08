package main

import (
	"net/http"

	"github.com/giannkbr/crud-employee-go/database"
	"github.com/giannkbr/crud-employee-go/routes"
)

func main() {
	db := database.InitDatabase()

	server := http.NewServeMux()

	routes.MapRoutes(server, db)

	http.ListenAndServe(":8080", server)
}
