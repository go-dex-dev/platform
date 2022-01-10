package main

import (
	"fmt"
	"platform/database"
)

func main() {
	path := "E:\\Go\\projects\\go-dex-dev\\examples\\project"
	fmt.Println("Read project definition from", path)
	fmt.Println("Generate ORM for entities")
	entities := database.GenerateEntities(path)
	fmt.Println("Generate database")
	database.CreateDatabaseScript(path, entities...)
}
