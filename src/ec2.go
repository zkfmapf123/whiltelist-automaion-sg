package src

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/ec2"
)

type ec2Params struct {
	co context.Context
	client *ec2.Client
}

func NewEC2(c context.Context) {
	
}

func (e ec2Params) SearchSG() {

}

func (e ec2Params) MakeSG() {

} 

func (e ec2Params) InjectSG() {

}

func (e ec2Params) DeleteIngressIpInSG() {
	
}