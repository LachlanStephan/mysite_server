package db

import (
	"database/sql"

	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

// like an interface
type Sections struct {
	ID      int
	Title   string
	Content string
}

func GetSections() []Sections {
	// connect to db -> this should be moved to its own file and imported
	db, err := sql.Open("mysql", "root:@/mysite")
	if err != nil {
		panic(err)
	}

	fmt.Println("connected")

	// get all sections
	rows, err := db.Query("SELECT section_id, title, description FROM section")

	// something goes wrong with sql query
	if err != nil {
		fmt.Println("broke")
	}

	fmt.Println("got em")

	// close connection -> automatically done below but a good practice to manually close
	defer rows.Close()

	// create array based on struct above
	var data []Sections

	// loop queried rows
	for rows.Next() {

		// assign a variable to each row based on struct
		s := new(Sections)

		err := rows.Scan(&s.ID, &s.Title, &s.Content)

		if err != nil {
			fmt.Println("broke 2")
		}

		// push data
		data = append(data, *s)
	}

	// not sure here -> maybe some redundant checking?
	err = rows.Err()
	if err != nil {
		fmt.Println("errrrrrr")
	}

	// return array of sections
	return data

}
