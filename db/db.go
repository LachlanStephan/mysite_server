package db

import (
	"database/sql"

	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var(
  title string
  description string
)


func GetSections() {
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

  output := make([]string, len(rows))
  for rows.Next() {
    err := rows.Scan(&title, &description)
    if err != nil {
		  fmt.Println("broke 2")
    }
    output.append(output, rows)
    fmt.Println(title, description)
  }

  err = rows.Err()
  if err != nil {
	  fmt.Println("errrrrrr")
  }


}
