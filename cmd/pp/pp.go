package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func main() {
	// Create the root command
	var rootCmd = &cobra.Command{
		Use:   "paperspls",
		Short: "CLI interface for ",
		Long:  `A simple CLI app built with Go and Cobra to demonstrate the basics.`,
		Run: func(cmd *cobra.Command, args []string) {
			// This will be called if no subcommands are provided
			fmt.Println("Welcome to the CLI app!")
		},
	}

	// Add your subcommands here
	rootCmd.AddCommand(initCmd())

	// Execute the root command
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
