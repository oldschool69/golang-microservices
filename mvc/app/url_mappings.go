package app

import (
	"golang-microservices/mvc/controllers"
)

//Cria as "routes" da API
func mapUrls() {
	router.GET("/users/:user_id", controllers.GetUser)
}