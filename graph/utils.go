package graph

import (
	"math/big"
	"strconv"
	"time"
)

func ParseUInt(value string) uint {
	parsed, err := strconv.ParseUint(value, 10, 64)
	if err != nil {
		return 0
	}
	return uint(parsed)
}

func ParseBigInt(value string) *big.Int {
	b := big.NewInt(0)
	v, ok := b.SetString(value, 10)
	if !ok {
		return big.NewInt(0)
	}
	return v
}

func ParseUnixTime(value string) time.Time {
	timestamp := ParseUInt(value)
	return time.Unix(int64(timestamp), 0)
}
