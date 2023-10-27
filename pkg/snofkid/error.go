package snofkid

import "errors"

var ErrInvalidComponents = errors.New("invalid snowflake components")
var ErrInvalidEpoch = errors.New("invalid epoch")
var ErrInvalidTimestamp = errors.New("invalid timestamp")
var ErrInvalidMachineID = errors.New("invalid machine ID")
var ErrInvalidSequence = errors.New("invalid sequence")
