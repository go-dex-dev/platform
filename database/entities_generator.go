package database

import (
	"bufio"
	"encoding/json"
	"fmt"
	"github.com/gobeam/stringy"
	"io/ioutil"
	"log"
	"os"
)

const BuildPath = "\\generated"
const BuildEntitiesPath = "\\generated\\domain\\entities"
const SrcEntitiesDefinitionPath = "\\resources\\domain\\entities"

var projectRootDir string

func GenerateEntities(rootDir string) []string {
	projectRootDir = rootDir
	if !createBuildFolders(rootDir) {
		return nil
	}

	return createEntities(rootDir)
}

func createEntities(rootDir string) []string {
	dir, _ := ioutil.ReadDir(rootDir + SrcEntitiesDefinitionPath)
	var entities []string
	for _, fileName := range dir {
		file, _ := ioutil.ReadFile(rootDir + SrcEntitiesDefinitionPath + "\\" + fileName.Name())
		entities = append(entities, createEntityDefinitionFile(file))
	}
	return entities
}

func createBuildFolders(rootDir string) bool {
	var err = os.RemoveAll(projectRootDir + BuildPath)
	if err != nil {
		log.Fatalf("failed removing current build folder: %s", err)
		return false
	}

	err = os.MkdirAll(rootDir+BuildEntitiesPath, os.FileMode(0755))
	if err != nil {
		log.Fatalf("failed creating current build folder and subFolders: %s", err)
		return false
	}
	return true
}

func createEntityDefinitionFile(file []byte) string {

	data := loadEntityDefinition(file)

	entityFile, err := createNewFile(data)
	if err != nil {
		return ""
	}

	dataWriter := bufio.NewWriter(entityFile)

	err = writeHeaders(data, dataWriter)
	err = writeFields(data, dataWriter)
	//err = writeRelations(data, dataWriter)
	err = writeFooter(dataWriter)
	err = writeConstructor(data, dataWriter)

	dataWriter.Flush()
	entityFile.Close()

	if err != nil {
		return ""
	}
	return data.Name
}

func createNewFile(data Entity) (*os.File, error) {
	newFileName := projectRootDir + BuildEntitiesPath + "\\" + data.Name + "Entity.go"
	entityFile, err := os.OpenFile(newFileName, os.O_CREATE, 0777)

	if err != nil {
		log.Fatalf("failed creating file: %s", err)
		return nil, nil
	}
	return entityFile, err
}

func loadEntityDefinition(file []byte) Entity {
	data := Entity{}

	_ = json.Unmarshal([]byte(file), &data)
	return data
}

func writeFooter(dataWriter *bufio.Writer) error {
	var err error
	// End of struct
	_, err = dataWriter.WriteString("} \n")
	return err
}

func writeRelations(data Entity, dataWriter *bufio.Writer) error {
	var err error
	for i := 0; i < len(data.Relations); i++ {
		fmt.Println("    |---| Name: ", data.Relations[i].Name)
		fmt.Println("        |-- Type: ", data.Relations[i].Type)
		fmt.Println("        |-- Relation: ", data.Relations[i].Relation)

		if data.Relations[i].Relation == "one2many" {
			//	_, err = dataWriter.WriteString("[]")
		} else if data.Relations[i].Relation == "many2one" {
			_, err = dataWriter.WriteString("\t" + toCamelCase(data.Relations[i].Name) + " ")
			_, err = dataWriter.WriteString("")
			_, err = dataWriter.WriteString(toCamelCase(data.Relations[i].Type) + "Entity \n")
		}
	}
	return err
}

func writeFields(data Entity, dataWriter *bufio.Writer) error {
	var err error
	fmt.Println("|---| Fields: ", data.Name)
	for i := 0; i < len(data.Fields); i++ {
		fmt.Println("    |---| Name: ", data.Fields[i].Name)
		fmt.Println("        |-- Type: ", data.Fields[i].Type)
		_, err = dataWriter.WriteString("\t" + toCamelCase(data.Fields[i].Name) + " " + data.Fields[i].Type + "\n")
	}
	return err
}

func writeHeaders(data Entity, dataWriter *bufio.Writer) error {
	var err error
	const fileHeader = "package entities  \n\n" +
		"import \"gorm.io/gorm\"  \n\n"
	const structHeader = "type %sEntity struct { \n" +
		"\tgorm.Model\n"

	fmt.Println("Entity Name: ", data.Name)
	_, err = dataWriter.WriteString(fileHeader)
	_, err = dataWriter.WriteString(fmt.Sprintf(structHeader, data.Name))
	return err
}

func writeConstructor(data Entity, dataWriter *bufio.Writer) error {
	var err error
	_, err = dataWriter.WriteString("\n")

	var parameters string
	for i, field := range data.Fields {
		if i != 0 {
			parameters += " "
		}
		parameters += field.Name + " " + field.Type
		if i < len(data.Fields)-1 {
			parameters += ","
		}
	}
	var constructorFunction = "func New" + data.Name + "Entity(" + parameters + ") *" + data.Name + "Entity {"
	_, err = dataWriter.WriteString(constructorFunction)

	_, err = dataWriter.WriteString("\n\t return &" + data.Name + "Entity{")
	var attributes string
	for i, field := range data.Fields {
		if i != 0 {
			attributes += "\n"
		}
		fieldName := field.Name
		attributes = "\n\t\t" + toCamelCase(fieldName) + ": " + fieldName + ","
		_, err = dataWriter.WriteString("\t" + attributes)
	}
	_, err = dataWriter.WriteString("\n\t}")

	_, err = dataWriter.WriteString("\n}")

	return err
}

func toCamelCase(fieldName string) string {
	return stringy.New(fieldName).CamelCase()
}
