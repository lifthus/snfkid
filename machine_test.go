package snofkid

import (
	"testing"
)

func TestNewMachine(t *testing.T) {
	// invalid arguments
	epochs := []int64{-10, -5, -1, 0, 1, 10}
	machineIDs := []int64{-5, 0, 1, -1, -5, MaxMachineID + 1}
	for i := 0; i < len(epochs); i++ {
		_, err := NewMachine(epochs[i], machineIDs[i])
		if err == nil {
			t.Errorf("NewMachine(%d, %d) is expected to fail but succeeded", epochs[i], machineIDs[i])
		}
	}
	// valid arguments
	epochs = []int64{0, 2, 1234, 123456, 1234654312, 10}
	machineIDs = []int64{0, 123, 0, 23, 555, MaxMachineID}
	for i := 0; i < len(epochs); i++ {
		_, err := NewMachine(epochs[i], machineIDs[i])
		if err != nil {
			t.Errorf("NewMachine(%d, %d) is expected to succeed but failed", epochs[i], machineIDs[i])
		}
	}
}

func TestMachineIDFrom(t *testing.T) {
}

func TestMachineNewID(t *testing.T) {

}
