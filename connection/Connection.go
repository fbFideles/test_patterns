package connection

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func New() *Connection {
	db, err := sql.Open("postgres", "host=localhost port=5432 user=postgres password=postgres dbname=exemplo sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	return &Connection{
		db: db,
	}
}

type Connection struct {
	db *sql.DB
}

func (c *Connection) Query(query string, statements ...any) (*sql.Rows, error) {
	return c.db.Query(query, statements...)
}
