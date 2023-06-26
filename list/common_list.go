package list

import "common"

type List[T any] interface {
	common.Container[T]

	PushFront(elem T)
	PopFront() (T, bool)
	PushBack(elem T)

	PopBack() (T, bool)
	Reverse()
}
