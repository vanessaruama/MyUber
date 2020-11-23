//Package drivers faz o cadastro e obtem a lista de motoristas
package drivers

import (
	"math/rand"
	"strconv"

	"github.com/gin-gonic/gin"
)

//DriverBody contém o conteúdo que deverá ir no body da requisição
type DriverBody struct {
	Username string `form:"username" json:"username" binding:"required"`
	Cpf      string `form:"cpf" json:"cpf" binding:"required"`
	Car      string `form:"car" json:"car" binding:"required"`
	Id       int    `form:"id" json:"id"`
}

//UpdateBody contém o conteúdo para atualizar a informação de um motorista cadastrado
type UpdateBody struct {
	Username string `form:"username" json:"username" binding:"required"`
	Car      string `form:"car" json:"car" binding:"required"`
}

//List contém a lista dos motoristas cadastrados
var List []DriverBody

//PostDrivers faz requisição para adicionar um novo motorista
func PostDrivers(c *gin.Context) {

	var body DriverBody
	var parsererror = c.BindJSON(&body)
	body.Id = rand.Intn(1000) //uuid.New().String()

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

//DeleteDrivers realiza a exclusão de um motorista cadastrado
func DeleteDrivers(c *gin.Context) {

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

//GetDriver realiza a busca de um motorista cadastrado de acordo com seu Id
func GetDriver(c *gin.Context) {

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
		"car":      List[getposition].Car,
	})
}

//UpdateDriver realiza a alteração do cadastro de um motorista
func UpdateDriver(c *gin.Context) {

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
	List[getposition].Car = body.Car

	c.JSON(200, gin.H{
		"id":       List[getposition].Id,
		"username": List[getposition].Username,
		"cpf":      List[getposition].Cpf,
		"car":      List[getposition].Car,
	})
}

//GetAllDrivers realiza uma get para retornar a lista de usuários cadastrados
func GetAllDrivers(c *gin.Context) {
	c.JSON(200, gin.H{
		"Itens": List,
	})
}
