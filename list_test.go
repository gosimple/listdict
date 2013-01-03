// Copyright 2012 Dobrosław Żybort
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package simpletype

import (
	//"errors"
	"reflect"
	"testing"
)

var listAppendTests = []struct {
	in  List
	val interface{}
	out List
}{
	{List{"one", "two"}, "three", List{"one", "two", "three"}},
	{List{"one", "two"}, 3, List{"one", "two", 3}},
	{List{}, 1, List{1}},
}

func TestListAppend(t *testing.T) {
	for index, lat := range listAppendTests {
		list := lat.in
		list.Append(lat.val)
		if !reflect.DeepEqual(list, lat.out) {
			t.Errorf(
				"%d. %v.Append(%v) => out list = %v, want %v",
				index, lat.in, lat.val, list, lat.out)
		}
	}
}

var listCountTests = []struct {
	in  List
	val interface{}
	out int
}{
	{List{"one", "two", "three", "two"}, "two", 2},
	{List{"one", "two", "three", "two"}, "zero", 0},
	{List{1, 2, 3, 2}, 3, 1},
	{List{}, 1, 0},
}

func TestListCount(t *testing.T) {
	for index, lct := range listCountTests {
		counted := lct.in.Count(lct.val)
		if counted != lct.out {
			t.Errorf(
				"%d. %v.Count(%v) => %d, want %d",
				index, lct.in, lct.val, counted, lct.out)
		}
	}
}

var listDeleteTests = []struct {
	in       List
	val      int
	outList  List
	outError error
}{
	{List{"one", "two", "three"}, 1, List{"one", "three"}, nil},
	{List{"one"}, 0, List{}, nil},
	{List{"one"}, 1, List{"one"}, ErrRange},
	{List{"one"}, -1, List{"one"}, ErrBounds},
	{List{}, 0, List{}, ErrRemoveFromEmptyList},
}

func TestListDelete(t *testing.T) {
	for index, ldt := range listDeleteTests {
		list := ldt.in
		err := list.Delete(ldt.val)
		if err != ldt.outError {
			t.Errorf(
				"%d. %v.Delete(%d) => %v, want %v",
				index, ldt.in, ldt.val, err, ldt.outError)
		}
		if !reflect.DeepEqual(list, ldt.outList) {
			t.Errorf(
				"%d. %v.Delete(%d) => out list = %v, want %v",
				index, ldt.in, ldt.val, list, ldt.outList)
		}
	}
}

var listExtendTests = []struct {
	in        List
	otherList List
	out       List
}{
	{List{"one", "two"}, List{"three"}, List{"one", "two", "three"}},
	{List{"one"}, List{}, List{"one"}},
	{List{}, List{"one"}, List{"one"}},
	{List{1}, List{"one"}, List{1, "one"}},
	{List{}, List{}, List{}},
}

func TestListExtend(t *testing.T) {
	for index, let := range listExtendTests {
		list := let.in
		list.Extend(let.otherList)
		if !reflect.DeepEqual(list, let.out) {
			t.Errorf(
				"%d. %v.Extend(%v) => out list = %v, want %v",
				index, let.in, let.otherList, list, let.out)
		}
	}
}

var listIndexTests = []struct {
	in       List
	val      interface{}
	out      int
	outError error
}{
	{List{"one", "two", "three", "two"}, "two", 1, nil},
	//{List{"one", "two"}, "zero", -1, errors.New("zero is not in list")},
	{List{1, 2, 3, 2}, 3, 2, nil},
	//{List{}, 1, -1, errors.New("1 is not in list")},
}

func TestListIndex(t *testing.T) {
	for index, lit := range listIndexTests {
		foundIndex, err := lit.in.Index(lit.val)
		if foundIndex != lit.out || err != lit.outError {
			t.Errorf(
				"%d. %v.Index(%v) => %v, %v, want %v, %v",
				index, lit.in, lit.val, foundIndex, err, lit.out, lit.outError)
		}
	}
}

var listInsertTests = []struct {
	in       List
	index    int
	val      interface{}
	outList  List
	outError error
}{
	{List{"one", "two", "three"},
		1, "inserted",
		List{"one", "inserted", "two", "three"}, nil},
	{List{"one", "two", "three"},
		10, "four",
		List{"one", "two", "three", "four"}, nil},
}

func TestListInsert(t *testing.T) {
	for index, lit := range listInsertTests {
		list := lit.in
		err := list.Insert(lit.index, lit.val)
		if err != lit.outError {
			t.Errorf(
				"%d. %v.Insert(%v, %v) => %v, want %v",
				index, lit.in, lit.index, lit.val, err, lit.outError)
		}
		if !reflect.DeepEqual(list, lit.outList) {
			t.Errorf(
				"%d. %v.Insert(%v, %v) => out list = %v, want %v",
				index, lit.in, lit.index, lit.val, list, lit.outList)
		}
	}
}

