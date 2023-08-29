package ChannelX

import "github.com/smallnest/chanx"

func NewChannelX[T any](initInCapacity int, initOutCapacity int, initBufCapacity int) *chanx.UnboundedChan[T] {
	return chanx.NewUnboundedChanSize[T](initInCapacity, initOutCapacity, initBufCapacity)
}
