package snofkid

const (
	TwitterEpoch int64 = 1288834974657

	TimestampBits = 41
	MachineIDBits = 10
)

func New(machineID int64) int64 {
	m := &SnowflakeMachine{
		Epoch: TwitterEpoch,
	}
	return m.New()
}

type SnowflakeMachine struct {
	Epoch     int64
	MachineID int
}

func (m *SnowflakeMachine) New() int64 {
	return 0
}
