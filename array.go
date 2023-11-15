package array

import (
	"sort"
)

// Append a value to the array
func Append(array []any, value any) []any {
	array = append(array, value)
	return array
}

// Contain test if the value is in the array return True if inside and False if not inside
func Contain(array []any, value any) bool {
	size := Length(array)
	for i := 0; i < size; i++ {
		if array[i] == value {
			return true
		}
	}
	return false
}

// Find returns the index of the value in the array
func Find(array []any, value any) int {
	size := Length(array)
	for i := 0; i < size; i++ {
		if array[i] == value {
			return i
		}
	}
	return -1
}

// IsEqual test two array and return true if same or false if different
func IsEqual(FirstArray, SecondArray []any) bool {
	size := Length(FirstArray)
	if Length(FirstArray) != Length(SecondArray) {
		return false
	} else {
		for i := 0; i < size; i++ {
			if FirstArray[i] != SecondArray[i] {
				return false
			}
		}
		return true
	}
}

// Length returns the length of the array
func Length(array []any) int {
	return len(array)
}

// Max returns the maximum value in the array
func Max(array []any) any {
	size := Length(array)
	if size == 0 {
		return nil
	}

	max := array[0]
	for i := 1; i < size; i++ {
		if lessThan(max, array[i]) {
			max = array[i]
		}
	}
	return max
}

// Min returns the minimum value in the array
func Min(array []any) any {
	size := Length(array)
	if size == 0 {
		return nil
	}
	min := array[0]
	for i := 1; i < size; i++ {
		if lessThan(array[i], min) {
			min = array[i]
		}
	}
	return min
}

func lessThan(a any, b any) bool {
	switch a.(type) {
	case int:
		return a.(int) < b.(int)
	case float64:
		return a.(float64) < b.(float64)
	case string:
		return len(a.(string)) < len(b.(string))
	default:
		return false
	}
}

// Remove the value at the index
func Remove(array []any, index int) []any {
	if index < 0 || index >= len(array) {
		return array
	}
	array = append(array[:index], array[index+1:]...)
	return array
}

// Slice the array from start to end included
func Slice(array []any, start, end int) []any {
	var arrRet []any

	if start < 0 {
		start = 0
	}
	if end > Length(array)-1 {
		end = Length(array) - 1
	}
	if start >= end {
		panic("Error | Start superior or equal to the end")
	}
	for i := start; i <= end; i++ {
		arrRet = append(arrRet, array[i])
	}
	return arrRet
}

// SortAsc sorts the array in ascending order
func SortAsc(array []any) {
	sort.Slice(array, func(i, j int) bool {
		switch array[i].(type) {
		case int:
			return array[i].(int) < array[j].(int)
		case float64:
			return array[i].(float64) < array[j].(float64)
		case string:
			return array[i].(string) < array[j].(string)
		default:
			return false
		}
	})
}

// SortDesc sorts the array in descending order
func SortDesc(array []any) {
	sort.Slice(array, func(i, j int) bool {
		switch array[i].(type) {
		case int:
			return array[i].(int) > array[j].(int)
		case float64:
			return array[i].(float64) > array[j].(float64)
		case string:
			return array[i].(string) > array[j].(string)
		default:
			return false
		}
	})
}
