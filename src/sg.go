package src

import (
	"context"
	"errors"
	"log"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
)

type ec2Params struct {
	ctx    context.Context
	client *ec2.Client
}

func NewEC2(c context.Context) *ec2Params {
	cfg, err := config.LoadDefaultConfig(c)
	if err != nil {
		log.Fatalln(err)
	}

	return &ec2Params{
		ctx:    c,
		client: ec2.NewFromConfig(cfg),
	}
}

func (c ec2Params) RetrieveSGIds(vpcId string, matchTag string) (string, error) {

	input := &ec2.DescribeSecurityGroupsInput{}

	sgInfos, err := c.client.DescribeSecurityGroups(c.ctx, input)
	if err != nil {
		log.Fatalln(err)
	}

	for _, sg := range sgInfos.SecurityGroups {
		sgVpcId, sgId, tags, sgIpPermissions := *sg.VpcId, *sg.GroupId, sg.Tags, sg.IpPermissions

		if sgVpcId != vpcId {
			continue
		}

		for _, v := range tags {
			key, value := *v.Key, *v.Value

			if key == "Properties" && value == matchTag && len(sgIpPermissions) < 50 {
				return sgId, nil
			}
		}
	}

	return "", errors.New("not Exists valid sg")
}

func (c ec2Params) CreateIpFromSG(vpcId string, sgId string, port string, ip string) error {

	return nil
}

func (c ec2Params) InjectIngressIpFromSG(vpcId string, sgId string, port string, ip string) error {

	return nil
}
