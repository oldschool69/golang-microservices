package app

import (
	"golang-microservices/src/api/controllers/polo"
	"golang-microservices/src/api/controllers/repositories"
)

// Rotas
func mapUrls() {
	// Para fazer deploy do serviço na AWS, é preciso uma rota GET para indicar
	// que o serviço está gerando tráfego
	router.GET("/marco", polo.Marco)
	router.POST("/repositories", repositories.CreateRepo)
}
