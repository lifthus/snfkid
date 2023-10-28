package snofkid

// NewMachine returns a new SnowflakeMachine based on the given epoch and machine ID.

func NewMachine(epoch int64, machineID int64) (*SnowflakeMachine, error) {
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

// New generates a new SnowflakeID based on the given epoch and machine ID.
func (m *SnowflakeMachine) New() int64 {
	return 0
}

// Validate validates the sign bit and the machine ID of the given SnowflakeID.
func (m *SnowflakeMachine) Validate(sfid int64) bool {
	return false
}

func (m SnowflakeMachine) Epoch() int64 {
	return m.epoch
}

func (m SnowflakeMachine) MachineID() int64 {
	return m.machineID
}
