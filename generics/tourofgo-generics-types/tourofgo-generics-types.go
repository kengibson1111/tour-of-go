package main

import "fmt"

// List represents a doubly-linked list that holds
// values of any type. Original requirement was singly-linked.
type List[T any] struct {
	head *ListElement[T]
	tail *ListElement[T]
	len  int
}

func (list *List[T]) String() string {
	if list == nil {
		return "nil List"
	}

	if list.head == nil {
		return "empty List"
	}

	var listStr string = fmt.Sprintf("{len=%v, head=%p, tail=%p, ", list.len, list.head, list.tail)
	for current := list.head; current != nil; current = current.next {
		listStr = listStr + current.String()
		if current.next != nil {
			listStr = listStr + ", "
		}
	}

	listStr = listStr + "}"
	return listStr
}

// separate List elements from the actual List
type ListElement[T any] struct {
	prev *ListElement[T]
	next *ListElement[T]
	val  T
}

func (listElement *ListElement[T]) String() string {
	if listElement == nil {
		return "nil"
	}

	return fmt.Sprintf("(val=%v, prev=%p, next=%p)", listElement.val, listElement.prev, listElement.next)
}

// private function for List use only.
func createElement[T any](t T) *ListElement[T] {
	var newElement = new(ListElement[T])
	newElement.val = t

	return newElement
}

func (list *List[T]) Size() int {
	// for the sake of the exercise, just return for now
	if list == nil {
		return 0
	}

	return list.len
}

func (list *List[T]) Append(t T) {
	// for the sake of the exercise, just return for now
	if list == nil {
		return
	}

	var newElement = createElement(t)

	if list.head == nil {
		list.head = newElement
		list.tail = newElement
	} else {
		newElement.prev = list.tail
		list.tail.next = newElement
		list.tail = newElement
	}

	list.len++
}

// private method.
func (list *List[T]) checkForEmptyorSingleValue() bool {
	// for the sake of the exercise, just return for now
	if list == nil {
		return true
	}

	var zeroT T

	// if nothing in the list, return for now. maybe later, an error
	if list.head == nil {
		return true
	} else if list.head == list.tail {
		//zero out existing list head which will also zero out the tail
		list.head.prev = nil
		list.head.next = nil
		list.head.val = zeroT

		// set both to nil. len should be zero
		list.head = nil
		list.tail = nil

		return true
	}

	return false
}

func (list *List[T]) RemoveHead() {
	// for the sake of the exercise, just return for now
	if list == nil {
		return
	}

	var zeroT T

	if list.checkForEmptyorSingleValue() {
		return
	}

	// establish a new list head and zero out the existing list head.
	var newHead = list.head.next
	newHead.prev = nil

	list.head.prev = nil
	list.head.next = nil
	list.head.val = zeroT

	list.head = newHead
	list.len--
}

func (list *List[T]) RemoveTail() {
	// for the sake of the exercise, just return for now
	if list == nil {
		return
	}

	var zeroT T

	if list.checkForEmptyorSingleValue() {
		return
	}

	// establish a new list tail and zero out the existing list tail.
	var newTail = list.tail.prev
	newTail.next = nil

	list.tail.prev = nil
	list.tail.next = nil
	list.tail.val = zeroT

	list.tail = newTail
	list.len--
}

func main() {
	fmt.Println("new...")
	var intList = new(List[int])
	fmt.Println("int list =", intList)

	fmt.Println("append...")
	intList.Append(0)
	fmt.Println("int list =", intList)

	fmt.Println("append...")
	intList.Append(1)
	fmt.Println("int list =", intList)

	fmt.Println("append...")
	intList.Append(2)
	fmt.Println("int list =", intList)

	fmt.Println("append...")
	intList.Append(3)
	fmt.Println("int list =", intList)

	fmt.Println("remove head...")
	intList.RemoveHead()
	fmt.Println("int list =", intList)

	fmt.Println("remove tail...")
	intList.RemoveTail()
	fmt.Println("int list =", intList)

	fmt.Println("remove tail...")
	intList.RemoveTail()
	fmt.Println("int list =", intList)

	fmt.Println("remove head...")
	intList.RemoveHead()
	fmt.Println("int list =", intList)
}
