/*
Copyright © 2024 MOHAMMED ELDOSOKY mohammedeldosoky0@gmail.com
*/
package cmd

import (
	"github.com/spf13/cobra"

	"fmt"

	"github.com/mo1878/cppProjectTemplate/core"
)

// initTemplateCmd represents the initTemplate command
var initTemplateCmd = &cobra.Command{
	Use:   "initTemplate",
	Short: "Creates Project Structure with base files.",
	Long: `This command will initialize the project structure and create a source directory, 
headers directory and include the relevent config files to build, compile, debug and run C++ projects.`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {

		p := core.Project{ProjectName: args[0]}
		projectDir, err := p.NewProjectDirectory(args[0])
		if err != nil {
			fmt.Printf("Failed to created a new project directory %v", projectDir)
		}

		sD := core.SubDirectory{}
		subDirs, err := sD.NewSubDirectory(projectDir)
		if err != nil {
			fmt.Printf("Could not created Subdirectories %v", subDirs)
		}

		f := core.TemplateFile{}
		files := f.NewFileCreation(subDirs)
		if err != nil {
			fmt.Printf("Failed to create files within project subdirectories %v", files)
		}

		ct := core.TextCopy{}
		success, err := ct.InsertBoilerPlateCode(files)
		if err != nil {
			fmt.Printf("Failed to insert boilerplate code into created files %v", success)
		}

	},
}

func init() {
	rootCmd.AddCommand(initTemplateCmd)

}
