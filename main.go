package main

import (
	"api-e-commerce/config"
	"api-e-commerce/controller/routes"
	"fmt"
)

func main() {
	app := routes.InitRoutes()
	config.LoadEnv()
	app.Listen(fmt.Sprintf(":%s", config.GetEnv("port_application")))
}
