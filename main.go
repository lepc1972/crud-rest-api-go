package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

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
	db, err = sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/go_crud_api")
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

	//start server port 8080
	log.Fatal(http.ListenAndServe(":8080", router))

}

// func get all users
// func get all users
func GetUsers(w http.ResponseWriter, r *http.Request) {
	// Declare an empty slice to store user objects
	var users []User

	// Execute a SQL query to select user data
	rows, err := db.Query("SELECT id, name, email, created_at FROM users")

	// Check for errors during query execution
	if err != nil {
		// Return an HTTP error with the error message and a 500 Internal Server Error status
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Defer closing the rows until the function returns
	defer rows.Close()

	// Iterate through each row returned by the query
	for rows.Next() {
		// Create a new User object to store the current row's data
		var user User

		// Scan the current row's data into the User object
		if err := rows.Scan(&user.ID, &user.Name, &user.Email, &user.CreateAt); err != nil {
			// Return an HTTP error if there's an issue scanning data
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Append the scanned User object to the users slice
		users = append(users, user)
	}

	// Encode the users slice as JSON and write it to the response writer
	json.NewEncoder(w).Encode(users)
}

// func get user by id
func GetUser(w http.ResponseWriter, r *http.Request) {
	// Get route parameters
	params := mux.Vars(r)
	id := params["id"]

	// Prepare the SQL query
	var user User

	// Execute the SQL query
	err := db.QueryRow("SELECT id, name, email, created_at FROM users WHERE id = ?", id).Scan(&user.ID, &user.Name, &user.Email, &user.CreateAt)

	// Handle errors
	if err != nil {
		// If there's an error, return an HTTP 500 error
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Encode the user as JSON and write it to the response
	json.NewEncoder(w).Encode(user)
}

// func create new user
func CreateUser(w http.ResponseWriter, r *http.Request) {
	// Parse the request body into a User object
	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Prepare the SQL statement
	stmt, err := db.Prepare("INSERT INTO users (name, email) VALUES (?, ?)")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer stmt.Close()

	// Execute the SQL statement
	_, err = stmt.Exec(user.Name, user.Email)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Return a success response
	w.WriteHeader(http.StatusCreated)
}

// func update user
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	// Get route parameters
	params := mux.Vars(r)
	id := params["id"]

	// Parse the request body into a User object
	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		// If there's an error parsing the JSON, return a bad request error
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Prepare the SQL statement
	stmt, err := db.Prepare("UPDATE users SET name = ?, email = ? WHERE id = ?")
	if err != nil {
		// If there's an error preparing the statement, return an internal server error
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// Clean up the statement when we're done
	defer stmt.Close()

	// Execute the SQL statement
	_, err = stmt.Exec(user.Name, user.Email, id)
	if err != nil {
		// If there's an error executing the statement, return an internal server error
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Return a success response
	w.WriteHeader(http.StatusOK)
}

// func delete user
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	// Get route parameters
	params := mux.Vars(r)
	id := params["id"]

	// Prepare the SQL statement
	stmt, err := db.Prepare("DELETE FROM users WHERE id = ?")
	if err != nil {
		// If there's an error preparing the statement, return an internal server error
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// Clean up the statement when we're done
	defer stmt.Close()

	// Execute the SQL statement
	_, err = stmt.Exec(id)
	if err != nil {
		// If there's an error executing the statement, return an internal server error
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Return a success response
	w.WriteHeader(http.StatusOK)
}
