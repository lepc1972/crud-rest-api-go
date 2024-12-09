# CRUD REST API with Go and MySQL

This project demonstrates a simple CRUD (Create, Read, Update, Delete) REST API built using Go and MySQL. The API allows for basic operations on user data stored in a local MySQL database.

## Features

- RESTful API endpoints for creating, reading, updating, and deleting users
- Uses Go for the backend
- Utilizes MySQL as the database
- Local database setup for easy testing and development

## Prerequisites

- Go installed on your system
- MySQL installed locally
- MySQL Workbench or similar tool for database management

## Setup

1. Clone the repository:

git clone https://github.com/lepc1972/crud-rest-api-go.git



2. Navigate to the project directory:

cd crud-rest-api-go


3. Install dependencies:

go mod tidy


4. Set up the MySQL database:
   - Create a new schema named `go_crud_api`
   - Create a table named `users` with columns: `id`, `name`, `email`, `created_at`

5. Configure the database connection in `main.go`:
   - Replace the placeholders in the database connection string with your actual MySQL credentials and host

6. Run the API:

go run main.go



## Usage

The API provides the following endpoints:

- GET /users - Retrieve all users
- GET /users/{id} - Retrieve a specific user
- POST /users - Create a new user
- PUT /users/{id} - Update a user
- DELETE /users/{id} - Delete a user

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.


