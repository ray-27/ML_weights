package handler

import (
	"github.com/gin-gonic/gin"
)

func Routes(r *gin.Engine) {
	//html pages
	r.GET("/", Home)
	r.GET("/signup", SignUP_page)
	r.GET("/login", Login_page)

	r.GET("/projects", Projects_page)
	r.GET("/run", Run_page)

	api := r.Group("/api")
	{
		api.GET("/signup", SignUp)
	}

	html_render := r.Group("/html")
	{
		html_render.GET("/log", Html_Log)
		html_render.GET("/charts", Html_Run_Charts)
		html_render.GET("/inference", Html_Run_Inference)
	}

	r.GET("/action", Action)
	r.POST("/poost", POOST)

	r.GET("/ws", ChartWs)

	// Serve the chart component dynamically
	r.GET("/chart-component.html", func(c *gin.Context) {
		c.HTML(200, "chart-component.html", nil)
	})
}
