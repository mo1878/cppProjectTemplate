package core

import (
	"log"
	"os"
	"path/filepath"
)

type project struct {
	name []string
}

type subDirectory struct {
	name string
}

type templateFile struct {
	fileName string 
	subDir subDirectory
}

func (p *project) newProjectDirectory() string {

	err := os.Mkdir(p.name[0], 0777)
	if err != nil {
		log.Fatal(err)
	}

	return p.name[0]
}

func (sD *subDirectory) newSubdirectory(parentPath string) []string {

	subDirNames := []string{"src", "headers", ".vscode"}

	subDirPathSlice := []string{}
	for _, dirName := range subDirNames {
		subDirectoryPath := filepath.Join(parentPath, dirName)
		err := os.Mkdir(subDirectoryPath, 0777)
		if err != nil {
			log.Fatalf("Failed to create subdirectory %s: %v", subDirectoryPath, err)
		}
		subDirPathSlice = append(subDirPathSlice, subDirectoryPath)
	}

	return subDirPathSlice
}


func (f *)
