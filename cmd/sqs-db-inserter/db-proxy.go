package main

import (
	"database/sql"
	"fmt"
	"strings"
	"time"

	// postgres
	//_ "github.com/lib/pq"

	// mysql
	_ "github.com/go-sql-driver/mysql"
)

// database timeout
var dbTimeout = 30

// maximum fields supported
var maxInsertFields = 10

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

	// connect to database (postgres)
	//connStr := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%d connect_timeout=%d",
	//	cfg.DbUser, cfg.DbPass, cfg.DbName, cfg.DbHost, cfg.DbPort, dbTimeout)
	//db, err := sql.Open("postgres", connStr)

	// connect to database (mysql)
	connStr := fmt.Sprintf("%s:%s@tcp(%s)/%s?timeout=%ds&parseTime=true",
		cfg.DbUser, cfg.DbPass, cfg.DbHost, cfg.DbName, dbTimeout)
	db, err := sql.Open("mysql", connStr)

	fatalIfError(err)

	err = db.Ping()
	fatalIfError(err)

	// specifically for mysql
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(cfg.Workers)
	db.SetMaxIdleConns(cfg.Workers)

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
	err := p.ensureFieldsExist(rec)
	if err != nil {
		return err
	}
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
	case 6:
		f1 := rec[p.fields[0].dbFieldName]
		f2 := rec[p.fields[1].dbFieldName]
		f3 := rec[p.fields[2].dbFieldName]
		f4 := rec[p.fields[3].dbFieldName]
		f5 := rec[p.fields[4].dbFieldName]
		f6 := rec[p.fields[5].dbFieldName]
		_, err = p.handle.Exec(p.insert, f1, f2, f3, f4, f5, f6)
	case 7:
		f1 := rec[p.fields[0].dbFieldName]
		f2 := rec[p.fields[1].dbFieldName]
		f3 := rec[p.fields[2].dbFieldName]
		f4 := rec[p.fields[3].dbFieldName]
		f5 := rec[p.fields[4].dbFieldName]
		f6 := rec[p.fields[5].dbFieldName]
		f7 := rec[p.fields[6].dbFieldName]
		_, err = p.handle.Exec(p.insert, f1, f2, f3, f4, f5, f6, f7)
	case 8:
		f1 := rec[p.fields[0].dbFieldName]
		f2 := rec[p.fields[1].dbFieldName]
		f3 := rec[p.fields[2].dbFieldName]
		f4 := rec[p.fields[3].dbFieldName]
		f5 := rec[p.fields[4].dbFieldName]
		f6 := rec[p.fields[5].dbFieldName]
		f7 := rec[p.fields[6].dbFieldName]
		f8 := rec[p.fields[7].dbFieldName]
		_, err = p.handle.Exec(p.insert, f1, f2, f3, f4, f5, f6, f7, f8)
	case 9:
		f1 := rec[p.fields[0].dbFieldName]
		f2 := rec[p.fields[1].dbFieldName]
		f3 := rec[p.fields[2].dbFieldName]
		f4 := rec[p.fields[3].dbFieldName]
		f5 := rec[p.fields[4].dbFieldName]
		f6 := rec[p.fields[5].dbFieldName]
		f7 := rec[p.fields[6].dbFieldName]
		f8 := rec[p.fields[7].dbFieldName]
		f9 := rec[p.fields[8].dbFieldName]
		_, err = p.handle.Exec(p.insert, f1, f2, f3, f4, f5, f6, f7, f8, f9)
	case 10:
		f1 := rec[p.fields[0].dbFieldName]
		f2 := rec[p.fields[1].dbFieldName]
		f3 := rec[p.fields[2].dbFieldName]
		f4 := rec[p.fields[3].dbFieldName]
		f5 := rec[p.fields[4].dbFieldName]
		f6 := rec[p.fields[5].dbFieldName]
		f7 := rec[p.fields[6].dbFieldName]
		f8 := rec[p.fields[7].dbFieldName]
		f9 := rec[p.fields[8].dbFieldName]
		f10 := rec[p.fields[9].dbFieldName]
		_, err = p.handle.Exec(p.insert, f1, f2, f3, f4, f5, f6, f7, f8, f9, f10)
	}

	return err
}

func (p *DbProxy) ensureFieldsExist(rec map[string]interface{}) error {

	// iterate through all the field definitions to ensure they appear in the map
	for _, f := range p.fields {
		_, exists := rec[f.dbFieldName]
		if exists == false {
			return fmt.Errorf("field '%s' does not appear in the message", f.dbFieldName)
		}
	}

	return nil
}

//
// end of file
//
