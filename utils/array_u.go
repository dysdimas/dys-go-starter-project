package utils

import "fmt"

func RemoveElementPreserveIndex(s []interface{}, i int) ([]interface{}, error) {
	if i >= len(s) || i < 0 {
		return nil, fmt.Errorf("Index is out of range. Index is %d with slice length %d", i, len(s))
	}
	newSlice := make([]interface{}, 0)
	newSlice = append(newSlice, s[:i]...)
	return append(newSlice, s[i+1:]...), nil
}

func RemoveElementMapStringPreserveIndex(s []map[string]interface{}, i int) ([]map[string]interface{}, error) {
	if i >= len(s) || i < 0 {
		return nil, fmt.Errorf("Index is out of range. Index is %d with slice length %d", i, len(s))
	}
	newSlice := make([]map[string]interface{}, 0)
	newSlice = append(newSlice, s[:i]...)
	return append(newSlice, s[i+1:]...), nil
}

func SliceContains[K int64 | string](elems []K, v K) bool {
	for _, s := range elems {
		if v == s {
			return true
		}
	}
	return false
}
