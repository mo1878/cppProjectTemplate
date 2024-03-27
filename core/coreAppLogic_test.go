package core

import (
	"os"
	"testing"
)

func TestProjectAndSubDirectoryCreation(t *testing.T) {
	// Setup
	dirName := "test_project_dir"
	p := project{name: []string{dirName}}

	// Execute first function
	projectDir, err := p.newProjectDirectory()
	if err != nil {
		t.Fatalf("Failed to create new project directory: %v", err)
	}

	// Assert project directory creation
	_, err = os.Stat(projectDir)
	if os.IsNotExist(err) {
		t.Errorf("Project directory %s was not created", projectDir)
	}

	// Setup for second function
	sD := subDirectory{}                           // Adjust according to your actual struct fields
	subDirs, err := sD.newSubDirectory(projectDir) // Adjust based on the actual signature and expected return

	if err != nil {
		t.Fatalf("Failed to create subdirectory within project directory: %v", err)
	}

	// Assert subdirectory creation
	for _, subDir := range subDirs {
		if _, err := os.Stat(subDir); os.IsNotExist(err) {
			t.Errorf("Subdirectory %s was not created within project directory %s", subDir, projectDir)
		}
	}

	f := templateFile{}
	files := f.newFileCreation(subDirs)

	if err != nil {
		t.Fatalf("Failed to create files within project directory: %v", err)
	}

	// Assert file creation
	for _, file := range files {
		if _, err := os.Stat(file); os.IsNotExist(err) {
			t.Errorf("File %s was not created within project directory %s", file, subDirs)
		}
	}

	// // Cleanup
	// defer os.RemoveAll(dirName)
}
