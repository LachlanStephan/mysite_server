package main

import (
	"mysite_server/db"

  "net/http"

  "github.com/gin-gonic/gin"

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

// func getSections(s *gin.Context) {
// s.IndentedJSON(http.StatusOK, db.GetSections)
// }

func getSections() {
  db.GetSections
}

func main() {
  router := gin.Default()
  router.GET("/sections", getSections)

  router.Run("localhost:8080")
}
