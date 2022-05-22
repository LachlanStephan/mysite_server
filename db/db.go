package db

import (
	"database/sql"

	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type Sections struct {
  Section_id int
  Title string
  Content string 
}


func GetSections() []Sections {
  db, err := sql.Open("mysql", "root:@/mysite")
  if err != nil {
    panic(err)
  }

  fmt.Println("connected")

  // test select stmt
  rows, err := db.Query("SELECT section_id, title, description FROM section")

  if err != nil {
	  fmt.Println("broke")
  }

  fmt.Println("got em")

  defer rows.Close()

  var data []Sections

  for rows.Next() {

    s := new(Sections)
    
    err := rows.Scan(&s.Section_id, &s.Title, &s.Content)

    if err != nil {
		  fmt.Println("broke 2")
    }

    data = append(data, *s)
      
    fmt.Println(s)

  }

  err = rows.Err()
  if err != nil {
	  fmt.Println("errrrrrr")
  }

  return data

}
