package main

import (
	"io/ioutil"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

var pass string = "Thu, 16 Sep 2021 06:30:49 UTC"
var imgData string

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET")
		c.Writer.Header().Set("Access-Control-Request-Method", "*")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	}
}

func setupRouter() *gin.Engine {
	// Disable Console Color
	// gin.DisableConsoleColor()
	r := gin.Default()
	r.Use(Cors())

	r.GET("/quiz/submit", submit)

	return r
}

func checkPasswordBCrypt(input string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(input), []byte(pass))
	if err != nil {
		return false
	} else {
		return true
	}
}

func submit(c *gin.Context) {
	input := c.Query("code")
	var base64data string = ""
	if !checkPasswordBCrypt(input) {
		base64data = ""
	} else {
		base64data = imgData
	}

	c.JSON(200, gin.H{"result": checkPasswordBCrypt(input), "data": base64data})
}

func main() {
	data, err := ioutil.ReadFile("imgData")
	if err != nil {
		panic(err)
	}
	imgData = string(data)
	r := setupRouter()
	r.Run(":8080")
}
