package main

import (
	"fmt"
	"github.com/go-dex-dev/platform/database"
)

func main() {
	path := "E:\\Go\\src\\github.com\\go-dex-dev\\platform\\example\\src\\project"
	fmt.Println("Read project definition from", path)
	fmt.Println("Generate ORM for entities")
	entities := database.GenerateEntities(path)
	fmt.Println("Generate database")
	database.CreateDatabaseScript(path, entities...)
}
