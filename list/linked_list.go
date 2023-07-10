package list

import "common"

type linkedListNode[T any] struct {
	value T
	prev  *linkedListNode[T]
	next  *linkedListNode[T]
}

type linkedListIterator[T any] struct {
	node *linkedListNode[T]
}

func (it *linkedListIterator[T]) Next() (T, bool) {
	var void T
	if it.node != nil {
		aux := it.node
		it.node = it.node.next
		return aux.value, true
	} else {
		return void, false
	}
}

type linkedListReverseIterator[T any] struct {
	node *linkedListNode[T]
}

func (it *linkedListReverseIterator[T]) Next() (T, bool) {
	var void T
	if it.node != nil {
		aux := it.node
		it.node = it.node.prev
		return aux.value, true
	} else {
		return void, false
	}
}

func (it *linkedListIterator[T]) Value() (T, bool) {
	var void T
	if it.node != nil {
		return it.node.value, true
	} else {
		return void, false
	}
}

type linkedListImpl[T any] struct {
	first *linkedListNode[T]
	last  *linkedListNode[T]
	size  uint
}

func (list *linkedListImpl[T]) Size() uint {
	return list.size
}

func (list *linkedListImpl[T]) GetFirst() (T, bool) {
	var void T
	if list.last != nil {
		return list.first.value, true
	} else {
		return void, false
	}
}

func (list *linkedListImpl[T]) GetLast() (T, bool) {
	var void T
	if list.last != nil {
		return list.last.value, true
	} else {
		return void, false
	}
}

func (list *linkedListImpl[T]) CreateIterator() common.Iterator[T] {
	return &linkedListIterator[T]{list.first}
}

func (list *linkedListImpl[T]) CreateReverseIterator() common.Iterator[T] {
	return &linkedListReverseIterator[T]{list.last}
}

func (list *linkedListImpl[T]) PushFront(value T) {
	node := &linkedListNode[T]{value, nil, list.first}

	if list.first != nil {
		list.first.prev = node
	}
	if list.last == nil {
		list.last = node
	}
	list.first = node
	list.size++
}

func (list *linkedListImpl[T]) PushBack(value T) {
	node := &linkedListNode[T]{value, list.last, nil}

	if list.last != nil {
		list.last.next = node
	}
	if list.first == nil {
		list.first = node
	}
	list.last = node
	list.size++
}

func (list *linkedListImpl[T]) PopFront() (T, bool) {
	var void T
	if list.first != nil {
		value, ok := list.GetFirst()
		removeLinkedListNode(list, list.first)
		return value, ok
	} else {
		return void, false
	}
}

func (list *linkedListImpl[T]) PopBack() (T, bool) {
	var void T
	if list.last != nil {
		value, ok := list.GetLast()
		removeLinkedListNode(list, list.last)
		return value, ok
	} else {
		return void, false
	}
}

func (list *linkedListImpl[T]) Reverse() {
	node := list.first
	var aux *linkedListNode[T]

	for node != nil {
		aux = node.next
		node.next = node.prev
		node.prev = aux
		node = aux
	}
	aux = list.first
	list.first = list.last
	list.last = aux
}

func (list *linkedListImpl[T]) Exists(cond common.Condition[T]) bool {
	exists := false
	for it := list.first; it != nil && !exists; it = it.next {
		if cond(it.value) {
			exists = true
		}
	}
	return exists
}

func (list *linkedListImpl[T]) ForAll(cond common.Condition[T]) bool {
	all_true := true
	for it := list.first; it != nil && all_true; it = it.next {
		if !cond(it.value) {
			all_true = false
		}
	}
	return all_true
}

func (list *linkedListImpl[T]) Filter(cond common.Condition[T]) {
	it := list.first
	for it != nil {
		if !cond(it.value) {
			aux := it
			it = it.next
			removeLinkedListNode(list, aux)
		} else {
			it = it.next
		}
	}
}

func (list *linkedListImpl[T]) FilterNot(cond common.Condition[T]) {
	var cond_not common.Condition[T] = func(item T) bool { return !cond(item) }
	list.Filter(cond_not)
}

func (list *linkedListImpl[T]) ForEach(action common.Action[T]) {
	for it := list.first; it != nil; it = it.next {
		action(it.value)
	}
}

func CreateLinkedList[T any]() List[T] {
	list := linkedListImpl[T]{}
	return &list
}

func removeLinkedListNode[T any](list *linkedListImpl[T], node *linkedListNode[T]) {
	if node != nil {
		if list.first != node {
			node.prev.next = node.next
		} else {
			list.first = node.next
		}
		if list.last != node {
			node.next.prev = node.prev
		} else {
			list.last = node.prev
		}
		node.next = nil
		node.prev = nil
		list.size--
	}
}
