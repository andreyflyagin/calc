package database

import (
	"context"
	"fmt"
	"github.com/jmoiron/sqlx"
	// postgres driver
	_ "github.com/lib/pq"
)

const (
	host     = "postgres"
	port     = 5432
	user     = "calc"
	password = "calc"
	dbname   = "calc"
)

// Connect - connect to database
func Connect(ctx context.Context) (*sqlx.DB, error) {
	db, err := sqlx.ConnectContext(ctx, "postgres",
		fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname))

	if err != nil {
		return nil, err
	}
	if err = db.PingContext(ctx); err != nil {
		return nil, err
	}
	return db, nil
}
