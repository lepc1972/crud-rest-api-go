package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

// User represents the user model for crud ops
type User struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	CreateAt string `json:"createAt"`
}

var db *sql.DB

func main() {

	// database connection
	var err error
	db, err = sql.Open("mysql", "root:12345@tcp(127.0.0.1:3306)/go_crud_api")
	if err != nil {
		log.Fatal(err)
	}
	// Verificar la conexión
	err = db.Ping()
	if err != nil {
		log.Fatalf("Error al verificar la conexión: %v", err)
	}
	log.Println("Conexión exitosa a la base de datos")

	defer db.Close()

	// initialize router
	router := mux.NewRouter()

	// api routes
	router.HandleFunc("/users", GetUsers).Methods("GET")
	router.HandleFunc("/users/{id}", GetUser).Methods("GET")
	router.HandleFunc("/users", CreateUser).Methods("POST")
	router.HandleFunc("/users/{id}", UpdateUser).Methods("PUT")
	router.HandleFunc("/users/{id}", DeleteUser).Methods("DELETE")

}
