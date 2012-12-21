// simpletype project simpletype.go
package simpletype

// Simple list
type List []interface{}

// Return new List with specified length
func NewList(length int) List {
	return make(List, length)
}

// Simple dict
type Dict map[string]interface{}

// Return new Dict
func NewDict() Dict {
	return make(Dict)
}

//=============================================================================

//Create a new dictionary with keys from list and values set to val.
func DictFromKeys(list List, val interface{}) Dict {
	newDict := NewDict()
	for _, value := range list {
		newDict[value.(string)] = val
	}
	return newDict
}
