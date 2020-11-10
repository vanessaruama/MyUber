package main

//criar packages separados, ex motorista, usuários

import (
	"github.com/gin-gonic/gin"
	"github.com/vanessaruama/myuber/drivers"
	"github.com/vanessaruama/myuber/namerequest"
	"github.com/vanessaruama/myuber/users"
)

func main() {
	router := gin.Default()

	// This handler will match /user/john but will not match /user/ or /user
	router.GET("/user/:name", namerequest.NameRequest)

	router.POST("/form_post", PostForm)

	router.POST("/v2/users", users.PostUsers)

	router.POST("/v2/drivers", drivers.PostDrivers)

	router.GET("/v2/alldrivers", drivers.GetAllDrivers)

	router.GET("/v2/allusers", users.GetAllUsers)

	router.Run(":3000") // listen and serve on 0.0.0.0:3000 (for windows "localhost:3000")
}

//PostForm exemplo de requisição com formdata
func PostForm(c *gin.Context) {
	id := c.PostForm("id")
	username := c.PostForm("username")
	email := c.PostForm("email")

	c.JSON(200, gin.H{
		"status":   "posted",
		"id":       id,
		"username": username,
		"email":    email,
	})
}
