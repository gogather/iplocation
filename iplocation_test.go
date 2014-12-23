package iplocation

import (
	"fmt"
	"testing"
)

func Test_IpLocation(t *testing.T) {
	// get api key from http://www.ipinfodb.com/
	il := NewIpLocation("xxxxxxxxxxxxxxxxxxxxxxxxx")

	info, err := il.Location("74.125.45.100")
	if err != nil {
		t.Error(err)
	} else {
		fmt.Println(info)
	}

}
