package handler

import (
	"fmt"
	"html/template"

	"github.com/gin-gonic/gin"
)

func Home(c *gin.Context) {
	c.HTML(200, "index.html", gin.H{
		"Head":        "this is the heading from gin",
		"HTMLContent": template.HTML(fmt.Sprintf(Chart_component, "one", "one")),
	})
}

func SignUP_page(c *gin.Context) {
	c.HTML(200, "signup.html", gin.H{})
}

func Login_page(c *gin.Context) {
	c.HTML(200, "login.html", gin.H{})
}

func Run_page(c *gin.Context) {
	c.HTML(200, "run_page.html", gin.H{})
}
func Projects_page(c *gin.Context) {
	c.HTML(200, "projects.html", gin.H{})
}

func Action_page(c *gin.Context) {
	c.HTML(200, "action.html", gin.H{
		"Head": "this is the heading from gin",
	})
}

// About handler
func About(c *gin.Context) {
	c.String(200, `<p>This is the "About" page content loaded via HTMX!</p>`)
}
