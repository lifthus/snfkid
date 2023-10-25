package snofkid

import (
	"testing"
)

func TestDefaultGenerateSnowflakeID(t *testing.T) {
	// curMs := time.Now().UnixMilli()
	// snfid := New(6)
	// if TimestampFromInt64(snfid) - TwitterEpoch - curMs > 1 {
	// 	t.Errorf()
	// }
}

func TestTimestampFromInt64(t *testing.T) {
	target := Int64Max
	validTimestamp := ClearInt64TopBit(target) >> 22
	if ts := TimestampFromInt64(target); ts != validTimestamp {
		t.Errorf("TimestampFromInt64 is expected to be %d but got %d", validTimestamp, ts)
	}
	target = 0
	validTimestamp = 0
	if ts := TimestampFromInt64(target); ts != validTimestamp {
		t.Errorf("TimestampFromInt64 is expected to be %d but got %d", validTimestamp, ts)
	}
}

func TestClearInt64TopBit(t *testing.T) {
	maxBit := Int64TopBit
	if result := ClearInt64TopBit(maxBit); result != 0 {
		t.Errorf("ClearInt64TopBit is expected to be %d but got %d", 0, result)
	}
}

func TestInt64TopBit(t *testing.T) {
	one := Int64One
	validInt64TopBit := one << 63
	if Int64TopBit != validInt64TopBit {
		t.Errorf("Int64TopBit is expected to be %d but got %d", validInt64TopBit, Int64TopBit)
	}
}

func TestInt64Max(t *testing.T) {
	one := Int64One
	validInt64Max := one<<63 - 1
	if Int64Max != validInt64Max {
		t.Errorf("Int64Max is expected to be %d but got %d", validInt64Max, Int64Max)
	}
}
