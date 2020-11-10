//Package users faz o cadastro e obtem a lista de usuários
package users

import (
	"github.com/gin-gonic/gin"
)

//UserBody contém o conteúdo que deverá ir no body da requisição
type UserBody struct {
	Username string `form:"username" json:"username" binding:"required"`
	Email    string `form:"email" json:"email" binding:"required"`
	Id       string `form:"id" json:"id" binding:"required"`
	Cpf      string `form:"cpf" json:"cpf" binding:"required"`
}

var List []UserBody

//PostUsers faz a requisação post para cadastro dos usuários
func PostUsers(c *gin.Context) {

	var body UserBody
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
		"email":    body.Email,
		"cpf":      body.Cpf,
	})
}

//GetAllUsers realiza uma get para retornar a lista de usuários cadastrados
func GetAllUsers(c *gin.Context) {
	c.JSON(200, gin.H{
		"Itens": List,
	})
}
