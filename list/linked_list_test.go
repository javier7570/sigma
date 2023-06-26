package list

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLinkedListPushFront(t *testing.T) {
	assert := assert.New(t)

	list := CreateLinkedList[int]()
	PushFront(list, assert)
}

func TestLinkedListPushBack(t *testing.T) {
	assert := assert.New(t)

	list := CreateLinkedList[int]()
	PushBack(list, assert)
}

func TestLinkedListPopFront(t *testing.T) {
	assert := assert.New(t)

	list := CreateLinkedList[int]()
	PopFront(list, assert)
}

func TestLinkedListPopBack(t *testing.T) {
	assert := assert.New(t)

	list := CreateLinkedList[int]()
	PopBack(list, assert)
}

func TestLinkedListIterate(t *testing.T) {
	assert := assert.New(t)

	list := CreateLinkedList[int]()
	Iterate(list, assert)
}

func TestLinkedListReverse(t *testing.T) {
	assert := assert.New(t)

	list := CreateLinkedList[int]()
	Reverse(list, assert)
}

func TestLinkedListExists(t *testing.T) {
	assert := assert.New(t)

	list := CreateLinkedList[int]()
	Exists(list, assert)

}

func TestLinkedListFilter(t *testing.T) {
	assert := assert.New(t)

	list := CreateLinkedList[int]()
	Filter(list, assert)
}

func TestLinkedListFilterNot(t *testing.T) {
	assert := assert.New(t)

	list := CreateLinkedList[int]()
	FilterNot(list, assert)
}

func TestLinkedListForAll(t *testing.T) {
	assert := assert.New(t)

	list := CreateLinkedList[int]()
	ForAll(list, assert)
}

func TestLinkedListForEach(t *testing.T) {
	assert := assert.New(t)

	list := CreateLinkedList[int]()
	ForEach(list, assert)
}
