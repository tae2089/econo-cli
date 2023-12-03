package util

import "github.com/spf13/cobra"

type GeneratorCmdFn func() *cobra.Command

func RegisterSubCommands(parent *cobra.Command, generateCommands ...GeneratorCmdFn) {
	for _, gc := range generateCommands {
		parent.AddCommand(gc())
	}
}

func GenerateGroups(parent *cobra.Command, groups ...*cobra.Group) {
	for _, g := range groups {
		parent.AddGroup(g)
	}
}
