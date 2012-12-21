// Copyright 2012 Dobrosław Żybort
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

/*
Package simpletype implement similar to Python types: list and dict for Go

This package bring list and dict to Go with most methods you can found 
in Python.

	dict := simpletype.Dict{"one": 1, "two": 2, "three": 3}

	if dict.HasKey("one") {
		// Do something if dict have key "one"
	}
*/
package simpletype
