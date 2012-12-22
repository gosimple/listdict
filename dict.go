// Copyright 2012 Dobrosław Żybort
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package simpletype

import "fmt"

// Simple dict
type Dict map[string]interface{}

// Return new Dict
func NewDict() Dict {
	return make(Dict)
}

//=============================================================================

//Create a new dictionary with keys from list and values set to val.
func DictFromKeys(list List, val interface{}) Dict {
	newDict := NewDict()
	for _, value := range list {
		newDict[fmt.Sprintf("%v", value)] = val
	}
	return newDict
}

//=============================================================================

// Removes all elements from the dictionary
func (dict Dict) Clear() {
	for key, _ := range dict {
		delete(dict, key)
	}
}

// Return value for the given key or val if key is not in the dictionary
func (dict Dict) Get(key string, val interface{}) interface{} {
	if dict.HasKey(key) {
		return dict[key]
	}
	return val
}

// Returns true if key is in the dictionary, false otherwise
func (dict Dict) HasKey(key string) bool {
	if _, ok := dict[key]; ok {
		return true
	}
	return false
}

// Returns a unordered list of the dictionary's [key, value] list pairs
func (dict Dict) Items() []List {
	list := []List{}
	for key, value := range dict {
		list = append(list, List{key, value})
	}
	return list
}

// Returns a list of the dictionary's keys, unordered
func (dict Dict) Keys() List {
	list := NewList(len(dict))
	i := 0
	for key, _ := range dict {
		list[i] = key
		i++
	}
	return list
}

// If the given key is in the dictionary, remove it and return its value,
// else return val.
//func (dict Dict) Pop(key string, val interface{}) interface{} {
//}

// Return and remove an random key-value pair as List from the dictionary.
//func (dict Dict) PopItem() List {
//}

// Similar to Get(), but will set dict[key]=val if key is not already in dict
func (dict Dict) SetDefault(key string, val interface{}) interface{} {
	if dict.HasKey(key) {
		return dict[key]
	}
	dict[key] = val
	return val
}

// Update the dictionary with the key-value pairs in the dict2 dictionary
// replacing current values and adding new if found.
func (dict Dict) Update(dict2 Dict) {
	for key, value := range dict2 {
		dict[key] = value
	}
}

//Returns a list of the dictionary's values, unordered
func (dict Dict) Values() List {
	list := NewList(len(dict))
	i := 0
	for _, value := range dict {
		list[i] = value
		i++
	}
	return list
}
