package controller

import (
	"database/sql"
	"net/http"
)

func NewDeleteEmployee(db *sql.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		// delete employee data from database
		id := r.URL.Query().Get("id")
		_, err := db.Exec("DELETE FROM employees WHERE id = ?", id)

		if err != nil {
			w.Write([]byte(err.Error()))
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		http.Redirect(w, r, "/employee", http.StatusMovedPermanently)
	}
}
