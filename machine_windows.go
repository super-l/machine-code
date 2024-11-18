//go:build windows
// +build windows

package machine_code

import "github.com/super-l/machine-code/goos"

func init() {
	Machine, MachineErr = goos.WindowsMachine{}.GetMachine()
}
