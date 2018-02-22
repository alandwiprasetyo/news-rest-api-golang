package main

import (
	"github.com/alandwiprasetyo/rest-api/src/routes"
	"github.com/alandwiprasetyo/rest-api/src/models/migrations"
	"github.com/alandwiprasetyo/rest-api/src/models/seeders"
	"github.com/alandwiprasetyo/rest-api/src/common"
)

func main() {
	migrations.Migration()
	seeders.Seeder()
	routes.Setup()
	common.GetRouter().Run(":9000")
}
