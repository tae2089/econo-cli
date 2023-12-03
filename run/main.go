package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/tae2089/econo-cli/cmd"
	"github.com/tae2089/econo-cli/pkg/util"
)

func main() {
	rootCmd := cmd.CreateRootCmd()
	util.GenerateGroups(rootCmd, []*cobra.Group{
		{
			ID:    "aws",
			Title: "AWS COMMAND",
		}, {
			ID:    "gcp",
			Title: "GCP COMMAND",
		},
		// {
		// 	ID:    "azure",
		// 	Title: "AZURE COMMAND",
		// }, {
		// 	ID:    "ncp",
		// 	Title: "NCP COMMAND",
		// },
	}...)
	util.RegisterSubCommands(rootCmd, cmd.InitVersionCmd, cmd.InitAwsCmd, cmd.InitGcpCmd)
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
