package main

import (
	// "os"

	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

var response = "Hello world"

func firstNamePost(firstName *gin.Context) {
	firstNames := firstName.PostForm("firstnames")
	fmt.Println("FirstName get from post: ", firstNames)
	//  curl -vvvv -XPOST  -d "firstnames=mamad" http://localhost/firstname
	// while true; do curl -XPOST  -d "firstnames=mamad" http://localhost/firstname && sleep 0.2 && clear ;done
}

func main() {
	server := gin.Default()
	server.GET("/hello", func(test *gin.Context) {
		test.String(http.StatusOK, "Hello world from string.")
	})

	server.POST("/firstname", firstNamePost)

	server.GET("/", func(r *gin.Context) {
		r.JSON(200, gin.H{
			"message": "Hellow World",
		})
	})

	server.Run(":80")
}
