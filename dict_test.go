// Copyright 2012 Dobrosław Żybort
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package simpletype

import (
	"reflect"
	"testing"
)

//=============================================================================

var dictFromKeysTests = []struct {
	in         List
	defaultVal interface{}
	out        Dict
}{
	{List{"one", "two", "three"}, nil,
		Dict{"one": nil, "two": nil, "three": nil}},
	{List{1, 2}, 1, Dict{"1": 1, "2": 1}},
	{List{}, nil, Dict{}},
}

func TestDictFromKeys(t *testing.T) {
	for index, dfkt := range dictFromKeysTests {
		dict := DictFromKeys(dfkt.in, dfkt.defaultVal)

		if !reflect.DeepEqual(dict, dfkt.out) {
			t.Errorf(
				"%d. DictFromKeys(%v, %v) => out dict = %v, want %v",
				index, dfkt.in, dfkt.defaultVal, dict, dfkt.out)
		}
	}
}

//=============================================================================

var dictClearTests = []struct {
	in  Dict
	out Dict
}{
	{Dict{"one": 1, "two": 2, "three": 3}, Dict{}},
	{Dict{"one": List{1, 2}}, Dict{}},
	{Dict{}, Dict{}},
}

func TestDictClear(t *testing.T) {
	for index, dct := range dictClearTests {
		dict := NewDict()
		for key, val := range dct.in {
			dict[key] = val
		}

		dict.Clear()

		if len(dict) != 0 {
			t.Fatalf("%d. len(%v.Clear()) => %v, should be 0",
				index, dct.in, len(dict))
		}
		if !reflect.DeepEqual(dict, dct.out) {
			t.Errorf(
				"%d. %v.Clear() => out dict = %v, want %v",
				index, dct.in, dict, dct.out)
		}
	}
}

//=============================================================================

var dictGetTests = []struct {
	in         Dict
	key        string
	defaultVal interface{}
	out        interface{}
}{
	{Dict{"one": 1, "two": 2, "three": 3}, "one", 0, 1},
	{Dict{"one": 1}, "two", 0, 0},
	{Dict{}, "one", 0, 0},
}

func TestDictGet(t *testing.T) {
	for index, dgt := range dictGetTests {
		dict := NewDict()
		for key, val := range dgt.in {
			dict[key] = val
		}

		myValue := dict.Get(dgt.key, dgt.defaultVal)

		switch {
		case myValue != dgt.out:
			t.Errorf("%d. %v.Get(%v, %v) => %v, want 0",
				index, dgt.in, dgt.key, dgt.defaultVal, myValue, dgt.out)
		case dict.HasKey(dgt.key) && !dgt.in.HasKey(dgt.key):
			t.Errorf("%d. %v.Get(%v, %v) => out dict should not have '%v' key",
				index, dgt.in, dgt.key, dgt.defaultVal, dgt.key)
		}
	}
}

//=============================================================================

var dictHasKeyTests = []struct {
	in  Dict
	key string
	out bool
}{
	{Dict{"one": 1, "two": 2, "three": 3}, "one", true},
	{Dict{"one": 1}, "two", false},
	{Dict{}, "one", false},
}

func TestDictHasKey(t *testing.T) {
	for index, dhkt := range dictHasKeyTests {
		var keyTest bool
		if _, ok := dhkt.in[dhkt.key]; ok {
			keyTest = true
		} else {
			keyTest = false
		}

		if keyTest != dhkt.out {
			t.Errorf("%d. %v.HasKey(%v) => %v, want %v",
				index, dhkt.in, dhkt.key, keyTest, dhkt.out)
		}
	}
}

//=============================================================================

var dictItemsTests = []struct {
	in  Dict
	out []List
}{
	{Dict{"one": 1, "two": 2, "three": 3},
		[]List{List{"one", 1}, List{"two", 2}, List{"three", 3}}},
	{Dict{"one": 1}, []List{List{"one", 1}}},
	{Dict{}, []List{}},
}

func TestDictItems(t *testing.T) {
	for index, dit := range dictItemsTests {

		itemList := dit.in.Items()

		founded := 0
		for _, listItem := range dit.out {
			for _, dictItem := range itemList {
				if reflect.DeepEqual(dictItem, listItem) {
					founded++
				}
			}
		}

		switch {
		case len(dit.out) != len(itemList):
			t.Errorf("%d. len(%v.Items()) => %d, should be %d",
				index, dit.in, len(itemList), len(dit.out))
		case founded != len(itemList):
			t.Errorf("%d. compare %v.Items() and %v => "+
				"found %v same list elements, shoud found %v",
				index, dit.in, dit.out, founded, len(itemList))
		}
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

func TestDictPop(t *testing.T) {
	dict := Dict{"one": 1, "two": 2, "three": 3}

	goodDict := Dict{"two": 2, "three": 3}

	val, err := dict.Pop("one", 0)
	if val != 1 || !reflect.DeepEqual(dict, goodDict) || err != nil {
		t.Errorf("Should be %v, got %v", goodDict, dict)
	}
	val, err = dict.Pop("four", 0)
	if val != 0 || !reflect.DeepEqual(dict, goodDict) || err != nil {
		t.Errorf("Should be %v, got %v", goodDict, dict)
	}
}

func TestDictPopItem(t *testing.T) {
	dict := Dict{"one": 1}

	goodDict := Dict{}
	goodVal := List{"one", 1}

	val, err := dict.PopItem()
	if !reflect.DeepEqual(val, goodVal) ||
		!reflect.DeepEqual(dict, goodDict) ||
		err != nil {
		t.Errorf(
			"Should be %v and %v, got %v and %v",
			goodVal, goodDict, val, dict,
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

func TestDictUpdate(t *testing.T) {
	dict := Dict{"one": 1, "two": 2, "three": 3}
	dict.Update(Dict{"one": 0, "four": 4})

	goodDict := Dict{"one": 0, "two": 2, "three": 3, "four": 4}

	if !reflect.DeepEqual(dict, goodDict) {
		t.Errorf("Error when updating, should be %v, got %v", goodDict, dict)
	}

	dict.Update(Dict{"one": 1, "five": 5})
	goodDict2 := Dict{"one": 1, "two": 2, "three": 3, "four": 4, "five": 5}

	if !reflect.DeepEqual(dict, goodDict2) {
		t.Errorf("Error when updating, should be %v, got %v", goodDict2, dict)
	}
}

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
