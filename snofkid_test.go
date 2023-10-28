package snofkid

import (
	"testing"
)

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
