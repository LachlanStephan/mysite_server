package blogs

import (
	"fmt"
	"mysite_server/db"
)

type Blog struct {
	ID      int
	Title   string
	Desc    string
	Content string
}

type NewBlog struct {
	Title       string
	Description string
	Content     string
}

func GetBlogs() []Blog {
	sql := "SELECT blog_id, content, title, description FROM blog ORDER BY blog_id DESC"

	// open db conn
	db.Connect()

	rows, err := db.Conn.Query(sql)

	if err != nil {
		fmt.Println("oops")
	}

	// manually close conn
	db.Close()

	var blogData []Blog

	for rows.Next() {
		b := new(Blog)

		err := rows.Scan(&b.ID, &b.Content, &b.Title, &b.Desc)

		if err != nil {
			fmt.Println("oops again")
		}

		blogData = append(blogData, *b)
	}

	err = rows.Err()
	if err != nil {
		fmt.Println("fuck")
	}

	return blogData
}

func PostBlogs(data []NewBlog) {
	// insert data
	fmt.Println(data)
}
