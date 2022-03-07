package db

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"simple-customer-service/config"
	"sync"
)

var once sync.Once
var dbInstance *sql.DB

func GetDatabaseConnection() *sql.DB {
	once.Do(func() {
		var err error
		log.Println("opendb")
		dsn := fmt.Sprintf(`%s:%s@tcp(%s:%s)/%s`,
			config.Env.CustomerDbUser, config.Env.CustomerDbPassword,
			config.Env.CustomerDbHost, config.Env.CustomerDbPort, config.Env.CustomerDbName)

		dbInstance, err = sql.Open("mysql", dsn)
		if err != nil {
			log.Fatal("error open db: ", err)
		}

		log.Println("ping db")
		err = dbInstance.Ping()
		if err != nil {
			log.Fatal("error ping db: ", err)
		}
	})

	return dbInstance
}

func OpenDatabase() {
	GetDatabaseConnection()
}
