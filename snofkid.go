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

// SnowflakeID is a type alias of int64 which represents a Twitter Snowflake ID.
//
// It is strongly recommended not to generate any SnowflakeID directly without this package,
// because we want the variables of this type to always be valid not to return error type from some functions.
// That is, lots of SnowflakeID's methods assume that the ID is in a valid range, which means it is greater than or equal to 0.
// Therefore, An ID that is deliberately generated to be invalid will cause unexpected behavior without any notification.
type SnowflakeID int64

// RawTimestamp returns the timestamp part of the ID in the unit of millisecond.
// Note that the returned timestamp doesn't reflect epoch time.
func (id SnowflakeID) RawTimestamp() int64 {
	return (int64(id) & TimestampMask >> (MachineIDBits + SequenceBits))
}

// Time returns the actual time with the reflection of given epoch.
func (id SnowflakeID) Time(epoch int64) time.Time {
	return time.UnixMilli(id.RawTimestamp() + epoch)
}

func (id SnowflakeID) MachineID() int64 {
	return (int64(id) & MachineIDMask >> SequenceBits)
}

func (id SnowflakeID) Sequence() int64 {
	return int64(id) & SequenceMask
}
