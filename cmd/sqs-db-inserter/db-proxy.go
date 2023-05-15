package main

import (
	"fmt"
	dbx "github.com/go-ozzo/ozzo-dbx"
	_ "github.com/lib/pq"
)

// database connection timeout
var connectionTimeout = 30

type DbProxy struct {
	handle *dbx.DB
	insert string
	//size   int
}

// NewDbProxy - the factory
func NewDbProxy(cfg ServiceConfig) *DbProxy {

	// connect to database
	connStr := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%d connect_timeout=%d sslmode=disable",
		cfg.DbUser, cfg.DbPass, cfg.DbName, cfg.DbHost, cfg.DbPort, connectionTimeout)

	db, err := dbx.MustOpen("postgres", connStr)
	fatalIfError(err)

	return &DbProxy{
		handle: db,
		insert: cfg.DbInsertStatement,
	}
}
