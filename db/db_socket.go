package db

import (
	"database/sql"
	"fmt"
	"github.com/giskook/express_delivery/conf"
	_ "github.com/lib/pq"
)

type DBSocket struct {
	db   *sql.DB
	conf *conf.DB
}

func NewDBSocket(conf *conf.DB) (*DBSocket, error) {
	conn_string := fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s sslmode=disable", conf.User, conf.Passwd, conf.Host, conf.Port, conf.DbName)

	db, err := sql.Open("postgres", conn_string)

	if err != nil {
		return nil, err
	}

	return &DBSocket{
		db:   db,
		conf: conf,
	}, nil
}

func (db *DBSocket) Close() {
	db.db.Close()
}
