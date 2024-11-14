package main

import (
	"database/sql"
	"sync"
	"time"

	"log/slog"

	_ "github.com/marcboeker/go-duckdb"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

type db_meta struct {
	instance  string
	db        *sql.DB
	lastFetch time.Time
}

var (
	db1    db_meta
	db2    db_meta
	dbLock sync.Mutex
)

type Entity struct {
	StopName       string
	StopSequence   int
	StartDatetime  string
	Platform       string
	ArrivalDelay   int
	DepartureDelay int
	RouteShortName string
	TripHeadsign   string
}

func get_db() *sql.DB {
	// Use a lock to ensure safe concurrent access if needed
	dbLock.Lock()
	defer dbLock.Unlock()

	// Compare lastFetch times to decide which db to use
	var mostRecent *db_meta
	if db1.lastFetch.After(db2.lastFetch) {
		mostRecent = &db1

		// Trigger db2 update asynchronously if needed
		if time.Since(db2.lastFetch) > 240*time.Second {
			go func() {
				dbLock.Lock() // Ensure safe update
				defer dbLock.Unlock()
				db2.db = fetch_db()
				db2.lastFetch = time.Now()
			}()
		}

	} else {
		mostRecent = &db2

		// Trigger db1 update asynchronously if needed
		if time.Since(db1.lastFetch) > 240*time.Second {
			go func() {
				dbLock.Lock() // Ensure safe update
				defer dbLock.Unlock()
				db1.db = fetch_db()
				db1.lastFetch = time.Now()
			}()
		}
	}

	slog.Info("Using db", "instance", mostRecent.instance)

	return mostRecent.db
}

func fetch_db() *sql.DB {
	slog.Info("Fetching new db")
	// Fetch new db connection
	newDb, err := sql.Open("duckdb", "")
	check(err)

	_, err = newDb.Exec("CREATE TABLE p1 AS SELECT * FROM read_parquet('https://gtfs-public-data.fsn1.your-objectstorage.com/current_feed.parquet');")
	check(err)

	return newDb
}
