//go:build darwin
// +build darwin

package goos

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"github.com/super-l/machine-code/types"
	"os"
	"os/exec"
	"strings"
)

type MacMachine struct{}

var macMachineData types.MachineInformation

type macXmlStruct struct {
	XMLName xml.Name           `xml:"plist"`
	Array   macDataArrayStruct `xml:"array"`
}

type macDataArrayStruct struct {
	Dict macDictStruct `xml:"dict"` // 读取user数组
}

type macDictStruct struct {
	Key    []string           `xml:"key"`
	Real   []string           `xml:"real"`
	String []string           `xml:"string"`
	Array  macDictArrayStruct `xml:"array"`
}

type macDictArrayStruct struct {
	Dict macDictItemStruct `xml:"dict"` // 读取user数组
}

type macDictItemStruct struct {
	Key     []string `xml:"key"`
	Integer []int    `xml:"integer"`
	String  []string `xml:"string"`
}

func (mac MacMachine) GetMachine() (res types.MachineInformation, err error) {
	platformUUID, err := mac.GetUUID()
	if err != nil {
		return res, err
	}
	boardSerialNumber, err := mac.GetBoardSerialNumber()
	if err != nil {
		return res, err
	}

	cpuSerialNumber, err := mac.GetCpuSerialNumber()
	if err != nil {
		return res, err
	}

	diskSerialNumber, err := mac.GetDiskSerialNumber()
	if err != nil {
		return res, err
	}

	macAddr, err := GetMACAddress()
	if err != nil {
		return res, err
	}

	machineData := types.MachineInformation{
		UUID:              platformUUID,
		BoardSerialNumber: boardSerialNumber,
		CpuSerialNumber:   cpuSerialNumber,
		DiskSerialNumber:  diskSerialNumber,
		Mac:               macAddr,
	}
	return machineData, nil
}

func (mac MacMachine) GetBoardSerialNumber() (data string, err error) {
	sysInfo, err := mac.GetMacSysInfo()
	if err != nil {
		return "", err
	}
	return sysInfo.BoardSerialNumber, err
}

func (mac MacMachine) GetUUID() (UUID string, err error) {
	sysInfo, err := mac.GetMacSysInfo()
	if err != nil {
		return "", err
	}
	return sysInfo.UUID, err
}

func (mac MacMachine) GetCpuSerialNumber() (cpuId string, err error) {
	sysInfo, err := mac.GetMacSysInfo()
	if err != nil {
		return "", err
	}
	return sysInfo.CpuSerialNumber, err
}

// 获取硬盘编号
func (mac MacMachine) GetDiskSerialNumber() (serialNumber string, err error) {
	// lsblk --nodeps -no serial /dev/sda
	var cmd *exec.Cmd
	cmd = exec.Command("lsblk", "--nodeps", "-no", "serial", "/dev/sda")

	var out bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = os.Stderr

	err = cmd.Start()
	if err != nil {
		return "", err
	}
	err = cmd.Wait()
	if err == nil {
		return out.String(), nil
	} else {
		return "", err
	}
}

func (mac MacMachine) GetMacSysInfo() (data types.MachineInformation, err error) {
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
		macMachineData, err = mac.macXmlToData(out.String())
		if err == nil {
			macMachineData.CpuSerialNumber, _ = mac.getCpuSerialNumberBase()
		}
		return macMachineData, nil
	} else {
		return types.MachineInformation{}, err
	}
}

func (MacMachine) macXmlToData(xmlcontent string) (types.MachineInformation, error) {
	x := macXmlStruct{}
	err := xml.Unmarshal([]byte(xmlcontent), &x)
	if err != nil {
		return types.MachineInformation{}, err
	} else {
		count := len(x.Array.Dict.Array.Dict.String)
		serialData := types.MachineInformation{
			UUID:              x.Array.Dict.Array.Dict.String[count-2],
			BoardSerialNumber: x.Array.Dict.Array.Dict.String[count-1],
			CpuSerialNumber:   "",
		}
		return serialData, nil
	}
}

func (mac MacMachine) getCpuSerialNumberBase() (cpuId string, err error) {
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
