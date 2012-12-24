// Copyright 2012 Dobrosław Żybort
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package simpletype

import (
	"testing"
)

func TestDictFromKeys(t *testing.T) {
	list := List{"one", "two", "three"}
	dict := DictFromKeys(list, nil)

	goodDict := Dict{"one": nil, "two": nil, "three": nil}
	if len(dict) != len(goodDict) {
		t.Errorf(
			"New dict length should be %v, got %v",
			len(goodDict),
			len(dict),
		)
	}
	founded := 0
	for goodKey, goodValue := range goodDict {
		for dictKey, dictValue := range dict {
			if goodKey == dictKey && goodValue == dictValue {
				founded++
			}
		}
	}
	if founded != len(dict) {
		t.Errorf(
			"Should found %v same dict elements, got %v",
			len(goodDict),
			len(dict),
		)
	}
}

func TestDictClear(t *testing.T) {
	dict := Dict{"one": 1, "two": 2, "three": 3}
	dict.Clear()

	if len(dict) != 0 {
		t.Errorf("Dict length should be 0, got %v", len(dict))
	}
}

func TestDictGet(t *testing.T) {
	dict := Dict{"one": 1, "two": 2, "three": 3}

	goodValue := dict.Get("one", 0)
	if goodValue != 1 {
		t.Errorf("Dict value should be 1, got %v", goodValue)
	}

	badValue := dict.Get("four", 0)
	switch {
	case badValue != 0:
		t.Errorf("Return value should be 0, got %v", badValue)
	case dict.HasKey("four"):
		t.Errorf("Dict should not have 'four' key")
	}
}

func TestDictHasKey(t *testing.T) {
	dict := Dict{"one": 1, "two": 2, "three": 3}

	goodValue := dict.HasKey("one")
	if goodValue != true {
		t.Errorf("Dict key check should be 'true', got '%v'", goodValue)
	}

	badValue := dict.HasKey("four")
	if badValue != false {
		t.Errorf("Dict key check should be 'true', got '%v'", badValue)
	}
}

func TestDictItems(t *testing.T) {
	dict := Dict{"one": 1, "two": 2, "three": 3}
	goodItems := []List{List{"one", 1}, List{"two", 2}, List{"three", 3}}
	itemList := dict.Items()

	if len(goodItems) != len(itemList) {
		t.Errorf(
			"Item list length should be %v, got %v",
			len(goodItems),
			len(itemList),
		)
	}
	founded := 0
	for _, listItem := range goodItems {
		for _, dictItem := range itemList {
			if dictItem[0] == listItem[0] && dictItem[1] == listItem[1] {
				founded++
			}
		}
	}
	if founded != len(itemList) {
		t.Errorf(
			"Should found %v same list elements, got %v",
			len(goodItems),
			len(itemList),
		)
	}
}

func TestDictKeys(t *testing.T) {
	dict := Dict{"one": 1, "two": 2, "three": 3}
	goodKeys := List{"one", "two", "three"}
	keyList := dict.Keys()

	if len(goodKeys) != len(keyList) {
		t.Errorf(
			"Value list length should be %v, got %v",
			len(goodKeys),
			len(keyList),
		)
	}

	founded := 0
	for _, listKey := range goodKeys {
		for _, dictKey := range keyList {
			if dictKey == listKey {
				founded++
			}
		}
	}
	if founded != len(keyList) {
		t.Errorf(
			"Should found %v same list elements, got %v",
			len(goodKeys),
			len(keyList),
		)
	}
}

func TestDictSetDefault(t *testing.T) {
	dict := Dict{"one": 1, "two": 2, "three": 3}

	goodValue := dict.SetDefault("one", 0)
	if goodValue != 1 {
		t.Errorf("Dict value should be 1, got %v", goodValue)
	}

	goodValue2 := dict.SetDefault("four", 4)

	switch {
	case goodValue2 != 4:
		t.Errorf("Dict value should be 4, got %v", goodValue2)
	case !dict.HasKey("four"):
		t.Errorf("Dict should have 'four' key")
	}
}

/*
func TestDictUpdate(t *testing.T) {
	dict := Dict{"one": 1, "two": 2, "three": 3}
	dict.Update(Dict{"one": 0, "four": 4})
	fmt.Println(dict)
	//dict = {'one': 0, 'two': 2, 'three': 3, 'four': 4}
	dict.Update(Dict{"one": 1, "five": 5})
	fmt.Println(dict)
	//dict = {'one': 1, 'two': 2, 'three': 3, 'four': 4, 'five': 5}
} */

func TestDictValues(t *testing.T) {
	dict := Dict{"one": 1, "two": 2, "three": 3}
	goodValues := List{1, 2, 3}
	valueList := dict.Values()

	if len(goodValues) != len(valueList) {
		t.Errorf(
			"Value list length should be %v, got %v",
			len(goodValues),
			len(valueList),
		)
	}

	founded := 0
	for _, listValue := range goodValues {
		for _, dictValue := range valueList {
			if dictValue == listValue {
				founded++
			}
		}
	}
	if founded != len(valueList) {
		t.Errorf(
			"Should found %v same list elements, got %v",
			len(goodValues),
			len(valueList),
		)
	}

}
