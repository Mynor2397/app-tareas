package mysql

import (
	"database/sql"
	"log"
	"sync"

	_ "github.com/go-sql-driver/mysql" // Este es el driver para mysql
)

var (
	once sync.Once
	db   *sql.DB
	err  error
)

// Connect esta funcion nos permite conectar a mysql
func Connect() *sql.DB {
	user := "root"
	password := "system"
	server := "localhost"
	database := "task"

	once.Do(func() {

		db, err = sql.Open("mysql", user+":"+password+"@tcp("+server+")/"+database)
		if err != nil {
			log.Println(err.Error())
		}

	})

	return db
}
