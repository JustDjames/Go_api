package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/ec2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	// creating general tags that the resources for this project will have
	general_tags := make(map[string]string)

	general_tags["project"] = "Go_api"
	general_tags["source_code"] = "https://github.com/JustDjames/Go_api"

	tags := make(map[string]string)

	pulumi.Run(func(ctx *pulumi.Context) error {
		// Create an AWS resource (S3 Bucket)
		// bucket, err := s3.NewBucket(ctx, "my-bucket", nil)
		// if err != nil {
		// 	return err
		// }

		// // Export the name of the bucket
		// ctx.Export("bucketName", bucket.ID())
		// return nil

		// vpc

		tags["Name"] = "Go_api_vpc"
		vpc, err := ec2.NewVpc(ctx, "go_api_vpc", &ec2.VpcArgs{
			CidrBlock: pulumi.String("10.0.0.0/16"),
			Tags:      pulumi.ToStringMap(MergeMaps(general_tags, tags)),
		})

		if err != nil {
			return err
		}

		// subnet
		tags["Name"] = "Go_api_subnet"
		sub, err := ec2.NewSubnet(ctx, "Go_api_subnet", &ec2.SubnetArgs{
			VpcId:     vpc.ID(),
			CidrBlock: pulumi.String("10.0.1.0/24"),
			Tags:      pulumi.ToStringMap(MergeMaps(general_tags, tags)),
		})

		if err != nil {
			return err

		}
		ctx.Export("vpc_id", vpc.ID())
		ctx.Export("Subnet_id", sub.ID())

		// vpc security group

		return nil
	})
}

// merges two string-string maps into one
func MergeMaps(m1 map[string]string, m2 map[string]string) map[string]string {
	m3 := make(map[string]string)

	for a, b := range m1 {
		m3[a] = b
	}

	for c, d := range m2 {
		m3[c] = d
	}

	return m3
}
