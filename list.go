// Copyright 2012 Dobrosław Żybort
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package simpletype

import (
	"errors"
	"fmt"
)

// Simple list
type List []interface{}

// Return new List with specified length
func NewList(length int) List {
	return make(List, length)
}

//=============================================================================

// Append adds an element to the end of the list.
func (list *List) Append(obj interface{}) { *list = append(*list, obj) }

// Count returns the number of times value appears in the list.
func (list List) Count(value interface{}) int {
	counter := 0
	for _, listValue := range list {
		if listValue == value {
			counter++
		}
	}
	return counter
}

// Extend one list with the contents of the other list.
func (list *List) Extend(otherList List) {
	for _, value := range otherList {
		*list = append(*list, value)
	}
}

// Index returns the index of the first item in the list whose value is val.
// It is -1 if there is no such item.
func (list List) Index(val interface{}) (int, error) {
	for index, listValue := range list {
		if listValue == val {
			return index, nil
		}
	}
	errorString := fmt.Sprintf("%v is not in list", val)
	return -1, errors.New(errorString)
}

// Insert an element at a given position. If the position is past the end 
// of the list, append to the end.
func (list *List) Insert(index int, val interface{}) {
	if index < 0 {
		panic("Index out of bounds")
	}

	if len(*list) > index {
		list2 := make(List, len(*list))
		copy(list2, *list)
		list2 = append(list2, 0)
		copy(list2[index+1:], list2[index:])
		list2[index] = val
		*list = list2
	} else {
		*list = append(*list, val)
	}
}

// Remove and returns the last element in the list.
func (list *List) Pop() interface{} {
	if len(*list) <= 0 {
		panic("Pop from empty list")
	}

	list2 := make(List, len(*list))
	copy(list2, *list)
	listLen := len(list2)
	val, list2 := list2[listLen-1], list2[:listLen-1]

	*list = list2
	return val
}

// Remove and returns the element at the given position in the list.
func (list *List) PopItem(index int) interface{} {
	if len(*list) <= 0 {
		panic("PopItem from empty list")
	}
	if index < 0 {
		panic("Index out of bounds")
	}

	list2 := make(List, len(*list))
	copy(list2, *list)
	listLen := len(list2)
	if index >= listLen {
		panic("Index out of range")
	}
	val := list2[index]

	copy(list2[index:], list2[index+1:])
	list2[listLen-1] = nil
	list2 = list2[:listLen-1]

	*list = list2
	return val
}

//// Remove the first element from the list whose value matches the given value. 
//// Error if no match is found.
//func (list *List) Remove(val interface{}) error {
//}

//// Reverse the elements of the list in place.
//func (list *List) Reverse() {
//}

//// Sort the list in place ordering elements from smallest to largest.
//func (list *List) Sort() {
//}
