/*
author: superl[N.S.T]
github: https://github.com/super-l/
*/
package machine

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

type MacMachine struct {}

var macMachineData MachineData

type macXmlStruct struct {
	XMLName  xml.Name `xml:"plist"`
	Array    macDataArrayStruct `xml:"array"`
}

type macDataArrayStruct struct {
	Dict macDictStruct `xml:"dict"` // 读取user数组
}

type macDictStruct struct {
	Key    []string  `xml:"key"`
	Real   []string  `xml:"real"`
	String []string  `xml:"string"`
	Array  macDictArrayStruct `xml:"array"`
}

type macDictArrayStruct struct {
	Dict macDictItemStruct `xml:"dict"` // 读取user数组
}

type macDictItemStruct struct {
	Key    []string  `xml:"key"`
	Integer []int    `xml:"integer"`
	String []string  `xml:"string"`
}

func (machine MacMachine) getSerialNumber()(data string, err error ){
	sysInfo, err := machine.getMacSysInfo()
	if err != nil {
		return "", err
	}
	return sysInfo.SerialNumber, err
}

func (machine MacMachine) getPlatformUUID() (UUID string, err error) {
	sysInfo, err := machine.getMacSysInfo()
	if err != nil {
		return "", err
	}
	return sysInfo.PlatformUUID, err
}

func (machine MacMachine) getCpuId() (cpuId string, err error) {
	sysInfo, err := machine.getMacSysInfo()
	if err != nil {
		return "", err
	}
	return sysInfo.CpuId, err
}

func (machine MacMachine) getMacSysInfo() (data MachineData, err error ){
	var cmd *exec.Cmd
	cmd = exec.Command("system_profiler", "SPHardwareDataType", "-xml")

	var out bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = os.Stderr

	err = cmd.Start()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	err = cmd.Wait()
	if err == nil {
		macMachineData, err = machine.macXmlToData(out.String())
		if err == nil {
			cpuid, _ := machine.getSysCpuId()
			macMachineData.CpuId = cpuid
		}
		return macMachineData, nil
	} else {
		return MachineData{}, err
	}
}

func (MacMachine) macXmlToData(xmlcontent string) (MachineData, error){
	x := macXmlStruct{}
	err := xml.Unmarshal([]byte(xmlcontent), &x)
	if err != nil {
		return MachineData{}, err
	} else {
		count := len(x.Array.Dict.Array.Dict.String)
		serialData := MachineData{
			PlatformUUID: x.Array.Dict.Array.Dict.String[count-2],
			SerialNumber: x.Array.Dict.Array.Dict.String[count-1],
			CpuId: "",
		}
		return serialData, nil
	}
}

func (machine MacMachine) getSysCpuId() (cpuId string, err error) {
	//sysctl -x machdep.cpu.signature
	var cmd *exec.Cmd
	cmd = exec.Command("sysctl", "-x", "machdep.cpu.signature")

	var out bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = os.Stderr

	err = cmd.Start()
	if err != nil {
		return "", err
	}
	err = cmd.Wait()
	if err == nil {
		cpuId = out.String()
		cpuId = strings.Replace(cpuId, " ", "", -1)
		cpuId = strings.Replace(cpuId, "\n", "", -1)
		cpuId = strings.Replace(cpuId, "machdep.cpu.signature:", "", -1)
		return cpuId, nil
	} else {
		return "", err
	}
}
