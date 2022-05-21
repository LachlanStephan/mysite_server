package db

import (
	"database/sql"

	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var(
  title string
  description string
  sections []string
)


func GetSections() ([] error) {
  db, err := sql.Open("mysql", "root:@/mysite")
  if err != nil {
    panic(err)
  }
  fmt.Println("connected")

  // test select stmt
  rows, err := db.Query("select title, description from section")

  if err != nil {
	  fmt.Println("broke")
  }

  fmt.Println("got em")

  defer rows.Close()

// var sections []section

  for rows.Next() {
    var sect sections
    err := rows.Scan(&sect.title, &sect.description)
    if err != nil {
		  fmt.Println("broke 2")
    }
    fmt.Println(title, description)
    sections = append(sections, rows)
  }

  err = rows.Err()
  if err != nil {
	  fmt.Println("errrrrrr")
  }

  return sections

}
