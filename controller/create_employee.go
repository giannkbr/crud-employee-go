package controller

import (
	"database/sql"
	"net/http"
	"path/filepath"
	"text/template"
)

func NewCreateEmployee(db *sql.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			r.ParseForm()

			name := r.Form["name"][0]
			address := r.Form["address"][0]
			npwp := r.Form["npwp"][0]

			// Save to database
			_, err := db.Exec("INSERT INTO employees (name, address, npwp) VALUES (?, ?, ?)", name, address, npwp)

			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			http.Redirect(w, r, "/employee", http.StatusSeeOther)
			return

		} else if r.Method == "GET" {
			fp := filepath.Join("views", "create.html")
			tmpl, err := template.ParseFiles(fp)

			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			tmpl.Execute(w, nil)
		}

	}
}
