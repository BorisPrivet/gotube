package main

import (
	"fmt"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.LoadHTMLGlob("frontend/templates/*")
	// r.Static("frontend/static", "")

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title":   "GoTube",
			"message": "Paste youtube url here to download a song",
		})
	})

	r.POST("/submit", func(c *gin.Context) {
		utubeURL := c.PostForm("utubeURL")
		stream, title, err := DownloadFromUrl(utubeURL)

		if err != nil {
			c.String(http.StatusInternalServerError, "Ошибка: %s", err)
			return
		}

		defer stream.Close()

		filename := fmt.Sprintf("attachment; filename=\"%s.mp3\"", title)

		c.Header("Content-Disposition", filename)
		c.Header("Content-Type", "audio/mpeg")

		io.Copy(c.Writer, stream)
	})

	r.Run(":8080")
}
