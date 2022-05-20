package main

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"

	"mysite_server/db"
)

type blog struct {
    ID string 
    TITLE string 
    DESC string 
    TIME string
    CONTENT string 
    TAGS string 
}

type section struct {
    ID string 
    TITLE string 
    CONTENT string 
}

func sayHi() {
  msg := db.Hello("hey")
  fmt.Println(msg)
}

// routes 
func main() {
  sayHi();
}
