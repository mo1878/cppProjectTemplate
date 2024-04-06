package core

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

type Project struct {
	ProjectName string
}

type SubDirectory struct {
	name []string
}

type TemplateFile struct {
	file   string
	subDir SubDirectory
}

type TextCopy struct {
	sourceFile os.File
}

func (p *Project) NewProjectDirectory(ProjectName string) (string, error) {

	err := os.Mkdir(p.ProjectName, 0777)
	if err != nil {
		return "", err
	}

	return p.ProjectName, nil
}

func (sD *SubDirectory) NewSubDirectory(ParentPath string) ([]string, error) {

	subDirNames := []string{"src", "headers", ".vscode"}

	subDirPathSlice := []string{}
	for _, dirName := range subDirNames {
		subDirectoryPath := filepath.Join(ParentPath, dirName)
		err := os.Mkdir(subDirectoryPath, 0777)
		if err != nil {
			log.Fatalf("Failed to create subdirectory %s: %v", subDirectoryPath, err)
		}
		subDirPathSlice = append(subDirPathSlice, subDirectoryPath)
	}

	return subDirPathSlice, nil
}

func (f *TemplateFile) NewFileCreation(SubDirPaths []string) []string {

	subDirFiles := []string{"main.cpp", "headers.h", "tasks.json", "launch.json"}
	subDirFileMap := make(map[string][]string)
	createdFiles := []string{}

	valueIndex := 0
	for _, key := range SubDirPaths {

		if key == SubDirPaths[2] {
			subDirFileMap[key] = append(subDirFileMap[key], subDirFiles[valueIndex:valueIndex+2]...)
			valueIndex += 2
			fmt.Printf("Map: %v ", subDirFileMap[key])
		} else {
			subDirFileMap[key] = append(subDirFileMap[key], subDirFiles[valueIndex])
			valueIndex++
		}

	}

	for dirPath, files := range subDirFileMap {

		for _, file := range files {
			fullPath := filepath.Join(dirPath, file)
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

func (ct *TextCopy) InsertBoilerPlateCode(CreatedFilePaths []string) (bool, error) {

	boilerPlateNamesToPath := make(map[string]string)
	targetToBoilerPlateMatches := make(map[string]string)
	boilerPlateTargetDir := "../boilerPlateCode/"

	// CREATE A MAP OF TARGETFILES:TESTFILES

	textFiles, err := os.ReadDir(boilerPlateTargetDir)
	if err != nil {
		log.Printf("Error reading Directory %v", boilerPlateTargetDir)
	}

	for _, file := range textFiles {

		fileName := file.Name()
		nameWithouExtension := strings.TrimSuffix(fileName, filepath.Ext(fileName))
		boilerPlateNamesToPath[nameWithouExtension] = filepath.Join(boilerPlateTargetDir, fileName)
	}

	for _, targetPath := range CreatedFilePaths {
		targetFileName := filepath.Base(targetPath)
		targetNameWithoutExtension := strings.TrimSuffix(targetFileName, filepath.Ext(targetFileName))

		if boilerPlatePath, exists := boilerPlateNamesToPath[targetNameWithoutExtension]; exists {
			targetToBoilerPlateMatches[targetPath] = boilerPlatePath
		}
	}

	// READ DATA IN TEXT FILE AND WRITE TO TARGETFILE FOR EACH LOOP
	for key, value := range targetToBoilerPlateMatches {

		text, ReadErr := os.ReadFile(value)
		if ReadErr != nil {
			log.Printf("Failed to Read from Boiler Plate File")
		}

		WriteErr := os.WriteFile(key, text, 0777)
		if err != nil {
			log.Printf("Failed to write to: %v", WriteErr)
		}

	}

	return true, nil

}
