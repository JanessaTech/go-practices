package generics

import (
	"fmt"
)

type Node[T any] struct {
	value T
	next  *Node[T]
}

type List[T any] struct {
	head, tail *Node[T]
}

func MapKeys[K comparable, V any](m map[K]V) []K {
	lst := make([]K, 0, len(m))
	for k := range m {
		lst = append(lst, k)
	}
	return lst
}

func (lst *List[T]) Push(v T) {
	if lst.tail == nil {
		lst.tail = &Node[T]{value: v}
		lst.head = lst.tail

	} else {
		lst.tail.next = &Node[T]{value: v}
		lst.tail = lst.tail.next
	}
}

func (lst *List[T]) GetAll() []T {
	var elems []T
	for e := lst.head; e != nil; e = e.next {
		elems = append(elems, e.value)
	}
	return elems
}

func Test() {
	var m = map[int]string{1: "aaa", 2: "bbb", 3: "ccc"}
	fmt.Println(MapKeys(m))

	list := List[int]{}
	list.Push(4)
	list.Push(4)
	list.Push(4)
	fmt.Println(list.GetAll())
}
