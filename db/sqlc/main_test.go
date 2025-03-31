package db

import (
	"context"
	"log"
	"os"
	"testing"

	"github.com/jackc/pgx/v5/pgxpool"
)

const dbSource = "postgresql://root:secret@localhost:5433/simple_bank?sslmode=disable"

var testQueries *Queries
var testDb *pgxpool.Pool

func TestMain(m *testing.M) {
	var err error
	ctx := context.Background()

	testDb, err = pgxpool.New(ctx, dbSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	testQueries = New(testDb)

	os.Exit(m.Run())
}
