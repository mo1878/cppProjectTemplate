package core

import (
	"fmt"
	"os"
	"testing"
)

func TestProjectAndSubDirectoryCreation(t *testing.T) {
	// Setup
	dirName := "test_project_dir"
	p := Project{ProjectName: dirName}

	// Execute first function
	projectDir, err := p.NewProjectDirectory(dirName)
	if err != nil {
		t.Fatalf("Failed to create new project directory: %v", err)
	}

	// Assert project directory creation
	_, err = os.Stat(projectDir)
	if os.IsNotExist(err) {
		t.Errorf("Project directory %s was not created", projectDir)
	}

	// Setup for second function
	sD := SubDirectory{}                           // Adjust according to your actual struct fields
	subDirs, err := sD.NewSubDirectory(projectDir) // Adjust based on the actual signature and expected return

	if err != nil {
		t.Fatalf("Failed to create subdirectory within project directory: %v", err)
	}

	// Assert subdirectory creation
	for _, subDir := range subDirs {
		if _, err := os.Stat(subDir); os.IsNotExist(err) {
			t.Errorf("Subdirectory %s was not created within project directory %s", subDir, projectDir)
		}
	}

	// Setup for third function
	f := TemplateFile{}
	files := f.NewFileCreation(subDirs)

	if err != nil {
		t.Fatalf("Failed to create files within project directory: %v", err)
	}

	// Assert file creation
	for _, file := range files {
		if _, err := os.Stat(file); os.IsNotExist(err) {
			t.Errorf("File %s was not created within project directory %s", file, subDirs)
		}
	}

	ct := TextCopy{}
	success, err := ct.InsertBoilerPlateCode(files)
	if err != nil {
		t.Fatalf("Function failed: %v", err)
	}
	fmt.Println(success)

	// Cleanup
	defer os.RemoveAll(dirName)
}
