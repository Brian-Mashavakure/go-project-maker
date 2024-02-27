package varsrepo

var MainWriteString string = `
package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/ping", pingHandler)

	router.Run("localhost:8080")

}

func pingHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
})
}`

var HandlersWriteString string = `
//Welcome Gopher!!!!!!!!!

//Move All Route Handlers Here For Better Code Structure
`
