package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
)

const (
	dbDriver = "postgres"
	dbSource = "postgresql://nada:koty123@localhost:5432/bank_db?sslmode=disable"
)

var testQueries *Queries

func TestMain(m *testing.M) {
	// Connect to test database
	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}
	testQueries = New(conn)
	// Run all tests
	os.Exit(m.Run())
}
