// Copyright 2012 Dobrosław Żybort
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package simpletype

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
func (list List) Index(val interface{}) int {
	for index, listValue := range list {
		if listValue == val {
			return index
		}
	}
	return -1
}

// Insert an element at a given position. If the position is past the end 
// of the list, append to the end.
//list.insert(index int, val interface{})

// Remove and returns the last element in the list.
//list.pop()

// Remove and return the element at the given position in the list.
//list.pop([i])

// Remove the first element from the list whose value matches the given value. 
// Error if no match is found.
//list.remove(val) error

// Reverse the elements of the list in place.
//list.reverse()

// Sort the list in place ordering elements from smallest to largest
//list.sort()
