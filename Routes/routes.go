//Package routes contém todas as rotas das requisições
package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/vanessaruama/myuber/drivers"
	"github.com/vanessaruama/myuber/namerequest"
	"github.com/vanessaruama/myuber/users"
)

//SetUpRouter obtém as rotas das requisições
func SetUpRouter() *gin.Engine {
	router := gin.Default()

	/*------------------ Router Teste --------------------------------------
	----------------------------------------------------------------------*/
	// This handler will match /user/john but will not match /user/ or /user
	router.GET("/user/:name", namerequest.NameRequest)

	//router.POST("/form_post", PostForm)

	/*------------------- Router Users -------------------------------------
	----------------------------------------------------------------------*/
	router.POST("/v2/users", users.PostUsers)

	router.GET("/v2/users", users.GetAllUsers)

	router.GET("/v2/users/:id", users.GetUser)

	router.DELETE("/v2/users/:id", users.DeleteUsers)

	router.PUT("/v2/users/:id", users.UpdateUser)

	/*------------------- Router Drivers -----------------------------------
	----------------------------------------------------------------------*/
	router.POST("/v2/drivers", drivers.PostDrivers)

	router.GET("/v2/drivers", drivers.GetAllDrivers) //recupera todos os motoristas

	router.GET("/v2/drivers/:id", drivers.GetDriver) //recupera um motorista especifico pelo ID

	router.DELETE("/v2/drivers/:id", drivers.DeleteDrivers) //deleta um motorista pelo ID

	router.PUT("/v2/drivers/:id", drivers.UpdateDriver) //deleta um motorista pelo ID

	return router
}
