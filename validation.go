package snofkid

func ValidateMachine(epoch, machineID int64) error {
	switch {
	case ValidateEpoch(epoch) != nil:
		return ErrInvalidEpoch
	case ValidateMachineID(machineID) != nil:
		return ErrInvalidMachineID
	default:
		return nil
	}
}

func ValidateEpoch(epoch int64) error {
	if epoch < 0 {
		return ErrInvalidEpoch
	}
	return nil
}

func ValidateSnowflakeComponents(timestamp, machineID, sequence int64) error {
	switch {
	case ValidateTimestamp(timestamp) != nil:
		return ErrInvalidTimestamp
	case ValidateMachineID(machineID) != nil:
		return ErrInvalidMachineID
	case ValidateSequence(sequence) != nil:
		return ErrInvalidSequence
	default:
		return nil
	}
}

func ValidateTimestamp(timestamp int64) error {
	if timestamp < 0 || timestamp > MaxTimestamp {
		return ErrInvalidTimestamp
	}
	return nil
}

func ValidateMachineID(machineID int64) error {
	if machineID < 0 || machineID > MaxMachineID {
		return ErrInvalidMachineID
	}
	return nil
}

func ValidateSequence(sequence int64) error {
	if sequence < 0 || sequence > MaxSequence {
		return ErrInvalidSequence
	}
	return nil
}
