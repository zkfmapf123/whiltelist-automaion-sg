package src

import (
	"context"
	"errors"
	"log"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
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

// func (e ec2Params) MakeSG() (string, error){

// } 

func (e ec2Params) InjectSG() {

}

func (e ec2Params) DeleteIngressIpInSG() {

}