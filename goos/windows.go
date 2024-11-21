//go:build windows
// +build windows

package goos

import (
	"errors"
	"fmt"
	"github.com/super-l/machine-code/types"
	"os/exec"
	"strings"
	"syscall"
)

type WindowsMachine struct{}

func (w WindowsMachine) GetMachine() (types.MachineInformation, []error) {
	var errs []error

	platformUUID, err := w.GetUUID()
	if err != nil {
		errs = append(errs, err)
	}
	boardSerialNumber, err := w.GetBoardSerialNumber()
	if err != nil {
		errs = append(errs, err)
	}

	cpuSerialNumber, err := w.GetCpuSerialNumber()
	if err != nil {
		errs = append(errs, err)
	}

	diskSerialNumber, err := w.GetDiskSerialNumber()
	if err != nil {
		errs = append(errs, err)
	}

	macAddr, err := GetMACAddress()
	if err != nil {
		errs = append(errs, err)
	}

	machineData := types.MachineInformation{
		UUID:              platformUUID,
		BoardSerialNumber: boardSerialNumber,
		CpuSerialNumber:   cpuSerialNumber,
		DiskSerialNumber:  diskSerialNumber,
		Mac:               macAddr,
	}
	return machineData, errs
}

// 获取主板编号
func (WindowsMachine) GetBoardSerialNumber() (string, error) {
	cmd := exec.Command("wmic", "baseboard", "get", "serialnumber")
	cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true} // @bestK 提交
	b, err := cmd.CombinedOutput()
	if err == nil {
		serialNumber := string(b)
		serialNumber = serialNumber[12 : len(serialNumber)-2]
		serialNumber = strings.ReplaceAll(serialNumber, "\n", "")
		serialNumber = strings.ReplaceAll(serialNumber, " ", "")
		serialNumber = strings.ReplaceAll(serialNumber, "\r", "")
		return serialNumber, nil
	}
	return "", err
}

// 获取硬盘编号
func (WindowsMachine) GetDiskSerialNumber() (serialNumber string, err error) {
	cmd := exec.Command("wmic", "diskdrive", "get", "serialnumber")
	cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true} // @bestK 提交
	output, err := cmd.Output()
	if err != nil {
		return "", err
	}
	result := strings.TrimSpace(string(output))
	lines := strings.Split(result, "\n")
	if len(lines) > 1 {
		return strings.TrimSpace(lines[1]), nil
	}
	return "", fmt.Errorf("diskdrive serial number not found")
}

// 获取系统UUID
func (WindowsMachine) GetUUID() (string, error) {
	// wmic csproduct get uuid   ||  wmic csproduct list full | findstr UUID
	var cmd *exec.Cmd
	cmd = exec.Command("wmic", "csproduct", "get", "uuid")
	cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true} // @bestK 提交
	b, err := cmd.CombinedOutput()

	var uuid string
	if err == nil {
		uuid = string(b)
		uuid = uuid[4 : len(uuid)-1]
		uuid = strings.ReplaceAll(uuid, "\n", "")
		uuid = strings.ReplaceAll(uuid, " ", "")
		uuid = strings.ReplaceAll(uuid, "\r", "")
		return uuid, nil
	}
	return "", errors.New("csproduct uuid not found")
}

// 获取CPU序列号
func (WindowsMachine) GetCpuSerialNumber() (string, error) {
	// wmic cpu get processorid
	var cpuid string
	cmd := exec.Command("wmic", "cpu", "get", "processorid")
	cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true} // @bestK 提交
	b, err := cmd.CombinedOutput()
	if err == nil {
		cpuid = string(b)
		cpuid = cpuid[12 : len(cpuid)-2]
		cpuid = strings.ReplaceAll(cpuid, "\n", "")
		cpuid = strings.ReplaceAll(cpuid, " ", "")
		cpuid = strings.ReplaceAll(cpuid, "\r", "")
		return cpuid, err
	}
	return "", errors.New("cpu processorid not found")
}
