package cmd

import (
	"github.com/spf13/cobra"
)

// func Execute() {
// 	if err := rootCmd.Execute(); err != nil {
// 		fmt.Println(err)
// 		os.Exit(1)
// 	}
// }

// func init() {
// }

func CreateRootCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "econo-cli",
		Short: "A brief description of your application",
		Long:  "econo-cli is a tool for managing FinOps",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}
}
