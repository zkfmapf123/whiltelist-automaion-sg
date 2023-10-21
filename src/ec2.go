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
	co     context.Context
	client *ec2.Client
}

func NewEC2(c context.Context) *ec2Params {
	cfg, err := config.LoadDefaultConfig(c)
	if err != nil {
		log.Fatalln(err)
	}

	return &ec2Params{
		co:     c,
		client: ec2.NewFromConfig(cfg),
	}
}

func (e ec2Params) RetriveSG(vpcId string, matchFn func(key string, value string, permissionLen int) bool) (string, error) {
	input := &ec2.DescribeSecurityGroupsInput{}

	sgInfos, err := e.client.DescribeSecurityGroups(e.co, input)
	if err != nil {
		log.Fatalln(err)
	}

	for _, sg := range sgInfos.SecurityGroups {
		sgVpcId, sgId, sgTags, sgPermissionIps := *sg.VpcId, *sg.GroupId, sg.Tags, sg.IpPermissions

		// verify vpcId
		if sgVpcId != vpcId {
			continue
		}

		for _, tag := range sgTags {
			k, v := *tag.Key, *tag.Value

			// check whitelist
			if matchFn(k, v, len(sgPermissionIps[0].IpRanges)) {
				return sgId, nil
			}
		}
	}

	return "", errors.New("not exists valid sg")
}

func (e ec2Params) MakeSG(vpcId string, name string, desc string, now string) (string, error) {
	sgName := fmt.Sprintf("%s-%s", name, now)

	input := &ec2.CreateSecurityGroupInput{
		VpcId:       aws.String(vpcId),
		Description: aws.String(desc),
		GroupName:   aws.String(sgName),
		TagSpecifications: []types.TagSpecification{
			{
				ResourceType: "security-group",
				Tags: []types.Tag{
					{
						Key:   aws.String("Name"),
						Value: aws.String(sgName),
					},
					{
						Key:   aws.String("Properties"),
						Value: aws.String(name),
					},
					{
						Key:   aws.String("CreatedAt"),
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

	input := &ec2.AuthorizeSecurityGroupIngressInput{
		IpProtocol: aws.String("tcp"),
		CidrIp:     aws.String(fmt.Sprintf("%s/32", ip)),
		FromPort:   aws.Int32(int32(port)),
		ToPort:     aws.Int32(int32(port)),
		GroupId:    aws.String(sgId),
	}

	_, err := e.client.AuthorizeSecurityGroupIngress(e.co, input)
	return err
}
