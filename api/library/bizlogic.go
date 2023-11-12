package api

import (
	"assignment2/dataservice"
	"database/sql"
	"net/http"
)

func CreateBookLogic(db *sql.DB, w http.ResponseWriter, r *http.Request) error {
	return dataservice.CreateBook(db, w, r)
}
