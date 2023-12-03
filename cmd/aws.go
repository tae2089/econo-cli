package cmd

import (
	"fmt"
	"os"

	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/spf13/cobra"
	"github.com/tae2089/econo-cli/pkg/aws"
	"github.com/tae2089/econo-cli/pkg/tui"
	"github.com/tae2089/econo-cli/pkg/util"
)

var profile string
var region string

func InitAwsCmd() *cobra.Command {
	awsCmd := createAwsCmd()
	util.RegisterSubCommands(awsCmd,
		createStopInstanceAwsCmd,
		createStartInstanceAwsCmd,
		createGetInstancesAwsCmd,
		createGetVolumesAwsCmd,
		createDeleteVolumesAwsCmd,
	)
	return awsCmd
}

func createAwsCmd() *cobra.Command {
	var awsCmd = &cobra.Command{
		Use:   "aws",
		Short: "Manage Aws Resoucres",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}
	awsCmd.GroupID = "aws"
	awsCmd.PersistentFlags().StringVarP(&profile, "profile", "p", "", "to use aws profile")
	awsCmd.PersistentFlags().StringVarP(&region, "region", "r", "", "to use aws region")
	return awsCmd
}

func createStopInstanceAwsCmd() *cobra.Command {
	var instancdIds []string
	var instanceStopCmd = &cobra.Command{
		Use:   "stop-instance",
		Short: "Stop EC2 Instance",
		Run: func(cmd *cobra.Command, args []string) {
			result := aws.StopInstances(profile, region, instancdIds)
			p := tea.NewProgram(tui.InitTextOutputModel(result))
			if _, err := p.Run(); err != nil {
				fmt.Println(err.Error())
				os.Exit(1)
			}
		},
	}
	instanceStopCmd.Flags().StringSliceVarP(&instancdIds, "instance-ids", "i", []string{}, "insert ec2 instance ids to stop")
	return instanceStopCmd
}

func createStartInstanceAwsCmd() *cobra.Command {
	var instancdIds []string
	var instanceStartCmd = &cobra.Command{
		Use:   "start-instance",
		Short: "Start EC2 Instance",
		Run: func(cmd *cobra.Command, args []string) {
			result := aws.RestartInstances(profile, region, instancdIds)
			p := tea.NewProgram(tui.InitTextOutputModel(result))
			if _, err := p.Run(); err != nil {
				fmt.Println(err.Error())
				os.Exit(1)
			}
		},
	}
	instanceStartCmd.Flags().StringSliceVarP(&instancdIds, "instance-ids", "i", []string{}, "insert ec2 instance ids to start")
	return instanceStartCmd
}

func createGetInstancesAwsCmd() *cobra.Command {
	var instanceListCmd = &cobra.Command{
		Use:   "get-instances",
		Short: "List EC2 Instance",
		Run: func(cmd *cobra.Command, args []string) {
			rows := aws.GetInstanceList(profile, region)
			columns := []table.Column{
				{Title: "No", Width: 4},
				{Title: "InstanceID", Width: 30},
				{Title: "InstanceName", Width: 30},
				// {Title: "RuntimeOfMonth", Width: 30},
				// {Title: "TotlaRuntime", Width: 30},
			}
			p := tea.NewProgram(tui.InitalTableModel(columns, rows))
			_, err := p.Run()
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
		},
	}
	return instanceListCmd
}

func createGetVolumesAwsCmd() *cobra.Command {
	var getVolumesCmd = &cobra.Command{
		Use:   "get-volumes",
		Short: "Get EC2 Volume of List",
		Run: func(cmd *cobra.Command, args []string) {
			rows := aws.GetEc2Volumes(profile, region)
			columns := []table.Column{
				{Title: "No", Width: 4},
				{Title: "VolumeID", Width: 30},
				{Title: "Size", Width: 15},
			}
			p := tea.NewProgram(tui.InitalTableModel(columns, rows))
			_, err := p.Run()
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
		},
	}
	return getVolumesCmd
}

func createDeleteVolumesAwsCmd() *cobra.Command {
	var volumeIDs []string
	var deleteVolumesCmd = &cobra.Command{
		Use:   "delete-volumes",
		Short: "Delete EC2 Volume of List",
		Run: func(cmd *cobra.Command, args []string) {
			result := aws.DeleteEc2Volumes(profile, region, volumeIDs)
			p := tea.NewProgram(tui.InitTextOutputModel(result))
			if _, err := p.Run(); err != nil {
				fmt.Println(err.Error())
				os.Exit(1)
			}
		},
	}
	deleteVolumesCmd.Flags().StringSliceVarP(&volumeIDs, "volume-ids", "i", []string{}, "insert ec2 volume ids to delete")
	return deleteVolumesCmd
}
