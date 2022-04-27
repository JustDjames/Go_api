package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/ec2"
	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/rds"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	// creating general tags that the resources for this project will have
	general_tags := make(map[string]string)

	general_tags["project"] = "Go_api"
	general_tags["source_code"] = "https://github.com/JustDjames/Go_api"

	tags := make(map[string]string)

	pulumi.Run(func(ctx *pulumi.Context) error {
		// vpc

		tags["Name"] = "Go_api_vpc"
		vpc, err := ec2.NewVpc(ctx, "vpc", &ec2.VpcArgs{
			CidrBlock: pulumi.String("10.0.0.0/16"),
			Tags:      pulumi.ToStringMap(MergeMaps(general_tags, tags)),
		})

		if err != nil {
			return err
		}

		// subnet
		tags["Name"] = "Go_api_subnet"
		sub, err := ec2.NewSubnet(ctx, "subnet", &ec2.SubnetArgs{
			VpcId:     vpc.ID(),
			CidrBlock: pulumi.String("10.0.1.0/24"),
			Tags:      pulumi.ToStringMap(MergeMaps(general_tags, tags)),
		})

		if err != nil {
			return err
		}

		// vpc security group

		// DB subnet group
		tags["Name"] = "Go_api_db_subnet"

		sb, err := rds.NewSubnetGroup(ctx, "subnet_group", &rds.SubnetGroupArgs{
			Name: pulumi.String(tags["Name"]),
			SubnetIds: pulumi.StringArray{
				pulumi.Any(sub.ID),	
			},
			Tags: pulumi.ToStringMap(MergeMaps(general_tags, tags)), 
		})
		
		if err != nil {
			return err
		}

		// RDS
		tags["Name"] = "Go_api_rds"
		rds, err := rds.NewInstance(ctx, "rds", &rds.InstanceArgs{
			Name:               pulumi.String("users"),
			InstanceClass:      pulumi.String("db.t3.micro"),
			AllocatedStorage:   pulumi.Int(20),
			Engine:             pulumi.String("mysql"),
			EngineVersion:      pulumi.String("8.0"),
			ParameterGroupName: pulumi.String("default.mysql8.0"),
			// need to create this
			DbSubnetGroupName: sb.Name,
			// also need to create this
			VpcSecurityGroupIds: ,
			Username: pulumi.String("root"),
			// kms encrypt this
			Password: ,
			SkipFinalSnapshot:  pulumi.Bool(true),
		})

		if err != nil {
			return err
		}

		ctx.Export("vpc_id", vpc.ID())
		ctx.Export("Subnet_cidr", sub.CidrBlock())

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