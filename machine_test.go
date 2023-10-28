package snofkid

import (
	"testing"
	"time"
)

const (
	TestEpoch int64 = TwitterEpoch + 123456789
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
	mch, err := NewMachine(TestEpoch, 123)
	if err != nil {
		panic(err)
	}

	snowflakes, actualMillis := generateSnowflakesWithMilliSecsGenerated(mch)

	for i := 0; i < len(snowflakes); i++ {
		if absoluteDiff(RawTimestamp(snowflakes[i])+TestEpoch, actualMillis[i]) > 2 {
			t.Errorf("inaccurate timestamp generated(diff > 2ms):%d and %d", snowflakes[i], actualMillis[i])
			t.FailNow()
		}
	}
}

func TestMachineNewSnowflakesMachineIDs(t *testing.T) {
}

func TestMachineNewSnowflakesSequences(t *testing.T) {
}

func TestMachineValidateSnowflake(t *testing.T) {
	mch, err := NewMachine(123456789, 123)
	if err != nil {
		panic(err)
	}
	sfid := from(123, 123, 10)
	if !mch.Validate(sfid) {
		t.Errorf("Snowflake %d should be valid with test machine", sfid)
	}
	sfid = from(0, 122, 123)
	if mch.Validate(sfid) {
		t.Errorf("Snowflake %d should be invalid with test machine", sfid)
	}
}

func TestMachineTimeParser(t *testing.T) {
	mch, err := NewMachine(TestEpoch, 123)
	if err != nil {
		panic(err)
	}

	curMs := time.Now().UnixMilli()

	sfid := from(curMs-TestEpoch, 123, 10)
	if mch.ParseTime(sfid).UnixMilli() != curMs {
		t.Errorf("unix ms of Snowflake %d should be parsed as %d", sfid, curMs)
	}
}
