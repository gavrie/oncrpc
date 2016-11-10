package types

import "errors"

var (
	ErrInvalidEnumOpt          = errors.New("Invalid enum option")
	ErrArrayTooLarge           = errors.New("Array too large")
	ErrArbitratorValueMismatch = errors.New("Union arbitrator does not match value type")
	ErrMessageDenied           = errors.New("RPC message denied by server")
)
