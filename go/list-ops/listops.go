// Package listops solves the List Ops side exercise on Exercism's Go track.
package listops

import (
//"log"
)

type IntList []int
type binFunc func(x, y int) int
type predFunc func(n int) bool
type unaryFunc func(x int) int

// Foldl
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

// Foldr
func (list IntList) Foldr(fn binFunc, start int) int {
	if list.Length() == 0 {
		return start
	}

	return fn(list[0], IntList(list[1:]).Foldr(fn, start))
}

// Filter
func (list IntList) Filter(fn predFunc) IntList {
	newList := IntList{}

	if list.Length() == 0 {
		return newList
	}

	for _, v := range list {
		if fn(v) == true {
			newList = newList.Append(IntList{v})
		}
	}

	return newList
}

// Length
func (list IntList) Length() (length int) {
	for range list {
		length++
	}

	return
}

// Map
func (list IntList) Map(fn unaryFunc) IntList {
	newList := IntList{}

	for _, v := range list {
		newList = newList.Append(IntList{fn(v)})
	}

	return newList
}

// Reverse
func (list IntList) Reverse() IntList {
	listLength := list.Length()

	newList := make(IntList, listLength)

	for i, j := 0, listLength-1; i < j; i, j = i+1, j-1 {
		newList[i], newList[j] = list[j], list[i]
	}

	return newList
}

// Append
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

// Concat
func (list IntList) Concat(concatList []IntList) IntList {
	newList := list

	for _, v := range concatList {
		newList = newList.Append(v)
	}

	return newList
}
