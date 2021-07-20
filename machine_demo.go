/*
author: superl[N.S.T]
github: https://github.com/super-l/
*/
package main

import (
	"encoding/json"
	"fmt"
	"github.com/super-l/machine-code/machine"
)

// https://www.icode9.com/content-3-710187.html  go 获取linux cpuId 的方法
func main() {
	serialNumber, err := machine.GetSerialNumber()
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println("serialNumber = ", serialNumber)

	uuid, err := machine.GetPlatformUUID()
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println("uuid = ", uuid)

	cpuid, err := machine.GetCpuId()
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println("cpuid = ", cpuid)

	macInfo, err := machine.GetMACAddress()
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println("mac = ", macInfo)


	machineData := machine.GetMachineData()
	result, err := json.Marshal(machineData)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(string(result))
}