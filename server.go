package main

import "database/sql"

func main() {

	var err error
	db, err := sql.Open("postgres", "user=zzibert dbname=postgres password=nekineki port=5432 sslmode=disable")

	if err != nil {
		panic(err)
	}
}
