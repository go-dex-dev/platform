package platform

import (
	"./generators"
)

func Start(path string) {
	entities := generators.GenerateEntities(path)
	generators.CreateDatabaseScript(path, entities...)
}
