package main

import (
	"context"
	"log"
	"os"
	"zkfmapf123/src"
)

var (
	SG_PREFIX_NAME = "whitelist"
	VPC_ID         = os.Getenv("VPC_ID")
	PORT           = "22791"
	IP             = "10.0.0.0/32"
)

func main() {
	ctx := context.TODO()
	ec2Options := src.NewEC2(ctx)
	sgId, err := ec2Options.RetrieveSGIds(VPC_ID, SG_PREFIX_NAME)

	// not exists valid sg
	if err != nil && sgId == "" {
		err = ec2Options.CreateIpFromSG(VPC_ID, sgId, PORT, IP)
	} else {
		err = ec2Options.InjectIngressIpFromSG(VPC_ID, sgId, PORT, IP)
	}

	if err != nil {
		log.Fatalln("Fail")
	}

	log.Fatalln("Success")
}

// make security group
// func makeSg(ctx context.Context, client *ec2.Client, num int) string {

// 	input := &ec2.CreateSecurityGroupInput{
// 		Description: aws.String("whitelist"),
// 		GroupName:   aws.String(fmt.Sprintf("%s-%d", SG_PREFIX_NAME, num)),
// 		TagSpecifications: []types.TagSpecification{
// 			{
// 				ResourceType: "security-group",
// 				Tags: []types.Tag{
// 					{
// 						Key:   aws.String("Name"),
// 						Value: aws.String(fmt.Sprintf("%s-%d", SG_PREFIX_NAME, num)),
// 					},
// 					{
// 						Key:   aws.String("Properties"),
// 						Value: aws.String(SG_PREFIX_NAME),
// 					},
// 				},
// 			},
// 		},
// 		VpcId: aws.String(VPC_ID),
// 	}

// 	sg, err := client.CreateSecurityGroup(ctx, input)

// 	if err != nil {
// 		log.Fatalln(err)
// 	}

// 	return *sg.GroupId
// }
