package src

import (
	"context"
	"errors"
	"fmt"
	"log"

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

func (e ec2Params) MakeSG(vpcId string, name string, now string) (string, error){

	input := &ec2.CreateSecurityGroupInput{
		VpcId: aws.String(vpcId),
		Description: aws.String("Whitelist SecurityGroup"),
		GroupName: aws.String("whitelist-sg"),
		TagSpecifications: []types.TagSpecification{
			{
				ResourceType: "security-group",
				Tags: []types.Tag{
					{
						Key : aws.String("Name"),
						Value: aws.String(fmt.Sprintf("%s-%s", name,now)),
					},
					{
						Key: aws.String("Properties"),
						Value: aws.String(name),
					},
					{
						Key: aws.String("CreatedAt"),
						Value: aws.String(now),
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

func (e ec2Params) InjectSG(sgId string, port int, ip string) error {	
	input := &ec2.UpdateSecurityGroupRuleDescriptionsIngressInput{
		GroupId: aws.String(sgId),
		IpPermissions: []types.IpPermission{
			{
				IpProtocol: aws.String("tcp"),
				FromPort: aws.Int32(int32(port)),
				ToPort: aws.Int32(int32(port)),
				IpRanges : []types.IpRange{
					{
						CidrIp: aws.String(fmt.Sprintf("%s/32", ip)),
					},
				},
			},
		},
	}

	_, err := e.client.UpdateSecurityGroupRuleDescriptionsIngress(e.co, input)
	if err != nil {
		return err
	}

	return nil

}

func (e ec2Params) DeleteIngressIpInSG() {

}