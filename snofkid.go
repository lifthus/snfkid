package snofkid

import "time"

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

// RawTimestamp returns the timestamp part of the ID in the unit of millisecond.
// Note that the returned timestamp doesn't reflect epoch time.
func RawTimestamp(snofkid int64) int64 {
	return (snofkid & TimestampMask >> (MachineIDBits + SequenceBits))
}

// Time returns the actual time with the reflection of given epoch.
func Time(snofkid, epoch int64) time.Time {
	return time.UnixMilli(RawTimestamp(snofkid) + epoch)
}

func MachineID(snofkid int64) int64 {
	return (snofkid & MachineIDMask >> SequenceBits)
}

func Sequence(snofkid int64) int64 {
	return snofkid & SequenceMask
}
