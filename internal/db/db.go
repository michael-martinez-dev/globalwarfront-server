package db

import "database/sql"

type DB interface {
	GetDB() *sql.DB
	Connect() error
	Healthcheck() error
	Close()
}
