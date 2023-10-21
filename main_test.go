package main

import (
	"fmt"
	"strconv"
	"testing"
)

var (
	TEST_VPC_ID    = "vpc-06151804b151e2c54"
	TEST_SG_PREFIX = "whitelist"
	TEST_PORT      = "27931"
)

func Test_SG(t *testing.T) {

	num, _ := strconv.Atoi(TEST_PORT)

	for i := 0; i < 200; i++ {
		start(TEST_VPC_ID, TEST_SG_PREFIX, num, fmt.Sprintf("10.0.0.%d", i))
	}
}
