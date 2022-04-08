package models

import (
	"math/big"
	"time"
)

type (
	Transfer struct {
		From        string
		To          string
		MethodId    string
		Contract    string
		Amount      *big.Int
		Timestamp   time.Time
		Transaction string
	}
)
