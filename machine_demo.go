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
	machineData := machine.GetMachineData()
	result, err := json.Marshal(machineData)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(string(result))
}
