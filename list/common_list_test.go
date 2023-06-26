package list

import (
	"common"

	"github.com/stretchr/testify/assert"
)

func checkListEmpty(list List[int], assert *assert.Assertions) {
	_, ok := list.First().Value()
	assert.False(ok, "Check list empty: First")
	_, ok = list.Last().Value()
	assert.False(ok, "Check list empty: Last")
	assert.Equal(uint(0), list.Size(), "Check list size")
}

func checkIteratorNext(it common.Iterator[int], value int, assert *assert.Assertions) {
	next_value, ok := it.Next()
	assert.True(ok, "Check not nil")
	assert.Equal(value, next_value, "Check value")
}

func checkIteratorPrev(it common.Iterator[int], value int, assert *assert.Assertions) {
	next_value, ok := it.Prev()
	assert.True(ok, "Check not nil")
	assert.Equal(value, next_value, "Check value")
}

func checkIteratorEndOfList(it common.Iterator[int], assert *assert.Assertions) {
	_, ok := it.Value()
	assert.False(ok, "Check no more elements")
}

func PushFront(list List[int], assert *assert.Assertions) {
	checkListEmpty(list, assert)

	for i := 1; i <= 10; i++ {
		list.PushFront(i)

		value, ok := list.First().Value()
		if assert.True(ok, "Check first not nil") {
			assert.Equal(i, value, "Check first value")
		}
	}
	assert.Equal(uint(10), list.Size(), "Check list size")

	value, ok := list.Last().Value()
	assert.True(ok, "Check last not nil")
	assert.Equal(1, value, "Check last value")
}

func PushBack(list List[int], assert *assert.Assertions) {
	checkListEmpty(list, assert)

	for i := 1; i <= 10; i++ {
		list.PushBack(i)

		value, ok := list.Last().Value()
		if assert.True(ok, "Check last not nil") {
			assert.Equal(i, value, "Check last value")
		}
	}
	assert.Equal(uint(10), list.Size(), "Check list size")

	value, ok := list.First().Value()
	assert.True(ok, "Check first not nil")
	assert.Equal(1, value, "Check first value")
}

func PopFront(list List[int], assert *assert.Assertions) {
	for i := 1; i <= 10; i++ {
		list.PushBack(i)
	}

	assert.Equal(uint(10), list.Size(), "Check list size")
	for i := 1; i <= 10; i++ {
		value, ok := list.PopFront()
		assert.True(ok, "Check pop not nil")
		assert.Equal(i, value, "Check value")
	}
	assert.Equal(uint(0), list.Size(), "Check list size")
	_, ok := list.PopFront()
	assert.False(ok, "Check no more elements")

	checkListEmpty(list, assert)
}

func PopBack(list List[int], assert *assert.Assertions) {
	for i := 1; i <= 10; i++ {
		list.PushFront(i)
	}

	assert.Equal(uint(10), list.Size(), "Check list size")
	for i := 1; i <= 10; i++ {
		value, ok := list.PopBack()
		assert.True(ok, "Check not nil")
		assert.Equal(i, value, "Check value")
	}
	assert.Equal(uint(0), list.Size(), "Check list size")

	_, ok := list.PopBack()
	assert.False(ok, "Check no more elements")

	checkListEmpty(list, assert)
}

func Iterate(list List[int], assert *assert.Assertions) {
	//Iterate forward over empty list
	checkListEmpty(list, assert)
	it := list.First()
	checkIteratorEndOfList(it, assert)

	//Iterate backward over empty list
	it = list.Last()
	checkIteratorEndOfList(it, assert)

	//Insert numbers from 1 to 200
	for i := 100; i >= 1; i-- {
		list.PushFront(i)
		list.PushBack(201 - i)
	}
	assert.Equal(uint(200), list.Size(), "Check list size")

	//Iterate forward
	it = list.First()
	for i := 1; i <= 200; i++ {
		checkIteratorNext(it, i, assert)
	}
	checkIteratorEndOfList(it, assert)

	//Iterate backward
	it = list.Last()
	for i := 200; i >= 1; i-- {
		checkIteratorPrev(it, i, assert)
	}
	checkIteratorEndOfList(it, assert)
}

