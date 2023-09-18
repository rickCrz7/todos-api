package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/jackc/pgx/v5/stdlib"
)

func main() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	conn, err := sql.Open("pgx", "postgres://todos:T0d05!@localhost:5432/todos?sslmode=disable")
	if err != nil {
		log.Fatal("Could not open database connection: ", err)
	}
	defer conn.Close()

	r := mux.NewRouter()

	NewTodoHandler(NewTodoDao(conn), r)

	log.Println("Starting server on port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
