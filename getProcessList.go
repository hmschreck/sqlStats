package main

import (
	"database/sql"
	"strings"
	"time"
)

import _ "github.com/go-sql-driver/mysql"

// Log into specified server and get the process list
func GetProcessList(db *sql.DB) (output []MySQLProcess) {
	start := time.Now()
	err := db.Ping()
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
		NewProcess.Date = start
		NewProcess.DatabaseHost = hostname
		output = append(output, NewProcess)
	}

	return
}
