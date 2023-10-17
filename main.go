package main

import (
	"context"
	"fmt"
	"strings"
	"zkfmapf123/whitelist/src"
)

var (
	VPC_ID = "vpc-06f59281921920a2c"
	WHITELIST_PREFIX = ""
)

func main() {
	ctx := context.TODO()
	ec2 := src.NewEC2(ctx)

	sgId ,err := ec2.RetriveSG(func(k string, v string, l int) bool{
		if k == "properties" && strings.Contains(v, "whitelist") && l < 50 {
			return true
		}
		return false
	})

	// if err != nil {

	// }

	fmt.Println(sgId, err)

}
