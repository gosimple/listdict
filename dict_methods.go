// simpletype project dict_methods.go
package simpletype

//=============================================================================

// Removes all elements from the dictionary
func (dict Dict) Clear() {
	for key, _ := range dict {
		delete(dict, key)
	}
}

// Return value for the given key or val if key is not in the dictionary
func (dict Dict) Get(key string, val interface{}) interface{} {
	if dict.HasKey(key) {
		return dict[key]
	}
	return val
}

// Returns true if key is in the dictionary, false otherwise
func (dict Dict) HasKey(key string) bool {
	if _, ok := dict[key]; ok {
		return true
	}
	return false
}

// Returns a list of the dictionary's [key, value] list pairs
func (dict Dict) Items() List {
	list := NewList(len(dict))
	i := 0
	for key, value := range dict {
		list[i] = List{key, value}
		i++
	}
	return list
}

// Returns a list of the dictionary's keys
func (dict Dict) Keys() List {
	list := NewList(len(dict))
	i := 0
	for key, _ := range dict {
		list[i] = key
		i++
	}
	return list
}

// Similar to Get(), but will set dict[key]=val if key is not already in dict
func (dict Dict) SetDefault(key string, val interface{}) interface{} {
	if dict.HasKey(key) {
		return dict[key]
	}
	dict[key] = val
	return val
}

// Update the dictionary with the key-value pairs in the dict2 dictionary
// replacing current values and adding new if found.
func (dict Dict) Update(dict2 Dict) {
	for key, value := range dict2 {
		dict[key] = value
	}
}

//Returns a list of the dictionary's values
func (dict Dict) Values() List {
	list := NewList(len(dict))
	i := 0
	for _, value := range dict {
		list[i] = value
		i++
	}
	return list
}
