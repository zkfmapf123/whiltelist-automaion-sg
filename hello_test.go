package main

import (
	"fmt"
	"strconv"
	"strings"
	"testing"
)

var (
	TEST_VPC_ID    = "vpc-06151804b151e2c54"
	TEST_SG_PREFIX = "whitelist"
	TEST_PORT      = "27931"
	TEST_IPRagne   = "10.0.0.1#10.0.0.2#10.0.0.3"
)

func Test_SG(t *testing.T) {

	num, _ := strconv.Atoi(TEST_PORT)

	for _, ip := range strings.Split(TEST_IPRagne, "#") {
		fmt.Println(TEST_VPC_ID, TEST_SG_PREFIX, num, ip)
		start(TEST_VPC_ID, TEST_SG_PREFIX, num, ip)
	}
}
