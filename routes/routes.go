package handler

import "github.com/gin-gonic/gin"

func Routes(r *gin.Engine) {
	r.GET("/", Home)
	r.GET("/about", About)
}
