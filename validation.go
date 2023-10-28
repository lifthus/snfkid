package snofkid

func IsEpochValid(epoch int64) bool {
	return epoch >= 0
}

func IsTimestampValid(timestamp int64) bool {
	return timestamp >= 0 && timestamp <= MaxTimestamp
}

func IsMachineIDValid(machineID int64) bool {
	return machineID >= 0 && machineID <= MaxMachineID
}

func IsSequenceValid(sequence int64) bool {
	return sequence >= 0 && sequence <= MaxSequence
}
