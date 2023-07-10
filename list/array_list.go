package list

import (
	"common"
)

type arrayListIterator[T any] struct {
	index uint
	list  *arrayListImpl[T]
}

func (it *arrayListIterator[T]) Next() (T, bool) {
	var void T
	if it.list.size > 0 && it.index < it.list.capacity {
		index := it.index
		if it.index == it.list.last {
			it.index = it.list.capacity
		} else {
			it.index = (it.index + 1) % it.list.capacity
		}
		return it.list.values[index], true
	} else {
		return void, false
	}
}

type arrayListReverseIterator[T any] struct {
	index uint
	list  *arrayListImpl[T]
}

func (it *arrayListReverseIterator[T]) Next() (T, bool) {
	var void T
	if it.list.size > 0 && it.index < it.list.capacity {
		index := it.index
		if it.index == it.list.first {
			it.index = it.list.capacity
		} else {
			if it.index == 0 {
				it.index = it.list.capacity - 1
			} else {
				it.index--
			}
		}
		return it.list.values[index], true
	} else {
		return void, false
	}
}

type arrayListImpl[T any] struct {
	values   []T
	first    uint
	last     uint
	size     uint
	capacity uint
}

func (list *arrayListImpl[T]) Size() uint {
	return list.size
}

func (list *arrayListImpl[T]) GetFirst() (T, bool) {
	var void T
	if list.size > 0 {
		return list.values[list.first], true
	} else {
		return void, false
	}
}

func (list *arrayListImpl[T]) GetLast() (T, bool) {
	var void T
	if list.size > 0 {
		return list.values[list.last], true
	} else {
		return void, false
	}
}

func (list *arrayListImpl[T]) CreateIterator() common.Iterator[T] {
	return &arrayListIterator[T]{list.first, list}
}

func (list *arrayListImpl[T]) CreateReverseIterator() common.Iterator[T] {
	return &arrayListReverseIterator[T]{list.last, list}
}

func (list *arrayListImpl[T]) PushFront(value T) {
	if list.size == list.capacity {
		resize(list)
	}
	if list.first == list.capacity {
		list.first = 0
		list.last = 0
	} else if list.first == 0 {
		list.first = list.capacity - 1
	} else {
		list.first--
	}
	list.values[list.first] = value
	list.size++
}

func (list *arrayListImpl[T]) PushBack(value T) {
	if list.size == list.capacity {
		resize(list)
	}
	if list.last == list.capacity {
		list.first = 0
		list.last = 0
	} else {
		list.last = (list.last + 1) % list.capacity
	}
	list.values[list.last] = value
	list.size++
}

func (list *arrayListImpl[T]) PopFront() (T, bool) {
	var void T
	if list.size > 0 {
		value := list.values[list.first]
		list.values[list.first] = void
		list.size--
		if list.size > 0 {
			list.first = (list.first + 1) % list.capacity
		}
		return value, true
	} else {
		return void, false
	}
}

func (list *arrayListImpl[T]) PopBack() (T, bool) {
	var void T
	if list.size > 0 {
		value := list.values[list.last]
		list.values[list.last] = void
		list.size--
		if list.size > 0 {
			if list.last == 0 {
				list.last = list.capacity - 1
			} else {
				list.last--
			}
		}
		return value, true
	} else {
		return void, false
	}
}

func (list *arrayListImpl[T]) Reverse() {
	forward := list.first
	backward := list.last
	steps := list.size / 2

	for i := uint(0); i < steps; i++ {
		aux := list.values[backward]
		list.values[backward] = list.values[forward]
		list.values[forward] = aux
		forward = (forward + 1) % list.capacity
		if backward == 0 {
			backward = list.capacity - 1
		} else {
			backward--
		}
	}
}

func (list *arrayListImpl[T]) Exists(cond common.Condition[T]) bool {
	exists := false
	if list.size > 0 {
		finish := false
		for i := list.first; !finish && !exists; i = (i + 1) % list.capacity {
			if cond(list.values[i]) {
				exists = true
			}
			finish = (i == list.last)
		}
	}
	return exists
}

func (list *arrayListImpl[T]) ForAll(cond common.Condition[T]) bool {
	all_true := true
	if list.size > 0 {
		finish := false
		for i := list.first; !finish && all_true; i = (i + 1) % list.capacity {
			if !cond(list.values[i]) {
				all_true = false
			}
			finish = (i == list.last)
		}
	}
	return all_true
}

func (list *arrayListImpl[T]) Filter(cond common.Condition[T]) {
	var void T
	if list.size > 0 {
		empty_index := list.capacity
		finish := false

		for i := list.first; !finish; i = (i + 1) % list.capacity {
			if !cond(list.values[i]) {
				if empty_index == list.capacity {
					empty_index = i
				}
				list.size--
			} else {
				if empty_index != list.capacity {
					list.values[empty_index] = list.values[i]
					empty_index = (empty_index + 1) % list.capacity
				}
			}
			finish = (i == list.last)
		}

		//Clean empty nodes
		finish = false
		for i := empty_index; !finish; i = (i + 1) % list.capacity {
			list.values[i] = void
			finish = (i == list.last)
		}
		if list.size == 0 {
			list.first = list.capacity
			list.last = list.capacity
		} else {
			list.last = (list.first + list.size - 1) % list.capacity
		}
	}
}

func (list *arrayListImpl[T]) FilterNot(cond common.Condition[T]) {
	var cond_not common.Condition[T] = func(item T) bool { return !cond(item) }
	list.Filter(cond_not)
}

func (list *arrayListImpl[T]) ForEach(action common.Action[T]) {
	if list.size > 0 {
		finish := false
		for i := list.first; !finish; i = (i + 1) % list.capacity {
			action(list.values[i])
			finish = (i == list.last)
		}
	}
}

func CreateArrayList[T any]() List[T] {
	list := arrayListImpl[T]{}
	list.values = make([]T, 16)
	list.capacity = 16
	list.first = list.capacity
	list.last = list.capacity
	return &list
}

func resize[T any](list *arrayListImpl[T]) {
	new_capacity := list.capacity * 2
	new_values := make([]T, new_capacity)

	old_values_index := list.first
	for i := uint(0); i < list.size; i++ {
		new_values[i] = list.values[old_values_index]
		old_values_index = (old_values_index + 1) % list.capacity
	}
	list.capacity = new_capacity
	list.first = 0
	list.last = list.size - 1
	list.values = new_values
}
