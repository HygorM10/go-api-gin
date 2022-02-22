package main

import (
	"github.com/hygorm10/go-api-gin/database"
	"github.com/hygorm10/go-api-gin/routes"
)

func main() {
	database.ConectaComBancoDeDados()
	routes.HandleRequest()
}
