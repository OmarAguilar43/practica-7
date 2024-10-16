package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type User struct {
	Name  string
	Email string
}

func main() {
	user := &User{
		Name:  "Jose",
		Email: "tilintilico@pilin.com",
	}

	users := []User{
		{
			Name:  "Omar",
			Email: "omarooo@pilin.com",
		},
		{
			Name:  "Miguel",
			Email: "Miguelon@pilin.com",
		},
	}
	router := gin.Default()
	fmt.Println("Running app")
	//Tomar archivos de la acarpeta templates
	router.LoadHTMLGlob("templates/*")
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title": "Main website",
			"user":  user,
			"users": users,
		})
	})
	router.Run(":8000") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
