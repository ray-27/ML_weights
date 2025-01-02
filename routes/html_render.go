package handler

import "github.com/gin-gonic/gin"

func Html_Log(c *gin.Context) {
	c.HTML(200, "log.html", gin.H{})
}

func Html_Run_Charts(c *gin.Context) {
	c.HTML(200, "run_chart.html", gin.H{})
}

func Html_Run_Inference(c *gin.Context) {
	c.HTML(200, "run_inference.html", gin.H{})
}
