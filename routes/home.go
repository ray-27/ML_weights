package handler

import "github.com/gin-gonic/gin"

func Home(c *gin.Context) {
	c.HTML(200, "index.html", gin.H{
		"Head": "this is the heading from gin",
	})
}

// About handler
func About(c *gin.Context) {
	c.String(200, `<p>This is the "About" page content loaded via HTMX!</p>`)
}
