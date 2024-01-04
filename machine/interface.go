package machine

import "github.com/super-l/machine-code/machine/types"

// 定义一个接口
type OsMachineInterface interface {
	GetMachine() types.Information
	GetBoardSerialNumber() (string, error)
	GetPlatformUUID() (string, error)
	GetCpuSerialNumber() (string, error)
}
