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
	name []string
}

type templateFile struct {
	fileName string
	subDir   subDirectory
}

func (p *project) newProjectDirectory() (string, error) {

	err := os.Mkdir(p.name[0], 0777)
	if err != nil {
		return "", err
	}

	return p.name[0], nil
}

func (sD *subDirectory) newSubDirectory(parentPath string) ([]string, error) {

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

	return subDirPathSlice, nil
}

func (f *templateFile) newFileCreation(subDirPaths []string) []string {

	subDirFiles := []string{"main.cpp", "header.h", "tasks.json", "launch.json"}
	subDirFileMap := make(map[string][]string)

	valueIndex := 0
	for _, key := range subDirPaths {

		if key == ".vscode" {
			subDirFileMap[key] = append(subDirFileMap[key], subDirFiles[valueIndex:valueIndex+2]...)
			valueIndex += 2
		} else {
			subDirFileMap[key] = append(subDirFileMap[key], subDirFiles[valueIndex])
			valueIndex++

		}

	}

	return subDirPaths
}
