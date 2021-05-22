package passengers

import (
	"math/rand"

	"github.com/gin-gonic/gin"
)

type PassengersService struct {
	Db *PassengersDb
}

func NewPassengersService(Db *PassengersDb) *PassengersService {
	instance := &PassengersService{}
	instance.Db = Db
	return instance
}

type PassengerBody struct {
	Id    int    `form:"id" json:"id"`
	Name  string `form:"name" json:"name" binding:"required"`
	Cpf   string `form:"cpf" json:"cpf" binding:"required"`
	Email string `form:"email" json:"email" binding:"required"`
}

type UpdatePassengerBody struct {
	Username string `form:"username" json:"username" binding:"required"`
	Email    string `form:"email" json:"email" binding:"required"`
}

func (p *PassengersService) CreatePassengers(c *gin.Context) {

	var body PassengerBody
	var parsererror = c.BindJSON(&body)
	var validdriver bool
	body.Id = rand.Intn(1000) //uuid.New().String()

	if parsererror != nil {
		c.JSON(400, gin.H{
			"status":      "error",
			"description": parsererror.Error(),
		})
		return
	}

	passenger := &PassengersModel{}
	passenger.Id = body.Id
	passenger.Name = body.Name
	passenger.Email = body.Email
	passenger.Cpf = body.Cpf

	//Adiciona um novo motorista na lista
	validdriver = p.Db.AddPassenger(passenger)

	if !validdriver {
		c.JSON(400, gin.H{
			"status":      "error",
			"description": "falhou na criação do motorista",
		})
		return
	}

	c.JSON(200, gin.H{
		"status": "posted",
		"id":     body.Id,
		"name":   body.Name,
		"cpf":    body.Cpf,
		"email":  body.Email,
	})
}
