package tests

import (
	"fmt"
	"regexp"
	"testing"
)

func  TestMatchImg(t *testing.T) {
	url := "accessory-11290f91a2a913975f57d3fb6b70bcc7.jpgw"
	if m, _ := regexp.MatchString(`^[a-zA-Z0-9].*\.(jpg|jpeg|png)$`, url); m == true {
		fmt.Println(m)
	}
	//fmt.Println(m)
}