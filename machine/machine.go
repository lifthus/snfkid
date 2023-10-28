package machine

import "github.com/lifthus/snofkid"

func NewDefaultMachine(machineID int64) (*SnowflakeMachine, error) {
	return NewMachine(snofkid.TwitterEpoch, machineID)
}

func NewMachine(epoch int64, machineID int64) (*SnowflakeMachine, error) {
	if err := ValidateMachine(epoch, machineID); err != nil {
		return nil, err
	}
	return &SnowflakeMachine{
		epoch:     epoch,
		machineID: machineID,
	}, nil
}

// SnowflakeMachine generates SnowflakeID based on the given epoch and machine ID.
type SnowflakeMachine struct {
	epoch     int64
	machineID int64
}

// From validates the given int64 value and converts it to a SnowflakeID if it is valid.
//
// The validity is determined by the sign and machine ID of the given int64 value.
func (m *SnowflakeMachine) From(sfid int64) (snofkid.SnowflakeID, error) {

	return 0, nil
}

// New generates a new SnowflakeID based on the given epoch and machine ID.
func (m *SnowflakeMachine) New() snofkid.SnowflakeID {
	return 0
}

func (m SnowflakeMachine) Epoch() int64 {
	return m.epoch
}

func (m SnowflakeMachine) MachineID() int64 {
	return m.machineID
}
