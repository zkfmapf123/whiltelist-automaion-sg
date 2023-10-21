package main

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"
	"zkfmapf123/whitelist/src"
)

var (
	VPC_ID    = "vpc-06151804b151e2c54"
	SG_PREFIX = "whitelist"
	PORT      = "27931"
	IPRange   = []string{"10.0.0.1"}
)

func main() {
	intPort, _ := strconv.Atoi(PORT)

	for _, ip := range IPRange {
		start(VPC_ID, SG_PREFIX, intPort, ip)
	}

}

func start(vpc_id string, whitelist string, port int, ip string) {
	ctx := context.TODO()
	ec2 := src.NewEC2(ctx)
	t := time.Now()
	formatted := fmt.Sprintf("%d-%02d-%02dT%02d:%02d:%02d",
		t.Year(), t.Month(), t.Day(),
		t.Hour(), t.Minute(), t.Second())

	sgId, err := ec2.RetriveSG(VPC_ID, func(k string, v string, l int) bool {
		if k == "Properties" && strings.Contains(v, SG_PREFIX) && l < 50 {
			fmt.Printf("sg ingress hit %d / 50 \n", l)
			return true
		}

		return false
	})

	if err != nil {
		fmt.Println(err)

		if sgId, err = ec2.MakeSG(VPC_ID, SG_PREFIX, fmt.Sprintf("%s-list", SG_PREFIX), formatted); err != nil {
			log.Fatalln(err)
		}
	}

	err = ec2.InjectSG(sgId, port, ip)
	if err != nil {
		log.Fatalln(err)
	}
}
