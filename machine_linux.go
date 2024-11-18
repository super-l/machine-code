//go:build linux
// +build linux

package machine_code

import "github.com/super-l/machine-code/goos"

func init() {
	Machine, MachineErr = goos.LinuxMachine{}.GetMachine()
}
