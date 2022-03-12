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
		vpc, err := ec2.NewVpc(ctx, "go_api_vpc", &ec2.VpcArgs{
			CidrBlock: pulumi.String("10.0.0.0/16"),
			Tags:      pulumi.ToStringMap(general_tags),
		})

		if err != nil {
			return err
		}
		ctx.Export("vpc_id", vpc.ID())
		// vpc security group
		return nil
	})
}
