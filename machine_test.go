/*
author: superl[N.S.T]
github: https://github.com/super-l/
*/
package main

import (
	"fmt"
	"github.com/super-l/machine-code/machine"
	"testing"
)

// https://www.icode9.com/content-3-710187.html  go 获取linux cpuId 的方法
func TestMac1(t *testing.T) {
	macInfo1, err := machine.GetMACAddress()
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println("Local MAC ADDR2 = ", macInfo1)
	return
}

func TestIp(t *testing.T) {
	macInfo2, err := machine.GetLocalIpAddr()
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println("Local Ip = ", macInfo2)
	return
}
