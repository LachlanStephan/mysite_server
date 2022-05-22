package main

import (
	db "mysite_server/db/sections"

	"net/http"

	"github.com/gin-gonic/gin"
)

// move this to db/blog when bothered
type blog struct {
	ID      string
	TITLE   string
	DESC    string
	TIME    string
	CONTENT string
	TAGS    string
}

// get sections from db/sections
func getSections(s *gin.Context) { // gin context is taken care of by gin engine -> used for handling http Req and Res
	sects := db.GetSections()
	s.IndentedJSON(http.StatusOK, sects) // returns 200 & json of queried data
}

func main() {
	router := gin.Default()
	router.GET("/sections", getSections)
	router.Run("localhost:8080")
}

/*

	gin will trust all proxies by default

		can SetTrustedProxies() to specify the correct client - worry about this later

*/
