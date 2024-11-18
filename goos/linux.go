//go:build linux
// +build linux

package goos

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/super-l/machine-code/types"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

type LinuxMachine struct{}

const DiskByIDPath = "/dev/disk/by-id/"
const DiskByUUIDPath = "/dev/disk/by-uuid/"

func (linux LinuxMachine) GetMachine() (res types.MachineInformation, err error) {
	platformUUID, err := linux.GetUUID()
	if err != nil {
		return res, err
	}
	boardSerialNumber, err := linux.GetBoardSerialNumber()
	if err != nil {
		return res, err
	}

	cpuSerialNumber, err := linux.GetCpuSerialNumber()
	if err != nil {
		return res, err
	}

	diskSerialNumber, err := linux.GetDiskSerialNumber()
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

// 主板序列码
func (linux LinuxMachine) GetBoardSerialNumber() (serialNumber string, err error) {
	// dmidecode -s system-serial-number
	var cmd *exec.Cmd
	cmd = exec.Command("dmidecode", "-s", "system-serial-number")
	var out bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = os.Stderr

	err = cmd.Start()
	if err != nil {
		return "", err
	}
	err = cmd.Wait()
	if err == nil {
		serial_number := out.String()
		serial_number = strings.Replace(serial_number, "\n", "", -1)
		return serial_number, nil
	} else {
		return "", err
	}
}

// CPU序列码
func (linux LinuxMachine) GetCpuSerialNumber() (cpuId string, err error) {
	// dmidecode -t processor |grep ID |head -1
	cmds := []*exec.Cmd{
		exec.Command("dmidecode", "-t", "processor"),
		exec.Command("grep", "ID"),
		exec.Command("head", "-1"),
	}
	cpuId, err = linux.execPipeLine(cmds...)
	cpuId = strings.TrimSpace(cpuId)
	cpuId = strings.Replace(cpuId, "ID: ", "", -1)
	cpuId = strings.Replace(cpuId, "\t", "", -1)
	cpuId = strings.Replace(cpuId, "\n", "", -1)
	cpuId = strings.Replace(cpuId, " ", "-", -1)
	return
}

// 获取硬盘编号
func (linux LinuxMachine) GetDiskSerialNumber() (serialNumber string, err error) {
	id, err1 := linux.GetDiskSerialNumberById()
	uuid, err2 := linux.GetDiskSerialNumberByUUID()

	if err1 != nil && err2 != nil {
		return "", err1
	}
	return fmt.Sprintf("id:%s uuid%s", id, uuid), nil
}

// 获取硬盘编号
func (linux LinuxMachine) GetDiskSerialNumberById() (serialNumber string, err error) {
	// ls /dev/disk/by-id -al
	entries, err := os.ReadDir(DiskByIDPath)
	if err != nil {
		return "", errors.New(fmt.Sprintf("Error reading disk directory: %s", err))
	}

	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}

		linkPath := filepath.Join(DiskByIDPath, entry.Name())
		_, err = os.Readlink(linkPath)
		if err != nil {
			continue
		}
		if strings.Contains(entry.Name(), "ata-") {
			parts := strings.Split(entry.Name(), "-")
			if len(parts) > 1 {
				serialNumber = parts[1]
			} else {
				serialNumber = entry.Name()
			}
		}
	}
	if serialNumber == "" {
		return "", errors.New("no data")
	}
	return serialNumber, nil
}

// 获取硬盘编号
func (linux LinuxMachine) GetDiskSerialNumberByUUID() (string, error) {
	var result []string
	entries, err := os.ReadDir(DiskByUUIDPath)
	if err != nil {
		return "", errors.New(fmt.Sprintf("Error reading disk directory:", err))
	}

	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}
		name := entry.Name()
		linkPath := filepath.Join(DiskByUUIDPath, entry.Name())
		_, err = os.Readlink(linkPath)
		if err != nil {
			continue
		}
		result = append(result, name)
	}

	if len(result) == 0 {
		return "", errors.New(fmt.Sprintf("no data"))
	}
	return strings.Join(result, "|"), nil
}

// 设备唯一UUID
func (linux LinuxMachine) GetUUID() (UUID string, err error) {
	// dmidecode -s system-uuid           UUID
	var cmd *exec.Cmd
	cmd = exec.Command("dmidecode", "-s", "system-uuid")
	var out bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = os.Stderr

	err = cmd.Start()
	if err != nil {
		return "", err
	}
	err = cmd.Wait()
	if err == nil {
		uuid := out.String()
		uuid = strings.Replace(uuid, "\n", "", -1)
		return uuid, nil
	} else {
		return "", err
	}
}

func (linux LinuxMachine) GetCpuSerialNumber2() (cpuId string, err error) {
	// dmidecode -t processor |grep ID |head -1
	var cmd *exec.Cmd
	cmd = exec.Command("dmidecode", "-t", "processor", "|grep ID |head -1")

	var out bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = os.Stderr

	err = cmd.Start()
	if err != nil {
		return "", err
	}
	err = cmd.Wait()
	if err == nil {
		uuid := out.String()
		//uuid = strings.Replace(uuid, "\n", "", -1)
		return uuid, nil
	} else {
		return "", err
	}
}

func (LinuxMachine) pipeline(cmds ...*exec.Cmd) (pipeLineOutput, collectedStandardError []byte, pipeLineError error) {
	if len(cmds) < 1 {
		return nil, nil, nil
	}

	var output bytes.Buffer
	var stderr bytes.Buffer

	last := len(cmds) - 1
	for i, cmd := range cmds[:last] {
		var err error
		if cmds[i+1].Stdin, err = cmd.StdoutPipe(); err != nil {
			return nil, nil, err
		}
		cmd.Stderr = &stderr
	}

	cmds[last].Stdout, cmds[last].Stderr = &output, &stderr

	for _, cmd := range cmds {
		if err := cmd.Start(); err != nil {
			return output.Bytes(), stderr.Bytes(), err
		}
	}

	for _, cmd := range cmds {
		if err := cmd.Wait(); err != nil {
			return output.Bytes(), stderr.Bytes(), err
		}
	}

	return output.Bytes(), stderr.Bytes(), nil
}

func (linux LinuxMachine) execPipeLine(cmds ...*exec.Cmd) (string, error) {
	output, stderr, err := linux.pipeline(cmds...)
	if err != nil {
		return "", err
	}

	if len(output) > 0 {
		return string(output), nil
	}

	if len(stderr) > 0 {
		return string(stderr), nil
	}
	return "", errors.New("no returns")
}
