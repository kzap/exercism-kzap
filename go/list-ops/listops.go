// Package listops solves the List Ops side exercise on Exercism's Go track.
package listops

// A list of integers.
type IntList []int

// A binary function for folding integers.
type binFunc func(x, y int) int

// A predicate function for filtering a list of integers.
type predFunc func(n int) bool

// A unary function for mapping a list of integers.
type unaryFunc func(x int) int

// Foldl takes in a starting integer and binary function of:
//
// func(x, y int) int
//
// It runs the given function over a slice of integers from a left to right direction.
func (list IntList) Foldl(fn binFunc, start int) int {
	if list.Length() == 0 {
		return start
	}

	acc := start

	for _, v := range list {
		acc = fn(acc, v)
	}

	return acc
	//return fn(IntList(list[1:]).Foldr(fn, start), list[0])
}

// Foldr takes in a starting integer and binary function of:
//
// func(x, y int) int
//
// It runs the given function over a slice of integers from a right to left direction.
func (list IntList) Foldr(fn binFunc, start int) int {
	if list.Length() == 0 {
		return start
	}

	return fn(list[0], IntList(list[1:]).Foldr(fn, start))
}

// Filter takes in a predicate function of:
//
// func(n int) bool
//
// It runs the given function over a slice of integers and returns all those that return true.
func (list IntList) Filter(fn predFunc) IntList {
	newList := make(IntList, 0, list.Length())

	for _, v := range list {
		if fn(v) == true {
			newList = newList.Append(IntList{v})
		}
	}

	return newList
}

// Length takes in a slice of integers and returns number of values in the slice.
func (list IntList) Length() (length int) {
	for range list {
		length++
	}

	return
}

// Map takes in a unary function of:
//
// unaryFunc func(x int) int
//
// It runs the given function and applies it to each value in the slice.
func (list IntList) Map(fn unaryFunc) IntList {
	newList := make(IntList, list.Length())

	for i, v := range list {
		newList[i] = fn(v)
	}

	return newList
}

// Reverse returns the slice in reverse order.
func (list IntList) Reverse() IntList {
	listLength := list.Length()

	newList := make(IntList, listLength)

	for i, j := 0, listLength-1; i < j; i, j = i+1, j-1 {
		newList[i], newList[j] = list[j], list[i]
	}

	return newList
}

// Append takes in a slice of integers and appends it to the current slice of integers.
func (list IntList) Append(appendList IntList) IntList {
	listLength := list.Length()
	appendLength := appendList.Length()

	newList := make(IntList, listLength+appendLength)

	for i, v := range list {
		newList[i] = v
	}

	for i, v := range appendList {
		newList[listLength+i] = v
	}

	return newList
}

// Concat takes in multiple slices integers and appends it to the current slice of integers.
func (list IntList) Concat(concatList []IntList) IntList {
	newList := IntList{}
	newList = newList.Append(list)

	for _, v := range concatList {
		newList = newList.Append(v)
	}

	return newList
}
