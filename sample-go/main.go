package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gopkg.in/Graylog2/go-gelf.v2/gelf"
)

func main() {
	r := gin.Default()

	writer, err := gelf.NewUDPWriter("localhost:12201") // replace with your graylog server address
	if err != nil {
		panic(err)
	}

	r.POST("/sample/test", func(c *gin.Context) {
		var json struct {
			Message string `json:"message" binding:"required"`
		}

		if err := c.ShouldBindJSON(&json); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		msg := gelf.Message{
			Version:  "1.1",
			Host:     "sample-go", // replace with your host
			Short:    json.Message,
			Full:     "",
			Level:    6, // info level
			Facility: "test",
		}

		if err := writer.WriteMessage(&msg); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"status": "logged"})
	})

	if err := r.Run(); err != nil {
		panic(err)
	}
}
