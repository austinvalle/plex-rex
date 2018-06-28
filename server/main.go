package main

import (
	"github.com/gin-gonic/gin"
	"github.com/moosebot/plex-rex/server/themoviedb"
)

func main() {
	r := gin.Default()
	r.GET("/search", func(c *gin.Context) {
		query := c.Query("q")
		if query == "" {
			c.JSON(400, gin.H{
				"message": "Query parameter 'q' is required",
			})
			return
		}

		c.JSON(200, themoviedb.Search(query))
	})
	r.Run()
}
