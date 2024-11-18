package machine_code

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestGetMachine(t *testing.T) {
	if MachineErr != nil {
		fmt.Println(MachineErr)
		return
	}

	machineJson, _ := json.Marshal(Machine)
	fmt.Println(string(machineJson))
}
