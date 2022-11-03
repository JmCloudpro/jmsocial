package db

import (
	"api/src/config"
	"database/sql"

	_ "github.com/go-sql-driver/mysql" //Driver mysql, needs to be imported manualy:  Go to main api package and run in cmd: go get github.com/go-sql-driver/mysql  and needs to start with _, because we'll not declare this driver in any funcion.
)

// Conn opens a connection with the db.
func Conn() (*sql.DB, error) {
	db, err := sql.Open("mysql", config.Dbcon)
	if err != nil {

		return nil, err
	}
	if err = db.Ping(); err != nil {

		db.Close()
		return nil, err
	}
	return db, nil
}
