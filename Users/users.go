//Package users faz o cadastro e obtem a lista de usuários
package users

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

//UserBody contém o conteúdo que deverá ir no body da requisição
type UserBody struct {
	Username string `form:"username" json:"username" binding:"required"`
	Email    string `form:"email" json:"email" binding:"required"`
	Id       int    `form:"id" json:"id" binding:"required"`
	Cpf      string `form:"cpf" json:"cpf" binding:"required"`
}

//UpdateBody contém o conteúdo para atualizar a informação de um usuário cadastrado
type UpdateBody struct {
	Username string `form:"username" json:"username" binding:"required"`
	Email    string `form:"email" json:"email" binding:"required"`
}

//List contém a lista dos usuários cadastrados
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

//DeleteUsers realiza a exclusão de um usuário cadastrado
func DeleteUsers(c *gin.Context) {

	var id int
	id, _ = strconv.Atoi(c.Param("id"))

	delposition := -1
	for i := range List {
		if List[i].Id == id {
			delposition = i
			break
		}
	}

	if delposition == -1 {
		c.JSON(200, gin.H{
			"status": "deleted",
			"id":     id,
		})
	}

	// Remove the element at index i from a.
	copy(List[delposition:], List[delposition+1:]) // Shift a[i+1:] left one index.
	//List[len(List)-1] =                         // Erase last element (write zero value).
	List = List[:len(List)-1] // Truncate slice.

	c.JSON(200, gin.H{
		"status": "deleted",
		"id":     id,
	})
}

//GetUser realiza a pesquisa de um usuário cadastrado de acordo com seu Id
func GetUser(c *gin.Context) {

	var id int
	id, _ = strconv.Atoi(c.Param("id"))

	getposition := -1
	for i := range List {
		if List[i].Id == id {
			getposition = i
			break
		}
	}

	if getposition == -1 {
		c.JSON(404, gin.H{
			"status": "not found",
			"id":     id,
		})
	}

	c.JSON(200, gin.H{
		"id":       List[getposition].Id,
		"username": List[getposition].Username,
		"cpf":      List[getposition].Cpf,
		"email":    List[getposition].Email,
	})
}

//UpdateUser realiza atualização/alteração de um usuário existente
func UpdateUser(c *gin.Context) {

	var id int
	var body UpdateBody
	id, _ = strconv.Atoi(c.Param("id"))

	getposition := -1
	for i := range List {
		if List[i].Id == id {
			getposition = i
			break
		}
	}

	if getposition == -1 {
		c.JSON(404, gin.H{
			"status": "not found",
			"id":     id,
		})
	}

	var parsererror = c.BindJSON(&body)

	if parsererror != nil {
		c.JSON(400, gin.H{
			"status":      "error",
			"description": parsererror.Error(),
		})
		return
	}

	List[getposition].Username = body.Username
	List[getposition].Email = body.Email

	c.JSON(200, gin.H{
		"id":       List[getposition].Id,
		"username": List[getposition].Username,
		"cpf":      List[getposition].Cpf,
		"email":    List[getposition].Email,
	})
}

//GetAllUsers realiza uma get para retornar a lista de usuários cadastrados
func GetAllUsers(c *gin.Context) {
	c.JSON(200, gin.H{
		"Itens": List,
	})
}
