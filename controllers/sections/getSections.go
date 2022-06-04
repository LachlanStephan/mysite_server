package sections

import (
	"fmt"

	db "mysite_server/db"
)

// like an interface
type Section struct {
	ID    int
	Title string
	Desc  string
}

func GetSections() []Section {
	// open connection from db package
	db.Connect()

	// get all sections
	rows, err := db.Conn.Query("SELECT section_id, title, description FROM section")

	// something goes wrong with sql query
	if err != nil {
		fmt.Println("broke")
	}

	fmt.Println("got em")

	// close connection -> automatically done below but a good practice to manually close
	db.Close()

	// create array based on struct above
	var data []Section

	// loop queried rows
	for rows.Next() {

		// assign a variable to each row based on struct
		s := new(Section)

		err := rows.Scan(&s.ID, &s.Title, &s.Desc)

		if err != nil {
			fmt.Println("broke 2")
		}

		// push data
		data = append(data, *s)
	}

	err = rows.Err()
	if err != nil {
		fmt.Println("errrrrrr")
	}

	// return array of sections
	return data

}
