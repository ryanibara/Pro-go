package middlewares

import (
	"database/sql"
	"log"
	"net/http"
	"os"
	"text/template"

	"github.com/joho/godotenv"
)

func VerifyUser(endpointHandler func(http.ResponseWriter, *http.Request)) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		username := req.PostFormValue("username")

		err := godotenv.Load()
		if err != nil {
			log.Fatal("Error loading .env file: %w", err)
		 
		}
		DBlink := os.Getenv("DB_LINK")
		dsn := DBlink
		db, err := sql.Open("postgres", dsn)
		if err != nil {
			log.Fatal("Failed to execute query: %w" ,err)
		}
		defer db.Close()
		rows, err := db.Query("SELECT * FROM Users WHERE username = $1", username)
		if err != nil {
			log.Fatal("Error executing SQL query: %w", err)
		}
		defer rows.Close()
	
		if rows.Next() {
			t := template.Must(template.ParseFiles("register-error.html"))
			t.Execute(w, nil)
		} else {
			log.Println("Username is not in the database.")
			endpointHandler(w, req)
		}

	})
}