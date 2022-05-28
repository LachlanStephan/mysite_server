package main

import (
	blogs "mysite_server/controllers/blogs"
	sections "mysite_server/controllers/sections"

	"github.com/gin-gonic/gin"

	cors "github.com/rs/cors/wrapper/gin"

	"net/http"
)

// get sections from db/sections
func getSections(context *gin.Context) { // gin context is taken care of by gin engine -> used for handling http req and res
	sects := sections.GetSections()
	context.IndentedJSON(http.StatusOK, sects) // returns 200 & json of queried data
}

// get sections from db/sections
func getBlogs(context *gin.Context) { // gin context is taken care of by gin engine -> used for handling http req and res
	bl := blogs.GetBlogs()
	context.IndentedJSON(http.StatusOK, bl) // returns 200 & json of queried data
}

func main() {
	router := gin.Default()

	router.Use(cors.Default())
	router.GET("/sections", getSections)
	router.GET("/blogs", getBlogs)

	router.Run("localhost:8080")
}

/*

	gin will trust all proxies by default

		can SetTrustedProxies() to specify the correct client - worry about this later

	NEED to learn what proper err handling looks like - change all Println stmts

*/
