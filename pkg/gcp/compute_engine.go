package gcp

import (
	"context"
	"fmt"
	"strconv"

	compute "cloud.google.com/go/compute/apiv1"
	"cloud.google.com/go/compute/apiv1/computepb"
	"github.com/charmbracelet/bubbles/table"
	"google.golang.org/api/iterator"
)

func GetInstanceList(projectID, zone string) []table.Row {
	instancesClient, ctx := getComputeEngineClient()
	defer instancesClient.Close()
	req := &computepb.ListInstancesRequest{
		Project: projectID,
		Zone:    zone,
	}
	it := instancesClient.List(ctx, req)
	idx := 0
	rows := make([]table.Row, 0)
	for {
		instance, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil
		}
		rows = append(rows, table.Row{strconv.Itoa(idx), fmt.Sprintf("%d", *instance.Id), instance.GetName()})
		idx++
	}
	return rows
}

func StopInstances(projectID, zone string, instanceNames []string) string {
	instancesClient, ctx := getComputeEngineClient()
	defer instancesClient.Close()
	success, failed := 0, 0
	for _, instanceName := range instanceNames {
		_, err := instancesClient.Stop(ctx, &computepb.StopInstanceRequest{
			Project:  projectID,
			Zone:     zone,
			Instance: instanceName,
		})
		if err != nil {
			failed++
			fmt.Println(err.Error())
		} else {
			success++
		}
	}
	return fmt.Sprintf("Successfully stopped %d instances and failed to stop %d compute engine", success, failed)
}

func StartInstance(projectID, zone string, instanceNames []string) string {
	instancesClient, ctx := getComputeEngineClient()
	defer instancesClient.Close()
	success, failed := 0, 0
	for _, instanceName := range instanceNames {
		_, err := instancesClient.Start(ctx, &computepb.StartInstanceRequest{
			Project:  projectID,
			Zone:     zone,
			Instance: instanceName,
		})
		if err != nil {
			failed++
			fmt.Println(err.Error())
		} else {
			success++
		}
	}
	return fmt.Sprintf("Successfully started %d instances and failed to start %d compute engine", success, failed)
}

func GetVolumes(projectID, zone string) []table.Row {
	disksClient, ctx := getDisksClient()
	defer disksClient.Close()
	req := &computepb.ListDisksRequest{
		Project: projectID,
		Zone:    zone,
	}
	it := disksClient.List(ctx, req)
	idx := 0
	rows := make([]table.Row, 0)
	for {
		volume, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			fmt.Println(err.Error())
			continue
		}
		rows = append(rows, table.Row{strconv.Itoa(idx), fmt.Sprintf("%d", *volume.Id), fmt.Sprintf("%d", *volume.SizeGb)})
		idx++
	}
	return rows
}

func DeleteDisks(projectID, zone string, diskNames []string) string {
	disksClient, ctx := getDisksClient()
	defer disksClient.Close()
	success, failed := 0, 0
	for _, diskName := range diskNames {
		_, err := disksClient.Delete(ctx, &computepb.DeleteDiskRequest{
			Project: projectID,
			Zone:    zone,
			Disk:    diskName,
		})
		if err != nil {
			failed++
			fmt.Println(err.Error())
		} else {
			success++
		}
	}
	return fmt.Sprintf("Successfully deleted %d instances and failed to delete %d compute engine", success, failed)
}

func getComputeEngineClient() (*compute.InstancesClient, context.Context) {
	ctx := context.Background()
	instancesClient, err := compute.NewInstancesRESTClient(ctx)
	if err != nil {
		panic(err)
	}
	return instancesClient, ctx
}

func getDisksClient() (*compute.DisksClient, context.Context) {
	ctx := context.Background()
	disksClient, err := compute.NewDisksRESTClient(ctx)
	if err != nil {
		panic(err)
	}
	return disksClient, ctx
}
