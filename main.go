package main

import (
	"flag"
	"fmt"
	"os"
	"time"
)

var databaseServer = flag.String("database", "127.0.0.1", "database server to get stats from")
var databaseUser = flag.String("dbuser", "root", "user to log in as (should have ability to INFORMATION_SCHEMA")
var databasePassword = flag.String("dbpass", "", "password to log in to database as dbuser")
var elkserver = flag.String("elasticserver", "http://127.0.0.1:9200", "full URL to Elastic Stack server")

func main() {
	start := time.Now()
	flag.Parse()
	hostname, err := os.Hostname()
	if err != nil {
		fmt.Println("Could not get hostname; defaulting to localhost")
	}
	processList := GetProcessList(*databaseServer, *databaseUser, *databasePassword)
	var fullProcessList MySQLProcessList
	fullProcessList.Date = start
	fullProcessList.DatabaseHost = hostname
	fullProcessList.Processes = processList
	SendToElk(*elkserver, "mysql", fullProcessList)
}
