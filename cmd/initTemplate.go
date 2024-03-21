/*
Copyright © 2024 MOHAMMED ELDOSOKY mohammedeldosoky0@gmail.com
*/
package cmd

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

// initTemplateCmd represents the initTemplate command
var initTemplateCmd = &cobra.Command{
	Use:   "initTemplate",
	Short: "Creates Project Structure with base files.",
	Long: `This command will initialize the project structure and create a source directory, 
headers directory and include the relevent config files to build, compile, debug and run C++ projects.`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {

		taskBoilerPlate := `{
"tasks": [
		{
		"type": "cppbuild",
		"label": "C/C++: clang build active file",
		"command": "/usr/bin/clang++",
		"args": [
				"-fcolor-diagnostics",
				"-fansi-escape-codes",
				"-g",
				"-ggdb",
				"-pedantic-errors",
				"-Wall",
				"-Weffc++",
				"-Wextra",
				"-Wconversion",
				"-Wsign-conversion",
				"-Werror",
				"-std=c++17",
				"${file}",
				"-o",
				"${fileDirname}/${fileBasenameNoExtension}"
			],
			"options": {
				"cwd": "${fileDirname}"
				},
				"problemMatcher": [
					"$gcc"
				],
				"group": {
					"kind": "build",
					"isDefault": true
				},
				"detail": "Task generated by Debugger."
				}
			],
			"version": "2.0.0"
}`

		launchBoilerPlate := `{
"version": "0.2.0",
	"configurations": [
		{
			"name": "Debug",
			"type": "cppdbg",
			"request": "launch",
			"program": "${workspaceFolder}/build/YourExecutable",
			"args": [],
			"stopAtEntry": false,
			"cwd": "${workspaceFolder}",
			"environment": [],
			"externalConsole": false,
			"MIMode": "lldb",
			"preLaunchTask": "build"
		}
	]
}`

		mainCppBoilerPlate := `#include <iostream>

int main() {

	std::cout << "Hello, World from the CppProjectTemplate" << std::endl;

	return 0;
}`

		headerBoilerPlate := `// MyClass.h
#ifndef MYCLASS_H
#define MYCLASS_H
		
class MyClass {
private:
	int myNumber;
		
public:
	MyClass(int initialNumber); // Constructor
	void setNumber(int newNumber); // Setter for myNumber
	int getNumber() const; // Getter for myNumber
};
		
#endif // MYCLASS_H`

		err := os.MkdirAll(args[0], 0777)
		if err != nil {
			log.Fatal("Could not create directory", err)
		}

		sourcePath := "src"
		mainCpp := "main.cpp"

		headerPath := "headers"
		headerFile := "header.h"

		dotVsCodePath := ".vscode"
		launchJSON := "launch.json"
		taskJSON := "task.json"

		sourceDir := filepath.Join(args[0], sourcePath)
		defaultSourceFile := filepath.Join(sourceDir, mainCpp)

		headerDir := filepath.Join(args[0], headerPath)
		defaultHeaderFile := filepath.Join(headerDir, headerFile)

		dotVsCodeDir := filepath.Join(args[0], dotVsCodePath)
		taskJSONFile := filepath.Join(dotVsCodeDir, taskJSON)
		launchJSONFile := filepath.Join(dotVsCodeDir, launchJSON)

		fileSlice := []string{}

		// Creating the src Directory
		dir1 := os.MkdirAll(sourceDir, 0777)
		if err != nil {
			log.Fatal("could not create source directory", dir1)
		} else {
			fmt.Println("created the src directory")
		}

		// Creating the main.cpp file
		src, err := os.Create(defaultSourceFile)
		if err != nil {
			log.Fatal("could not create file", src)
		} else {
			fmt.Println("created main.cpp file")
			fileSlice = append(fileSlice, mainCpp)
		}

		// Creating the header directory
		dir2 := os.Mkdir(headerDir, 0777)
		if err != nil {
			log.Fatal("could not create headers directory", dir2)
		}

		// Creating the header.h file
		header, err := os.Create(defaultHeaderFile)
		if err != nil {
			log.Fatal("Could not create file", header)
		} else {
			fmt.Println("created header file")
			fileSlice = append(fileSlice, headerFile)
		}

		// Creating the .vscode directory
		dir3 := os.Mkdir(dotVsCodeDir, 0777)
		if err != nil {
			log.Fatal("Could not create .vscode dir", dir3)
		}

		// Crearing the task.json file
		task, err := os.Create(taskJSONFile)
		if err != nil {
			log.Fatal("could not create file", task)
		} else {
			fmt.Println("task.json created")
			fileSlice = append(fileSlice, taskJSON)
		}

		// Creating the launch.json file
		launch, err := os.Create(launchJSONFile)
		if err != nil {
			log.Fatal("could not create file", launch)
		} else {
			fmt.Println("launch.json created")
			fileSlice = append(fileSlice, launchJSON)
		}

		for _, file := range fileSlice {

			fmt.Println("filename:", file)

			if file == "launch.json" {

				nf1, err := launch.WriteString(launchBoilerPlate)
				if err != nil {
					log.Fatal(err)
				} else {
					log.Println("launch.json file populated", nf1)
				}

			} else if file == "task.json" {

				nf2, err := task.WriteString(taskBoilerPlate)
				if err != nil {
					log.Fatal(err)
				} else {
					log.Println("task.json file populated", nf2)
				}
			} else if file == "header.h" {

				nf3, err := header.WriteString(headerBoilerPlate)
				if err != nil {
					log.Fatal(err)
				} else {
					log.Println("header.h file populated", nf3)
				}
			} else if file == "main.cpp" {
				nf4, err := src.WriteString(mainCppBoilerPlate)
				if err != nil {
					log.Fatal(err)
				} else {
					log.Println("main.cpp file populated", nf4)
				}
			} else {
				fmt.Print("finish")
			}

		}
	},
}

func init() {
	rootCmd.AddCommand(initTemplateCmd)

}
