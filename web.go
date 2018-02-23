package main

import (
	"github.com/alandwiprasetyo/rest-api/src/routes"
	"github.com/alandwiprasetyo/rest-api/src/models/migrations"
	"github.com/alandwiprasetyo/rest-api/src/models/seeders"
)

func main() {
	migrations.Migration()
	seeders.Seeder()
	routes.GetRouter().Run(":9000")
}
