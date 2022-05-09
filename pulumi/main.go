package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v5/go/aws/ec2"
	"github.com/pulumi/pulumi-aws/sdk/v5/go/aws/rds"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
)

func main() {
	// creating general tags that the resources for this project will have
	general_tags := make(map[string]string)

	general_tags["project"] = "Go_api"
	general_tags["source_code"] = "https://github.com/JustDjames/Go_api"

	tags := make(map[string]string)

	pulumi.Run(func(ctx *pulumi.Context) error {
		// setting secrets

		c := config.New(ctx, "")

		ip := c.Require("my_ip")

		db_pass := c.RequireSecret("db_pass")

		// vpc

		tags["Name"] = "Go_api_vpc"
		vpc, err := ec2.NewVpc(ctx, "vpc", &ec2.VpcArgs{
			CidrBlock: pulumi.String("10.0.0.0/16"),
			Tags:      pulumi.ToStringMap(MergeMaps(general_tags, tags)),
		})

		if err != nil {
			return err
		}

		// subnet1
		tags["Name"] = "Go_api_subnet1"
		sub1, err := ec2.NewSubnet(ctx, "subnet1", &ec2.SubnetArgs{
			VpcId:            vpc.ID(),
			CidrBlock:        pulumi.String("10.0.1.0/24"),
			Tags:             pulumi.ToStringMap(MergeMaps(general_tags, tags)),
			AvailabilityZone: pulumi.String("eu-west-2a"),
		})

		if err != nil {
			return err
		}

		// subnet2
		tags["Name"] = "Go_api_subnet2"
		sub2, err := ec2.NewSubnet(ctx, "subnet2", &ec2.SubnetArgs{
			VpcId:            vpc.ID(),
			CidrBlock:        pulumi.String("10.0.2.0/24"),
			Tags:             pulumi.ToStringMap(MergeMaps(general_tags, tags)),
			AvailabilityZone: pulumi.String("eu-west-2b"),
		})

		if err != nil {
			return err
		}

		// internet gateway
		tags["Name"] = "Go_api_ig"
		ig, err := ec2.NewInternetGateway(ctx, "ig", &ec2.InternetGatewayArgs{
			VpcId: vpc.ID(),
			Tags:  pulumi.ToStringMap(MergeMaps(general_tags, tags)),
		})

		if err != nil {
			return err
		}

		// route table
		tags["Name"] = "Go_api_route_table"

		rt, err := ec2.NewRouteTable(ctx, "example", &ec2.RouteTableArgs{
			VpcId: vpc.ID(),
			Routes: ec2.RouteTableRouteArray{
				&ec2.RouteTableRouteArgs{
					CidrBlock: pulumi.String("0.0.0.0/0"),
					GatewayId: ig.ID(),
				},
			},
			Tags: pulumi.ToStringMap(MergeMaps(general_tags, tags)),
		})

		if err != nil {
			return err
		}

		rta1, err := ec2.NewRouteTableAssociation(ctx, "routeTableAssociation1", &ec2.RouteTableAssociationArgs{
			SubnetId:     sub1.ID(),
			RouteTableId: rt.ID(),
		})

		if err != nil {
			return err
		}

		rta2, err := ec2.NewRouteTableAssociation(ctx, "routeTableAssociation2", &ec2.RouteTableAssociationArgs{
			SubnetId:     sub2.ID(),
			RouteTableId: rt.ID(),
		})

		if err != nil {
			return err
		}
		// nacl
		tags["Name"] = "Go_api_nacl"

		nacl, err := ec2.NewNetworkAcl(ctx, "nacl", &ec2.NetworkAclArgs{
			VpcId: vpc.ID(),

			Ingress: ec2.NetworkAclIngressArray{
				&ec2.NetworkAclIngressArgs{
					Protocol:  pulumi.String("tcp"),
					RuleNo:    pulumi.Int(100),
					Action:    pulumi.String("allow"),
					CidrBlock: pulumi.String(ip),
					FromPort:  pulumi.Int(3306),
					ToPort:    pulumi.Int(3306),
				},
			},

			Egress: ec2.NetworkAclEgressArray{
				&ec2.NetworkAclEgressArgs{
					Protocol:  pulumi.String("tcp"),
					RuleNo:    pulumi.Int(100),
					Action:    pulumi.String("allow"),
					CidrBlock: pulumi.String(ip),
					FromPort:  pulumi.Int(1024),
					ToPort:    pulumi.Int(65535),
				},
			},
			Tags: pulumi.ToStringMap(MergeMaps(general_tags, tags)),
		})

		if err != nil {
			return err
		}

		nacla1, err := ec2.NewNetworkAclAssociation(ctx, "naclAssociation1", &ec2.NetworkAclAssociationArgs{
			NetworkAclId: nacl.ID(),
			SubnetId:     sub1.ID(),
		})

		if err != nil {
			return err
		}

		nacla2, err := ec2.NewNetworkAclAssociation(ctx, "naclAssociation2", &ec2.NetworkAclAssociationArgs{
			NetworkAclId: nacl.ID(),
			SubnetId:     sub2.ID(),
		})

		if err != nil {
			return err
		}

		// vpc security group
		tags["Name"] = "Go_api_sg"
		sg, err := ec2.NewSecurityGroup(ctx, "sg", &ec2.SecurityGroupArgs{
			Name:        pulumi.String(tags["Name"]),
			Description: pulumi.String("Security Group to allow access to MySql RDS from my ip"),
			VpcId:       vpc.ID(),
			Ingress: ec2.SecurityGroupIngressArray{
				&ec2.SecurityGroupIngressArgs{
					FromPort: pulumi.Int(3306),
					ToPort:   pulumi.Int(3306),
					Protocol: pulumi.String("tcp"),
					CidrBlocks: pulumi.StringArray{
						pulumi.String(ip),
					},
				},
			},
			Egress: ec2.SecurityGroupEgressArray{
				&ec2.SecurityGroupEgressArgs{
					FromPort: pulumi.Int(0),
					ToPort:   pulumi.Int(0),
					Protocol: pulumi.String("-1"),
					CidrBlocks: pulumi.StringArray{
						pulumi.String("0.0.0.0/0"),
					},
				},
			},
			Tags: pulumi.ToStringMap(MergeMaps(general_tags, tags)),
		})

		if err != nil {
			return err
		}

		// DB subnet group
		tags["Name"] = "go_api_db_subnet"

		sb, err := rds.NewSubnetGroup(ctx, "subnet_group", &rds.SubnetGroupArgs{
			Name: pulumi.String(tags["Name"]),
			SubnetIds: pulumi.StringArray{
				sub1.ID(),
				sub2.ID(),
			},
			Tags: pulumi.ToStringMap(MergeMaps(general_tags, tags)),
		})

		if err != nil {
			return err
		}

		// RDS
		tags["Name"] = "Go_api_rds"
		rds, err := rds.NewInstance(ctx, "rds", &rds.InstanceArgs{
			DbName:             pulumi.String("users"),
			InstanceClass:      pulumi.String("db.t3.micro"),
			AllocatedStorage:   pulumi.Int(20),
			Engine:             pulumi.String("mysql"),
			EngineVersion:      pulumi.String("8.0"),
			ParameterGroupName: pulumi.String("default.mysql8.0"),
			DbSubnetGroupName:  sb.Name,
			VpcSecurityGroupIds: pulumi.StringArray{
				sg.ID(),
			},
			Username: pulumi.String("root"),

			Password:          db_pass,
			SkipFinalSnapshot: pulumi.Bool(true),
			Tags:              pulumi.ToStringMap(MergeMaps(general_tags, tags)),
		})

		if err != nil {
			return err
		}

		ctx.Export("vpc_id", vpc.ID())
		ctx.Export("route_table_association1_id", rta1.ID())
		ctx.Export("route_table_association2_id", rta2.ID())
		ctx.Export("nacl_association1_id", nacla1.ID())
		ctx.Export("nacl_association2_id", nacla2.ID())
		ctx.Export("rds_endpoint", rds.Endpoint)
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
