package sqlc

import (
	"database/sql"
	"log"
	"os"
	"simple-bank/config"
	"testing"

	_ "github.com/lib/pq"
)

var testQueries *Queries
var testDB *sql.DB

func TestMain(t *testing.M) {
	config, err := config.LoadENVConfig("../..")
	if err != nil {
		log.Fatal(err)
	}
	testDB, err = sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	testQueries = New(testDB)

	os.Exit(t.Run())
}
