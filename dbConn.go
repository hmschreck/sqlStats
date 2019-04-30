package main

import (
	"database/sql"
	"fmt"
	"os"
)
import _ "github.com/go-sql-driver/mysql"

func connectToDB(databaseServer string, user string, password string) (dbConn *sql.DB, err error) {
	dbConn, err = sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s)/", user, password, databaseServer))
	return
}

func checkFatal(err error) {
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Received error in DB connect: %s", err.Error())
		os.Exit(1)
	}
}
