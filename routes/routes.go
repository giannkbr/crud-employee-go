package routes

import (
	"database/sql"
	"net/http"

	"github.com/giannkbr/crud-employee-go/controller"
)

func MapRoutes(server *http.ServeMux, db *sql.DB) {
	server.HandleFunc("/hello", controller.NewHelloWorldController())
	server.HandleFunc("/employee", controller.NewIndexEmployee(db))
	server.HandleFunc("/employee/create", controller.NewCreateEmployee(db))
	server.HandleFunc("/employee/update", controller.NewUpdateEmployee(db))
	server.HandleFunc("/employee/delete", controller.NewDeleteEmployee(db))
}
