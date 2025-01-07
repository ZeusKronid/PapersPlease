package main

import (
	"fmt"

	"github.com/spf13/cobra"
)

func initCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "init",
		Short: "Initialize something",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Initialization complete!")
		},
	}
}
