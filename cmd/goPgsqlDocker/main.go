package main

import (
	"github.com/dusanbre/goPgsqlDocker/database"
	"github.com/dusanbre/goPgsqlDocker/internal/routes"
)

func main() {
	r := routes.New()
	db := database.Connect()

	database.Migrate(db)

	routes.Routes(r)

	r.Logger.Fatal(r.Start(":8080"))
}
