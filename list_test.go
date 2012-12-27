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

func TestListAppend(t *testing.T) {
	list := List{"one", "two"}
	list.Append("three")

	goodList := List{"one", "two", "three"}
	if !reflect.DeepEqual(list, goodList) {
		t.Errorf("Error when appending, should be %v, got %v", goodList, list)
	}
}

func TestListCount(t *testing.T) {
	list := List{"one", "two", "three", "two"}

	if counted := list.Count("two"); counted != 2 {
		t.Errorf("Error when counting, should be 2, got %v", counted)
	}
	if counted := list.Count("zero"); counted != 0 {
		t.Errorf("Error when counting, should be 0, got %v", counted)
	}
}

func TestListDelete(t *testing.T) {
	list := List{"one", "two", "three"}
	list.Delete(1)

	goodList := List{"one", "three"}
	if !reflect.DeepEqual(list, goodList) {
		t.Errorf("Error when deleting, should be %v, got %v", goodList, list)
	}
}

func TestListExtend(t *testing.T) {
	list1 := List{"one", "two", "three"}
	list2 := List{"four", "five", "six"}
	list1.Extend(list2)

	goodList := List{"one", "two", "three", "four", "five", "six"}
	if !reflect.DeepEqual(list1, goodList) {
		t.Errorf("Error when extending, should be %v, got %v", goodList, list1)
	}
}

func TestListIndex(t *testing.T) {
	list := List{"one", "two", "three"}

	if index, err := list.Index("two"); index != 1 || err != nil {
		t.Errorf("Error when indexing, should be 1, got %v", index)
	}
	if index, err := list.Index("zero"); index != -1 || err == nil {
		t.Errorf("Error when indexing, should be -1, got %v", index)
	}
}

func TestListInsert(t *testing.T) {
	list := List{"one", "two", "three"}
	list.Insert(1, "inserted")

	goodList := List{"one", "inserted", "two", "three"}
	if !reflect.DeepEqual(list, goodList) {
		t.Errorf("Error when inserting, should be %v, got %v", goodList, list)
	}

	list2 := List{"one", "two", "three"}
	list2.Insert(10, "four")

	goodList2 := List{"one", "two", "three", "four"}
	if !reflect.DeepEqual(list, goodList) {
		t.Errorf(
			"Error when inserting, should be %v, got %v",
			goodList2,
			list2,
		)
	}
}

func TestListPop(t *testing.T) {
	list := List{"one", "two", "three"}
	item := list.Pop()

	goodList := List{"one", "two"}
	if item != "three" || !reflect.DeepEqual(list, goodList) {
		t.Errorf("Error when poping, should be 'three', got %v", item)
	}
}

func TestListPopItem(t *testing.T) {
	list := List{"one", "two", "three"}
	item := list.PopItem(0)

	goodList := List{"two", "three"}
	if item != "one" || !reflect.DeepEqual(list, goodList) {
		t.Errorf("Error when poping item, should be 'one', got %v", item)
	}
}

func TestListRemove(t *testing.T) {
	list := List{"one", "two", "three", "two"}
	err := list.Remove("two")

	goodList := List{"one", "three", "two"}
	if err != nil || !reflect.DeepEqual(list, goodList) {
		t.Errorf("Error when removing, should be %v, got %v", goodList, list)
	}

	list2 := List{"one", "two", "three"}
	err2 := list.Remove("zero")

	goodList2 := List{"one", "two", "three"}
	if err2 == nil || !reflect.DeepEqual(list2, goodList2) {
		t.Errorf("Error when removing, should be %v, got %v", goodList2, list2)
	}
}

func TestListReverse(t *testing.T) {
	list := List{"one", "two", "three", 1, 2, 3}
	list.Reverse()

	goodList := List{3, 2, 1, "three", "two", "one"}
	if !reflect.DeepEqual(list, goodList) {
		t.Errorf("Error when reversing, should be %v, got %v", goodList, list)
	}
}

func TestListSort(t *testing.T) {
	list := List{2, 1, 3}
	list.Sort()

	goodList := List{1, 2, 3}
	if !reflect.DeepEqual(list, goodList) {
		t.Errorf("Error when sorting, should be %v, got %v", goodList, list)
	}

	list2 := List{"two", 1, 3}
	list2.Sort()

	goodList2 := List{1, 3, "two"}
	if !reflect.DeepEqual(list2, goodList2) {
		t.Errorf("Error when sorting, should be %v, got %v", goodList2, list2)
	}
}
