package src

import (
	"context"
	"errors"
	"log"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
)

type ec2Params struct {
	co context.Context
	client *ec2.Client
}

func NewEC2(c context.Context) *ec2Params{
	cfg, err := config.LoadDefaultConfig(c)
	if err != nil {
		log.Fatalln(err)
	}

	return &ec2Params{
		co: c,
		client: ec2.NewFromConfig(cfg),
	}
}

func (e ec2Params) RetriveSG(matchFn func(key string, value string, permissionLen int) bool) (string, error){	
	input := &ec2.DescribeSecurityGroupsInput{}

	sgInfos, err := e.client.DescribeSecurityGroups(e.co, input)
	if err != nil {
		log.Fatalln(err)
	}

	for _, sg := range sgInfos.SecurityGroups {
		sgVpcId, sgId,sgTags,  sgPermissionIps := *sg.VpcId, *sg.GroupId, sg.Tags, *&sg.IpPermissions

		// verify vpcId
		if sgVpcId != sgVpcId {
			continue
		}

		for _, tag := range sgTags {
			k ,v := *tag.Key, *tag.Value
			
			// check whitelist
			if matchFn(k, v, len(sgPermissionIps)) {
				return sgId, nil
			}
		}
	}

	return "", errors.New("Not Exists Valid SG")
}

func (e ec2Params) MakeSG(vpcId string) (string, error){
	input := &ec2.CreateSecurityGroupInput{
		VpcId: aws.String(vpcId),
		Description: aws.String("Whitelist SecurityGroup"),
		GroupName: aws.String("whitelist-sg"),
		TagSpecifications: []types.TagSpecification{
			{
				Tags: []types.Tag{
					{
						Key: aws.String("Properties"),
						Value: aws.String("whitelist"),
					},
					{
						Key: aws.String("CreatedAt"),
						Value: aws.String(time.Now().Format("YYYY-MM-DD")),
					},
				},
			},
		},
	}

	sgInfo, err := e.client.CreateSecurityGroup(e.co, input)
	if err != nil {
		return "", err
	}

	return *sgInfo.GroupId, nil
} 

func (e ec2Params) InjectSG() {

}

func (e ec2Params) DeleteIngressIpInSG() {

}