package passengers

import (
	"fmt"

	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
)

type PassengersDb struct {
	Db *pg.DB
}

func NewPassengersDb(DbContext *pg.DB) *PassengersDb {
	instance := &PassengersDb{}
	instance.Db = DbContext
	return instance
}

type PassengersModel struct {
	Id    int
	Name  string
	Cpf   string
	Email string
}

func (p *PassengersDb) CreateTab() {

	model := &PassengersModel{}
	err := p.Db.Model(model).CreateTable(&orm.CreateTableOptions{
		Temp:        false,
		IfNotExists: true,
	})

	if err != nil {
		fmt.Println(err)
	}
}

func (p *PassengersDb) AddPassenger(data *PassengersModel) bool {
	p.Db.Model(data).Insert()
	return true
}
