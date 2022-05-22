package db

import (
	"database/sql"

	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var (
	Conn *sql.DB
)

func Connect() {
	var err error
	// connect to db -> this should be moved to its own file and imported
	Conn, err = sql.Open("mysql", "root:@/mysite")
	if err != nil {
		fmt.Println("shit", err)
	}
}

func Close() {
	Conn.Close()
}
