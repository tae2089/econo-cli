package cmd

import (
	"fmt"
	"os"

	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/spf13/cobra"
	"github.com/tae2089/econo-cli/pkg/gcp"
	"github.com/tae2089/econo-cli/pkg/tui"
	"github.com/tae2089/econo-cli/pkg/util"
)

var projectID string
var zone string

func InitGcpCmd() *cobra.Command {
	gcpCmd := createGcpCmd()
	util.RegisterSubCommands(gcpCmd, createStartInstanceGcpCmd, createStopInstanceGcpCmd, createGetInstancesGcpCmd, createGetVolumesGcpCmd, createDeleteVolumesGcpCmd)
	return gcpCmd
}

func createGcpCmd() *cobra.Command {
	var gcpCmd = &cobra.Command{
		Use:   "gcp",
		Short: "Manage Gcp Resoucres",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}
	gcpCmd.GroupID = "gcp"
	gcpCmd.PersistentFlags().StringVarP(&projectID, "project-id", "p", "", "to use gcp projecgt id")
	gcpCmd.PersistentFlags().StringVarP(&zone, "zone", "z", "", "to use gcp zone")
	return gcpCmd
}

func createStopInstanceGcpCmd() *cobra.Command {
	var computeEngineIDs []string
	var computeEngineStopCmd = &cobra.Command{
		Use:   "stop-instances",
		Short: "Stop Compute Engine",
		Run: func(cmd *cobra.Command, args []string) {
			result := gcp.StopInstances(projectID, zone, computeEngineIDs)
			fmt.Println(result)
		},
	}
	computeEngineStopCmd.Flags().StringSliceVarP(&computeEngineIDs, "instance-ids", "i", []string{}, "insert ec2 instance ids to stop")
	return computeEngineStopCmd
}

func createStartInstanceGcpCmd() *cobra.Command {
	var computeEngineNames []string
	var computeEngineStartCmd = &cobra.Command{
		Use:   "start-instances",
		Short: "Start Compute Engine",
		Run: func(cmd *cobra.Command, args []string) {
			result := gcp.StartInstance(projectID, zone, computeEngineNames)
			fmt.Println(result)
		},
	}
	computeEngineStartCmd.Flags().StringSliceVarP(&computeEngineNames, "instancenames", "i", []string{}, "insert ec2 instance ids to start")
	return computeEngineStartCmd
}

func createGetInstancesGcpCmd() *cobra.Command {
	var instanceListCmd = &cobra.Command{
		Use:   "get-instances",
		Short: "List Compute Engine",
		Run: func(cmd *cobra.Command, args []string) {
			rows := gcp.GetInstanceList(projectID, zone)
			columns := []table.Column{
				{Title: "No", Width: 4},
				{Title: "InstanceID", Width: 30},
				{Title: "InstanceName", Width: 30},
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

func createGetVolumesGcpCmd() *cobra.Command {
	var getVolumesCmd = &cobra.Command{
		Use:   "get-disks",
		Short: "Get Compute Engine Disk of List",
		Run: func(cmd *cobra.Command, args []string) {
			rows := gcp.GetVolumes(projectID, zone)
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

func createDeleteVolumesGcpCmd() *cobra.Command {
	var volumeIDs []string
	var deleteVolumesCmd = &cobra.Command{
		Use:   "delete-disks",
		Short: "Delete Compute Engine Disk of List",
		Run: func(cmd *cobra.Command, args []string) {
			result := gcp.DeleteDisks(projectID, zone, volumeIDs)
			fmt.Println(result)
		},
	}
	deleteVolumesCmd.Flags().StringSliceVarP(&volumeIDs, "disk-ids", "v", []string{}, "insert compute engine disk ids to delete")
	return deleteVolumesCmd
}
