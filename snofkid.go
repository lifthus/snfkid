package snofkid

const (
	TwitterEpoch int64 = 1288834974657

	TimestampBits = 41
	MachineIDBits = 10
	SequenceBits  = 12

	MaxTimestamp = 2199023255551 // 2^41 - 1
	MaxMachineID = 1023          // 2^10 - 1
	MaxSequence  = 4095          // 2^12 - 1

	TimestampMask = MaxTimestamp << (MachineIDBits + SequenceBits)
	MachineIDMask = MaxMachineID << SequenceBits
	SequenceMask  = MaxSequence
)

// from generates a snowflake ID optimistically, that is, without validation of the given arguments.
// Note that each argument must be within the specified range following this to generate a valid Snowflake ID:
// - timestamp: 0 <= timestamp <= 2^41 - 1 (2199023255551)
// - machineID: 0 <= machineID <= 2^10 - 1 (1023)
// - sequence: 0 <= sequence <= 2^12 - 1 (4095)
func from(timestamp, machineID, sequence int64) int64 {
	return (timestamp << (MachineIDBits + SequenceBits)) | (machineID << SequenceBits) | sequence
}

// RawTimestamp returns the timestamp part of the ID in the unit of milliseconds, without any validation.
// Note that the returned timestamp doesn't reflect epoch time.
func RawTimestamp(snofkid int64) int64 {
	return (snofkid & TimestampMask >> (MachineIDBits + SequenceBits))
}

// MachineID returns the machine ID part of the Snowflake ID, without any validation.
func MachineID(snofkid int64) int64 {
	return (snofkid & MachineIDMask >> SequenceBits)
}

// Sequence returns the sequence part of the Snowflake ID, without any validation.
func Sequence(snofkid int64) int64 {
	return snofkid & SequenceMask
}
