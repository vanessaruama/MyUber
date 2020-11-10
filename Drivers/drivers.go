//Package drivers faz o cadastro e obtem a lista de motoristas
package drivers

import (
	"github.com/gin-gonic/gin"
)

//DriverBody contém o conteúdo que deverá ir no body da requisição
type DriverBody struct {
	Username string `form:"username" json:"username" binding:"required"`
	Cpf      string `form:"cpf" json:"cpf" binding:"required"`
	Car      string `form:"car" json:"car" binding:"required"`
	Id       string `form:"id" json:"id" binding:"required"`
}

var List []DriverBody

func PostDrivers(c *gin.Context) {

	var body DriverBody
	var parsererror = c.BindJSON(&body)

	if parsererror != nil {
		c.JSON(400, gin.H{
			"status":      "error",
			"description": parsererror.Error(),
		})
		return
	}

	//Lista
	List = append(List, body)

	c.JSON(200, gin.H{
		"status":   "posted",
		"id":       body.Id,
		"username": body.Username,
		"cpf":      body.Cpf,
		"car":      body.Car,
	})
}

//GetAllDrivers realiza uma get para retornar a lista de usuários cadastrados
func GetAllDrivers(c *gin.Context) {
	c.JSON(200, gin.H{
		"Itens": List,
	})
}
