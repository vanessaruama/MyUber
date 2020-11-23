package main

//criar packages separados, ex motorista, usuários

import (
	"github.com/vanessaruama/myuber/routes"
)

func main() {

	router := routes.SetUpRouter()

	router.Run(":3000") // listen and serve on 0.0.0.0:3000 (for windows "localhost:3000")
}

//PostForm exemplo de requisição com formdata
/*func PostForm(c *gin.Context) {
	id := c.PostForm("id")
	username := c.PostForm("username")
	email := c.PostForm("email")

	c.JSON(200, gin.H{
		"status":   "posted",
		"id":       id,
		"username": username,
		"email":    email,
	})
}*/
