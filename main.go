package main

import (
	"api-e-commerce/config"
	"api-e-commerce/controller/routes"
	"fmt"
)

//func main() {
//	app := routes.InitRoutes()
//	config.LoadEnv()
//	app.Listen(fmt.Sprintf(":%s", config.GetEnv("port_application")))
//}

func main() {
	app := routes.InitRoutes()
	config.LoadEnv()
	port := config.GetEnv("port_application")
	if err := app.Listen(fmt.Sprintf(":%s", port)); err != nil {
		fmt.Printf("Erro ao iniciar o servidor: %v\n", err)
	}
}
