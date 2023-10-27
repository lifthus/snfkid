package machine

import (
	"testing"

	"github.com/lifthus/snofkid/pkg/snofkid"
)

func TestNewMachine(t *testing.T) {
	// invalid arguments
	epochs := []int64{-10, -5, -1, 0, 1, 10}
	machineIDs := []int64{-5, 0, 1, -1, -5, snofkid.MaxMachineID + 1}
	for i := 0; i < len(epochs); i++ {
		_, err := NewMachine(epochs[i], machineIDs[i])
		if err == nil {
			t.Errorf("NewMachine(%d, %d) is expected to fail but succeeded", epochs[i], machineIDs[i])
		}
	}
	// valid arguments
	epochs = []int64{0, 2, 1234, 123456, 1234654312, 10}
	machineIDs = []int64{0, 123, 0, 23, 555, snofkid.MaxMachineID}
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

// MockID generates a mock snowflake ID.
// Note that each argument must be within the range specified below:
// - timestamp: 0 <= timestamp <= 2^41 - 1 (2199023255551)
// - machineID: 0 <= machineID <= 2^10 - 1 (1023)
// - sequence: 0 <= sequence <= 2^12 - 1 (4095)
func MockSnowflakeID(timestamp, machineID, sequence int64) snofkid.SnowflakeID {
	return snofkid.SnowflakeID((timestamp << (snofkid.MachineIDBits + snofkid.SequenceBits)) | (machineID << snofkid.SequenceBits) | sequence)
}
