package main

import (
	"context"
	"zkfmapf123/whitelist/src"
)

var (
	VPC_ID = "vpc-06f59281921920a2c"
	WHITELIST_PREFIX = ""
)

func main() {
	ctx := context.TODO()
	ec2 := src.NewEC2(ctx)

	sgId ,err := ec2.RetriveSG()
	if err != nil {

	}

}