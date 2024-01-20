package concurrent

import (
	"sync/atomic"

	"github.com/smallnest/goroutine"
)

type Exchanger[T any] struct {
	leftGoID, rightGoID int64
	left, right         chan T
}

// NewExchanger creates a new exchanger.
func NewExchanger[T any]() *Exchanger[T] {
	return &Exchanger[T]{
		leftGoID:  -1,
		rightGoID: -1,
		left:      make(chan T, 1),
		right:     make(chan T, 1),
	}
}

func (e *Exchanger[T]) Exchange(value T) T {
	goid := goroutine.ID()

	// left goroutine
	isLeft := atomic.CompareAndSwapInt64(&e.leftGoID, -1, goid)
	if !isLeft {
		isLeft = atomic.LoadInt64(&e.leftGoID) == goid
	}
	if isLeft {
		e.right <- value // send value to right
		return <-e.left  // wait for value from right
	}

	// right goroutine
	isRight := atomic.CompareAndSwapInt64(&e.rightGoID, -1, goid)
	if !isRight {
		isRight = atomic.LoadInt64(&e.rightGoID) == goid
	}
	if isRight {
		e.left <- value  // send value to left
		return <-e.right // wait for value from left
	}

	// other goroutine
	panic("sync: exchange called from neither left nor right goroutine")
}
