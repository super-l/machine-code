package main

import (
	"encoding/json"
	"fmt"
	machine_code "github.com/super-l/machine-code"
	"strings"
)

func main() {
	if machine_code.MachineErr != nil {
		fmt.Println("获取机器码信息错误:" + machine_code.MachineErr.Error())
		return
	}

	machineJson, _ := json.Marshal(machine_code.Machine)
	fmt.Println("机器码信息汇总:" + string(machineJson))

	trafficIp, _ := machine_code.GetIpAddr()
	fmt.Println("当前活跃IP：" + trafficIp)

	allIp, _ := machine_code.GetIpAddrAll()
	fmt.Println("所有IP：" + strings.Join(allIp, " "))
}
