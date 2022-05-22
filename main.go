package main

import (
	"mysite_server/db"

  "fmt"

// "net/http"

// "github.com/gin-gonic/gin"

)

type blog struct {
    ID string 
    TITLE string 
    DESC string 
    TIME string
    CONTENT string 
    TAGS string 
}


// func getSections(s *gin.Context) {
// s.IndentedJSON(http.StatusOK, db.GetSections)
// }

func getSections() {
  x := db.GetSections()
  fmt.Println(x, "heeeey")
}

func main() {
  getSections()
// router := gin.Default()
// router.GET("/sections", getSections)

// router.Run("localhost:8080")
}
