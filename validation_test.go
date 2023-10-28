package snofkid

import "testing"

func TestIsEpochValid(t *testing.T) {
	// invalid cases
	if IsEpochValid(-100) || IsEpochValid(-1) {
		t.Errorf("epoch lesser than 0 is judged as valid")
	}
	// valid cases
	if !IsEpochValid(0) {
		t.Errorf("valid epoch is judged as invalid")
	}
}

func TestIsTimestampValid(t *testing.T) {
	invalidTimestamps := []int64{-100, -1, MaxTimestamp + 1}
	for _, ts := range invalidTimestamps {
		if IsTimestampValid(ts) {
			t.Errorf("invalid timestamp %d is judged as valid", ts)
		}
	}
	validTimestamps := []int64{0, MaxTimestamp}
	for _, ts := range validTimestamps {
		if !IsTimestampValid(ts) {
			t.Errorf("valid timestamp %d is judged as invalid", ts)
		}
	}
}

func TestIsMachineIDValid(t *testing.T) {
	invalidMachineIDs := []int64{-100, -1, MaxMachineID + 1}
	for _, mid := range invalidMachineIDs {
		if IsMachineIDValid(mid) {
			t.Errorf("invalid macine ID %d is judged as valid", mid)
		}
	}
	validMachineIDs := []int64{0, MaxMachineID}
	for _, mid := range validMachineIDs {
		if !IsMachineIDValid(mid) {
			t.Errorf("valid machine ID %d is judged as invalid", mid)
		}
	}
}

func TestIsSequenceValid(t *testing.T) {
	invalidSequence := []int64{-100, -1, MaxSequence + 1}
	for _, seq := range invalidSequence {
		if IsSequenceValid(seq) {
			t.Errorf("invalid sequence %d is judged as valid", seq)
		}
	}
	validSequence := []int64{0, MaxSequence}
	for _, seq := range validSequence {
		if !IsSequenceValid(seq) {
			t.Errorf("valid sequence %d is judged as invalid", seq)
		}
	}
}
