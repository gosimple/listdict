ListDict
==========

Package ListDict implement Python types: list and dict for Go.

[Documentation online](http://godoc.org/bitbucket.org/matrixik/listdict)

This package bring list and dict to Go with most methods you can find 
in Python.

	dict := listdict.Dict{"one": 1, "two": 2, "three": 3}

	if dict.HasKey("one") {
		// Do something if dict have key "one"
	}
	
	keys := dict.Keys()		// keys = [one two three]
	val := dict.Values()	// val = [3 1 2]
	// Keys() and Values() are unordered, same as in Python

### Requests or bugs? 
<https://bitbucket.org/matrixik/listdict/issues>

## Installation

	go get bitbucket.org/matrixik/listdict

## License

The source files are distributed under the 
[Mozilla Public License, version 2.0](http://mozilla.org/MPL/2.0/),
unless otherwise noted.  
Please read the [FAQ](http://www.mozilla.org/MPL/2.0/FAQ.html)
if you have further questions regarding the license.
