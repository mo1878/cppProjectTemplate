/*
Copyright © 2024 MOHAMMED ELDOSOKY mohammedeldosoky0@gmail.com
*/
package cmd

import (
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "cppProjectTemplate [command]",
	Short: "An application that allows you to build a C++ project template",
	Long: `This CLI application will allow you to build a project template
for a C++ application. It includes the relevant directories and an inital template
to get started with, alongside any config files required for building and debugging
your application code.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}
