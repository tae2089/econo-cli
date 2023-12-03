package aws

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/charmbracelet/bubbles/table"
)

func GetInstanceList(profile, region string) []table.Row {
	ec2Client := getEc2Client(profile, region)
	resp, err := ec2Client.DescribeInstances(context.TODO(), &ec2.DescribeInstancesInput{
		Filters: []types.Filter{
			{
				Name:   aws.String("instance-state-name"),
				Values: []string{"running"},
			},
		},
	})
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}
	rows := make([]table.Row, 0)
	// currentTime := time.Now().UTC()
	for idx, reservation := range resp.Reservations {
		for _, instance := range reservation.Instances {
			// launchTime := instance.LaunchTime.UTC()
			// startOfMonth := time.Date(currentTime.Year(), currentTime.Month(), 1, 0, 0, 0, 0, currentTime.Location())
			name := getInstacnceName(instance)

			// duration := time.Since(launchTime)
			// runningTimeOfMonth := getRunningTimeOfMonth(launchTime, startOfMonth, currentTime)
			// rows = append(rows, table.Row{strconv.Itoa(idx), *instance.InstanceId, name, runningTimeOfMonth.Truncate(time.Hour).String(), duration.Truncate(time.Hour).String()})
			rows = append(rows, table.Row{strconv.Itoa(idx), *instance.InstanceId, name})
		}
	}
	return rows
}

func StopInstances(profile, region string, instanceIds []string) string {
	ec2Client := getEc2Client(profile, region)
	success, failed := 0, 0
	output, err := ec2Client.StopInstances(context.TODO(), &ec2.StopInstancesInput{
		InstanceIds: instanceIds,
	})

	for _, stoppingInstance := range output.StoppingInstances {
		if stoppingInstance.CurrentState.Code == aws.Int32(16) {
			failed++
		} else {
			success++
		}
	}
	if err != nil {
		fmt.Println(err.Error())
	}
	return fmt.Sprintf("Successfully stopped %d instances and failed to stop %d instances", success, failed)
}

func RestartInstances(profile, region string, instanceIds []string) string {
	cfg := getAwsConfig(profile, region)
	ec2Client := ec2.NewFromConfig(cfg)
	success, failed := 0, 0
	output, err := ec2Client.StartInstances(context.TODO(), &ec2.StartInstancesInput{
		InstanceIds: instanceIds,
	})

	for _, startingInstance := range output.StartingInstances {
		if startingInstance.CurrentState.Code == aws.Int32(80) {
			failed++
		} else {
			success++
		}
	}
	if err != nil {
		fmt.Println(err.Error())
	}
	return fmt.Sprintf("Successfully restarted %d instances and failed to restart %d instances", success, failed)
}

// getInstanceName returns the name of the EC2 instance.
// if EC2 instance name is not set, it returns the empty string.
func getInstacnceName(instance types.Instance) string {
	name := ""
	for _, tag := range instance.Tags {
		if *tag.Key == "Name" {
			name = *tag.Value
		}
	}
	return name
}

// getRunningTimeOfMonth returns the running time of the given instance in the given month
func getRunningTimeOfMonth(launchTime time.Time, startOfMonth time.Time, currentTime time.Time) time.Duration {
	var runningTimeOfMonth time.Duration
	if launchTime.Before(startOfMonth) {
		runningTimeOfMonth = currentTime.Sub(startOfMonth)
	} else {

		runningTimeOfMonth = currentTime.Sub(launchTime)
	}
	return runningTimeOfMonth
}

func GetEc2Volumes(profile, region string) []table.Row {
	ec2Client := getEc2Client(profile, region)
	resp, err := ec2Client.DescribeVolumes(context.TODO(), &ec2.DescribeVolumesInput{
		Filters: []types.Filter{
			{
				Name:   aws.String("status"),
				Values: []string{"available"},
			},
		},
	})
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}
	rows := make([]table.Row, 0)
	for idx, volume := range resp.Volumes {
		rows = append(rows, table.Row{strconv.Itoa(idx), *volume.VolumeId, fmt.Sprintf("%dGB", aws.ToInt32(volume.Size))})
	}
	return rows
}

func DeleteEc2Volumes(profile, region string, volumeIds []string) string {
	ec2Client := getEc2Client(profile, region)
	success, failed := 0, 0
	for _, volumeID := range volumeIds {
		_, err := ec2Client.DeleteVolume(context.TODO(), &ec2.DeleteVolumeInput{
			VolumeId: aws.String(volumeID),
		})
		if err != nil {
			fmt.Println(err.Error())
			failed++
		} else {
			success++
		}
	}
	return fmt.Sprintf("Successfully deleted %d volumes and failed to delete %d volumes", success, failed)
}
