//go:build darwin
// +build darwin

package machine_code

import "github.com/super-l/machine-code/goos"

func init() {
	Machine, MachineErr = goos.MacMachine{}.GetMachine()
}
