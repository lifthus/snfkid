package machine

import "github.com/lifthus/snofkid/pkg/snofkid"

func DefaultMachine(machineID int64) *SnowflakeMachine {
	return NewMachine(snofkid.TwitterEpoch, machineID)
}

func NewMachine(epoch int64, machineID int64) *SnowflakeMachine {
	return &SnowflakeMachine{
		epoch:     epoch,
		machineID: machineID,
	}
}

type SnowflakeMachine struct {
	epoch     int64
	machineID int64
}

func (m *SnowflakeMachine) From(sfid int64) (snofkid.SnowflakeID, error) {
	return 0, nil
}

func (m *SnowflakeMachine) New() snofkid.SnowflakeID {
	return 0
}
