// Copyright 2012 Dobrosław Żybort
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package listdict

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

var dictIsEqualTests = []struct {
	in         Dict
	secondDict Dict
	out        bool
}{
	{Dict{"one": 1, "two": 2}, Dict{"one": 1, "two": 2}, true},
	{Dict{"one": 1, "two": 2}, Dict{"one": 1, "two": 2, "three": 3}, false},
	{Dict{"one": 1, "two": 2}, Dict{"two": 2, "one": 1}, true},
	{Dict{}, Dict{}, true},
}

func TestDictIsEqual(t *testing.T) {
	for index, diet := range dictIsEqualTests {
		val := diet.in.IsEqual(diet.secondDict)
		if val != diet.out {
			t.Errorf(
				"%d. %v.IsEqual(%v) => %v, want %v",
				index, diet.in, diet.secondDict, val, diet.out)
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

//=============================================================================

var dictKeysTests = []struct {
	in  Dict
	out List
}{
	{Dict{"one": 1, "two": 2, "three": 3}, List{"one", "two", "three"}},
	{Dict{"one": 1, "two": 2, "three": 3}, List{"two", "one", "three"}},
	{Dict{"one": 1}, List{"one"}},
	{Dict{}, List{}},
}

func TestDictKeys(t *testing.T) {
	for index, dkt := range dictKeysTests {

		keyList := dkt.in.Keys()

		founded := 0
		for _, listItem := range dkt.out {
			for _, dictItem := range keyList {
				if reflect.DeepEqual(dictItem, listItem) {
					founded++
				}
			}
		}

		switch {
		case len(dkt.out) != len(keyList):
			t.Errorf("%d. len(%v.Keys()) => %d, should be %d",
				index, dkt.in, len(keyList), len(dkt.out))
		case founded != len(keyList):
			t.Errorf("%d. compare %v.Keys() and %v => "+
				"found %v same elements, shoud found %v",
				index, dkt.in, dkt.out, founded, len(keyList))
		}
	}
}

//=============================================================================

var dictPopTests = []struct {
	in         Dict
	key        string
	defaultVal interface{}
	out        interface{}
	outError   error
	outDict    Dict
}{
	{Dict{"one": 1, "two": 2, "three": 3},
		"one", 0,
		1, nil, Dict{"two": 2, "three": 3}},
	{Dict{"one": 1, "two": 2, "three": 3},
		"four", 0,
		0, nil, Dict{"one": 1, "two": 2, "three": 3}},
	{Dict{"one": 1},
		"one", 0,
		1, nil, Dict{}},
	{Dict{},
		"one", 0,
		0, ErrRemoveFromEmptyDict, Dict{}},
}

func TestDictPop(t *testing.T) {
	for index, dpt := range dictPopTests {
		dict := NewDict()
		for key, val := range dpt.in {
			dict[key] = val
		}
		val, err := dict.Pop(dpt.key, dpt.defaultVal)

		switch {
		case val != dpt.out || err != dpt.outError:
			t.Errorf("%d. %v.Pop(%v, %v) => %v, %v, want %v, %v",
				index, dpt.in, dpt.key, dpt.defaultVal,
				val, err, dpt.out, dpt.outError)
		case !reflect.DeepEqual(dict, dpt.outDict):
			t.Errorf("%d. %v.Pop(%v, %v) => out dict = %v, want %v",
				index, dpt.in, dpt.key, dpt.defaultVal, dict, dpt.outDict)
		}
	}
}

//=============================================================================

var dictPopItemTests = []struct {
	in       Dict
	out      List
	outError error
	outDict  Dict
}{
	{Dict{"one": 1}, List{"one", 1}, nil, Dict{}},
	{Dict{}, List{}, ErrRemoveFromEmptyDict, Dict{}},
}

func TestDictPopItem(t *testing.T) {
	for index, dpit := range dictPopItemTests {
		dict := NewDict()
		for key, val := range dpit.in {
			dict[key] = val
		}
		list, err := dict.PopItem()

		switch {
		case !reflect.DeepEqual(list, dpit.out) || err != dpit.outError:
			t.Errorf("%d. %v.PopItem() => %v, %v, want %v, %v",
				index, dpit.in, list, dpit.out, dpit.outError)
		case !reflect.DeepEqual(dict, dpit.outDict):
			t.Errorf("%d. %v.PopItem() => out dict = %v, want %v",
				index, dpit.in, dict, dpit.outDict)
		}
	}
}

//=============================================================================

var dictSetDefaultTests = []struct {
	in         Dict
	key        string
	defaultVal interface{}
	out        interface{}
	outDict    Dict
}{
	{Dict{"one": 1, "two": 2}, "one", 0, 1, Dict{"one": 1, "two": 2}},
	{Dict{"one": 1}, "two", 0, 0, Dict{"one": 1, "two": 0}},
	{Dict{}, "one", 0, 0, Dict{"one": 0}},
}

func TestDictSetDefault(t *testing.T) {
	for index, dsdt := range dictSetDefaultTests {
		dict := NewDict()
		for key, val := range dsdt.in {
			dict[key] = val
		}

		myValue := dict.SetDefault(dsdt.key, dsdt.defaultVal)

		switch {
		case myValue != dsdt.out:
			t.Errorf("%d. %v.SetDefault(%v, %v) => %v, want %v",
				index, dsdt.in, dsdt.key, dsdt.defaultVal, myValue, dsdt.out)
		case !reflect.DeepEqual(dict, dsdt.outDict) || !dict.HasKey(dsdt.key):
			t.Errorf("%d. %v.SetDefault(%v, %v) => out dict = %v, want %v",
				index, dsdt.in, dsdt.key, dsdt.defaultVal, dict, dsdt.outDict)
		}
	}
}

//=============================================================================

var dictUpdateTests = []struct {
	in        Dict
	otherDict Dict
	out       Dict
}{
	{Dict{"one": 1}, Dict{"one": 0, "two": 2}, Dict{"one": 0, "two": 2}},
	{Dict{"one": 1}, Dict{"two": 2}, Dict{"one": 1, "two": 2}},
	{Dict{"one": 1}, Dict{"one": 0}, Dict{"one": 0}},
	{Dict{"one": 1}, Dict{}, Dict{"one": 1}},
	{Dict{}, Dict{"one": 1}, Dict{"one": 1}},
	{Dict{}, Dict{}, Dict{}},
}

func TestDictUpdate(t *testing.T) {
	for index, dut := range dictUpdateTests {
		dict := NewDict()
		for key, val := range dut.in {
			dict[key] = val
		}
		dict.Update(dut.otherDict)

		switch {
		case !reflect.DeepEqual(dict, dut.out):
			t.Errorf("%d. %v.Update(%v) => %v, want %v",
				index, dut.in, dut.otherDict, dict, dut.out)
		}
	}
}

//=============================================================================

var dictValuesTests = []struct {
	in  Dict
	out List
}{
	{Dict{"one": 1, "two": 2, "three": 3}, List{1, 2, 3}},
	{Dict{"one": 1, "two": 2, "three": 3}, List{3, 1, 2}},
	{Dict{"one": 1}, List{1}},
	{Dict{}, List{}},
}

func TestDictValues(t *testing.T) {
	for index, dvt := range dictValuesTests {

		valueList := dvt.in.Values()

		founded := 0
		for _, listItem := range dvt.out {
			for _, dictItem := range valueList {
				if reflect.DeepEqual(dictItem, listItem) {
					founded++
				}
			}
		}

		switch {
		case len(dvt.out) != len(valueList):
			t.Errorf("%d. len(%v.Values()) => %d, should be %d",
				index, dvt.in, len(valueList), len(dvt.out))
		case founded != len(valueList):
			t.Errorf("%d. compare %v.Values() and %v => "+
				"found %v same elements, shoud found %v",
				index, dvt.in, dvt.out, founded, len(valueList))
		}
	}
}
