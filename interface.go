package machine_code

import "github.com/super-l/machine-code/types"

type OsMachineInterface interface {
	GetMachine() (types.MachineInformation, []error)
	GetBoardSerialNumber() (string, error)
	GetCpuSerialNumber() (string, error)
	GetDiskSerialNumber() (string, error)
	GetUUID() (string, error)
}
