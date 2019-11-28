package database

import (
	"database/sql"
	"os"
	"sync"

	_ "github.com/go-sql-driver/mysql"
	"github.com/sirupsen/logrus"
)

var err error
var db *sql.DB
var once sync.Once

func GetDB() *sql.DB {
	once.Do(initMysql)
	return db
}

func initMysql() {
	db, err = sql.Open("mysql", "root:123456@tcp(database:3306)/leaf?charset=utf8mb4&parseTime=true&&loc=Asia%2FShanghai")
	if err != nil {
		logrus.Errorf("open database err:%s", err)
		os.Exit(1)
	}

	err = db.Ping()
	if err != nil {
		logrus.Errorf("ping database err:%s", err)
		os.Exit(1)
	}
}
