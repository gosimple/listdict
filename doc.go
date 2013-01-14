// Copyright 2012 Dobrosław Żybort
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

/*
Python List and Dict for Go

This package bring list and dict to Go with most methods you can find
in Python.

	dict := listdict.Dict{"one": 1, "two": 2, "three": 3}

	if dict.HasKey("one") {
		// Do something if dict have key "one"
	}

	keys := dict.Keys()		// keys = [one two three]
	val := dict.Values()	// val = [3 1 2]
	// Keys() and Values() are unordered, same as in Python

Requests or bugs?
https://bitbucket.org/gosimple/listdict/issues
*/
package listdict
