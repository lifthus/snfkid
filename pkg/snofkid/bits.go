package snofkid

const (
	Int64TopBit   int64 = -9223372036854775808 // 1 << 63
	Int64Max      int64 = 9223372036854775807  // 1 << 63 - 1
	Int64One      int64 = 1                    // 1
	Int64MinusOne int64 = -1                   // 0 - 1
)

func TimestampFromInt64(sfid int64) int64 {
	return sfid >> 22
}

func ClearInt64TopBit(sfid int64) int64 {
	return sfid &^ Int64TopBit
}
