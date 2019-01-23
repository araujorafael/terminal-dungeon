package main

import (
	"math/rand"
	"net/http"
	"time"

	"terminal-adventure/objects"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/room", func(c *gin.Context) {
		// define a seed to a rand
		seed := rand.NewSource(time.Now().UnixNano())
		r := rand.New(seed)

		objectsList := objects.GetObjectsList()

		c.JSON(http.StatusOK, gin.H{
			"doors":   r.Intn(5) + 1,
			"objects": objectsList,
			"message": objects.GetMessage(objectsList),
		})
	})

	r.Run() // listen and serve on 0.0.0.0:8080
}
