package main

import "time"

type MySQLProcess struct {
	ID           *int      `json:"ID"`
	User         *string   `json:"User"`
	Host         *string   `json:"Host"`
	Database     *string   `json:"Database"`
	Command      *string   `json:"Command"`
	Time         *int      `json:"Time"`
	State        *string   `json:"State"`
	Info         *string   `json:"Info"`
	Date         time.Time `json:"Date"`
	DatabaseHost string    `json:"DatabaseHost"`
}

type MySQLProcessList struct {
	Date         time.Time      `json:"Date"`
	Processes    []MySQLProcess `json:"Processes"`
	DatabaseHost string         `json:"DatabaseHost"`
}
