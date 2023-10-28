package snofkid

import (
	"testing"
)

func TestRawTimestampGetter(t *testing.T) {
	target := MockSnowflakeID(123456789, 1, 10)
	if res := target.RawTimestamp(); res != 123456789 {
		t.Errorf("Timestamp is expected to be %d but got %d", 123456789, res)
	}
}

func TestTimeGetter(t *testing.T) {
	epoch := TwitterEpoch
	target := MockSnowflakeID(123456789, 1, 10)
	if res := target.Time(epoch); res.UnixMilli() != 123456789+epoch {
		t.Errorf("Time UnixMilli is expected to be %d but got %d", 123456789+epoch, res.UnixMilli())
	}
}

func TestMachineIDGetter(t *testing.T) {
	target := MockSnowflakeID(123456789, 123, 10)
	if res := target.MachineID(); res != 123 {
		t.Errorf("MachineID is expected to be %d but got %d", 123, res)
	}
}

func TestSequenceGetter(t *testing.T) {
	target := MockSnowflakeID(123456789, 1, 123)
	if res := target.Sequence(); res != 123 {
		t.Errorf("Sequence is expected to be %d but got %d", 123, res)
	}
}

// MockID generates a mock snowflake ID.
// Note that each argument must be within the range specified below:
// - timestamp: 0 <= timestamp <= 2^41 - 1 (2199023255551)
// - machineID: 0 <= machineID <= 2^10 - 1 (1023)
// - sequence: 0 <= sequence <= 2^12 - 1 (4095)
func MockSnowflakeID(timestamp, machineID, sequence int64) SnowflakeID {
	return SnowflakeID((timestamp << (MachineIDBits + SequenceBits)) | (machineID << SequenceBits) | sequence)
}
