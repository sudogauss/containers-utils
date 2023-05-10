/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)


var rootCmd = &cobra.Command{
	Use:   "stabu-cli",
	Short: "stabu-cli (standalone-builder cli) allows you to build cross-compilers and standalone executables inside docker alpine container",
	Long: `You can configure your stabu-cli by providing it a workspace path or customize it to provide directly the path to cross compilers or to the project
	which allows to build cross-compilers. You can do the same for standalone executables. However, be aware that you have full responsability to choose
	a correct compiler and executable as stabu-cli is used only to detect a build system and compile.`,
	// Run: func(cmd *cobra.Command, args []string) { },
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
}


