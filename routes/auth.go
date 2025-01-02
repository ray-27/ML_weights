package handler

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SignUp(c *gin.Context) {
	type SignupRequest struct {
		Name     string `form:"name" binding:"required"`
		Email    string `form:"email" binding:"required,email"`
		Password string `form:"password" binding:"required,min=6"`
	}

	var signupData SignupRequest

	// Method 1: Parse form data into struct with automatic validation
	if err := c.ShouldBind(&signupData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// For demonstration, just sending back the received data
	c.JSON(http.StatusOK, gin.H{
		"message": "Signup successful",
		"user": gin.H{
			"name":  signupData.Name,
			"email": signupData.Email,
			// Don't send password back in real applications
		},
	})

	log.Printf("got a signup with name %s", signupData.Name)
}
