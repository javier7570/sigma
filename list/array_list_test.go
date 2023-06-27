package list

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestArrayListPushFront(t *testing.T) {
	assert := assert.New(t)

	list := CreateArrayList[int]()
	PushFront(list, assert)
}

func TestArrayListPushBack(t *testing.T) {
	assert := assert.New(t)

	list := CreateArrayList[int]()
	PushBack(list, assert)
}

func TestArrayListPopFront(t *testing.T) {
	assert := assert.New(t)

	list := CreateArrayList[int]()
	PopFront(list, assert)
}

func TestArrayListPopBack(t *testing.T) {
	assert := assert.New(t)

	list := CreateArrayList[int]()
	PopBack(list, assert)
}

func TestArrayListIterate(t *testing.T) {
	assert := assert.New(t)

	list := CreateArrayList[int]()
	Iterate(list, assert)
}

func TestArrayListReverse(t *testing.T) {
	assert := assert.New(t)

	list := CreateArrayList[int]()
	Reverse(list, assert)
}

func TestArrayListExists(t *testing.T) {
	assert := assert.New(t)

	list := CreateArrayList[int]()
	Exists(list, assert)

}

func TestArrayListForAll(t *testing.T) {
	assert := assert.New(t)

	list := CreateArrayList[int]()
	ForAll(list, assert)
}

func TestArrayListFilter(t *testing.T) {
	assert := assert.New(t)

	list := CreateArrayList[int]()
	Filter(list, assert)
}

func TestArrayListFilterNot(t *testing.T) {
	assert := assert.New(t)

	list := CreateArrayList[int]()
	FilterNot(list, assert)
}

func TestArrayListForEach(t *testing.T) {
	assert := assert.New(t)

	list := CreateArrayList[int]()
	ForEach(list, assert)
}
