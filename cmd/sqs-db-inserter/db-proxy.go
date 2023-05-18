package main

import (
	"database/sql"
	"fmt"
	"strings"

	_ "github.com/lib/pq"
)

// database timeout
var dbTimeout = 30

// maximum fields supported
var maxInsertFields = 5

type DbProxy struct {
	handle *sql.DB
	insert string
	fields []DbField
}

type DbField struct {
	dbFieldName string
}

// NewDbProxy - the factory
func NewDbProxy(cfg ServiceConfig) *DbProxy {

	// parse the configuration data to be sure it is formatted correctly
	fmapper := strings.Split(cfg.DbInsertFields, ",")
	fields := make([]DbField, len(fmapper))
	for ix, f := range fmapper {
		fields[ix].dbFieldName = f
	}

	if len(fields) > maxInsertFields {
		fatalIfError(fmt.Errorf("unsupported number of insert fields (max is %d)", maxInsertFields))
	}

	// connect to database
	connStr := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%d connect_timeout=%d sslmode=disable",
		cfg.DbUser, cfg.DbPass, cfg.DbName, cfg.DbHost, cfg.DbPort, dbTimeout)

	db, err := sql.Open("postgres", connStr)
	fatalIfError(err)

	return &DbProxy{
		handle: db,
		insert: cfg.DbInsertStatement,
		fields: fields,
	}
}

func (p *DbProxy) Insert(rec map[string]interface{}) error {

	//
	// if you add more cases here to support more insert fields, adjust
	// maxInsertFields above as necessary
	//
	var err error
	switch len(p.fields) {
	case 1:
		f1 := rec[p.fields[0].dbFieldName]
		_, err = p.handle.Exec(p.insert, f1)
	case 2:
		f1 := rec[p.fields[0].dbFieldName]
		f2 := rec[p.fields[1].dbFieldName]
		_, err = p.handle.Exec(p.insert, f1, f2)
	case 3:
		f1 := rec[p.fields[0].dbFieldName]
		f2 := rec[p.fields[1].dbFieldName]
		f3 := rec[p.fields[2].dbFieldName]
		_, err = p.handle.Exec(p.insert, f1, f2, f3)
	case 4:
		f1 := rec[p.fields[0].dbFieldName]
		f2 := rec[p.fields[1].dbFieldName]
		f3 := rec[p.fields[2].dbFieldName]
		f4 := rec[p.fields[3].dbFieldName]
		_, err = p.handle.Exec(p.insert, f1, f2, f3, f4)
	case 5:
		f1 := rec[p.fields[0].dbFieldName]
		f2 := rec[p.fields[1].dbFieldName]
		f3 := rec[p.fields[2].dbFieldName]
		f4 := rec[p.fields[3].dbFieldName]
		f5 := rec[p.fields[4].dbFieldName]
		_, err = p.handle.Exec(p.insert, f1, f2, f3, f4, f5)
	}

	return err
}

//
// end of file
//
