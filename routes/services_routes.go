package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func ChartWs(c *gin.Context) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Println("Error upgrading to websocket:", err)
		return
	}
	defer conn.Close()

	// Simulate sending data to the client every second
	for {
		data := map[string]float64{
			"timestamp": float64(time.Now().Unix()),
			"value":     float64(10 + time.Now().Unix()%20), // Example dynamic value
		}

		if err := conn.WriteJSON(data); err != nil {
			log.Println("Error writing JSON:", err)
			break
		}

		time.Sleep(1 * time.Second)
	}
}

func Action(c *gin.Context) {
	// var body struct {
	// 	Email string `json:"email"`
	// }

	// if err := c.ShouldBindJSON(&body); err != nil {
	// 	log.Println("Error binding JSON:", err) // Print the error for debugging
	// 	c.JSON(400, gin.H{
	// 		"message": "Unable to bind input data",
	// 		"error":   err.Error(),
	// 	})
	// 	return
	// }

	// log.Println(body.Email)

	c.HTML(200, "button.html", gin.H{})
}

func POOST(c *gin.Context) {
	var body struct {
		Email string `json:"email"`
	}

	// if err := c.ShouldBindJSON(&body); err != nil {
	// 	log.Println("Error binding JSON:", err) // Print the error for debugging
	// 	c.JSON(400, gin.H{
	// 		"message": "Unable to bind input data",
	// 		"error":   err.Error(),
	// 	})
	// 	return
	// }
	if err := c.ShouldBindJSON(&body); err != nil {
		var e *json.UnmarshalTypeError
		if errors.As(err, &e) {
			// If the error arises from an invalid character
			c.JSON(400, gin.H{
				"message": fmt.Sprintf("Invalid character "),
				"error":   err.Error(),
			})
			log.Println(body.Email)
			return
		}
		// Handle other binding errors
		c.JSON(400, gin.H{
			"message": "Unable to bind input data",
			"error":   err.Error(),
		})
		return
	}

	log.Println(body.Email)

}

func Te(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "okay",
	})
}
