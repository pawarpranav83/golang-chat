package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	// postgres driver for connection, blank identifier (otherwise would be removed)
	_ "github.com/lib/pq"
)

const (
	dbDriver = "postgres"
	dbSource = "postgresql://root:mysecretpassword@localhost:5432/chat-db?sslmode=disable"
)

// Queries struct contains db DBTX - which can either be a db conn or a db transaction.
var testQueries *Queries

// Doubt - rewatch #6
var testDB *sql.DB

func TestMain(m *testing.M) {
	// conn to db
	var err error
	testDB, err = sql.Open(dbDriver, dbSource)

	if err != nil {
		log.Fatal("Cannot connect to db: ", err)
	}

	// New func in db.go
	testQueries = New(testDB)

	// Run returns an exit code that is to be passed to os.Exit()
	// Run runs the unit tests.
	os.Exit(m.Run())
}
