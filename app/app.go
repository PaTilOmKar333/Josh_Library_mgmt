package app

import (
	"fmt"
	"project/config"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var (
	db *sqlx.DB
)

func Init() {

	err := initDB()
	if err != nil {
		panic(err)
	}
}

func initDB() (err error) {
	dbConfig := config.Database()

	db, err = sqlx.Open(dbConfig.Driver(), dbConfig.ConnectionURL())
	if err != nil {
		return
	}
	fmt.Println("database connected")
	if err = db.Ping(); err != nil {
		return
	}

	return
}

func GetDB() *sqlx.DB {
	return db
}

func Close() {
	db.Close()
}
