package machine

import "github.com/lifthus/snofkid"

func ValidateMachine(epoch, machineID int64) error {
	switch {
	case snofkid.ValidateEpoch(epoch) != nil:
		return snofkid.ErrInvalidEpoch
	case snofkid.ValidateMachineID(machineID) != nil:
		return snofkid.ErrInvalidMachineID
	default:
		return nil
	}
}
