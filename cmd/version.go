package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func InitVersionCmd() *cobra.Command {
	return createVersionCmd()
}

func createVersionCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "version",
		Short: "Print the version number of Econocli",
		Long:  `All software has versions. This is Econocli's`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Econocli FinOps of Cli v0.1 -- HEAD")
		},
	}
}