func Reverse(list List[int], assert *assert.Assertions) {
	//Reverse empty list
	checkListEmpty(list, assert)
	list.Reverse()
	checkListEmpty(list, assert)

	//Insert numbers from 1 to 200
	for i := 100; i >= 1; i-- {
		list.PushFront(i)
		list.PushBack(201 - i)
	}

	//Reverse list
	list.Reverse()
	assert.Equal(uint(200), list.Size(), "Check list size")

	//Iterate forward
	it := list.First()
	for i := 200; i >= 1; i-- {
		checkIteratorNext(it, i, assert)
	}
	_, ok := it.Next()
	assert.False(ok, "Check no more elements")

	//Iterate backward
	it = list.Last()
	for i := 1; i <= 200; i++ {
		checkIteratorPrev(it, i, assert)
	}
	checkIteratorEndOfList(it, assert)
}

func Exists(list List[int], assert *assert.Assertions) {
	//If the list is empty, it should return false
	assert.Equal(false, list.Exists(func(x int) bool {
		if x%2 == 0 {
			return true
		} else {
			return false
		}
	}), "Check empty list condition")

	//Insert numbers from 1 to 10
	for i := 1; i <= 10; i++ {
		list.PushBack(i)
	}

	//Check if there is an even number
	assert.Equal(true, list.Exists(func(x int) bool {
		if x%2 == 0 {
			return true
		} else {
			return false
		}
	}), "Check true condition")

	//Check if there is a number bigger that 200
	assert.Equal(false, list.Exists(func(x int) bool {
		if x >= 200 {
			return true
		} else {
			return false
		}
	}), "Check false condition")
}

func ForAll(list List[int], assert *assert.Assertions) {
	//If the list is empty, it should return true
	assert.Equal(true, list.ForAll(func(x int) bool {
		if x%2 == 0 {
			return true
		} else {
			return false
		}
	}), "Check empty list condition")

	//Insert numbers from 1 to 10
	for i := 1; i <= 10; i++ {
		list.PushBack(i)
	}

	//Check if all numbers are even
	assert.Equal(false, list.ForAll(func(x int) bool {
		if x%2 == 0 {
			return true
		} else {
			return false
		}
	}), "Check false condition")

	//Check if all numbers are smaller than 200
	assert.Equal(true, list.ForAll(func(x int) bool {
		if x < 200 {
			return true
		} else {
			return false
		}
	}), "Check true condition")
}

func Filter(list List[int], assert *assert.Assertions) {
	//Filter empty list
	list.Filter(func(x int) bool {
		if x%2 == 0 {
			return true
		} else {
			return false
		}
	})
	checkListEmpty(list, assert)

	//Insert numbers from 1 to 10
	for i := 1; i <= 10; i++ {
		list.PushBack(i)
	}

	//Leave only even numbers
	list.Filter(func(x int) bool {
		if x%2 == 0 {
			return true
		} else {
			return false
		}
	})
	assert.Equal(uint(5), list.Size(), "Check list size")
	it := list.First()
	for i := 2; i <= 10; i += 2 {
		checkIteratorNext(it, i, assert)
	}
	checkIteratorEndOfList(it, assert)

	//Remove all elements
	list.Filter(func(x int) bool {
		if x%2 != 0 {
			return true
		} else {
			return false
		}
	})
	//Check list is empty
	checkListEmpty(list, assert)
}

func FilterNot(list List[int], assert *assert.Assertions) {
	//Filter empty list
	list.FilterNot(func(x int) bool {
		if x%2 == 0 {
			return true
		} else {
			return false
		}
	})
	checkListEmpty(list, assert)

	//Insert numbers from 1 to 10
	for i := 1; i <= 10; i++ {
		list.PushBack(i)
	}

	//Leave only even numbers
	list.FilterNot(func(x int) bool {
		if x%2 != 0 {
			return true
		} else {
			return false
		}
	})
	assert.Equal(uint(5), list.Size(), "Check list size")
	it := list.First()
	for i := 2; i <= 10; i += 2 {
		checkIteratorNext(it, i, assert)
	}
	checkIteratorEndOfList(it, assert)

	//Remove all elements
	list.FilterNot(func(x int) bool {
		if x%2 == 0 {
			return true
		} else {
			return false
		}
	})
	//Check list is empty
	checkListEmpty(list, assert)
}

func ForEach(list List[int], assert *assert.Assertions) {
	//Test with empty list
	a := 0
	list.ForEach(func(x int) {
		a += x
	})

	assert.Equal(0, a, list.Size(), "Check result for empty list")

	//Insert numbers from 1 to 10
	for i := 1; i <= 10; i++ {
		list.PushBack(i)
	}

	a = 0
	list.ForEach(func(x int) {
		a += x
	})

	assert.Equal(55, a, list.Size(), "Check result for non empty list")
}