var listPopTests = []struct {
	in       List
	out      interface{}
	outError error
	outList  List
}{
	{List{"one", "two", "three", "two"},
		"two", nil, List{"one", "two", "three"}},
	{List{1, 2, 3, 2}, 2, nil, List{1, 2, 3}},
	{List{}, nil, ErrRemoveFromEmptyList, List{}},
}

func TestListPop(t *testing.T) {
	for index, lpt := range listPopTests {
		list := lpt.in
		item, err := list.Pop()
		if item != lpt.out || err != lpt.outError {
			t.Errorf(
				"%d. %v.Pop() => %v, %v, want %v, %v",
				index, lpt.in, item, err, lpt.out, lpt.outError)
		}
		if !reflect.DeepEqual(list, lpt.outList) {
			t.Errorf(
				"%d. %v.Pop() => out list = %v, want %v",
				index, lpt.in, list, lpt.outList)
		}
	}
}

var listPopItemTests = []struct {
	in       List
	index    int
	out      interface{}
	outError error
	outList  List
}{
	{List{"one", "two", "three"}, 1, "two", nil, List{"one", "three"}},
	{List{1, 2, 3, 2}, 2, 3, nil, List{1, 2, 2}},
	{List{}, 0, nil, ErrRemoveFromEmptyList, List{}},
}

func TestListPopItem(t *testing.T) {
	for index, lpt := range listPopItemTests {
		list := lpt.in
		item, err := list.PopItem(lpt.index)
		if item != lpt.out || err != lpt.outError {
			t.Errorf(
				"%d. %v.PopItem(%d) => %v, %v, want %v, %v",
				index, lpt.in, lpt.index, item, err, lpt.out, lpt.outError)
		}
		if !reflect.DeepEqual(list, lpt.outList) {
			t.Errorf(
				"%d. %v.PopItem(%d) => out list = %v, want %v",
				index, lpt.in, list, lpt.outList)
		}
	}
}

var listRemoveTests = []struct {
	in      List
	val     interface{}
	out     error
	outList List
}{
	{List{"one", "two", "one"}, "one", nil, List{"two", "one"}},
	//{List{"one", "two", "three"}, "zero",
	//	0, errors.New("zero is not in list"), List{"one", "two", "three"}},
	{List{1, 2, 3, 2}, 3, nil, List{1, 2, 2}},
	//{List{}, 1, errors.New("1 is not in list"), List{}},
}

func TestListRemove(t *testing.T) {
	for index, lrt := range listRemoveTests {
		list := lrt.in
		err := list.Remove(lrt.val)
		if err != lrt.out {
			t.Errorf(
				"%d. %v.Remove(%v) => %v, want %v",
				index, lrt.in, lrt.val, err, lrt.out)
		}
		if !reflect.DeepEqual(list, lrt.outList) {
			t.Errorf(
				"%d. %v.Remove(%v) => out list = %v, want %v",
				index, lrt.in, lrt.val, list, lrt.outList)
		}
	}
}

var listReverseTests = []struct {
	in  List
	out List
}{
	{List{"one", "two", "three", "two"}, List{"two", "three", "two", "one"}},
	{List{"one", "two", 2, 1}, List{1, 2, "two", "one"}},
	{List{2}, List{2}},
	{List{}, List{}},
}

func TestListReverse(t *testing.T) {
	for index, lrt := range listReverseTests {
		list := lrt.in
		list.Reverse()
		if !reflect.DeepEqual(list, lrt.out) {
			t.Errorf(
				"%d. %v.Reverse() => out list = %v, want %v",
				index, lrt.in, list, lrt.out)
		}
	}
}

//var listSortTests = []struct {
//	in  List
//	out List
//}{
//	{List{"o", "ze", "a", "two"}, List{"a", "o", "two", "ze"}},
//	{List{2, 3, 1, -2}, List{-2, 1, 2, 3}},
//	{List{1, "one", 2}, List{1, 2, "one"}},
//	{List{2, 1.3, 1}, List{1, 1.3, 2}},
//	{List{1, "1.2", "a", "ola", 34, 2, "1", "12", 1.2, "2", "ala"},
//		List{1, 1.2, 2, 34, "1", "1.2", "12", "2", "a", "ala", "ola"}},
//}

//func TestListSort(t *testing.T) {
//	for index, lst := range listSortTests {
//		list := lst.in
//		list.Sort()
//		if !reflect.DeepEqual(list, lst.out) {
//			t.Errorf(
//				"%d. %v.Sort() => out list = %v, want %v",
//				index, lst.in, list, lst.out)
//		}
//	}
//}
