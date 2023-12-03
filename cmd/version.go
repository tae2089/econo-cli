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
		Short: "Print the version number of econo-cli",
		Long:  `All software has versions. This is econo-cli's`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("econo-cli FinOps of Cli v0.1 -- HEAD")
		},
	}
}
