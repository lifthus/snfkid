package snofkid

import (
	"testing"
	"time"
)

const (
	TestEpoch     int64 = TwitterEpoch + 123456789
	TestMachineID int64 = 257
)

func TestNewMachine(t *testing.T) {
	_, err := NewMachine(-1, 0)
	if err == nil {
		t.Errorf("NewMachine should return error when epoch is invalid")
	}
	_, err = NewMachine(-1, -1)
	if err == nil {
		t.Errorf("NewMachine should return error when epoch and machine id are invalid")
	}
	_, err = NewMachine(0, -1)
	if err == nil {
		t.Errorf("NewMachine should return error when machine id is invalid")
	}
	_, err = NewMachine(0, MaxMachineID+1)
	if err == nil {
		t.Errorf("NewMachine should return error when machine id is invalid")
	}
	_, err = NewMachine(0, 0)
	if err != nil {
		t.Errorf("NewMachine shouldn't return error when epoch and machine id are valid")
	}
	_, err = NewMachine(0, MaxMachineID)
	if err != nil {
		t.Errorf("NewMachine shouldn't return error when epoch and machine id are valid")
	}
}

func TestMachineNewSnowflakesTimestamps(t *testing.T) {
	snowflakes, actualMillis := snowflakes4096, actualMillis4096
	for i := 0; i < len(snowflakes); i++ {
		if absoluteDiff(RawTimestamp(snowflakes[i])+TestEpoch, actualMillis[i]) > 2 {
			t.Errorf("inaccurate timestamp generated(diff > 2ms):%d and %d", snowflakes[i], actualMillis[i])
			t.FailNow()
		}
	}
}

func TestMachineNewSnowflakesMachineIDs(t *testing.T) {
	snowflakes := snowflakes4096
	for i := 0; i < len(snowflakes); i++ {
		if MachineID(snowflakes[i]) != TestMachineID {
			t.Errorf("Snowflake %d is generated with wrong machine ID", snowflakes[i])
			t.FailNow()
		}
	}
}

func TestMachineNewSnowflakesDuplicates(t *testing.T) {
	snowflakes := snowflakes500ms
	if HasDuplicates(snowflakes) {
		t.Errorf("duplicate Snowflakes are generated in 500ms")
	}
}

func TestMachineValidateSnowflake(t *testing.T) {
	mch, err := NewMachine(TestEpoch, TestMachineID)
	if err != nil {
		panic(err)
	}
	sfid := from(123, TestMachineID, 10)
	if !mch.Validate(sfid) {
		t.Errorf("Snowflake %d should be valid with test machine", sfid)
	}
	sfid = from(0, TestMachineID-1, 123)
	if mch.Validate(sfid) {
		t.Errorf("Snowflake %d should be invalid with test machine", sfid)
	}
}

func TestMachineTimeParser(t *testing.T) {
	mch, err := NewMachine(TestEpoch, TestMachineID)
	if err != nil {
		panic(err)
	}

	curMs := time.Now().UnixMilli()

	sfid := from(curMs-TestEpoch, 123, 10)
	if mch.ParseTime(sfid).UnixMilli() != curMs {
		t.Errorf("unix ms of Snowflake %d should be parsed as %d", sfid, curMs)
	}
}
