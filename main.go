package main

import (
	"database/sql"
	"flag"
	"github.com/marcsauter/single"
	"github.com/olivere/elastic/v7"
	"log"
	"os"
	"sync"
	"time"
)

var tickerTime time.Duration
var databaseServer = flag.String("database", "127.0.0.1", "database server to get stats from")
var databaseUser = flag.String("dbuser", "root", "user to log in as (should have ability to INFORMATION_SCHEMA")
var databasePassword = flag.String("dbpass", "", "password to log in to database as dbuser")
var elkserver = flag.String("elasticserver", "http://127.0.0.1:9200", "full URL to Elastic Stack server")

var hostname, _ = os.Hostname()
var wg sync.WaitGroup

func main() {
	s := single.New("sqlstats")
	if err := s.CheckLock(); err != nil && err == single.ErrAlreadyRunning {
		log.Fatal("another instance of the app is already running, exiting")
	} else if err != nil {
		// Another error occurred, might be worth handling it as well
		log.Fatalf("failed to acquire exclusive app lock: %v", err)
	}
	defer s.TryUnlock()
	flag.DurationVar(&tickerTime, "ticker", 60*time.Second, "time to wait before starting next query")
	flag.Parse()
	dbConn, err := connectToDB(*databaseServer, *databaseUser, *databasePassword)
	checkFatal(err)
	defer dbConn.Close()
	elasticConn, err := elastic.NewClient(elastic.SetURL(*elkserver))
	checkFatal(err)
	processLoopTicker := time.NewTicker(tickerTime)
	wg.Add(1)
	// LOOP
	go func() {
		wg.Add(1)
		for range processLoopTicker.C {
			ProcessLoop(dbConn, elasticConn)
		}
		wg.Done()
	}()
	wg.Wait()
	// ENDLOOP
}

func ProcessLoop(dbConn *sql.DB, client *elastic.Client) {
	processList := GetProcessList(dbConn)
	var fullProcessList MySQLProcessList
	fullProcessList.Processes = processList
	SendToElk(client, "mysql", fullProcessList)
}
