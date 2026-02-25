package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.LoadHTMLGlob("frontend/templates/*")
	// r.Static("frontend/static", "")

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title":   "Main page",
			"message": "Privet blyat LOL",
		})
	})

	r.Run(":8080")
}
