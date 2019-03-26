package main

import (
	"database/sql"
	"fmt"
	"strings"
	"time"
)

import _ "github.com/go-sql-driver/mysql"

// Log into specified server and get the process list
func GetProcessList(databaseServer string, user string, password string, date time.Time) (output []MySQLProcess) {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s)/", user, password, databaseServer))
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err.Error())
	}

	rows, err := db.Query("SELECT ID, USER, HOST, DB, COMMAND, TIME, STATE, INFO FROM INFORMATION_SCHEMA.PROCESSLIST")
	if err != nil {
		panic(err.Error())
	}
	for rows.Next() {
		var NewProcess MySQLProcess
		_ = rows.Scan(&NewProcess.ID,
			&NewProcess.User,
			&NewProcess.Host,
			&NewProcess.Database,
			&NewProcess.Command,
			&NewProcess.Time,
			&NewProcess.State,
			&NewProcess.Info,
		)
		hostStringSplit := strings.Split(*NewProcess.Host, ":")
		NewProcess.Host = &hostStringSplit[0]
		NewProcess.Date = date
		NewProcess.DatabaseHost = hostname
		output = append(output, NewProcess)
	}

	return
}
