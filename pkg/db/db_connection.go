package db

import (
	"database/sql"
	"fmt"
	"net/url"
	"os"

	_ "github.com/jackc/pgx/v4/stdlib"
)

var db_static *sql.DB

func Connect() {
	dbpass := os.Getenv("DB_PASS")
	dsn := url.URL{
		Scheme: "postgres",
		Host:   "foreveryours-db.cie6oavuia6e.us-east-2.rds.amazonaws.com:5432",
		User:   url.UserPassword("root", dbpass),
		Path:   "initial_db",
	}
	q := dsn.Query()
	q.Add("sslmode", "prefer")
	dsn.RawQuery = q.Encode()
	db, err := sql.Open("pgx", dsn.String())
	if err != nil {
		fmt.Println("sql.Open", err)
		defer func() {
			_ = db.Close()
			fmt.Println("closed")
		}()
	}
	db_static = db

}

func db() *sql.DB {
	return db_static
}
