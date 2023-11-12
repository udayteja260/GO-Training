package main

import (
	api "assignment2/api/library"
	"database/sql"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

// "assignment2/oper"
// "assignment2/assignment3"
// "assignment2/calci"
// "fmt"

func main() {
	// oper.Operation1()
	// assignment3.Sli()
	// i1 := calci.Inputs{Input1: 1, Input2: 6}
	// fmt.Println(i1.Add())
	dsn := "root:admin@tcp(127.0.0.1:3306)/sys?parseTime=true"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	api.RegisterRoutes(db)

	log.Println("server starting on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
