package common

import "golang.org/x/exp/constraints"

type Iterator[T any] interface {
	Next() (T, bool)
	Prev() (T, bool)
	Value() (T, bool)
}

type Condition[T any] func(item T) bool

type Action[T any] func(item T)

type Container[T any] interface {
	Size() uint

	First() Iterator[T]
	Last() Iterator[T]

	Exists(cond Condition[T]) bool
	ForAll(cond Condition[T]) bool

	Filter(cond Condition[T])
	FilterNot(cond Condition[T])

	ForEach(action Action[T])
}

func defaultcmp[T constraints.Ordered](a, b T) int {
	if a < b {
		return -1
	} else if a > b {
		return 1
	}
	return 0
}
