package main

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"
	"zkfmapf123/whitelist/src"
)

var (
	VPC_ID = "vpc-06f59281921920a2c"
	WHITELIST_PREFIX = "whitelist"
)

func main() {
	ctx := context.TODO()
	ec2 := src.NewEC2(ctx)
	t := time.Now()
	formatted := fmt.Sprintf("%d-%02d-%02dT%02d:%02d:%02d",
        t.Year(), t.Month(), t.Day(),
        t.Hour(), t.Minute(), t.Second())

	sgId ,err := ec2.RetriveSG(func(k string, v string, l int) bool{
		if k == "Properties" && strings.Contains(v, WHITELIST_PREFIX) && l < 50 {
			return true
		}
		return false
	})

	if err != nil {
		fmt.Println(err)

		if sgId, err = ec2.MakeSG(VPC_ID, WHITELIST_PREFIX, formatted); err != nil {
			log.Fatalln(err)
		}
	}

	err = ec2.InjectSG(sgId, 22732,"10.0.0.1")
	if err != nil {
		log.Fatalln(err)
	}

}
