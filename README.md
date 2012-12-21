simpletype
==========

Package simpletype implement similar to Python types: list and dict for Go

This package bring list and dict to Go with most methods you can found 
in Python.

	dict := simpletype.Dict{"one": 1, "two": 2, "three": 3}

	if dict.HasKey("one") {
		// Do something if dict have key "one"
	}