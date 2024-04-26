package db

import (
	"database/sql"
	"fmt"
	log "github.com/sirupsen/logrus"

	_ "github.com/lib/pq"
)

type postgresDB struct {
	uri string
	db  *sql.DB
}

func NewPostgresDB(host string, port int, user, password, database string) DB {
	return &postgresDB{
		uri: fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
			host, port, user, password, database)}
}

func (psql *postgresDB) Connect() error {
	db, err := sql.Open("postgres", psql.uri)
	psql.db = db
	if err == nil {
		log.Println("connected successfully")
	}
	return err
}

func (psql *postgresDB) Healthcheck() error {
	return psql.db.Ping()
}

func (psql *postgresDB) Close() {
	psql.db.Close()
}

func (psql *postgresDB) GetDB() *sql.DB {
	return psql.db
}
