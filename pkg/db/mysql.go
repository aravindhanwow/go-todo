package db

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

type DBConnect interface {
	NewDBConnection() (*sql.DB, error)
}

type dbConn struct {
}

func NewMysqlConnect() DBConnect {
	return &dbConn{}
}

func (d dbConn) NewDBConnection() (*sql.DB, error) {
	dsn := "root:root@tcp(127.0.0.1:3306)/todoapp"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	fmt.Println("Connected to the database!")
	return db, nil
}
