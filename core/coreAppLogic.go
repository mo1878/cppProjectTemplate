package core

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

type project struct {
	projectName string
}

type subDirectory struct {
	name []string
}

type templateFile struct {
	fileName string
	subDir   subDirectory
}

type textCopy struct {
	sourceFile os.File
}

func (p *project) newProjectDirectory(projectName string) (string, error) {

	err := os.Mkdir(p.projectName, 0777)
	if err != nil {
		return "", err
	}

	return p.projectName, nil
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
	createdFiles := []string{}

	valueIndex := 0
	for _, key := range subDirPaths {

		if key == subDirPaths[2] {
			subDirFileMap[key] = append(subDirFileMap[key], subDirFiles[valueIndex:valueIndex+2]...)
			valueIndex += 2
			fmt.Printf("Map: %v ", subDirFileMap[key])
		} else {
			subDirFileMap[key] = append(subDirFileMap[key], subDirFiles[valueIndex])
			valueIndex++
		}

	}

	for dirPath, files := range subDirFileMap {

		for _, fileName := range files {
			fullPath := filepath.Join(dirPath, fileName)
			file, err := os.Create(fullPath)
			if err != nil {
				log.Printf("Failed to create %s: %v\n", fullPath, err)
				continue
			}
			defer file.Close()
			createdFiles = append(createdFiles, fullPath)
		}

	}

	return createdFiles
}

func (ct *textCopy) insertBoilerPlateCode(createdFilePaths []string) (bool, error) {

	boilerPlateTargetDir := "./boilerPlateCode"

	files, err := os.ReadDir(boilerPlateTargetDir)
	if err != nil {
		log.Printf("Error reading Directory %v", boilerPlateTargetDir)
	}

	for _, file := range files {

		fileName := file.Name()

		text, err := os.ReadFile(fileName)
		if err != nil {
			log.Printf("Failed to read file: %v", fileName)
		}

		for _, path := range createdFilePaths {

			err := os.WriteFile(path, text, 0777)
			if err != nil {
				log.Printf("Failed to write to: %v", path)
			}

		}

	}

	return true, nil
}
