package balance

import "errors"

var (
	ErrInputNotSlice   = errors.New("input not slice")
	ErrNoAvailableNode = errors.New("no available node")
)
