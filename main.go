package main

import (
	"testcompany/config"
	"testcompany/database"
	"testcompany/routes"
)

func main() {
	config.Init()

	database.StartDatabase()
	s := routes.NewServer()

	s.Run()
}
