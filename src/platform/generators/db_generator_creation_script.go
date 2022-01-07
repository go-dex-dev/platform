package generators

import (
	"io/ioutil"
	"log"
	"os"
	"strings"
)

const BuildDbPath = "\\build\\db"

var lines []string

func CreateDatabaseScript(projectRootDir string, entities ...string) error {
	createDbBuildFolders(projectRootDir)
	currentFileName := "E:\\Go\\godex-platform\\src\\platform\\generators\\db_generator.go"
	newFileName := projectRootDir + BuildDbPath + "\\db_generator.go"

	err := loadFile(currentFileName)
	if err != nil {
		return err
	}
	setPackage()
	copyTemplate(entities...)

	err = writeNewFile(newFileName)
	if err != nil {
		return err
	}

	return nil
}

func createDbBuildFolders(rootDir string) bool {
	var err = os.RemoveAll(projectRootDir + BuildDbPath)
	if err != nil {
		log.Fatalf("failed removing current build folder: %s", err)
		return false
	}

	err = os.MkdirAll(rootDir+BuildDbPath, os.FileMode(0755))
	if err != nil {
		log.Fatalf("failed creating current build folder and subFolders: %s", err)
		return false
	}
	return true
}

func copyTemplate(entities ...string) {
	for i, line := range lines {
		if strings.Contains(line, "//--Entities Import--//") {
			lines[i] = "\t\"../domain/entities\""
		}

		if strings.Contains(line, "//--Entities Hook--//") {
			var newLine string
			for _, entity := range entities {
				newLine = newLine + "\t\t&entities." + entity + "Entity{},\n"
			}
			lines[i] = newLine + "\t)"
			lines[i+1] = "" // Remove already existing ´)´ character
		}
	}
}

func writeNewFile(destination string) error {
	output := strings.Join(lines, "\n")
	return ioutil.WriteFile(destination, []byte(output), 0777)
}

func loadFile(source string) error {
	var data, err1 = ioutil.ReadFile(source)
	if err1 != nil {
		return err1
	}
	lines = strings.Split(string(data), "\n")
	return nil
}

func setPackage() {
	lines[0] = "package db"
}

/*
func copyTemplate(source, destination string) error {
	var err = filepath.Walk(source, func(path string, info os.FileInfo, err error) error {
		var relPath = strings.Replace(path, source, "", 1)
		if relPath == "" {
			return nil
		}
		if info.IsDir() {
			return os.Mkdir(filepath.Join(destination, relPath), 0755)
		} else {
			var data, err1 = ioutil.ReadFile(filepath.Join(source, relPath))
			if err1 != nil {
				return err1
			}
			return ioutil.WriteFile(filepath.Join(destination, relPath), data, 0777)
		}
	})
	return err
}
*/
