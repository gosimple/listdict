// Copyright 2012 Dobrosław Żybort
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package simpletype

// Simple list
type List []interface{}

// Return new List with specified length
func NewList(length int) List {
	return make(List, length)
}

//=============================================================================
