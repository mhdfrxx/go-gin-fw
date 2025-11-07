package main

import (
	"go-gin-fw/database"
	"go-gin-fw/routes"
)

func main() {
	database.Connect()

	routes.Routes()
}
