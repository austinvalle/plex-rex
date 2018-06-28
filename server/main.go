package main

import (
	"github.com/gin-gonic/gin"
	"github.com/moosebot/plex-rek/server/internal/api"
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

		c.JSON(200, api.Search(query))
	})
	r.Run()
}
