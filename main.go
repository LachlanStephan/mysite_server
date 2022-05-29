package main

import (
	"fmt"
	blogs "mysite_server/controllers/blogs"
	sections "mysite_server/controllers/sections"

	"github.com/gin-gonic/gin"

	cors "github.com/rs/cors/wrapper/gin"

	"net/http"
)

// get sections from db/sections
func getSections(c *gin.Context) { // gin context is taken care of by gin engine -> used for handling http req and res
	sects := sections.GetSections()
	c.IndentedJSON(http.StatusOK, sects) // returns 200 & json of queried data
}

// get sections from db/sections
func getBlogs(c *gin.Context) { // gin context is taken care of by gin engine -> used for handling http req and res
	bl := blogs.GetBlogs()
	c.IndentedJSON(http.StatusOK, bl) // returns 200 & json of queried data
}

// wip
type NewBlog struct {
	Content     string `json:"Content"`
	Description string `json:"Description"`
	Title       string `json:"Title"`
}

func postBlogs(c *gin.Context) {
	var blog []NewBlog

	err := c.BindJSON(&blog)
	if err != nil {
		fmt.Println(err, "oops")
	}

	fmt.Println(blog)

	// err := blogs.PostBlogs(blog)
	// if err != nil {
	// fmt.Println(err, "insert broken")
	// exit here also
	// }

	// c.(http.StatusOK)
}

func checkIfAmdin() {
	// this needs to be done first - before other cruds - finish newBlog later
}

func main() {
	router := gin.Default()

	router.Use(cors.Default())
	router.GET("/sections", getSections)
	router.GET("/blogs", getBlogs)
	router.POST("/newBlog", postBlogs)

	router.Run("localhost:8080")
}

/*

	gin will trust all proxies by default

		can SetTrustedProxies() to specify the correct client - worry about this later

	NEED to learn what proper err handling looks like - change all Println stmts

*/
