// Package classification of API
//
// Documentation for API
//
// Schemes: http
// BasePath: /
// Version: 1.0.0
//
// Consumes:
// - application/json
//
// Produces:
// - application/json
// swagger:meta
package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	. "github.com/zzibert/rest-api/api"
	. "github.com/zzibert/rest-api/data"
)

func main() {
	godotenv.Load(".env")

	var err error
	connectionString := fmt.Sprintf("user=%s dbname=%s password=%s port=%s sslmode=%s", os.Getenv("DB_USER"), os.Getenv("DB"), os.Getenv("PASSWORD"), os.Getenv("PORT"), os.Getenv("SSLMODE"))
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		panic(err)
	}

	server := http.Server{
		Addr: "127.0.0.1:",
	}
	http.HandleFunc("/group", HandleGroupRequest(&Group{Db: db}))
	http.HandleFunc("/user", HandleUserRequest(&User{Db: db}))
	server.ListenAndServe()
}
