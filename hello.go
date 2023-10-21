package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
	"zkfmapf123/whitelist/src"

	"github.com/aws/aws-lambda-go/lambda"
)

// TEST
// var (
// 	VPC_ID    = "vpc-06151804b151e2c54"
// 	SG_PREFIX = "whitelist"
// 	PORT      = "27931"
// 	IPRange   = "10.0.0.1#10.0.0.2#10.0.0.3"
// )

var (
	VPC_ID    = os.Getenv("VPC_ID")
	SG_PREFIX = os.Getenv("SG_PREFIX")
	PORT      = os.Getenv("PORT")
	IPRange   = os.Getenv("IPRange")
)

func main() {
	lambda.Start(handler)
}

func handler() {
	intPort, _ := strconv.Atoi(PORT)
	ipRange := strings.Split(IPRange, "#")
	errMsgs := make([]string, len(ipRange))

	for _, ip := range ipRange {
		errMsgs = append(errMsgs, start(VPC_ID, SG_PREFIX, intPort, ip))
	}

	fmt.Println(errMsgs)
}

func start(vpcID string, sgName string, port int, ip string) string {

	ctx := context.TODO()
	ec2 := src.NewEC2(ctx)
	t := time.Now()
	formatted := fmt.Sprintf("%d-%02d-%02dT%02d:%02d:%02d",
		t.Year(), t.Month(), t.Day(),
		t.Hour(), t.Minute(), t.Second())

	sgId, err := ec2.RetriveSG(vpcID, func(k string, v string, l int) bool {
		if k == "Properties" && strings.Contains(v, sgName) && l < 50 {
			fmt.Printf("sg ingress hit %d / 50 \n", l)
			return true
		}

		return false
	})

	if err != nil {
		fmt.Println(err)

		if sgId, err = ec2.MakeSG(vpcID, sgName, fmt.Sprintf("%s-list", sgName), formatted); err != nil {
			log.Fatalln(err)
		}
	}

	err = ec2.InjectSG(sgId, port, ip)

	if strings.Contains(err.Error(), "already exists") {
		return fmt.Sprintf("already exists %s", ip)
	}

	if err == nil {
		return fmt.Sprintf("hit from %s", ip)
	}

	return err.Error()
}
