package database

import (
	"database/sql"
	"github.com/go-sql-driver/mysql"
	"sync"
)

var (
	db *sql.DB
	once sync.Once
	dbErr error
)

func Connect() (*sql.DB, error) {
	once.Do(func() {
		var err error
		cfg := mysql.Config{
			User:   "root",
			Passwd: "!23qwe",
			Net:    "tcp",
			Addr:   "127.0.0.1:3306",
			DBName: "recordings",
		}

		db, err = sql.Open("mysql", cfg.FormatDSN())
		if err != nil {
			dbErr = err
		}
		err = db.Ping()
		if err != nil {
			dbErr = err
		}
	})

	if dbErr != nil {
		return nil, dbErr
	}	
	return db, nil
}
