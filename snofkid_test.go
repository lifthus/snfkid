package snofkid

import (
	"testing"
)

func TestSnowflakeIDFrom(t *testing.T) {
	sfid := SnowflakeIDFrom(123456789, 123, 10)
	switch {
	case RawTimestamp(sfid) != 123456789:
		fallthrough
	case MachineID(sfid) != 123:
		fallthrough
	case Sequence(sfid) != 10:
		t.Errorf("SnowflakeIDFrom generated invalid Snowflake ID")
	}
}

func TestRawTimestampGetter(t *testing.T) {
	target := SnowflakeIDFrom(123456789, 1, 10)
	if res := RawTimestamp(target); res != 123456789 {
		t.Errorf("Timestamp is expected to be %d but got %d", 123456789, res)
	}
}

func TestMachineIDGetter(t *testing.T) {
	target := SnowflakeIDFrom(123456789, 123, 10)
	if res := MachineID(target); res != 123 {
		t.Errorf("MachineID is expected to be %d but got %d", 123, res)
	}
}

func TestSequenceGetter(t *testing.T) {
	target := SnowflakeIDFrom(123456789, 1, 123)
	if res := Sequence(target); res != 123 {
		t.Errorf("Sequence is expected to be %d but got %d", 123, res)
	}
}
func TestMaxTimestampConst(t *testing.T) {
	calcedMaxTs := int64(-1) ^ (int64(-1) << TimestampBits)
	if MaxTimestamp != calcedMaxTs {
		t.Errorf("MaxTimestamp is expected to be %d but got %d", calcedMaxTs, MaxTimestamp)
	}
}

func TestMaxMachineIDConst(t *testing.T) {
	calcedMaxMID := -1 ^ (-1 << MachineIDBits)
	if MaxMachineID != calcedMaxMID {
		t.Errorf("MaxMachineID is expected to be %d but got %d", calcedMaxMID, MaxMachineID)
	}
}

func TestMaxSequenceConst(t *testing.T) {
	calcedMaxSeq := -1 ^ (-1 << SequenceBits)
	if MaxSequence != calcedMaxSeq {
		t.Errorf("MaxSequence is expected to be %d but got %d", calcedMaxSeq, MaxSequence)
	}
}
