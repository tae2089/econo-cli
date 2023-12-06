package aws

import (
	"context"
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func getEc2Client(profile, region string) *ec2.Client {
	cfg := getAwsConfig(profile, region)
	ec2Client := ec2.NewFromConfig(cfg)
	return ec2Client
}

func getS3Client(profile, region string) *s3.Client {
	cfg := getAwsConfig(profile, region)
	s3Client := s3.NewFromConfig(cfg)
	return s3Client
}

func getAwsConfig(profile, region string) aws.Config {
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithSharedConfigProfile(profile), config.WithRegion(region))
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	return cfg
}
